package tag

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	yaORMTag = "yaorm"
	tagSep   = ","
)

type (
	// ParsedTags is a struct that store all information needed
	// to get the yaORM tags from a struct
	ParsedTags struct {
		parsedSlice []parsedTag
		rawSlice    []string
		valueOf     reflect.Value
		typeOf      reflect.Type
		Error       error
	}

	parsedTag struct {
		tagType tag
		key     string
		value   string
		valueOf reflect.Value
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
	p.rawSlice = []string{}

	var field reflect.StructField

	for i := 0; i < p.valueOf.NumField(); i++ {
		field = p.typeOf.Field(i)
		p.rawSlice = append(p.rawSlice, field.Tag.Get(yaORMTag))
	}

	return p
}

func (p *ParsedTags) parseRawTags() *ParsedTags {
	if p.Error != nil {
		return p
	}

	p.parsedSlice = []parsedTag{}

	for _, tag := range p.rawSlice {
		parsed := parsedTag{
			tagType: strTagMap[tag],
		}
		tagSlice := strings.Split(tag, tagSep)
		if len(tagSlice) == 1 {
			parsed.key = tagSlice[0]
		} else {
			parsed.key = tagSlice[0]
			parsed.value = tagSlice[1]
		}

		p.parsedSlice = append(p.parsedSlice, parsed)
	}

	return p
}
