package alipay

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"pay/common"
	"time"
)

const (
	//gateway="https://openapi.alipay.com/gateway.do"
	gateway = "https://openapi.alipaydev.com/gateway.do"
)

type commonParams struct {
	AppID        string `json:"app_id"`
	Method       string `json:"method"`
	Format       string `json:"format"`
	Charset      string `json:"charset"`
	Sign         string `json:"sign"`
	Timestamp    string `json:"timestamp"`
	Version      string `json:"version"`
	AppAuthToken string `json:"app_auth_token"`
	BizContent   string `json:"biz_content"`
	NotifyUrl    string `json:"notify_url"`
	ReturnUrl    string `json:"return_url"`
	SignType     string `json:"sign_type"`
}

type Request struct {
	commonParams
	client     *http.Client
	privateKey *rsa.PrivateKey
	pubKey     *rsa.PublicKey
}

//参数签名
func (req *Request) sign(param []byte) (string, error) {

	if req.SignType == "RSA2" {

		return common.SHA256Sign(param, req.privateKey)

	}

	return common.SHASign(param, req.privateKey)
}

//签名验证

func (req *Request) verifySign(param []byte, sign string) (bool, error) {
	//map
	respParam := make(map[string]string, 2)
	if err := json.Unmarshal(param, &respParam); err != nil {
		return false, err
	}
	delete(respParam, "sign")
	//delete(respParam,"sign_type")
	asciiStr, _ := json.Marshal(respParam)
	_, sortStr, err := common.AsciiSort(asciiStr)
	if err != nil {
		return false, err
	}
	signstr, err := req.sign([]byte(sortStr))
	if err != nil {
		return false, err
	}
	if sign == signstr {
		return true, nil
	}

	return false, common.ErrMsg(fmt.Sprintf("alipay:%s,self:%s", sign, signstr))

}

func (req *Request) createRequest(params interface{}, method string) (*http.Response, error) {
	//构建request请求
	content, err := json.Marshal(params)
	if err != nil {
		return nil, common.ErrMsg("handleRequest|Marshal")
	}
	req.commonParams.BizContent = string(content)
	req.commonParams.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	req.commonParams.Method = method
	//asicc排序
	data, _ := json.Marshal(req.commonParams)
	arr, sortStr, err := common.AsciiSort(data)
	if err != nil {
		return nil, err
	}
	//签名
	signParam, err := req.sign([]byte(sortStr))
	if err != nil {
		return nil, err
	}
	req.Sign = signParam
	arr["sign"] = req.Sign
	requestUrl := url.Values{}
	for k, v := range arr {
		requestUrl.Add(k, v)
	}
	request, err := http.NewRequest(http.MethodGet, gateway+"?"+requestUrl.Encode(), nil)
	if err != nil {
		return nil, err
	}
	return req.client.Do(request)
}
