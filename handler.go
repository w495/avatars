package main
import (
    "strings"
    "net/http"
    "github.com/zenazn/goji/web"
    "encoding/json"
    "log"
    "io/ioutil"
    //"gopkg.in/mgo.v2"
    "io"
    //"fmt"
    "mime"
    "mime/multipart"
)


const (
    FILE_FORM_NAME = `file`
    MASK_MIN_X_FORM_NAME = `mask-min-x`
    MASK_MIN_Y_FORM_NAME = `mask-min-y`
    MASK_MAX_X_FORM_NAME = `mask-max-x`
    MASK_MAX_Y_FORM_NAME = `mask-max-y`
)


// Сохраняет файл из формы `multipart/form-data`.
// Кажется, что для единообразия протоколов хорошо бы использовать
// обычный POST, где в теле передается изображение,
// а маска передается или GET-параметрами того же запроса,
func UploadFile (c web.C, w http.ResponseWriter, request *http.Request) {

    // Требуется обработка ошибок
    mediaType, paramList, err := mime.ParseMediaType(request.Header.Get("Content-Type"))
    if err != nil {
        log.Fatal(err)
    }

    formFileParts := new(ImageFile)

    if strings.HasPrefix(mediaType, "multipart/") {
        reader := multipart.NewReader(request.Body, paramList["boundary"])
        for {
            part, err := reader.NextPart()
            if err == io.EOF {
                // При нормальном исходе функция завершится тут.
                formFileParts.Save()
                w.WriteHeader(http.StatusOK)
                json.NewEncoder(w).Encode(map[string]string{
                    "id"    : formFileParts.IdStr(),
                    "type"  : formFileParts.ContentType(),
                    "name"  : formFileParts.Name(),
                })

                return
            }
            if err != nil {
                log.Fatal(err)
            }
            formFileParts = ParseFormFile(formFileParts, part)
        }
    }
    return
}


func ParseFormFile (formFileParts *ImageFile, part *multipart.Part) *ImageFile {
    //
    // TODO нужен котроль типа и размера.
    //
    slurp, err := ioutil.ReadAll(part)
    if err != nil {
        log.Fatal(err)
    }

    switch part.FormName() {
        case FILE_FORM_NAME: {
            formFileParts.UpdateData(slurp)
            formFileParts.SetName(part.FileName())
            formFileParts.SetContentType(part.Header.Get("Content-Type"))
        }
        case MASK_MIN_X_FORM_NAME: {
            // TODO: завернуть эту красоту в цикл,
            // если захочется.
            formFileParts.SetFromBytes(`Min/X`, slurp)
        }
        case MASK_MIN_Y_FORM_NAME: {
            formFileParts.SetFromBytes(`Min/Y`, slurp)
        }
        case MASK_MAX_X_FORM_NAME: {
            formFileParts.SetFromBytes(`Max/X`, slurp)
        }
        case MASK_MAX_Y_FORM_NAME: {
            formFileParts.SetFromBytes(`Max/Y`, slurp)
        }
    }
    return formFileParts
}

func ChangeMask (c web.C, w http.ResponseWriter, r *http.Request) {
    return
}

func GetOriginalFile (c web.C, w http.ResponseWriter, r *http.Request) {
    //получаем файл из env
    file,_ := c.Env["file"].(*DbFile)
    log.Print("file name: ",file.Name())
    log.Print("file name: ",file.ContentType())
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", file.ContentType())
    w.Write(file.Body())
    return
}

func GetResizeFile (c web.C,w http.ResponseWriter,r *http.Request) {
    //todo: resize used - https://github.com/nfnt/resize
//    file,_ := c.Env["file"].(*mgo.GridFile)
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"msg":`not found avatar for id - `})
    return
}
