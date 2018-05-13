package core

import (
	"github.com/funxdata/commons/rpc/cache"
	"github.com/funxdata/commons/rpc/constants"
	"github.com/funxdata/commons/rpc/metadata"
	"golang.org/x/net/context"
)

type InjectionService struct {
	session cache.Cache
}

func mustGet(ctx context.Context) interface{} {
	if value := ctx.Value(constants.NEVIS_CTX); value != nil {
		return value
	}
	panic("Key \"" + constants.NEVIS_CTX + "\" does not exist")
}

func (i InjectionService) Session() cache.Cache {
	if i.session == nil {
		i.session = cache.DefaultSession()
	}
	return i.session
}

func (i InjectionService) Option(key string, ctx context.Context) interface{} {
	return metadata.Option(key, ctx)
}
