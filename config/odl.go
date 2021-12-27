package config

import (
	ripplegosdk "github.com/n-shaburoff/ripple-go-sdk"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type ODL interface {
	ODL() ripplegosdk.Client
}

type odler struct {
	getter kv.Getter
	once   comfig.Once
}

func NewODlClient(getter kv.Getter) ODL {
	return &odler{
		getter: getter,
	}
}

func (c *odler) ODL() ripplegosdk.Client {
	return c.once.Do(func() interface{} {

		var cfg struct {
			GrantType    string `fig:"grant_type"`
			ClientID     string `fig:"client_id"`
			ClientSecret string `fig:"client_secret"`
			Audience     string `fig:"audience"`
			AuthUrl      string `fig:"auth_url"`
			BaseUrl      string `fig:"base_url"`
		}

		if err := figure.Out(&cfg).From(kv.MustGetStringMap(c.getter, "odl")).Please(); err != nil {
			panic(err)
		}

		servicer := ripplegosdk.NewServicer(cfg.AuthUrl, cfg.BaseUrl, cfg.GrantType, cfg.ClientID, cfg.ClientSecret,
			cfg.Audience)

		odlClient, err := ripplegosdk.NewClient(servicer)
		if err != nil {
			panic(err)
		}

		return odlClient
	}).(ripplegosdk.Client)
}
