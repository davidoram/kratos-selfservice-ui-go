package handlers

import (
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
)

// Home displays a simple homepage
func Home(w http.ResponseWriter, r *http.Request) {
	session := sessions.GetSession(r)
	kratosSession := session.Get("kratosSession")
	dataMap := map[string]interface{}{
		"kratosSession": kratosSession,
		"headers":       []string{},
	}
	if err := GetTemplate(homePage).Render("layout", w, r, dataMap); err != nil {
		ErrorHandler(w, r, err)
	}
}
