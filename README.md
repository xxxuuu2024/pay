# pay
###
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
		NotifyUrl:        "http://8.210.250.185/hello",
		CallBackUrl:      "http://8.210.250.185/hello",
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
	t.Log(output)

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

func TestPayUrl(t *testing.T) {
	outOrder := fmt.Sprint(time.Now().UnixNano())
	output, err := trade.TradePagePayReq(
		alipay.TradePagePayInput{
			ProductCode: "FAST_INSTANT_TRADE_PAY",
			OutTradeNo:  outOrder,
			TotalAmount: "26.09",
			Subject:     "测试demo",
		})
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(output)
}

func TestNotify(t *testing.T) {
	x := "gmt_create=2020-12-31+17%3A30%3A23&charset=utf-8&gmt_payment=2020-12-31+17%3A30%3A37&notify_time=2020-12-31+17%3A30%3A38&subject=%E6%B5%8B%E8%AF%95demo&sign=yM2Y2GBDona3%2BX8NL2QZQM6x71OdjvEVIA%2B23rIeDY8ggJ94WbsawMI78asMglqGzuO9cP7EuKvDGFYStp1bMq0JCZiB7SGEXJ6M6KV%2FAVLbiN3BqkJPEJ%2F0ju2rXNxfHqnwIWiudjTWrR670q%2B0CknZfxH53qliWyKdC6FUIbrl2%2BQbzQrjZtEpHjsPCeSf04UOyk9WoRJbbjr5tYg5lG%2FjKWQwaUkJ7Sqe1hJGi9axo0oHxjfPirwE0C6XkXn2El2%2BvP6GVodGw3SvLD12HnvF%2BDCUHwOvhjL2vxu%2BXEvq8CYvpffHJxLOishUoI2yBLv1729xS6kFrwUVTry5Iw%3D%3D&buyer_id=2088102169360587&invoice_amount=26.09&version=1.0&notify_id=2020123100222173038060580511975558&fund_bill_list=%5B%7B%22amount%22%3A%2226.09%22%2C%22fundChannel%22%3A%22ALIPAYACCOUNT%22%7D%5D&notify_type=trade_status_sync&out_trade_no=1609406994128446000&total_amount=26.09&trade_status=TRADE_SUCCESS&trade_no=2020123122001460580501348783&auth_app_id=2016093000631941&receipt_amount=26.09&point_amount=0.00&app_id=2016093000631941&buyer_pay_amount=26.09&sign_type=RSA2&seller_id=2088102177960249"

	t.Log(trade.TradeNotify([]byte(x)))

}

func TestRefund(t *testing.T) {

	output, err := trade.TradeRefund(alipay.RefundInput{
		OutTradeNo:   "1609406994128446000",
		RefundAmount: "26.09",
	})
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(output)

}
