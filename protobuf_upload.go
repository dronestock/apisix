package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
)

func (p *protobuf) upload(plugin *plugin) (err error) {
	req := plugin.Http()
	if content, re := os.ReadFile(p.Source); nil != re {
		err = re
	} else {
		pr := new(protobufReq)
		pr.Content = base64.StdEncoding.EncodeToString(content)
		req.SetBody(pr)
	}
	if nil != err {
		return
	}

	rsp := new(protobufRsp)
	url := fmt.Sprintf("%s/apisix/admin/protos/%s", plugin.Endpoint, p.Id)
	if hr, he := req.SetResult(rsp).Put(url); nil != he {
		err = he
	} else if hr.IsError() {
		err = exc.NewFields(
			"Apisix返回错误",
			field.New("url", hr.Request.URL),
			field.New("code", hr.StatusCode()),
			field.New("body", string(hr.Body())),
		)
	} else {
		plugin.Debug("上传Protobuf文件到Apisix成功", field.New("rsp", rsp))
	}

	return
}
