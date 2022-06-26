package tag

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	yaORMTag    = "yaorm"
	tagSep      = ","
	keyValueSep = ":"
)

type (
	// ParsedTags is a struct that store all information needed
	// to get the yaORM tags from a struct
	ParsedTags struct {
		parsedTags map[string][]parsedTag
		rawSlice   []rawTag
		valueOf    reflect.Value
		typeOf     reflect.Type
		Error      error
	}

	parsedTag struct {
		tagType tag
		value   string
	}

	rawTag struct {
		rawStr string
		field  reflect.StructField
		value  reflect.Value
	}
)

// ParseTags receives a struct, and parse all yaORM tags.
func ParseTags(v any) (p *ParsedTags) {
	p = &ParsedTags{}
	typeOf := reflect.TypeOf(v)

	if typeOf.Kind() != reflect.Struct {
		err := fmt.Errorf(invalidTypeErr, reflect.Struct.String(), typeOf.Kind().String())
		p.Error = err
		return
	}

	p.typeOf = typeOf
	p.valueOf = reflect.ValueOf(v)

	return p.
		getRawTags().
		parseRawTags()
}

func (p *ParsedTags) getRawTags() *ParsedTags {
	if p.Error != nil {
		return p
	}
	p.rawSlice = []rawTag{}

	var field reflect.StructField

	for i := 0; i < p.valueOf.NumField(); i++ {
		field = p.typeOf.Field(i)
		p.rawSlice = append(p.rawSlice, rawTag{
			rawStr: field.Tag.Get(yaORMTag),
			field:  field,
			value:  p.valueOf.Field(i),
		})
	}
	return p
}

func (p *ParsedTags) parseRawTags() *ParsedTags {
	if p.Error != nil {
		return p
	}

	p.parsedTags = map[string][]parsedTag{}

	for _, wholeTag := range p.rawSlice {
		wholeTagSlice := strings.Split(wholeTag.rawStr, tagSep)
		parsedTagSlice := []parsedTag{}
		for _, tag := range wholeTagSlice {
			splittedTag := strings.Split(tag, ":")
			pTag := parsedTag{}
			if len(splittedTag) == 1 {
				pTag.tagType = strTagMap[splittedTag[0]]
			} else {
				pTag.tagType = strTagMap[splittedTag[0]]
				pTag.value = splittedTag[1]
			}
			parsedTagSlice = append(parsedTagSlice, pTag)
		}
		p.parsedTags[wholeTag.field.Name] = parsedTagSlice
	}

	return p
}
