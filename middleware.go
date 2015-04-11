package main

import (
    "net/http"
    "github.com/zenazn/goji/web"
    "gopkg.in/mgo.v2/bson"
    "encoding/json"
)

func SetFile(c *web.C, h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fileId := c.URLParams["id"]
        if !bson.IsObjectIdHex(fileId) {
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"msg":`"Id" must be bson.ObjectId of type`})
            return
        }
        file,err := GetImageById(bson.ObjectIdHex(fileId))
        if err != nil {
            w.WriteHeader(http.StatusNotFound)
            json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
            return
        }

        if c.Env == nil {
            c.Env = make(map[string]interface{})
        }
        c.Env["file"] = file
        h.ServeHTTP(w, r)
    })
}