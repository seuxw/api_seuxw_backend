package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "regexp"
	// "seuxw/embrice/constant/constuser"
	"seuxw/embrice/entity"
	"seuxw/embrice/entity/user"
	"seuxw/embrice/extension"
)

func (svr *server) GetUserByUUID(w http.ResponseWriter, r *http.Request) {

	var (
		response entity.Response
		err      error
		data     []byte
		uuid     string
		reqData  user.GetUserByUUIDReq
		respData *user.GetUserByUUIDResp
	)

	seuxwRequest := entity.GetSeuxwRequest(r)
	processName := "GetUserByUUID"

	body, _ := extension.HandlerRequestLog(seuxwRequest, processName)

	if len(body) == 0 {
		goto GET
	} else {
		goto POST
	}

GET:
	uuid = r.FormValue("uuid")
	goto DB

POST:
	if err = json.Unmarshal(body, &reqData); err != nil {
		goto END
	}
	uuid = reqData.UUID

DB:
	if len(uuid) < 10 {
		err = fmt.Errorf("UUID 长度非法")
		goto END
	}

	if respData, err = svr.db.GetUserByUUIDDB(uuid); err != nil {
		goto END
	}
	response.Data = respData

END:
	if err != nil {
		response.Code = 1
		response.Message = fmt.Sprintf("%s", err)
	}

	data = extension.HandlerResponseLog(seuxwRequest, response, processName, true)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(data))
}
