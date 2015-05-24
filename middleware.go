package main

import (
    "log"
    "net/http"
    "github.com/zenazn/goji/web"
    "gopkg.in/mgo.v2/bson"
    "encoding/json"
)


// Зачем ???
// Кажется, что не очень хорошо,
// что в middleware вызывается функция GetImageById
// Как минимум, это надо делаеть или в моделе (или dao) или контроллере.
func SetFile(c *web.C, h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fileId := c.URLParams["id"]
        if !bson.IsObjectIdHex(fileId) {
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"msg":`"Id" must be bson.ObjectId of type`})
            return
        }
        log.Print("fileId = ", fileId)
        file, err := GetImageById(bson.ObjectIdHex(fileId))

        log.Print("fileId = ", bson.ObjectIdHex(fileId), err)
        if err != nil {

            log.Print("2 fileId = ", bson.ObjectIdHex(fileId), err)
            w.WriteHeader(http.StatusNotFound)

            log.Print("3 fileId = ", bson.ObjectIdHex(fileId), err)

            json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
            return
        }

        c.Env = make(map[interface {}]interface {})
        c.Env["file"] = file

        h.ServeHTTP(w, r)
    })
}
