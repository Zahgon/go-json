package benchmark

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson21677a1cDecodeBenchmark(in *jlexer.Lexer, out *SmallPayloadEasyJson) {
	_ = "STUB: not implemented"
	return
}

func easyjson21677a1cEncodeBenchmark(out *jwriter.Writer, in SmallPayloadEasyJson) {
	_ = "STUB: not implemented"
	return
}

func (v SmallPayloadEasyJson) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v SmallPayloadEasyJson) MarshalEasyJSON(w *jwriter.Writer) { _ = "STUB: not implemented"; return }

func (v *SmallPayloadEasyJson) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *SmallPayloadEasyJson) UnmarshalEasyJSON(l *jlexer.Lexer) {
	_ = "STUB: not implemented"
	return
}
