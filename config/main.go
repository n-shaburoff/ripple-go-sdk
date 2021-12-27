package config

import (
	"gitlab.com/distributed_lab/kit/kv"
)

type Config interface {
	ODL
}

type config struct {
	ODL
	getter kv.Getter
}

func NewConfig(getter kv.Getter) Config {
	return &config{
		getter: getter,
		ODL:    NewRippleODlClient(getter),
	}
}
