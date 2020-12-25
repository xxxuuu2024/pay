package alipay

const (
	format  = "json"
	charset = "utf-8"
	version = "1.0"
)

type Config struct {
	AppID    string `json:"app_id"`
	SignType string `json:"sign_type"`
}
