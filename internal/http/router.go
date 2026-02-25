package http

import (
	"github.com/valyala/fasthttp"
)

func Router(h *Handler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/session/start":
			h.StartSession(ctx)
		default:
			ctx.SetStatusCode(fasthttp.StatusNotFound)
		}
	}
}
