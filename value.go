package main

type value struct {
	Content     string `json:"content"`
	UpdateTime  int64  `json:"update_time"`
	CreateTime  int64  `json:"create_time"`
	Description string `json:"desc"`
}
