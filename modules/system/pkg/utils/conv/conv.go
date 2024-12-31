// Package conv
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package conv

import (
	"devinggo/modules/system/pkg/utils/slice"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
)

var (
	descTags  = []string{"excel", "description", "dc", "json"} // 实体描述标签
	fieldTags = []string{"json"}                               // 实体字段名称映射
)

func ConvIntMap(m map[string]interface{}) map[string]int {
	ret := make(map[string]int, len(m))
	for k, v := range m {
		ret[k] = gconv.Int(v)
	}
	return ret
}

// GetEntityFieldTags 获取实体中的字段名称
func GetEntityFieldTags(entity interface{}) (tags []string, err error) {
	var formRef = reflect.TypeOf(entity)
	for i := 0; i < formRef.NumField(); i++ {
		field := formRef.Field(i)
		if field.Type.Kind() == reflect.Struct {
			addTags, err := reflectTag(field.Type, fieldTags, []string{})
			if err != nil {
				return nil, err
			}
			tags = append(tags, addTags...)
			continue
		}
		tags = append(tags, reflectTagName(field, fieldTags, true))
	}
	return
}

// GetEntityDescTags 获取实体中的描述标签
func GetEntityDescTags(entity interface{}) (tags []string, err error) {
	var formRef = reflect.TypeOf(entity)
	for i := 0; i < formRef.NumField(); i++ {
		field := formRef.Field(i)
		if field.Type.Kind() == reflect.Struct {
			addTags, err := reflectTag(field.Type, descTags, []string{})
			if err != nil {
				return nil, err
			}
			tags = append(tags, addTags...)
			continue
		}
		tags = append(tags, reflectTagName(field, descTags, true))
	}
	return
}

// reflectTag 层级递增解析tag
func reflectTag(reflectType reflect.Type, filterTags []string, tags []string) ([]string, error) {
	if reflectType.Kind() == reflect.Ptr {
		return nil, gerror.Newf("reflect type do not support reflect.Ptr yet, reflectType:%+v", reflectType)
	}
	if reflectType.Kind() != reflect.Struct {
		return nil, nil
	}
	for i := 0; i < reflectType.NumField(); i++ {
		tag := reflectTagName(reflectType.Field(i), filterTags, false)
		if tag == "" {
			addTags, err := reflectTag(reflectType.Field(i).Type, filterTags, tags)
			if err != nil {
				return nil, err
			}
			tags = append(tags, addTags...)
			continue
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

// reflectTagName 解析实体中的描述标签，优先级：excel>description > dc > json > Name
func reflectTagName(field reflect.StructField, filterTags []string, isDef bool) string {
	if slice.Contains(filterTags, "excel") {
		if excel, ok := field.Tag.Lookup("excel"); ok && excel != "" {
			return excel
		}
	}

	if slice.Contains(filterTags, "description") {
		if description, ok := field.Tag.Lookup("description"); ok && description != "" {
			return description
		}
	}

	if slice.Contains(filterTags, "dc") {
		if dc, ok := field.Tag.Lookup("dc"); ok && dc != "" {
			return dc
		}
	}

	if slice.Contains(filterTags, "json") {
		if jsonName, ok := field.Tag.Lookup("json"); ok && jsonName != "" {
			return jsonName
		}
	}

	if !isDef {
		return ""
	}
	return field.Name
}
