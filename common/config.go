package common

import (
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/etcd"
)

type ConfigServer struct {
	source.Source
}

func NewConfigServer(address string, prefix string) *ConfigServer {
	return &ConfigServer{
		Source: etcd.NewSource(
			etcd.WithAddress(address),
			etcd.WithPrefix(prefix),
		),
	}
}
