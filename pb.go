package main

type pb struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled"`
	// 编号
	Id string `json:"id" validate:"required"`
	// 源文件
	Source string `json:"source" validate:"required"`
}

func (p *plugin) pb() (undo bool, err error) {
	if undo = !*p.Pb.Enabled || 0 == len(p.Pbs); undo {
		return
	}

	for _, _pb := range p.Pbs {
		if err = _pb.upload(p); nil != err {
			return
		}
	}

	return
}
