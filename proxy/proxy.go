package proxy

import (
    "fmt"
    "strings"
    "net/http"
    "bytes"
    "log"
    "encoding/json"

    "io/ioutil"
    "app/configurations"
    "github.com/jmoiron/jsonq"

	   "app/models"
)

func GetRedirectedResponse(w http.ResponseWriter, r *http.Request, c RedirectAble) {
    body := redirectRequest(w, r, c)
    uj := c.ProcessResponse(w, body)

    fmt.Fprintf(w, "%s\n", uj)
}

func redirectRequest(w http.ResponseWriter, req *http.Request, c RedirectAble) (respBody []byte) {

    httpClient := &http.Client{}

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    req.Body = ioutil.NopCloser(bytes.NewReader(body))
    getConfDataFromBody(body, c)
    // create a new url from the raw RequestURI sent by the client
    conf := configurations.GetConfiguration()
    proxyReq, err := http.NewRequest(req.Method, conf.Proxy.BaseUrl + c.GetRouteUrl(), bytes.NewReader(body))

    proxyReq.Header = make(http.Header)
    for h, val := range req.Header {
        // Avoid the keep-alive connection
        if(!strings.Contains(h, "Connection")) {
            proxyReq.Header[h] = val
        }
    }

    resp, err := httpClient.Do(proxyReq)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadGateway)
        return
    }

    respBody,err = ioutil.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    defer resp.Body.Close()

    return
}

func getConfDataFromBody(body []byte, c RedirectAble) {

    data := map[string]interface{}{}
    dec := json.NewDecoder(strings.NewReader(string(body)))
    dec.Decode(&data)
    jq := jsonq.NewQuery(data)
    str, err := jq.String("Dummy")
    if err != nil {
        log.Println("Could not get name from path : ", err)
        str = "unknown"
    }
    models.Req.Dummy = str
}
