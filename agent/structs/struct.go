package structs

type MessageType uint8

const (
	NamespaceAddRequestType MessageType = iota
	NamespaceDelRequestType
	NamespaceAddUserRequestType
	NamespaceDelUserRequestType
	NamespaceAddAccessKeyRequestType
	NamespaceDelAccessKeyRequestType
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
	AccessKeyGenRequestType
	AccessKeyDelRequestType
	AccessKeyAddNamespaceRequestType
	AccessKeyDelNamespaceRequestType
)

const (
	KVsBucketNamePrefix      = "kvs_"
	ServicesBucketNamePrefix = "services_"

	NamespacesBucketName           = "namespaces"
	NamespacesUsersBucketName      = "namespace_users"
	NamespacesAccessKeysBucketName = "namespace_access_keys"
	LeasesBucketName               = "leases"
	LeasesServicesBucketName       = "leases_services"
	UsersBucketName                = "user"
	UserNamespacesBucketName       = "user_namespaces"
	AccessKeysBucketName           = "access_key"
	AccessKeyNamespacesBucketName  = "access_key_namespaces"
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
	case NamespaceAddAccessKeyRequestType:
		return "NamespaceAddAccessKeyRequestType"
	case NamespaceDelAccessKeyRequestType:
		return "NamespaceDelAccessKeyRequestType"
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
	case AccessKeyGenRequestType:
		return "AccessKeyGenRequestType"
	case AccessKeyDelRequestType:
		return "AccessKeyDelRequestType"
	case AccessKeyAddNamespaceRequestType:
		return "AccessKeyAddNamespaceRequestType"
	case AccessKeyDelNamespaceRequestType:
		return "AccessKeyDelNamespaceRequestType"
	default:
		return "unknown"
	}
}
