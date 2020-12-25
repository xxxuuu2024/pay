package alipay

import (
	"crypto"
	"net/http"
	"time"
)

const (
	gateway = "https://openapi.alipay.com/gateway.do"
)

type Config struct {
	AppID    string      `json:"app_id"`
	SignType crypto.Hash `json:"sign_type"`
	//私钥路径
	PrivateKey []byte `json:"private_key"`
}
type Alipay struct {
	Config
	client *http.Client
	*commonReqParam
}

func New(config Config) Alipay {
	return Alipay{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		commonReqParam: &commonReqParam{
			AppID: config.AppID,
			//SignType: "RSA2",
			Format:  "json",
			Version: "1.0",
			Charset: "utf-8",
		},
	}
}
