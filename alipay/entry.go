package alipay

import (
	"net/http"
	"time"
)

const (
	//app 支付
	aliAppPay = "alipay.trade.app.pay"
	//h5支付
	aliH5Pay = "alipay.trade.wap.pay"
	//交易创建
	aliTradeCreate = "alipay.trade.create"
	//订单查询
	aliTradeQuery = "alipay.trade.query"
	//关闭订单
	aliTradeCancel = "alipay.trade.cancel"
	//交易退款接口
	aliTradeRefund = "alipay.trade.refund"
)

type commonResponse struct {
	http.Response
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
	Sign    string `json:"sign"`
}

//交易创建
type GoodsDetails struct {
	// required 商品的编号
	GoodsID string `json:"goods_id"`
	//required 商品名称
	GoodsName string `json:"goods_name"`
	//required 商品数量
	Quantity string `json:"quantity"`
	//required 价格
	Price string `json:"price"`
	// 商品类目
	GoodsCategory string `json:"goods_category"`
	//子类目
	CategoriesTree string `json:"categories_tree"`
	//商品描述信息
	Body string `json:"body"`
	//商品展示地址
	ShowUrl string `json:"show_url"`
}
type CreateTradeRequest struct {
	//required 商户订单号
	OutTradeNo string `json:"out_trade_no"`
	//卖家支付宝用户ID
	SellerID string `json:"seller_id"`
	//required 订单总金额，单位为元，精确到小数点后两位
	TotalAmount string `json:"total_amount"`
	//可打折金额
	DiscountableAmount string `json:"discountable_amount"`
	//required 商品标题/交易标题/订单标题/订单关键字
	Subject string `json:"subject"`
	//交易商品描述
	Body string `json:"body"`
	//销售产品码 OFFLINE_PAYMENT=当面付快捷版 其它支付宝当面付产品=FACE_TO_FACE_PAYMENT , 默认 FACE_TO_FACE_PAYMENT
	ProductCode string `json:"product_code"`
	//商户操作员编号
	OperatorID string `json:"operator_id"`
	//商户门店编号
	StoreID string `json:"store_id"`
	//商户机具终端编号
	TerminalID string `json:"terminal_id"`
	// 订单包含的商品列表信息
	GoodsDetails []GoodsDetails `json:"goods_details"`
}
type CreateTradeResponse struct {
	commonResponse
	//商户订单号
	OutTradeNo string `json:"out_trade_no"`
	//支付宝交易号
	TradeNo string `json:"trade_no"`
}

//订单查询
type TradeQueryRequest struct {
	OutTradeNo   string `json:"out_trade_no"`
	TradeNo      string `json:"trade_no"`
	OrgPid       string `json:"org_pid"`
	QueryOptions string `json:"query_options"`
}
type TradeQueryResponse struct {
	//支付宝交易号
	TradeNo string `json:"trade_no"`
	//商家订单号
	OutTradeNo string `json:"out_trade_no"`
	//买家支付宝账号
	BuyerLogonID string `json:"buyer_logon_id"`
	//交易状态：WAIT_BUYER_PAY（交易创建，等待买家付款）、TRADE_CLOSED（未付款交易超时关闭，或支付完成后全额退款）、TRADE_SUCCESS（交易支付成功）、TRADE_FINISHED（交易结束，不可退款）
	TradeStatus string `json:"trade_status"`
	//交易的订单金额，单位为元，两位小数
	TotalAmount string `json:"total_amount"`
	//标价币种
	TransCurrency string `json:"trans_currency,omitempty"`
	//订单结算币种，对应支付接口传入的settle_currency
	SettleCurrency string `json:"settle_currency,omitempty"`
	//结算币种订单金额
	SettleAmount string `json:"settle_amount"`
	//订单支付币种
	PayCurrency string `json:"pay_currency"`
	//标价币种兑换支付币种汇率
	SettleTransRate string `json:"settle_trans_rate,omitempty"`
	//标价币种兑换支付币种汇率
	TransPayRate string `json:"trans_pay_rate"`
	//买家实付金额，单位为元，两位小数。该金额代表该笔交易买家实际支付的金额，不包含商户折扣等金额
	BuyerPayAmount string `json:"buyer_pay_amount"`
	//积分支付的金额，单位为元，两位小数
	PointAmount string `json:"point_amount"`
	//交易中用户支付的可开具发票的金额，单位为元
	InvoiceAmount string `json:"invoice_amount"`
	//本次交易打款给卖家的时间
	SendPayDate time.Time `json:"send_pay_date"`
	//实收金额，单位为元
	ReceiptAmount string `json:"receipt_amount"`
	//商户门店编号
	StoreID string `json:"store_id"`
	//商户机具终端编号
	TerminalID   string `json:"terminal_id"`
	FundBillList struct {
		//支付渠道 https://alipay.open.taobao.com/doc2/detail?treeId=26&articleId=103259&docType=1
		FundChannel string `json:"fund_channel"`
		//该支付工具类型所使用的金额
		Amount string `json:"amount"`
		//渠道实际付款金额
		RealAmount string `json:"real_amount"`
	} `json:"fund_bill_list"`
	//请求交易支付中的商户店铺的名称
	StoreName string `json:"store_name"`
	//该笔交易针对收款方的收费金额；
	ChargeAmount string `json:"charge_amount,omitempty"`
	//买家在支付宝的用户id
	BuyerUserID string `json:"buyer_user_id"`
	//费率活动标识，当交易享受活动优惠费率时，返回该活动的标识
	ChargeFlags string `json:"charge_flags,omitempty"`
	//支付清算编号，用于清算对账使用
	SettlementID    string `json:"settlement_id"`
	TradeSettleInfo struct {
		//交易结算明细信息
		TradeSettleDetailList struct {
			//结算操作类型 replenish、replenish_refund、transfer、transfer_refund等类型
			OperationType string `json:"operation_type"`
			//商户操作序列号。商户发起请求的外部请求号
			OperationSerialNo string `json:"operation_serial_no"`
			//操作日期
			OperationDt time.Time `json:"operation_dt"`
			//转出账号
			TransOut string `json:"trans_out"`
			//实际操作金额
			Amount string `json:"amount"`
			//转入账号
			TransIn string `json:"trans_in"`
		} `json:"trade_settle_detail_list"`
	} `json:"trade_settle_info"`
	//预授权支付模式，该参数仅在信用预授权支付场景下返回 CREDIT_PREAUTH_PAY
	AuthTradePayMode string `json:"auth_trade_pay_mode"`
	//买家用户类型。CORPORATE:企业用户；PRIVATE:个人用户
	BuyerUserType string `json:"buyer_user_type"`
	//商家优惠金额
	MdiscountAmount string `json:"mdiscount_amount"`
	//平台优惠金额
	DiscountAmount string `json:"discount_amount"`
	//订单标题
	Subject string `json:"subject"`
	//订单描述
	Body string `json:"body"`
	//间连商户在支付宝端的商户编号
	AlipaySubMerchantID string `json:"alipay_sub_merchant_id"`
	//交易额外信息
	ExtInfos string `json:"ext_infos"`
}

//发起支付
type GetPayRequest struct {
}

//发起退款
type CreateRefundRequest struct {
}

//关闭订单
type CancelTradeRequest struct {
}

//type AliPay interface {
//	//创建支付订单
//	CreatePay()
//	//创建退款请求
//	CreateRefund()
//	//关闭订单
//	QueryTrade()
//	//订单查询
//	CancelTrade()
//}
