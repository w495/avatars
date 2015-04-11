package main
import (
    "net/http"
    "github.com/zenazn/goji/web"
    "encoding/json"
    "log"
    "gopkg.in/mgo.v2"
)


func UploadFile (c web.C,w http.ResponseWriter,r *http.Request) {
    return
}

func ChangeMask (c web.C,w http.ResponseWriter,r *http.Request) {
    return
}

func GetOriginalFile (c web.C,w http.ResponseWriter,r *http.Request) {
    //получаем файл из env
    file,_ := c.Env["file"].(*mgo.GridFile)
    log.Print("file name: ",file.Name())
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"msg":`not found avatar`})
    return
}

func GetResizeFile (c web.C,w http.ResponseWriter,r *http.Request) {
    //todo: resize used - https://github.com/nfnt/resize
//    file,_ := c.Env["file"].(*mgo.GridFile)
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"msg":`not found avatar for id - `})
    return
}
