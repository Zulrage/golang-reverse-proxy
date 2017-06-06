package configurations

import (
    "encoding/json"
    "io/ioutil"
    "path/filepath"
    "log"
    "path"
    "runtime"
)

type CServer struct {
    Url    string
    Port   string
    Root   string
}

type CProxy struct {
    BaseUrl    string
}

type Configuration struct {
    Server    CServer
    Proxy     CProxy
}

var config Configuration

func InitConfiguration() {
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        log.Fatal("Could not find path to configuration ")
    }

    absPath, _ := filepath.Abs(path.Dir(path.Dir(filename)) + "/resources/config.json")
    b, err := ioutil.ReadFile(absPath)
    if err != nil {
        log.Fatal("Configuration IO error ", err)
    }
    err = json.Unmarshal(b, &config)
    if err != nil {
        log.Fatal("Configuration fetching error ", err)
    }
}

func GetConfiguration() Configuration {
    return config
}
