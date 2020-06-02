package wxctx

import (
	"errors"
)

type Context struct {
	OpenID      string
	IsInCmdMode bool
	LastJoke    int
	LastOutput  string
}

var contextid = make(map[string]*Context)

func GetContextByOpenID(openID string) *Context {
	if ctx, ok := contextid[openID]; ok {
		return ctx
	} else {
		ctx := &Context{OpenID: openID}
		contextid[openID] = ctx
		return ctx
	}
}

func PutContext(context *Context) error {
	if context.OpenID == "" {
		return errors.New("openid can't be empty")
	}
	contextid[context.OpenID] = context
	return nil
}
