package main

import (
	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	// 端点
	Endpoint string `default:"{ENDPOINT}}" validate:"required"`
	// 接口密钥
	ApiKey string `default:"${API_KEY}" validate:"required"`

	// Protobuf文件上传
	Protobuf *protobuf `default:"PROTOBUF"`
	// Protobuf文件上传列表
	Protobufs []*protobuf `default:"PROTOBUFS"`
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(p.protobuf, drone.Name(`启动守护进程`)),
	}
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New(`endpoint`, p.Endpoint),
	}
}
