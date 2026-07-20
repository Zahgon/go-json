package benchmark

import "github.com/francoispqt/gojay"

var SmallFixture = []byte(`{"st": 1,"sid": 486,"tt": "active","gr": 0,"uuid": "de305d54-75b4-431b-adb2-eb6b9e546014","ip": "127.0.0.1","ua": "user_agent","tz": -6,"v": 1}`)

type SmallPayload struct {
	St   int
	Sid  int
	Tt   string
	Gr   int
	Uuid string
	Ip   string
	Ua   string
	Tz   int
	V    int
}

type SmallPayloadFFJson struct {
	St   int
	Sid  int
	Tt   string
	Gr   int
	Uuid string
	Ip   string
	Ua   string
	Tz   int
	V    int
}

//easyjson:json
type SmallPayloadEasyJson struct {
	St   int
	Sid  int
	Tt   string
	Gr   int
	Uuid string
	Ip   string
	Ua   string
	Tz   int
	V    int
}

func (t *SmallPayload) MarshalJSONObject(enc *gojay.Encoder) { _ = "STUB: not implemented"; return }

func (t *SmallPayload) IsNil() bool { _ = "STUB: not implemented"; return false }

func (t *SmallPayload) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *SmallPayload) NKeys() int { _ = "STUB: not implemented"; return 0 }

func NewSmallPayload() *SmallPayload { _ = "STUB: not implemented"; return nil }

func NewSmallPayloadEasyJson() *SmallPayloadEasyJson { _ = "STUB: not implemented"; return nil }

func NewSmallPayloadFFJson() *SmallPayloadFFJson { _ = "STUB: not implemented"; return nil }
