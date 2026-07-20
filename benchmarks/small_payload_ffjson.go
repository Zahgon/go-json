package benchmark

import (
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (j *SmallPayloadFFJson) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *SmallPayloadFFJson) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	ffjtSmallPayloadFFJsonbase = iota
	ffjtSmallPayloadFFJsonnosuchkey

	ffjtSmallPayloadFFJsonSt

	ffjtSmallPayloadFFJsonSid

	ffjtSmallPayloadFFJsonTt

	ffjtSmallPayloadFFJsonGr

	ffjtSmallPayloadFFJsonUuid

	ffjtSmallPayloadFFJsonIp

	ffjtSmallPayloadFFJsonUa

	ffjtSmallPayloadFFJsonTz

	ffjtSmallPayloadFFJsonV
)

var ffjKeySmallPayloadFFJsonSt = []byte("St")

var ffjKeySmallPayloadFFJsonSid = []byte("Sid")

var ffjKeySmallPayloadFFJsonTt = []byte("Tt")

var ffjKeySmallPayloadFFJsonGr = []byte("Gr")

var ffjKeySmallPayloadFFJsonUuid = []byte("Uuid")

var ffjKeySmallPayloadFFJsonIp = []byte("Ip")

var ffjKeySmallPayloadFFJsonUa = []byte("Ua")

var ffjKeySmallPayloadFFJsonTz = []byte("Tz")

var ffjKeySmallPayloadFFJsonV = []byte("V")

func (j *SmallPayloadFFJson) UnmarshalJSON(input []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *SmallPayloadFFJson) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	_ = "STUB: not implemented"
	return nil
}
