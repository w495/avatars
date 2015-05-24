package main

// TODO: произвольный тип файла у которого будут
//  все необходимые методы для работы с ним.

import (
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
func (mask *Mask) accessByPath(field_path string, value_opt ... int) (int, error) {
    str_list := strings.Split(field_path, "/")
    if len(str_list) < 2 {
        err := errors.New("use {Min|Max}/{X|Y}")
        return -1, err
    }
    point_field := str_list[0]
    xy_field    := str_list[1]
    return  mask.accessByNames(point_field, xy_field, value_opt...), nil
}

// Возвращает и устанавливает  методы image.Rectangle.{Min|Max}.{X|Y} по сигнатурам
func (m *Mask) accessByNames(point_field string, xy_field string, value_opt ... int) int {
    mask := reflect.ValueOf(m)
    point_value := reflect.Indirect(mask).FieldByName(point_field)
    xy_value := reflect.Indirect(point_value).FieldByName(xy_field)
    if len(value_opt) > 0 {
        value := value_opt[0]
        xy_value.SetInt(int64(value))
    }
    return int(xy_value.Int())
}



/*
func (ff *FormFile) setMinX () (string)  {
    return ff.name
}*/


type FormFile struct {
    name string
    data []byte
}

func (ff *FormFile) Name () (string)  {
    return ff.name
}

func (ff *FormFile) Data () ([]byte)  {
    return ff.data
}

func (ff *FormFile) SetName (name string) (*FormFile, error)  {
    ff.name = name
    return ff, nil
}

func (ff *FormFile) UpdateData (slurp []byte) (*FormFile, error)  {
    ff.data = append(ff.data, slurp...)
    return ff, nil
}




type FormFileParts struct {
    FormFile
    Mask
}
