package json

import (
	"context"
	"io"
	"unsafe"

	"github.com/goccy/go-json/internal/decoder"
	"github.com/goccy/go-json/internal/runtime"
)

type Decoder struct {
	s *decoder.Stream
}

const (
	nul = '\000'
)

type emptyInterface struct {
	typ *runtime.Type
	ptr unsafe.Pointer
}

func unmarshal(data []byte, v interface{}, optFuncs ...DecodeOptionFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalContext(ctx context.Context, data []byte, v interface{}, optFuncs ...DecodeOptionFunc) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	pathDecoder = decoder.NewPathDecoder()
)

func extractFromPath(path *Path, data []byte, optFuncs ...DecodeOptionFunc) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unmarshalNoEscape(data []byte, v interface{}, optFuncs ...DecodeOptionFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEndBuf(src []byte, cursor int64) error { _ = "STUB: not implemented"; return nil }

//nolint:staticcheck
//go:nosplit
func noescape(p unsafe.Pointer) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

func validateType(typ *runtime.Type, p uintptr) error { _ = "STUB: not implemented"; return nil }

func NewDecoder(r io.Reader) *Decoder { _ = "STUB: not implemented"; return nil }

func (d *Decoder) Buffered() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func (d *Decoder) Decode(v interface{}) error { _ = "STUB: not implemented"; return nil }

func (d *Decoder) DecodeContext(ctx context.Context, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Decoder) DecodeWithOption(v interface{}, optFuncs ...DecodeOptionFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Decoder) More() bool { _ = "STUB: not implemented"; return false }

func (d *Decoder) Token() (Token, error) { _ = "STUB: not implemented"; return *new(Token), nil }

func (d *Decoder) DisallowUnknownFields() { _ = "STUB: not implemented"; return }

func (d *Decoder) InputOffset() int64 { _ = "STUB: not implemented"; return 0 }

func (d *Decoder) UseNumber() { _ = "STUB: not implemented"; return }
