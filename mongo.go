package main

import (
    "gopkg.in/mgo.v2/bson"
    "github.com/drone/config"
    "gopkg.in/mgo.v2"
)

var (
    mongoUrl = config.String("mongo-url","mongodb://localhost/avatars")
    GridFsPrefix = config.String("mongo-gridfs-prefix","avatars")
)

func GetImageById (id bson.ObjectId) (*mgo.GridFile,error) {
    var (
        file *mgo.GridFile
        err error
    )
    sess,err := connect()
    if err != nil {
        // TODO: Почему на panic(err.String()) ?
        return file,err
    }
    defer sess.Close()
    file,err = sess.DB("").GridFS(*GridFsPrefix).OpenId(id)
    if err != nil {
        return file,err
    }
    return file,nil
}


func SaveFile (name string, data []byte) (interface {}, error) {
    var (
        file *mgo.GridFile
        err error
    )
    sess,err := connect()
    if err != nil {
        // TODO: Почему на panic(err.String()) ?
        return file,err
    }
    defer sess.Close()
    file,err = sess.DB("").GridFS(*GridFsPrefix).Create(name)
    if err != nil {
        return file,err
    }
    _, err = file.Write(data)

    return file.Id(),nil
}


// коннект к монге
func connect() (*mgo.Session,error) {
    return mgo.Dial(*mongoUrl)
}
