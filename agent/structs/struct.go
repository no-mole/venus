package structs

type MessageType uint8

const (
	AddNamespaceRequestType MessageType = iota
	AddKVRequestType
	LeaseGrantRequestType
	LeaseRevokeRequestType
	ServiceRegisterRequestType
	ServiceUnRegisterRequestType
)
