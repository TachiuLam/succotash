package testcase

// go;generate mockgen -source=server.go -destination=./mock/mock_server.go -package=mock
type WebServer interface {
	Login(username, password string) int
	Logout() int
}
