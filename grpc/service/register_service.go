package service

import (
	"log"
	"time"
)

// Info 服务描述信息
type Info struct {
	//服务名称
	ServiceName string
	//ip地址
	Host string
	//端口
	Port int
	//心跳间隔 秒
	IntervalTime time.Duration
}

// RegisterService 服务注册和下线的接口
type RegisterService interface {
	// Register 服务注册
	Register(info Info) error
	// UnRegister 服务下线
	UnRegister(info Info) error
}

type registerService struct {
	Address string
}

func NewRegisterService(address string) RegisterService {
	return &registerService{
		Address: address,
	}
}

func (s *registerService) Register(info Info) error {
	log.Printf("service %s:%s register...", info.ServiceName, info.Host)
	return nil
}

func (s *registerService) UnRegister(info Info) error {
	log.Printf("service %s:%s unregister...", info.ServiceName, info.Host)
	return nil
}
