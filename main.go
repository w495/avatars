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
    GetFileRouter := web.New()
    GetFileRouter.Use(SetFile)

    GetFileRouter.Get(BaseApiUrl+"/file/:id",       GetResizeFile)
    GetFileRouter.Get(BaseApiUrl+"/file/:id/raw",   GetOriginalFile)

    mux.Handle(BaseApiUrl+"/file/:id",              GetFileRouter)
    mux.Handle(BaseApiUrl+"/file/:id/*",            GetFileRouter)
    mux.Post(BaseApiUrl+"/file",                    UploadFile)
    mux.Put(BaseApiUrl+"/file",                     ChangeMask)
    http.Handle(BaseApiUrl+"/",                     mux)

    http.Handle("/", http.FileServer(http.Dir("app")))
    
    panic(http.ListenAndServe(*Listen, nil))
}
