package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"seuxw/embrice/entity/test"
	"seuxw/embrice/extension"
)

func (self *server) TestGet(w http.ResponseWriter, r *http.Request) {
	var (
		getPara string
		Test    test.Test
		err     error
	)

	// 获取date参数
	getPara = r.FormValue("date")
	if len(getPara) == 0 {
		getPara = extension.CurrentDateInStr()
	}

	self.PrintTrace(r, fmt.Sprintf("date:%s", getPara), "GET：开始调用测试接口")

	// 数据库操作
	if Test, err = self.db.Test(getPara); err != nil {
		err = fmt.Errorf("[Test]数据库调用错误！ %s", err)
		goto END
	}
END:
	extension.EndLabel(self.log, err, w, nil, Test)
}

func (self *server) TestPost(w http.ResponseWriter, r *http.Request) {
	var (
		getPara string
		Test    test.Test
		err     error
	)

	// 获取date参数
	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&Test); err != nil {
		err = fmt.Errorf("[Test] Json decode err: %v", err)
		goto END
	}

	getPara = Test.Date
	if len(getPara) == 0 {
		getPara = extension.CurrentDateInStr()
	}

	self.PrintTrace(r, fmt.Sprintf("date:%s", getPara), "POST：开始调用测试接口")

	// 数据库操作
	if Test, err = self.db.Test(getPara); err != nil {
		err = fmt.Errorf("[Test]数据库调用错误！ %s", err)
		goto END
	}
END:
	extension.EndLabel(self.log, err, w, nil, Test)
}
