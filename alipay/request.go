package alipay

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"hash"
	"net/http"
	"pay/common"
	"time"
)

type commonReqParam struct {
	AppID   string `json:"app_id"`
	Method  string `json:"method"`
	Format  string `json:"format"`
	Charset string `json:"charset"`
	//SignType string `json:"sign_type"`
	Sign         string `json:"sign"`
	Timestamp    string `json:"timestamp"`
	Version      string `json:"version"`
	AppAuthToken string `json:"app_auth_token"`
	BizContent   string `json:"biz_content"`
	NotifyUrl    string `json:"notify_url"`
	ReturnUrl    string `json:"return_url"`
}

//数据组装
func (alipay *Alipay) assemble() {

}

//参数签名
func (alipay *Alipay) sign(param []byte) (string, error) {
	block, _ := pem.Decode(alipay.PrivateKey)
	if block == nil {
		return "", common.ErrMsg("pem decode error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	//var h
	var h hash.Hash
	if alipay.SignType == crypto.SHA256 {
		//采用RSA256
		h = sha256.New()
		h.Write(param)
	} else {
		//默认采用RSA1签名方式
		h = sha1.New()
		h.Write(param)
	}
	digest := h.Sum(nil)
	s, err := rsa.SignPKCS1v15(nil, privateKey, alipay.SignType, digest)
	if err != nil {
		return "", err
	}
	data := base64.StdEncoding.EncodeToString(s)
	return data, nil
}

func (alipay *Alipay) handleRequest(params interface{}) (*http.Response, error) {
	//构建request请求
	content, err := json.Marshal(params)
	if err != nil {
		return nil, common.ErrMsg("handleRequest|Marshal")
	}
	alipay.commonReqParam.BizContent = string(content)
	alipay.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	//asicc排序
	req, _ := json.Marshal(alipay.commonReqParam)
	arr, sortStr, err := common.AsciiSort(req)
	if err != nil {
		return nil, err
	}
	//签名
	signParam, err := alipay.sign([]byte(sortStr))
	if err != nil {
		return nil, err
	}
	alipay.Sign = signParam
	request, err := http.NewRequest(http.MethodPost, gateway, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for key, value := range arr {
		request.Form.Add(key, value)
	}
	return alipay.client.Do(request)
}
