package alipay

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"pay/common"
	"time"
)

type AlipayTrade struct {
	request *Request
}
type AlipayConf struct {
	AppID            string
	SignType         string
	PriveKeyPath     string
	Gateway          string
	AlipayPubKeyPath string
	NotifyUrl        string
	CallBackUrl      string
}

func New(conf AlipayConf) *AlipayTrade {

	switch conf.SignType {
	case "RSA2":
	case "RSA":
		panic("RSA is not supported because it is not secure")
	default:
		panic("not found " + conf.SignType)
	}
	rsaKey, err := common.PrivateKeyDecode(conf.PriveKeyPath)
	if err != nil {
		panic(err)
	}
	rsaPubKey, err := common.PubKeyDecode(conf.AlipayPubKeyPath)
	if err != nil {
		panic(err)
	}
	req := Request{
		client: &http.Client{
			Timeout: time.Second * 60,
		},
		privateKey: rsaKey,
		pubKey:     rsaPubKey,
	}
	req.commonParams.AppID = conf.AppID
	req.commonParams.SignType = conf.SignType
	req.commonParams.Format = "json"
	req.commonParams.Version = "1.0"
	req.commonParams.Charset = "utf-8"
	req.commonParams.NotifyUrl = conf.NotifyUrl
	req.commonParams.ReturnUrl = conf.CallBackUrl
	return &AlipayTrade{
		request: &req,
	}
}

func (trade *AlipayTrade) CreatePreTradeRequest(input CreateTradeInput) (CreatePreTradeOutPut, error) {

	resp, err := trade.tradeReq(input, aliTradeCreate)
	if err != nil {
		return CreatePreTradeOutPut{}, err
	}
	var output CreatePreTradeOutPut
	err = resp.handle(&output)
	if err != nil {
		return CreatePreTradeOutPut{}, err
	}
	output.commonResponse = resp
	return output, nil
}
func (trade *AlipayTrade) CancelTradeReq(input CancelTradeInput) (CancelTradeOutInput, error) {
	resp, err := trade.tradeReq(input, aliTradeCancel)
	if err != nil {
		return CancelTradeOutInput{}, err
	}
	var output CancelTradeOutInput
	err = resp.handle(&output)
	if err != nil {
		return CancelTradeOutInput{}, err
	}
	output.commonResponse = resp
	return output, nil

}

func (trade *AlipayTrade) TradePagePayReq(input TradePagePayInput) (TradePagePayOutInput, error) {

	uri, err := trade.request.assemble(input, "alipay.trade.page.pay")
	if err != nil {
		return TradePagePayOutInput{}, err
	}

	return TradePagePayOutInput{PayUrl: uri}, nil

}

//异步通知处理
func (trade *AlipayTrade) TradeNotify(req []byte) (Notification, error) {
	urlVal, err := url.ParseQuery(string(req))
	if err != nil {
		return Notification{}, err
	}
	result := make(map[string]interface{}, 2)
	for key, _ := range urlVal {
		result[key] = urlVal.Get(key)
	}
	sign := result["sign"].(string)
	delete(result, "sign")
	delete(result, "sign_type")
	notifyByte, _ := json.Marshal(result)
	_, wsign, err := common.AsciiSort(notifyByte)
	if err != nil {
		return Notification{}, err
	}
	err = common.SHA256SignVerify([]byte(wsign), trade.request.pubKey, sign)
	if err != nil {
		return Notification{}, err
	}
	var notifyMsg Notification
	if err = json.Unmarshal(notifyByte, &notifyMsg); err != nil {
		return Notification{}, err
	}
	return notifyMsg, nil
}

func (trade *AlipayTrade) tradeReq(param interface{}, method string) (*commonResponse, error) {
	resp, err := trade.request.createRequest(param, method)
	if err != nil {
		return nil, err
	}
	bodyByte, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return &commonResponse{
		Response: resp,
		RespByte: bodyByte,
	}, nil
}
