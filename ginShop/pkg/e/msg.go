package e

/*
   编码消息
*/
var MsgFlags = map[int]string{
	200:"success",
	9001:"响应错误信息",
}

/*
   根据传入的编码。获取对应的编码消息
*/
func GetMsg(code int)string  {
	msg,ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[9001]
}