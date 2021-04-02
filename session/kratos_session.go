package session

import (
	"bytes"
	"encoding/json"

	"github.com/tidwall/gjson"
)

// KratosSession is used to access information from a Kratos 'Session'
// JSON payload
type KratosSession struct {
	session *string
}

func NewKratosSession(s string) KratosSession {
	return KratosSession{&s}
}

func (ks KratosSession) Email() string {
	return gjson.Get(*ks.session, "identity.traits.email").String()
}

func (ks KratosSession) Id() string {
	return gjson.Get(*ks.session, "identity.id").String()
}

func (ks KratosSession) FirstName() string {
	return gjson.Get(*ks.session, "identity.traits.name.first").String()
}

func (ks KratosSession) LastName() string {
	return gjson.Get(*ks.session, "identity.traits.name.last").String()
}

func (ks KratosSession) AddressVerified() bool {
	return gjson.Get(*ks.session, "identity.verifiable_addresses.0.verified").Bool()
}

func (ks KratosSession) Json() string {
	return *ks.session
}

func (ks KratosSession) JsonPretty() string {
	buf := new(bytes.Buffer)
	err := json.Indent(buf, []byte(*ks.session), "", "  ")
	if err != nil {
		return *ks.session
	}
	return buf.String()
}
