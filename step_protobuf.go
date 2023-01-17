package main

import (
	"context"
)

type protobufStep struct {
	*plugin
}

func newProtobufStep(plugin *plugin) *protobufStep {
	return &protobufStep{
		plugin: plugin,
	}
}

func (s *protobufStep) Runnable() bool {
	return 0 != len(s.Pbs)
}

func (s *protobufStep) Run(_ context.Context) (err error) {
	for _, _pb := range s.Pbs {
		if nil != _pb.Enabled && *_pb.Enabled {
			err = _pb.upload(s.plugin)
		}

		if nil != err {
			return
		}
	}

	return
}
