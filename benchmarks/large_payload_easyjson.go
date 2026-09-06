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

func easyjsonD519278DecodeBenchmark(in *jlexer.Lexer, out *LargePayloadEasyJson) {
	_ = "STUB: not implemented"
	return
}

func easyjsonD519278EncodeBenchmark(out *jwriter.Writer, in LargePayloadEasyJson) {
	_ = "STUB: not implemented"
	return
}

func (v LargePayloadEasyJson) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v LargePayloadEasyJson) MarshalEasyJSON(w *jwriter.Writer) { _ = "STUB: not implemented"; return }

func (v *LargePayloadEasyJson) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *LargePayloadEasyJson) UnmarshalEasyJSON(l *jlexer.Lexer) {
	_ = "STUB: not implemented"
	return
}

func easyjsonD519278DecodeBenchmark2(in *jlexer.Lexer, out *DSTopicsList) {
	_ = "STUB: not implemented"
	return
}

func easyjsonD519278EncodeBenchmark2(out *jwriter.Writer, in DSTopicsList) {
	_ = "STUB: not implemented"
	return
}

func easyjsonD519278DecodeBenchmark3(in *jlexer.Lexer, out *DSTopic) {
	_ = "STUB: not implemented"
	return
}

func easyjsonD519278EncodeBenchmark3(out *jwriter.Writer, in DSTopic) {
	_ = "STUB: not implemented"
	return
}

func easyjsonD519278DecodeBenchmark1(in *jlexer.Lexer, out *DSUser) {
	_ = "STUB: not implemented"
	return
}

func easyjsonD519278EncodeBenchmark1(out *jwriter.Writer, in DSUser) {
	_ = "STUB: not implemented"
	return
}
