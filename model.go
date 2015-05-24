package main

// TODO: произвольный тип файла у которого будут
//  все необходимые методы для работы с ним.

// Кажется, структур становится достаточно много.
// Хочется разнести все по отдельным файлам.
// Но я крайне не уверен, что это будет хорошо в контектсе Golang

import (
    "gopkg.in/mgo.v2/bson"
    "gopkg.in/mgo.v2"
    "log"
    "strconv"
    "image"
    "strings"
    "errors"
    "reflect"
)

// Определение маски
type Mask struct {
    image.Rectangle
}

func (mask *Mask) SetFromBytes(field_path string, buf []byte) (*Mask, error) {
    return mask.SetFromStr(field_path, string(buf))
}

func (mask *Mask) SetFromStr(field_path string, buf string) (*Mask, error) {
    value, err := strconv.Atoi(buf)
    if err != nil {
        return mask, err
    }
    return mask.Set(field_path, value)
}

// Возвращает image.Rectangle.{Min|Max}.{X|Y}
// по сигнатурам, записанным в виде пути {Min|Max}/{X|Y}
func (mask *Mask) Get(field_path string) (int, error) {
    return mask.accessByPath(field_path)
}

// Устанавливает image.Rectangle.{Min|Max}.{X|Y}
// по сигнатурам, записанным в виде пути {Min|Max}/{X|Y}
func (mask *Mask) Set(field_path string, value int) (*Mask, error) {
    _, err := mask.accessByPath(field_path, value)
    return mask, err
}

// Возвращает и устанавливает image.Rectangle.{Min|Max}.{X|Y}
// по сигнатурам, записанным в виде пути {Min|Max}/{X|Y}
func (mask *Mask) accessByPath(
    field_path string,
    value_opt ... int
) (int, error) {
    str_list := strings.Split(field_path, "/")
    if len(str_list) < 2 {
        err := errors.New("use {Min|Max}/{X|Y}")
        return -1, err
    }
    point_field := str_list[0]
    xy_field    := str_list[1]
    return  mask.accessByNames(point_field, xy_field, value_opt...), nil
}

// Возвращает и устанавливает методы
// image.Rectangle.{Min|Max}.{X|Y} по сигнатурам
func (m *Mask) accessByNames(
    point_field string,
    xy_field string,
    value_opt ... int
) int {
    mask := reflect.ValueOf(m)
    point_value := reflect.Indirect(mask).FieldByName(point_field)
    xy_value := reflect.Indirect(point_value).FieldByName(xy_field)
    if len(value_opt) > 0 {
        value := value_opt[0]
        xy_value.SetInt(int64(value))
    }
    return int(xy_value.Int())
}

type FormFile struct {
    id  interface {}
    contentType string
    name string
    data []byte
}

func (ff *FormFile) Name () string  {
    return ff.name
}

func (ff *FormFile) Data () []byte  {
    return ff.data
}

func (ff *FormFile) ContentType () string  {
    return ff.contentType
}

func (ff *FormFile) Id () interface {}  {
    return ff.id
}

func (ff *FormFile) IdStr () string  {
    oid, _ := ff.id.(bson.ObjectId)
    return oid.Hex()
}

func (ff *FormFile) SetId (id interface {}) (*FormFile, error)  {
    ff.id = id
    return ff, nil
}

func (ff *FormFile) SetName (name string) (*FormFile, error)  {
    ff.name = name
    return ff, nil
}

func (ff *FormFile) SetContentType (contentType string) (*FormFile, error)  {
    ff.contentType = string(contentType)
    return ff, nil
}

func (ff *FormFile) UpdateData (slurp []byte) (*FormFile, error)  {
    ff.data = append(ff.data, slurp...)
    return ff, nil
}


func (ff *FormFile) Save() (interface {}, error)  {
    var (
        err error
    )
    ff.id, err = PutImageToDb(ff.name, ff.contentType, ff.data)

    log.Print("ff.id = ", ff.id)
    return ff.id, err
}

type DbFile struct {
    mgo.GridFile
    body []byte
}

func (df *DbFile) Body () []byte   {
    return df.body
}

func (df *DbFile) SetBody (body []byte) (*DbFile, error)  {
    df.body = body
    return df, nil
}

func (df *DbFile) BuildBody () (*DbFile, error)  {
    df.body = make([]byte, int(df.Size()))
    _, err := df.Read(df.body)
    if err != nil {
        return df,err
    }
    return df, nil
}

type ImageFile struct {
    FormFile
    Mask
}
