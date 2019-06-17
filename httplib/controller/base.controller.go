package controller

import (
	"anagramma"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

//Controller
type Controller struct {
	HashMapAlgorithm *anagramma.HashMap
}

//FPrintGoodResponse
func (ctrl *Controller) FPrintGoodResponse(ctx *fasthttp.RequestCtx, a interface{}) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprint(ctx, a)
	return
}

//sendGoodResponse
func (ctrl *Controller) sendGoodResponse(ctx *fasthttp.RequestCtx, message interface{}) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")

	data, err := json.Marshal(message)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}
	fmt.Fprint(ctx, string(data))
}

//sendFailResponse
func (ctrl *Controller) sendFailResponse(ctx *fasthttp.RequestCtx, status int, message interface{}) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.SetStatusCode(status)
	ctx.SetContentType("application/json")

	data, err := json.Marshal(message)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}
	fmt.Fprint(ctx, string(data))
}
