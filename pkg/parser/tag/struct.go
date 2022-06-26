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
		ParsedTags map[string][]ParsedTag
		rawSlice   []rawTag
		ValueOf    reflect.Value
		TypeOf     reflect.Type
		Error      error
	}

	// ParsedTag is one of the yaORM tags already parsed.
	ParsedTag struct {
		TagType tag
		Value   string
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

	p.TypeOf = typeOf
	p.ValueOf = reflect.ValueOf(v)

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

	for i := 0; i < p.ValueOf.NumField(); i++ {
		field = p.TypeOf.Field(i)
		p.rawSlice = append(p.rawSlice, rawTag{
			rawStr: field.Tag.Get(yaORMTag),
			field:  field,
			value:  p.ValueOf.Field(i),
		})
	}
	return p
}

func (p *ParsedTags) parseRawTags() *ParsedTags {
	if p.Error != nil {
		return p
	}

	p.ParsedTags = map[string][]ParsedTag{}

	for _, wholeTag := range p.rawSlice {
		wholeTagSlice := strings.Split(wholeTag.rawStr, tagSep)
		parsedTagSlice := []ParsedTag{}
		for _, tag := range wholeTagSlice {
			splittedTag := strings.Split(tag, ":")
			pTag := ParsedTag{}
			if len(splittedTag) == 1 {
				pTag.TagType = strTagMap[splittedTag[0]]
			} else {
				pTag.TagType = strTagMap[splittedTag[0]]
				pTag.Value = splittedTag[1]
			}
			parsedTagSlice = append(parsedTagSlice, pTag)
		}
		p.ParsedTags[wholeTag.field.Name] = parsedTagSlice
	}

	return p
}
