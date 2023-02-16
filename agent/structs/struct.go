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

func GenBucketName(prefix, namespace string) []byte {
	return []byte(prefix + namespace)
}

func (l MessageType) String() string {
	switch l {
	case NamespaceAddRequestType:
		return "NamespaceAddRequestType"
	case NamespaceDelRequestType:
		return "NamespaceDelRequestType"
	case NamespaceAddUserRequestType:
		return "NamespaceAddUserRequestType"
	case NamespaceDelUserRequestType:
		return "NamespaceDelUserRequestType"
	case UserRegisterRequestType:
		return "UserRegisterRequestType"
	case UserUnregisterRequestType:
		return "UserUnregisterRequestType"
	case UserAddNamespaceRequestType:
		return "UserAddNamespaceRequestType"
	case UserDelNamespaceRequestType:
		return "UserDelNamespaceRequestType"
	case KVAddRequestType:
		return "KVAddRequestType"
	case KVDelRequestType:
		return "KVDelRequestType"
	case LeaseGrantRequestType:
		return "LeaseGrantRequestType"
	case LeaseRevokeRequestType:
		return "LeaseRevokeRequestType"
	case ServiceRegisterRequestType:
		return "ServiceRegisterRequestType"
	case ServiceUnRegisterRequestType:
		return "ServiceUnRegisterRequestType"
	default:
		return "unknown"
	}
}
