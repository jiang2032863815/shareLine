package clientMessage
type CommandMessage struct {
	Version byte
	Command byte
	Rsv byte
	AddressType byte
	DstAddr string
	DstPort int
}
func ParseToCommandMessage(bs []byte)CommandMessage{
}