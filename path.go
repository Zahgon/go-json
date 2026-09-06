package json

import (
	"github.com/goccy/go-json/internal/decoder"
)

func CreatePath(p string) (*Path, error) { _ = "STUB: not implemented"; return nil, nil }

type Path struct {
	path *decoder.Path
}

func (p *Path) RootSelectorOnly() bool { _ = "STUB: not implemented"; return false }

func (p *Path) UsedSingleQuotePathSelector() bool { _ = "STUB: not implemented"; return false }

func (p *Path) UsedDoubleQuotePathSelector() bool { _ = "STUB: not implemented"; return false }

func (p *Path) Extract(data []byte, optFuncs ...DecodeOptionFunc) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Path) PathString() string { _ = "STUB: not implemented"; return "" }

func (p *Path) Unmarshal(data []byte, v interface{}, optFuncs ...DecodeOptionFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Path) Get(src, dst interface{}) error { _ = "STUB: not implemented"; return nil }
