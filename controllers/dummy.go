package controllers

import (
    "net/http"

    "app/models"
    "app/proxy"
    "github.com/julienschmidt/httprouter"
)

type (
    ProgramRankingController struct{}
)

func NewDummyController() *DummyController {
    return &DummyController{}
}

func (c *DummyController) GetDummy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    proxy.GetRedirectedResponse(w, r, c)
}

func (c *DummyController) GetRouteUrl() string {
  return "/dummy"
}

func (c *DummyController) ProcessResponse(w http.ResponseWriter, body []byte) (jsonResult []byte) {
    jsonResult = body

    return
}
