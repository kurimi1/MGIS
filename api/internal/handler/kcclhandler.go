package handler

import (
	"net/http"

	"LGIS/api/internal/logic"
	"LGIS/api/internal/svc"
	"LGIS/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func KcclHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewKcclLogic(r.Context(), svcCtx)
		resp, err := l.Kccl(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
