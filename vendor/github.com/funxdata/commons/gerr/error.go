package gerr

import (
	"fmt"

	"github.com/ckeyer/logrus"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrNotFound spb.Status

func (e *ErrNotFound) Error() string {
	p := (*spb.Status)(e)
	return fmt.Sprintf("rpc error: code = %s desc = %s", codes.Code(p.GetCode()), p.GetMessage())
}

func New(c codes.Code, format string, a ...interface{}) error {
	logrus.Errorf(format, a...)
	return status.Errorf(c, format, a...)
}

func InvalidArgument(format string, a ...interface{}) error {
	return New(codes.InvalidArgument, format, a...)
}

func Internal(format string, a ...interface{}) error {
	return New(codes.Internal, format, a...)
}

func NotFound(format string, a ...interface{}) error {
	return &ErrNotFound{int32(codes.NotFound), fmt.Sprintf(format, a...), nil}
}

func DataError(format string, a ...interface{}) error {
	return New(codes.DataLoss, "内部数据发错错误，请联系管理员！！！ "+format, a...)
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	if _, ok := err.(*ErrNotFound); ok {
		return true
	}
	if e, ok := status.FromError(err); ok {
		return e.Code() == codes.NotFound
	}
	return false
}
