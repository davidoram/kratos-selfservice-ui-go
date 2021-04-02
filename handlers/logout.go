package handlers

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/benbjohnson/hashfs"
	sessions "github.com/goincremental/negroni-sessions"
)

// LogoutParams configure the Logout http handler
type LogoutParams struct {
	// FS provides access to static files
	FS *hashfs.FS

	// FlowRedirectURL is the kratos URL to redirect the browser to,
	// when the user wishes to logout, and the 'flow' query param is missing
	FlowRedirectURL string
}

// Logout handler clears the session & logs the user out
func (lp LogoutParams) Logout(w http.ResponseWriter, r *http.Request) {

	// Start the logout flow with Kratos if required
	flow := r.URL.Query().Get("flow")
	if flow == "" {
		log.Printf("No flow ID found in URL, initializing logout flow, redirect to %s", lp.FlowRedirectURL)
		http.Redirect(w, r, lp.FlowRedirectURL, http.StatusMovedPermanently)
		return
	}
	// Clear the session
	session := sessions.GetSession(r)
	session.Clear()

	dataMap := map[string]interface{}{
		"fs":          lp.FS,
		"pageHeading": "Logged Out",
	}
	if err := GetTemplate(logoutPage).Render("layout", w, r, dataMap); err != nil {
		ErrorHandler(w, r, err)
	}
}
