package proxy

import "net/http"

type RedirectAble interface {
    GetRouteUrl() string
    ProcessResponse(w http.ResponseWriter, body []byte) (jsonResult []byte)
}
