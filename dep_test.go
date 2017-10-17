package dig

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Config struct {
	Prefix string
	Port   int
}

type Server struct {
}

func NewLogger(cfg *Config) *log.Logger {
	return log.New(os.Stdout, cfg.Prefix, 0)
}

func NewConfig() *Config {
	return &Config{"[test]", 80}
}

func TestDepContainer(t *testing.T) {
	base := New()
	require.NoError(t, base.Provide(NewLogger), "provide failed")
	require.NoError(t, base.Provide(NewConfig), "provide failed")

	dNil := NewWithParent(nil)
	require.NotNil(t, dNil)
	assert.Equal(t, 0, len(dNil.parents))

	derived := NewWithParent(base)
	assert.NotNil(t, derived, "nil container")
	testProvide(t, base, derived)

	anotherDerived := NewWithParent(base)
	assert.NotNil(t, derived, "nil container")
	testProvide(t, base, anotherDerived)

	// Test multiple parents
	mpc1 := NewWithParents([]*Container{base, derived, anotherDerived})
	assert.NotNil(t, mpc1, "nil multi-parent container")
	assert.Equal(t, 3, len(mpc1.parents))

	mpc2 := NewWithParents([]*Container{nil, base, base})
	assert.NotNil(t, mpc2, "nil multi-parent container")
	assert.Equal(t, 1, len(mpc2.parents))
}

func testProvide(t *testing.T, base *Container, derived *Container) {
	require.Equal(t, base, derived.parents[0], "wrong parent")
	getServer := func(cfg *Config, l *log.Logger) *Server {
		assert.NotNil(t, cfg, "nil config")
		assert.NotNil(t, l, "nil logger")
		return &Server{}
	}
	require.NoError(t, derived.Provide(getServer), "derived provide failed")
	useServer := func(s *Server) {}
	require.NoError(t, derived.Invoke(useServer), "Server invoke failed")

	// Check the reprovide case
	require.Error(t, derived.Provide(func() *Config { return nil }), "reprovide must fail")
}
