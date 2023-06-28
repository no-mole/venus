package errors

import (
	"errors"
	"testing"
)

func TestToGrpcError(t *testing.T) {
	for errMsg, grpcError := range stringToGrpcErrorMap {
		if ToGrpcError(errors.New(errMsg)) != grpcError {
			t.Error("ToGrpcError(errors.New(errMsg)) != grpcError")
		}
	}
}
