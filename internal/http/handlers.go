package http

import (
	"sessions/internal/service"

	"github.com/valyala/fasthttp"
)

type Handler struct {
	sessions *service.SessionService
}

func New(s *service.SessionService) *Handler {
	return &Handler{sessions: s}
}

func (h *Handler) StartSession(ctx *fasthttp.RequestCtx) {
	id := string(ctx.FormValue("id"))
	user := string(ctx.FormValue("user"))

	if id == "" || user == "" {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	_ = h.sessions.Start(ctx, id, user)
	ctx.SetStatusCode(fasthttp.StatusCreated)
}
