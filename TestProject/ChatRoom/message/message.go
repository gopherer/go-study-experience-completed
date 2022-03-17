package message

const (
	LoginType    = "LoginMes"
	LoginResType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

//用户登入消息
type LoginMes struct {
	UserId   int    `json:"user_id"`   //用户编号
	UserPwd  string `json:"user_pwd"`  //用户密码
	UserName string `json:"user_name"` //用户名称
}

//服务器回复用户登入消息
type LoginResMes struct {
	Code  int    `json:"code"`  //404表示用户未注册 100表示用户登入成功
	Error string `json:"error"` //返回的错误信息
}
