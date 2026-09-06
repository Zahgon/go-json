package json

import (
	"io"

	"github.com/goccy/go-json/internal/decoder"
	"github.com/goccy/go-json/internal/encoder"
)

type EncodeOption = encoder.Option
type EncodeOptionFunc func(*EncodeOption)

func UnorderedMap() EncodeOptionFunc { _ = "STUB: not implemented"; return *new(EncodeOptionFunc) }

func DisableHTMLEscape() EncodeOptionFunc { _ = "STUB: not implemented"; return *new(EncodeOptionFunc) }

func DisableNormalizeUTF8() EncodeOptionFunc {
	_ = "STUB: not implemented"
	return *new(EncodeOptionFunc)
}

func Debug() EncodeOptionFunc { _ = "STUB: not implemented"; return *new(EncodeOptionFunc) }

func DebugWith(w io.Writer) EncodeOptionFunc {
	_ = "STUB: not implemented"
	return *new(EncodeOptionFunc)
}

func DebugDOT(w io.WriteCloser) EncodeOptionFunc {
	_ = "STUB: not implemented"
	return *new(EncodeOptionFunc)
}

func Colorize(scheme *ColorScheme) EncodeOptionFunc {
	_ = "STUB: not implemented"
	return *new(EncodeOptionFunc)
}

type DecodeOption = decoder.Option
type DecodeOptionFunc func(*DecodeOption)

func DecodeFieldPriorityFirstWin() DecodeOptionFunc {
	_ = "STUB: not implemented"
	return *new(DecodeOptionFunc)
}
