package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Settings struct {
	AuthUrl string `fig:"auth_url"`
	BaseUrl string `fig:"base_url"`
}

type URL interface {
	URL() *Settings
}

type urler struct {
	getter kv.Getter
	once   comfig.Once
}

func NewUrler(getter kv.Getter) URL {
	return &urler{
		getter: getter,
	}
}

func (c *urler) URL() *Settings {
	return c.once.Do(func() interface{} {
		var cfg Settings

		if err := figure.Out(&cfg).From(kv.MustGetStringMap(c.getter, "settings")).Please(); err != nil {
			panic(err)
		}

		return &cfg
	}).(*Settings)
}
