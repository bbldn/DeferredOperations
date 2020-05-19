package response

import (
	"encoding/json"
)

type Response struct {
	Ok     bool        `json:"ok"`
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}

func (r Response) ToJson() string {
	response := make(map[string]interface{})
	response["ok"] = r.Ok

	if nil != r.Data {
		response["data"] = r.Data
	}

	if nil != r.Errors {
		response["errors"] = r.Errors
	}

	result, _ := json.Marshal(response)

	return string(result)
}
