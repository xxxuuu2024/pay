package alipay

import (
	"io/ioutil"
	"net/http"
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
}

func New(conf AlipayConf) *AlipayTrade {

	switch conf.SignType {
	case "RSA2", "RSA":
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
