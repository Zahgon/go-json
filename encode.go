package json

import (
	"context"
	"io"

	"github.com/goccy/go-json/internal/encoder"
)

type Encoder struct {
	w                 io.Writer
	enabledIndent     bool
	enabledHTMLEscape bool
	prefix            string
	indentStr         string
}

func NewEncoder(w io.Writer) *Encoder { _ = "STUB: not implemented"; return nil }

func (e *Encoder) Encode(v interface{}) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) EncodeWithOption(v interface{}, optFuncs ...EncodeOptionFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) EncodeContext(ctx context.Context, v interface{}, optFuncs ...EncodeOptionFunc) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint: contextcheck

func (e *Encoder) encodeWithOption(ctx *encoder.RuntimeContext, v interface{}, optFuncs ...EncodeOptionFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) SetEscapeHTML(on bool) { _ = "STUB: not implemented"; return }

func (e *Encoder) SetIndent(prefix, indent string) { _ = "STUB: not implemented"; return }

func marshalContext(ctx context.Context, v interface{}, optFuncs ...EncodeOptionFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint: contextcheck

func marshal(v interface{}, optFuncs ...EncodeOptionFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalNoEscape(v interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func marshalIndent(v interface{}, prefix, indent string, optFuncs ...EncodeOptionFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encode(ctx *encoder.RuntimeContext, v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeNoEscape(ctx *encoder.RuntimeContext, v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeIndent(ctx *encoder.RuntimeContext, v interface{}, prefix, indent string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeRunCode(ctx *encoder.RuntimeContext, b []byte, codeSet *encoder.OpcodeSet) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeRunIndentCode(ctx *encoder.RuntimeContext, b []byte, codeSet *encoder.OpcodeSet, prefix, indent string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
