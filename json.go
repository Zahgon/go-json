package json

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/goccy/go-json/internal/encoder"
)

type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

type MarshalerContext interface {
	MarshalJSON(context.Context) ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

type UnmarshalerContext interface {
	UnmarshalJSON(context.Context, []byte) error
}

func Marshal(v interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func MarshalNoEscape(v interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func MarshalContext(ctx context.Context, v interface{}, optFuncs ...EncodeOptionFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MarshalWithOption(v interface{}, optFuncs ...EncodeOptionFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MarshalIndentWithOption(v interface{}, prefix, indent string, optFuncs ...EncodeOptionFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Unmarshal(data []byte, v interface{}) error { _ = "STUB: not implemented"; return nil }

func UnmarshalContext(ctx context.Context, data []byte, v interface{}, optFuncs ...DecodeOptionFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func UnmarshalWithOption(data []byte, v interface{}, optFuncs ...DecodeOptionFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func UnmarshalNoEscape(data []byte, v interface{}, optFuncs ...DecodeOptionFunc) error {
	_ = "STUB: not implemented"
	return nil
}

type Token = json.Token

type Number = json.Number

type RawMessage = json.RawMessage

type Delim = json.Delim

func Compact(dst *bytes.Buffer, src []byte) error { _ = "STUB: not implemented"; return nil }

func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
	_ = "STUB: not implemented"
	return nil
}

func HTMLEscape(dst *bytes.Buffer, src []byte) { _ = "STUB: not implemented"; return }

func Valid(data []byte) bool { _ = "STUB: not implemented"; return false }

func init() {
	encoder.Marshal = Marshal
	encoder.Unmarshal = Unmarshal
}
