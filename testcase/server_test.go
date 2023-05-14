package testcase

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/TachiuLam/succotash/testcase/mock"
)

func TestMockServeLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockServer := mock.NewMockWebServer(ctrl)
	user, pass := "Li", "123456"
	mockServer.EXPECT().Login(user, pass).Return(0)

	{
		output := mockServer.Login("Li", "123456")
		t.Errorf("%d", output)
	}
	{
		output := mockServer.Login("Li", "222")
		t.Errorf("%d", output)
	}
}

func TestMockServeLogout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockServer := mock.NewMockWebServer(ctrl)

	user, pass := "Li", "123456"
	sub1 := mockServer.EXPECT().Login(user, pass).Return(0)
	sub2 := mockServer.EXPECT().Logout().Return(0)
	gomock.InOrder(sub1, sub2)

	output := mockServer.Logout()
	t.Errorf("%d", output)
	output = mockServer.Login("Li", "123456")
	t.Errorf("%d", output)
}
