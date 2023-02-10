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
	ErrorLeaseExist                     = errors.New(ErrorDesc(ErrGRPCLeaseExist))
	ErrorLeaseNotExist                  = errors.New(ErrorDesc(ErrGRPCLeaseNotExist))
	ErrorLeaseExpired                   = errors.New(ErrorDesc(ErrGRPCLeaseExpired))
	ErrorUserNotExist                   = errors.New(ErrorDesc(ErrGrpcUserNotExist))
	ErrorUserNotExistOrPasswordNotMatch = errors.New(ErrorDesc(ErrGrpcUserNotExistOrPasswordNotMatch))
)

// https://skyao.gitbooks.io/learning-grpc/content/server/status/status_code_definition.html
var (
	ErrGRPCLeaseExist                     = status.New(codes.AlreadyExists, "venus-server:grant lease exist").Err()
	ErrGRPCLeaseNotExist                  = status.New(codes.NotFound, "venus-server:lease not exist").Err()
	ErrGRPCLeaseExpired                   = status.New(codes.NotFound, "venus-server:lease expired").Err()
	ErrGrpcUserNotExist                   = status.New(codes.NotFound, "venus-server:user not exit").Err()
	ErrGrpcUserNotExistOrPasswordNotMatch = status.New(codes.NotFound, "venus-server:user not exit or password not match").Err()
)

var stringToGrpcErrorMap = map[string]error{
	ErrorLeaseExist.Error():                     ErrGRPCLeaseExist,
	ErrorLeaseNotExist.Error():                  ErrGRPCLeaseNotExist,
	ErrorLeaseExpired.Error():                   ErrGRPCLeaseExpired,
	ErrorUserNotExist.Error():                   ErrGrpcUserNotExist,
	ErrorUserNotExistOrPasswordNotMatch.Error(): ErrGrpcUserNotExistOrPasswordNotMatch,
}
