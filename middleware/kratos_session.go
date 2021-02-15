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
