package main

type pb struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled"`
	// 编号
	Id string `json:"id" validate:"required"`
	// 源文件
	Source string `json:"source" validate:"required"`
	// 描述信息
	Description string `default:"这家伙很懒，没配置，相关Leader应该严肃对待此事，保证不要乱整！" json:"description"`
}
