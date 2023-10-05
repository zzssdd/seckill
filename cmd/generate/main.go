package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"seckill/conf"
	"strings"
)

var MysqlUserDSN = conf.MysqlUserDSN

func main() {
	db, err := gorm.Open(mysql.Open(MysqlUserDSN))
	if err != nil {
		panic(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:           "../../dao/db/",
		Mode:              gen.WithDefaultQuery,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
	})
	g.UseDB(db)
	// 自定义字段的数据类型
	dataMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"varchar": func(columnType gorm.ColumnType) (dataType string) {
			return "string"
		},
		"bigint": func(columnType gorm.ColumnType) (dataType string) {
			return "int64"
		},
		"int": func(columnType gorm.ColumnType) (dataType string) {
			return "int32"
		},
		"double": func(columnType gorm.ColumnType) (dataType string) {
			return "float64"
		},
		"text": func(columnType gorm.ColumnType) (dataType string) {
			return "string"
		},
		"tinytext": func(columnType gorm.ColumnType) (dataType string) {
			return "string"
		},
	}
	g.WithDataTypeMap(dataMap)

	// 自定义模型结体字段的标签
	// 将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		toStringField := `balance, `
		if strings.Contains(toStringField, columnName) {
			return columnName + ",string"
		}
		return columnName
	})

	// 将非默认字段名的字段定义为自动时间戳和软删除字段;
	// 自动时间戳默认字段名为:`updated_at`、`created_at, 表字段数据类型为: INT 或 DATETIME
	// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME
	softDeleteField := gen.FieldType("deleted_at", "soft_delete.DeletedAt")
	fieldOpts := []gen.ModelOpt{jsonField, softDeleteField}
	// 创建全部模型文件, 并覆盖前面创建的同名模型
	g.GenerateAllTable(fieldOpts...)
	g.Execute()
}
