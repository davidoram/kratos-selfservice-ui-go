package handlers

import (
	"net/http"

	"github.com/davidoram/kratos-selfservice-ui-go/session"
)

// HomeParams configure the Home http handler
type HomeParams struct {
	session.SessionStore
}

// Home displays a simple homepage
func (p HomeParams) Home(w http.ResponseWriter, r *http.Request) {
	dataMap := map[string]interface{}{
		"kratosSession": p.GetKratosSession(r),
		"headers":       []string{},
	}
	if err := GetTemplate(homePage).Render("layout", w, r, dataMap); err != nil {
		ErrorHandler(w, r, err)
	}
}
