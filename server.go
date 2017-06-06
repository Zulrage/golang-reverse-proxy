package main

import (
    // Standard library packages
    "net/http"
    "log"
    "fmt"

    // Third party packages
    "github.com/julienschmidt/httprouter"
    "app/controllers"
    "app/configurations"

)

func main() {
  log.Println("Server up")
  configurations.InitConfiguration()
  conf := configurations.GetConfiguration()
  r := httprouter.New()

  r.GET(conf.Server.Root + "/_ping", ping)

  dummyController := controllers.NewDsummyController()
  r.POST(conf.Server.Root + "/dummy", dummyController.GetDummys)

  // Fire up the server
  log.Println(http.ListenAndServe(conf.Server.Url + ":" + conf.Server.Port, r))
}

func ping(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintf(w, "ping")
}
