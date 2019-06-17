package controller

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

//Load
func (ctrl *Controller) Load(ctx *fasthttp.RequestCtx) {
	var wslice []string
	data := ctx.PostBody()

	err := json.Unmarshal(data, &wslice)
	if err != nil {
		ctrl.sendFailResponse(ctx, 400, map[string]string{
			"error": "body format invalid, must be json array",
		})
		return
	}

	tsize := ctrl.HashMapAlgorithm.Store(wslice...)
	ctrl.sendGoodResponse(ctx, map[string]int{
		"match_anagramm": tsize,
	})
}

//Get
func (ctrl *Controller) Get(ctx *fasthttp.RequestCtx) {
	word := ctx.QueryArgs().Peek("word")
	matchSlice := ctrl.HashMapAlgorithm.Load(string(word))
	ctrl.sendGoodResponse(ctx, matchSlice)
}
