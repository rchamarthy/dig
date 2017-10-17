package dig

import (
	"log"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Service interface {
	OnConfig() error
	OnStart() error
	OnStop() error
}

type LogService struct {
	l *log.Logger
}

func NewLogService(l *log.Logger) *LogService {
	return &LogService{l}
}

func (l *LogService) OnConfig() error {
	return nil
}

func (l *LogService) OnStart() error {
	return nil
}

func (l *LogService) OnStop() error {
	return nil
}

type HTTPService struct {
	l *log.Logger
}

func NewHTTPService(l *log.Logger) *HTTPService {
	return &HTTPService{l}
}

func (h *HTTPService) OnConfig() error {
	return nil
}

func (h *HTTPService) OnStart() error {
	return nil
}

func (h *HTTPService) OnStop() error {
	return nil
}

type OptHandler struct {
	services []Service
}

func (oh *OptHandler) ProcessResults(results []reflect.Value) error {
	if oh.services == nil {
		oh.services = []Service{}
	}
	for _, v := range results {
		if s, ok := v.Interface().(Service); ok {
			oh.services = append(oh.services, s)
		}
	}
	return nil
}

func TestOptContainer(t *testing.T) {
	opt := &OptHandler{}
	base := New(opt)
	assert.NotNil(t, base)
	require.NoError(t, base.Provide(NewConfig))
	require.NoError(t, base.Provide(NewLogger))
	require.NoError(t, base.Provide(NewLogService))
	require.NoError(t, base.Invoke(func(ls *LogService) {}))
	assert.Equal(t, 1, len(opt.services))

	derived := NewWithParent(base, opt)
	assert.NotNil(t, derived)
	require.NoError(t, derived.Provide(NewHTTPService))

	require.NoError(t, derived.Invoke(func(hs *HTTPService) {}))
	assert.Equal(t, 2, len(opt.services))

	for _, s := range opt.services {
		s.OnConfig()
		s.OnStart()
		s.OnStop()
	}
}
