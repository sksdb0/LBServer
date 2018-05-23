package common

import (
	"encoding/json"
	"logger"
	"net/http"
)

func GetBuffer(req *http.Request, buf []byte) {
	for {
		_, err := req.Body.Read(buf)
		if err != nil {
			break
		}
	}
}

func Unmarshal(buf []byte, i interface{}) bool {
	err := json.Unmarshal(buf, i)
	if err != nil {
		logger.PRINTLINE("Unmarshal error: ", err)
		return false
	}
	return true
}
