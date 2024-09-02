package eval

import "strconv"

var RespNIL []byte = []byte("$-1\r\n")
var RespOK []byte = []byte("+OK\r\n")
var RespZero []byte = []byte(":0\r\n")
var RespOne []byte = []byte(":1\r\n")

func GetIntResponse(val int64) []byte {
	return []byte(":" + strconv.FormatInt(val, 10) + "\r\n")
}
