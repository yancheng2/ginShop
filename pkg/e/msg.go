package e

/*
   编码消息
*/
var MsgFlags = map[int]string{
	200:   "success",
	1001:  "商品不存在",
	9001:  "响应错误信息",
	9002:  "系统错误",
	9003:  "请先登录",
	9004:  "登录信息错误",
	9005:  "登录信息过期",
	9006:  "无数据",
	10000: "请求参数错误",
}

/*
   根据传入的编码。获取对应的编码消息
*/
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[9001]
}
