package string

import (
	"math/rand"
	"time"
)

func GetNonceStr(num int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var res []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<num ;i++  {
		res = append(res,bytes[r.Intn(len(bytes))])
	}
	return string(res)

}
