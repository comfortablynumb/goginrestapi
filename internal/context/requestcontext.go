package context

import (
	"time"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
)

// Structs

type RequestContext struct {
	ginContext *gin.Context
	translator *ut.UniversalTranslator
}

func (r *RequestContext) GetAcceptLanguage() string {
	return r.ginContext.GetHeader("Accept-Language")
}

func (r *RequestContext) GetTranslator() *ut.Translator {
	trans, found := r.translator.GetTranslator(r.GetAcceptLanguage())

	if !found {
		trans = r.translator.GetFallback()
	}

	return &trans
}

func (r *RequestContext) Deadline() (deadline time.Time, ok bool) {
	return r.ginContext.Deadline()
}

func (r *RequestContext) Done() <-chan struct{} {
	return r.ginContext.Done()
}

func (r *RequestContext) Err() error {
	return r.ginContext.Err()
}

func (r *RequestContext) Value(key interface{}) interface{} {
	return r.Value(key)
}
