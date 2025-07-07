package otel

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func (p *Provider) GinMiddleware() gin.HandlerFunc {
	return otelgin.Middleware(p.cfg.ServerName)
}
