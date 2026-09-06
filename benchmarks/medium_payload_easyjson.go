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

func easyjson8ca7813eDecodeBenchmark(in *jlexer.Lexer, out *MediumPayloadEasyJson) {
	_ = "STUB: not implemented"
	return
}

func easyjson8ca7813eEncodeBenchmark(out *jwriter.Writer, in MediumPayloadEasyJson) {
	_ = "STUB: not implemented"
	return
}

func (v MediumPayloadEasyJson) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v MediumPayloadEasyJson) MarshalEasyJSON(w *jwriter.Writer) {
	_ = "STUB: not implemented"
	return
}

func (v *MediumPayloadEasyJson) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *MediumPayloadEasyJson) UnmarshalEasyJSON(l *jlexer.Lexer) {
	_ = "STUB: not implemented"
	return
}

func easyjson8ca7813eDecodeBenchmark1(in *jlexer.Lexer, out *CBPerson) {
	_ = "STUB: not implemented"
	return
}

func easyjson8ca7813eEncodeBenchmark1(out *jwriter.Writer, in CBPerson) {
	_ = "STUB: not implemented"
	return
}

func easyjson8ca7813eDecodeBenchmark4(in *jlexer.Lexer, out *CBGravatar) {
	_ = "STUB: not implemented"
	return
}

func easyjson8ca7813eEncodeBenchmark4(out *jwriter.Writer, in CBGravatar) {
	_ = "STUB: not implemented"
	return
}

func easyjson8ca7813eDecodeBenchmark5(in *jlexer.Lexer, out *CBAvatar) {
	_ = "STUB: not implemented"
	return
}

func easyjson8ca7813eEncodeBenchmark5(out *jwriter.Writer, in CBAvatar) {
	_ = "STUB: not implemented"
	return
}

func easyjson8ca7813eDecodeBenchmark3(in *jlexer.Lexer, out *CBGithub) {
	_ = "STUB: not implemented"
	return
}

func easyjson8ca7813eEncodeBenchmark3(out *jwriter.Writer, in CBGithub) {
	_ = "STUB: not implemented"
	return
}

func easyjson8ca7813eDecodeBenchmark2(in *jlexer.Lexer, out *CBName) {
	_ = "STUB: not implemented"
	return
}

func easyjson8ca7813eEncodeBenchmark2(out *jwriter.Writer, in CBName) {
	_ = "STUB: not implemented"
	return
}
