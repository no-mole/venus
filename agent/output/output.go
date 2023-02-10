package output

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

/*
| HTTP Status Code           | gRPC Status Code   |
|----------------------------|--------------------|
| 400 Bad Request            | INTERNAL           |
| 401 Unauthorized           | UNAUTHENTICATED    |
| 403 Forbidden              | PERMISSION\_DENIED |
| 404 Not Found              | UNIMPLEMENTED      |
| 429 Too Many Requests      | UNAVAILABLE        |
| 502 Bad Gateway            | UNAVAILABLE        |
| 503 Service Unavailable    | UNAVAILABLE        |
| 504 Gateway Timeout        | UNAVAILABLE        |
| _All other codes_          | UNKNOWN            |
*/

var GrpcCodeToHttpCode = map[codes.Code]int{
	codes.OK:                 http.StatusOK,
	codes.Canceled:           http.StatusGone,
	codes.Unknown:            http.StatusBadRequest,
	codes.InvalidArgument:    http.StatusBadRequest,
	codes.DeadlineExceeded:   http.StatusGatewayTimeout,
	codes.NotFound:           http.StatusNotFound,
	codes.PermissionDenied:   http.StatusForbidden,
	codes.ResourceExhausted:  http.StatusInsufficientStorage,
	codes.FailedPrecondition: http.StatusBadRequest,
	codes.Aborted:            http.StatusBadRequest,
	codes.OutOfRange:         http.StatusBadRequest,
	codes.Unimplemented:      http.StatusNotImplemented,
	codes.Internal:           http.StatusInternalServerError,
	codes.Unavailable:        http.StatusServiceUnavailable,
	codes.DataLoss:           http.StatusInternalServerError,
	codes.Unauthenticated:    http.StatusUnauthorized,
}

func Json(ctx *gin.Context, err error, data interface{}) {
	if err != nil {
		grpcErr := status.Convert(err)
		ctx.JSON(
			GrpcCodeToHttpCode[grpcErr.Code()],
			&Result{
				Code: int(grpcErr.Code()),
				Msg:  grpcErr.Code().String(),
			},
		)
		return
	}
	ctx.JSON(http.StatusOK, &Result{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}
