package clientMessage
const(
	NoAuthMethod=iota
	GSSAPIAuthMethod=iota
	PasswordAuthMethod=iota
	IANAAuthMethod=iota
	PrivateAuthMethod=iota
	NotSupportAuth=iota
)