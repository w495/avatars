package main

import (
    "github.com/drone/config"
    "github.com/zenazn/goji/web"
    "net/http"
)

const BaseApiUrl = `/api/v1`

var (
    Listen = config.String("http","0.0.0.0:4567")
)

func main () {
    config.SetPrefix("AV_")
    config.Parse("")
    mux := web.New()
    getter := web.New()
    getter.Use(SetFile)
    getter.Get(BaseApiUrl+"/file/:id",GetResizeFile)
    getter.Get(BaseApiUrl+"/file/:id/raw",GetOriginalFile)
    mux.Handle(BaseApiUrl+"/file/:id", getter)
    mux.Handle(BaseApiUrl+"/file/:id/*", getter)
    mux.Post(BaseApiUrl+"/file",UploadFile)
    mux.Put(BaseApiUrl+"/file",ChangeMask)
    http.Handle("/", http.FileServer(http.Dir("app")))
    http.Handle(BaseApiUrl+"/", mux)
    panic(http.ListenAndServe(*Listen, nil))
}