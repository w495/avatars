package main
import (
    "strings"
    "net/http"
    "github.com/zenazn/goji/web"
    "encoding/json"
    "log"
    "io/ioutil"
    "gopkg.in/mgo.v2"
    "io"
    "fmt"
    "mime"
    "mime/multipart"
)


func UploadFile (c web.C, w http.ResponseWriter, request *http.Request) {

    // Требуется обработка ошибок
    mediaType, paramList, err := mime.ParseMediaType(request.Header.Get("Content-Type"))
    if err != nil {
        log.Fatal(err)
    }
    if strings.HasPrefix(mediaType, "multipart/") {
        reader := multipart.NewReader(request.Body, paramList["boundary"])
        for {
            part, err := reader.NextPart()
            if err == io.EOF {
                return
            }
            if err != nil {
                log.Fatal(err)
            }
            slurp, err := ioutil.ReadAll(part)
            if err != nil {
                log.Fatal(err)
            }
            fmt.Printf("Part %q: %q\n", part.FormName(), slurp)
        }
    }


//     body, _ := ioutil.ReadAll(r.Body)
//     log.Print("file name: ",body    )


    return
}

func ChangeMask (c web.C, w http.ResponseWriter, r *http.Request) {
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
