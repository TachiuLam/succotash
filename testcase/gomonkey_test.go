package testcase

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
)

type Server struct {
}

func (s *Server) Logout() int {
	return 0
}

//go:noinline 取消内联
func Logout() int {
	return 0
}

func TestGoMonkeyMethod(t *testing.T) {
	s := &Server{}
	patch := gomonkey.ApplyMethod(reflect.TypeOf(&Server{}), "Logout", func(s *Server) int {
		return 111
	})
	defer patch.Reset()
	t.Errorf("logout: %d", s.Logout())
}

func TestGoMonkeyFunc(t *testing.T) {
	patch := gomonkey.ApplyFunc(Logout, func() int {
		return 111
	})
	defer patch.Reset()
	t.Errorf("logout: %d", Logout())
}
