package main

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetMetadata(c *gin.Context) {

	c.JSON(200, Resp("Success", nil, Metadata))
}

var Metadata = make(map[string]*TableMeta)

type TableMeta = struct {
	Fields   map[string]FieldMeta
	Preloads []string
}

type FieldMeta struct {
	Type    string `json:"type"`
	Comment string `json:"comment"`
}

func GetMeta() {
	for _, struPtr := range Tables {
		tableVal := reflect.ValueOf(struPtr).Elem()
		tableTyp := tableVal.Type()
		tableName := strings.ToLower(tableTyp.Name())
		tableMeta := &TableMeta{
			Fields:   make(map[string]FieldMeta),
			Preloads: make([]string, 0),
		}
		Metadata[tableName] = tableMeta
		for i := 0; i < tableTyp.NumField(); i++ {
			outField := tableTyp.Field(i)
			outFieldVal := tableVal.Field(i)
			outFieldTyp := outFieldVal.Type()
			outJsonTag := outField.Tag.Get("json")
			if outField.Tag.Get("extend") == "true" {
				tableMeta.Preloads = append(tableMeta.Preloads, outField.Name)
				if outFieldTyp.Kind() == reflect.Pointer {
					outFieldVal = outFieldVal.Elem()
					outFieldTyp = outFieldTyp.Elem()
				}
				for j := 0; j < outFieldTyp.NumField(); j++ {
					inField := outFieldTyp.Field(j)
					Store(inField, outJsonTag+".", false, tableMeta.Fields)
				}
			} else if outField.Anonymous {
				for j := 0; j < outFieldTyp.NumField(); j++ {
					inField := outFieldTyp.Field(j)
					Store(inField, "", false, tableMeta.Fields)
				}
			} else {
				Store(outField, "", true, tableMeta.Fields)
			}
		}
	}
}

func Store(field reflect.StructField, prefix string, withExtends bool, m map[string]FieldMeta) {

	jsonTag := field.Tag.Get("json")
	if jsonTag == "" || jsonTag == "-" {
		return
	}
	key := prefix + jsonTag
	typ := field.Type
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Slice {
		return
	}
	if !withExtends && field.Tag.Get("extend") == "true" {
		return
	}
	m[key] = FieldMeta{
		Type:    typ.Kind().String(),
		Comment: ReadComment(field),
	}
}

func ReadComment(field reflect.StructField) string {
	if gormTag := field.Tag.Get("gorm"); gormTag != "" {
		for _, part := range strings.Split(gormTag, ";") {
			part = strings.TrimSpace(part)
			if strings.HasPrefix(part, "comment:") {
				comment := strings.TrimPrefix(part, "comment:")
				return strings.Trim(comment, "'\"")
			}
		}
	}
	return ""
}
