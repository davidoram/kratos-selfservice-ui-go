package handlers

import (
	"log"
	"net/http"

	"github.com/davidoram/kratos-selfservice-ui-go/middleware"
	sessions "github.com/goincremental/negroni-sessions"
)

// Dashboard page is accessible to logged in users only
func Dashboard(w http.ResponseWriter, r *http.Request) {
	session := sessions.GetSession(r)
	kratosSession := session.Get("kratosSession").(middleware.KratosSession)
	log.Print(kratosSession.JsonPretty())
	dataMap := map[string]interface{}{
		"kratosSession": kratosSession,
		"headers":       []string{},
	}
	if err := GetTemplate(dashboardPage).Render("layout", w, r, dataMap); err != nil {
		ErrorHandler(w, r, err)
	}
}
