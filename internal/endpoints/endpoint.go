package endpoints

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Endpoint interface {
	Handle(ctx *gin.Context)

	Path() string
	Method() string
}

func AsEndpoint(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Endpoint)),
		fx.ResultTags(`group:"endpoints"`),
	)
}
