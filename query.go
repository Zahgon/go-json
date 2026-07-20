package json

import (
	"github.com/goccy/go-json/internal/encoder"
)

type (
	FieldQuery       = encoder.FieldQuery
	FieldQueryString = encoder.FieldQueryString
)

var (
	FieldQueryFromContext = encoder.FieldQueryFromContext

	SetFieldQueryToContext = encoder.SetFieldQueryToContext
)

func BuildFieldQuery(fields ...FieldQueryString) (*FieldQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func BuildSubFieldQuery(name string) *SubFieldQuery { _ = "STUB: not implemented"; return nil }

type SubFieldQuery struct {
	name string
}

func (q *SubFieldQuery) Fields(fields ...FieldQueryString) FieldQueryString {
	_ = "STUB: not implemented"
	return *new(FieldQueryString)
}
