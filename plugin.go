package main

import (
	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	// 端点
	Endpoint string `default:"${ENDPOINT}" validate:"required"`
	// 接口密钥
	ApiKey string `default:"${API_KEY}" validate:"required"`

	// Protobuf文件上传
	Pb *pb `default:"${PB}"`
	// Protobuf文件上传列表
	Pbs []*pb `default:"${PBS}"`
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Setup() (unset bool, err error) {
	if nil != p.Pb {
		if nil == p.Pb {
			p.Pbs = make([]*pb, 0, 1)
		}
		p.Pbs = append(p.Pbs, p.Pb)
	}

	return
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(p.pb, drone.Name(`Protobuf`)),
	}
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New(`endpoint`, p.Endpoint),
	}
}
