package middleware

import (
	"time"

	"github.com/ckeyer/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func Logger() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		ret, err := handler(ctx, req)
		consume := time.Now().Sub(start)

		fds := logrus.Fields{
			"method":  info.FullMethod,
			"consume": consume.String(),
		}
		if err != nil {
			fds["error"] = err
			logrus.WithFields(fds).Error("request over.")
		} else {
			logrus.WithFields(fds).Info("request over.")
		}

		return ret, err
	}
}
