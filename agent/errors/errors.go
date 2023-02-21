package errors

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorDesc fetch grpc error message
func ErrorDesc(err error) string {
	if s, ok := status.FromError(err); ok {
		return s.Message()
	}
	return err.Error()
}

// ToGrpcError convert error to grpc error status err
func ToGrpcError(err error) error {
	if err == nil {
		return nil
	}
	if grpcErr, ok := stringToGrpcErrorMap[err.Error()]; ok {
		return grpcErr
	}
	return status.New(codes.Unknown, err.Error()).Err()
}

var (
	ErrorAccessKeyNotExist                   = errors.New(ErrorDesc(ErrorGrpcAccessKeyNotExist))
	ErrorAccessKeyNotExistOrPasswordNotMatch = errors.New(ErrorDesc(ErrorGrpcAccessKeyNotExistOrPasswordNotMatch))
	ErrorLeaseExist                          = errors.New(ErrorDesc(ErrorGRPCLeaseExist))
	ErrorLeaseNotExist                       = errors.New(ErrorDesc(ErrorGRPCLeaseNotExist))
	ErrorLeaseExpired                        = errors.New(ErrorDesc(ErrorGRPCLeaseExpired))
	ErrorUserNotExist                        = errors.New(ErrorDesc(ErrorGrpcUserNotExist))
	ErrorUserNotExistOrPasswordNotMatch      = errors.New(ErrorDesc(ErrorGrpcUserNotExistOrPasswordNotMatch))
	ErrorTokenUnexpectedSigningMethod        = errors.New(ErrorDesc(ErrorGrpcTokenUnexpectedSigningMethod))
	ErrorTokenNotValid                       = errors.New(ErrorDesc(ErrorGrpcTokenNotValid))
	ErrorTokenUnexpectedTokenType            = errors.New(ErrorDesc(ErrorGrpcTokenUnexpectedTokenType))
	ErrorNotLogin                            = errors.New(ErrorDesc(ErrorGrpcTokenUnexpectedTokenType))
)

// https://skyao.gitbooks.io/learning-grpc/content/server/status/status_code_definition.html
var (
	ErrorGrpcAccessKeyNotExist                   = status.New(codes.NotFound, "venus-server:access key not exit").Err()
	ErrorGrpcAccessKeyNotExistOrPasswordNotMatch = status.New(codes.NotFound, "venus-server:access key not exit or password not match").Err()
	ErrorGRPCLeaseExist                          = status.New(codes.AlreadyExists, "venus-server:grant lease exist").Err()
	ErrorGRPCLeaseNotExist                       = status.New(codes.NotFound, "venus-server:lease not exist").Err()
	ErrorGRPCLeaseExpired                        = status.New(codes.NotFound, "venus-server:lease expired").Err()
	ErrorGrpcUserNotExist                        = status.New(codes.NotFound, "venus-server:user not exit").Err()
	ErrorGrpcUserNotExistOrPasswordNotMatch      = status.New(codes.NotFound, "venus-server:user not exit or password not match").Err()
	ErrorGrpcTokenUnexpectedSigningMethod        = status.New(codes.InvalidArgument, "venus-server:unexpected signing method").Err()
	ErrorGrpcTokenNotValid                       = status.New(codes.InvalidArgument, "venus-server:token not valid").Err()
	ErrorGrpcTokenUnexpectedTokenType            = status.New(codes.InvalidArgument, "venus-server:unexpected token type").Err()
	ErrorGrpcNotLogin                            = status.New(codes.Unauthenticated, "venus-server:unexpected token type").Err()
)

var stringToGrpcErrorMap = map[string]error{
	ErrorLeaseExist.Error():                              ErrorGRPCLeaseExist,
	ErrorLeaseNotExist.Error():                           ErrorGRPCLeaseNotExist,
	ErrorLeaseExpired.Error():                            ErrorGRPCLeaseExpired,
	ErrorUserNotExist.Error():                            ErrorGrpcUserNotExist,
	ErrorUserNotExistOrPasswordNotMatch.Error():          ErrorGrpcUserNotExistOrPasswordNotMatch,
	ErrorTokenUnexpectedSigningMethod.Error():            ErrorGrpcTokenUnexpectedSigningMethod,
	ErrorTokenNotValid.Error():                           ErrorGrpcTokenNotValid,
	ErrorTokenUnexpectedTokenType.Error():                ErrorGrpcTokenUnexpectedTokenType,
	ErrorNotLogin.Error():                                ErrorGrpcNotLogin,
	ErrorLeaseNotExist.Error():                           ErrorGRPCLeaseNotExist,
	ErrorLeaseExpired.Error():                            ErrorGRPCLeaseExpired,
	ErrorUserNotExist.Error():                            ErrorGrpcUserNotExist,
	ErrorUserNotExistOrPasswordNotMatch.Error():          ErrorGrpcUserNotExistOrPasswordNotMatch,
	ErrorAccessKeyNotExist.Error():                       ErrorGrpcAccessKeyNotExist,
	ErrorGrpcAccessKeyNotExistOrPasswordNotMatch.Error(): ErrorGrpcAccessKeyNotExistOrPasswordNotMatch,
	ErrorTokenUnexpectedSigningMethod.Error():            ErrorGrpcTokenUnexpectedSigningMethod,
	ErrorTokenNotValid.Error():                           ErrorGrpcTokenNotValid,
	ErrorTokenUnexpectedTokenType.Error():                ErrorGrpcTokenUnexpectedTokenType,
}
