package alipay

import "net/http"

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

type CommonResponse struct {
	http.Response
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

type TradeCreateRequest struct {
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

//订单查询

//发起支付

//发起退款

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
