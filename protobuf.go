package main

type protobuf struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled"`
	// 编号
	Id string `json:"id" validate:"required"`
	// 源文件
	Source string `json:"source" validate:"required"`
}

func (p *plugin) protobuf() (undo bool, err error) {
	if undo = !*p.Protobuf.Enabled || 0 == len(p.Protobufs); undo {
		return
	}

	for _, _protobuf := range p.Protobufs {
		if err = _protobuf.upload(p); nil != err {
			return
		}
	}

	return
}
