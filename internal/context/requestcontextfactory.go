package context

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
)

// Structs

type RequestContextFactory struct {
	translator *ut.UniversalTranslator
}

func (r *RequestContextFactory) NewRequestContext(ginContext *gin.Context) *RequestContext {
	return &RequestContext{
		ginContext: ginContext,
		translator: r.translator,
	}
}

// Static functions

func NewRequestContextFactory(translator *ut.UniversalTranslator) *RequestContextFactory {
	return &RequestContextFactory{
		translator: translator,
	}
}
