package http

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func Router(h *Handler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		// Добавьте эту строку, чтобы видеть запросы в 'kubectl logs'
		fmt.Printf("Request: %s %s\n", ctx.Method(), ctx.Path())

		switch string(ctx.Path()) {
		case "/session/start":
			h.StartSession(ctx)
		default:
			ctx.SetStatusCode(fasthttp.StatusNotFound)
		}
	}
}
