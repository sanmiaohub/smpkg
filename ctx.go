package smpkg

import (
	"context"
	"github.com/satori/go.uuid"
)

const XRequestId = "x-request-id"

type CtxContext struct {
	context.Context
	requestId string
}

func (c *CtxContext) RequestID() string {
	return c.requestId
}

func NewCtx() context.Context {
	return CtxContext{
		requestId: uuid.NewV4().String(),
	}
}

func WrapCtx(ctx context.Context) CtxContext {
	xid := ctx.Value(XRequestId)
	if xid == nil {
		xid = uuid.NewV4().String()
	}

	return CtxContext{
		requestId: xid.(string),
	}
}
