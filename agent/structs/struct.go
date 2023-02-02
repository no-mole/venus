package structs

type MessageType uint8

const (
	AddNamespaceRequestType MessageType = iota
	DelNamespaceRequestType
	AddKVRequestType
	DelKVRequestType
	LeaseGrantRequestType
	LeaseRevokeRequestType
	ServiceRegisterRequestType
	ServiceUnRegisterRequestType
)

const (
	KVsBucketNamePrefix      = "kvs_"
	ServicesBucketNamePrefix = "services_"

	NamespacesBucketName = "namespaces"
	LeasesBucketName     = "leases"
	UsersBucketName      = "users"
	AccessKeysBucketName = "access_keys"
)
