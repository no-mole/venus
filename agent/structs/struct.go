package structs

type MessageType uint8

const (
	NamespaceAddRequestType MessageType = iota
	NamespaceDelRequestType
	NamespaceAddUserRequestType
	NamespaceDelUserRequestType
	UserRegisterRequestType
	UserUnregisterRequestType
	UserAddNamespaceRequestType
	UserDelNamespaceRequestType
	KVAddRequestType
	KVDelRequestType
	LeaseGrantRequestType
	LeaseRevokeRequestType
	ServiceRegisterRequestType
	ServiceUnRegisterRequestType
)

const (
	KVsBucketNamePrefix      = "kvs_"
	ServicesBucketNamePrefix = "services_"

	NamespacesBucketName      = "namespaces"
	NamespacesUsersBucketName = "namespace_users"
	LeasesBucketName          = "leases"
	UsersBucketName           = "user"
	UserNamespacesBucketName  = "user_namespaces"
)
