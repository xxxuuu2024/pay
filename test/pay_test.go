package test

import (
	"fmt"
	"pay/alipay"
	"testing"
	"time"
)

var trade *alipay.AlipayTrade

func init() {

	trade = alipay.New(alipay.AlipayConf{
		AppID:            "2016093000631941",
		PriveKeyPath:     "/Users/x/Downloads/app_private_key.pem", //应用私钥
		AlipayPubKeyPath: "/Users/x/Downloads/app_public_key.csr",  //支付宝公钥
		SignType:         "RSA2",
	})
}

func TestCreatePay(t *testing.T) {
	outOrder := fmt.Sprint(time.Now().UnixNano())
	output, err := trade.CreatePreTradeRequest(
		alipay.CreateTradeInput{
			OutTradeNo:  outOrder,
			TotalAmount: "26.09",
			Subject:     "测试demo",
		})
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(output.RespByte))

}

//1609320593954340000 撤销交易
func TestCancelTrade(t *testing.T) {

	output, err := trade.CancelTradeReq(
		alipay.CancelTradeInput{
			OutTradeNo: "1609320708997073000",
		})
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(output)
}

//签名验证
func TestSignVerfy(t *testing.T) {

}

//
func TestDemo(t *testing.T) {
	var totalPrice uint32 = 500
	const couponPrice = 550

	fmt.Println("用户需要支付金额: ", totalPrice)
}
