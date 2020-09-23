package socks5

import "errors"

var(
	UnknownAuthTypeError=errors.New("未知认证类型")
)
func New(listenAddr string,authType int)(*Server,error){
	switch authType{
	case NoAuthTypeServer:
		{
		var s=&Server{}
		s.listenAddr=listenAddr
		s.authType=authType
		return s,nil
		}
	default:
		return nil,UnknownAuthTypeError
	}
}