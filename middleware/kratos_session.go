package middleware

import (
	"github.com/tidwall/gjson"
)

type KratosSession struct {
	session *string
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
