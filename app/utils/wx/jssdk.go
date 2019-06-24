package wx

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
	"js_ticket_service/app/utils/request"
	string2 "js_ticket_service/app/utils/string"
	"js_ticket_service/database"
	"strconv"
	"strings"
	"sync"
	"time"
)

const GET_TICKET_URL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"

var Jsconfig JsConfig

type Wx struct {
	AppId       string
	AccessToken string
	Lock        sync.Mutex
}

type JsConfig struct {
	NonceStr  string `json:"nonce_str"`
	TimeStamp int64  `json:"timestamp"`
	AppId     string `json:"appid"`
	Signature string `json:"signature"`
}

func (wx *Wx) GetJsConfig(url string) (JsConfig, error) {
	jsTicket, err := wx.getTicket()
	if err != nil {
		return Jsconfig, err
	}
	nonceStr := string2.GetNonceStr(10)
	now := time.Now().Unix()
	signature := wx.WxConfigSign(string(jsTicket), nonceStr, strconv.Itoa(int(now)), url)

	Jsconfig = JsConfig{
		AppId:     wx.AppId,
		NonceStr:  nonceStr,
		TimeStamp: now,
		Signature: signature,
	}
	return Jsconfig, nil
}

func (wx *Wx) getTicket() ([]byte, error) {
	key := wx.getTicketByAppId(wx.AppId)
	wx.Lock.Lock()
	defer wx.Lock.Unlock()
	res, err := database.Get(key)
	if err != nil || res == nil {
		res, err = wx.refreshTicket()
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (wx *Wx) getTicketByAppId(appId string) string {
	return "s:js_ticket:ticket:" + appId
}

func (wx *Wx) refreshTicket() ([]byte, error) {
	key := wx.getTicketByAppId(wx.AppId)
	res, err := request.HttpGet(fmt.Sprintf(GET_TICKET_URL, wx.AccessToken))
	if err != nil {
		return nil, err
	}
	errcode, err := jsonparser.GetInt(res, "errcode")
	if err != nil {
		return nil, err
	}
	if errcode != 0 {
		return nil, errors.New(string(res))
	}
	jsTicket, err := jsonparser.GetString(res, "ticket")

	err = database.Set(key, jsTicket)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (wx *Wx) WxConfigSign(jsapiTicket, nonceStr, timestamp, url string) (signature string) {
	if i := strings.IndexByte(url, '#'); i >= 0 {
		url = url[:i]
	}

	n := len("jsapi_ticket=") + len(jsapiTicket) +
		len("&noncestr=") + len(nonceStr) +
		len("&timestamp=") + len(timestamp) +
		len("&url=") + len(url)
	buf := make([]byte, 0, n)

	buf = append(buf, "jsapi_ticket="...)
	buf = append(buf, jsapiTicket...)
	buf = append(buf, "&noncestr="...)
	buf = append(buf, nonceStr...)
	buf = append(buf, "&timestamp="...)
	buf = append(buf, timestamp...)
	buf = append(buf, "&url="...)
	buf = append(buf, url...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}
