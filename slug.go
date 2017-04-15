package core

import (
	"reflect"
	"strings"
)

const tag_value = "slug"

type slug struct {
	tag   string
	value string
	field reflect.Value
}

func NewSlug(v interface{}) *slug {
	mirrorValue := reflect.Indirect(reflect.ValueOf(v))
	mirrorType := mirrorValue.Type()

	for i := 0; i < mirrorType.NumField(); i++ {
		field := mirrorType.Field(i)
		if field.Type.Kind() != reflect.String {
			continue
		}

		val, ok := field.Tag.Lookup(TAG_KEY)
		if !ok || !strings.Contains(val, tag_value) {
			continue
		}

		s := &slug{}
		s.tag = field.Name
		s.field = mirrorValue.Field(i)
		s.value = s.field.String()
		return s
	}

	panic(`No slug field found! Use tag gobly:"slug" to define one!`)
}

func (s *slug) SetValue(value string) {
	s.value = strings.ToLower(value)
	s.field.SetString(s.value)
}

func (s *slug) Tag() string {
	return s.tag
}

func (s *slug) Value() (value string, ok bool) {
	return s.value, len(strings.TrimSpace(s.value)) > 0
}
