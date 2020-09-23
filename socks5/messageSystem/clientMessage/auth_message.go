package clientMessage

import "errors"

type AuthMessage struct{
	Version byte
	MethodCount int
	Methods []int
}
var(
	UnknownAuthMethodError=errors.New("未知认证方式")
	IllegalAuthBytes=errors.New("认证字节数组不合法")
)
func ParseToAuthMessage(bs []byte,length int)(*AuthMessage,error){
	if length<2||(int(bs[1])!=length-2){
		return nil,IllegalAuthBytes
	}
	var msg AuthMessage
	msg.Version=bs[0]
	msg.MethodCount=int(bs[1])
	for i:=0;i<msg.MethodCount&&i+2<length;i++{
		switch bs[i+2]{
		case 0x00:
			{
			msg.Methods[i]=NoAuthMethod
			}
		case 0x01:
			{
			msg.Methods[i]=GSSAPIAuthMethod
			}
		case 0x02:
			{
			msg.Methods[i]=PasswordAuthMethod
			}
		case 0x03:
			{
			msg.Methods[i]=IANAAuthMethod
			}
		case 0x80:
			{
			msg.Methods[i]=PrivateAuthMethod
			}
		case 0xff:
			{
			msg.Methods[i]=NotSupportAuth
			}
		default:
			return nil,UnknownAuthMethodError
		}
	}
	return &msg,nil
}