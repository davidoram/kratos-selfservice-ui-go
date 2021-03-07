package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// HomeParams configure the Home http handler
type HomeParams struct {

	// Session store
	Store *sessions.CookieStore
}

// Home displays a simple homepage
func (p HomeParams) Home(w http.ResponseWriter, r *http.Request) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := p.Store.Get(r, "my-app-session")
	dataMap := map[string]interface{}{
		"kratosSession": session.Values["kratosSession"],
		"headers":       []string{},
	}
	if err := GetTemplate(homePage).Render("layout", w, r, dataMap); err != nil {
		ErrorHandler(w, r, err)
	}
}
