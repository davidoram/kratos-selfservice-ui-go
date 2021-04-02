package handlers

import (
	"log"
	"net/http"

	"github.com/benbjohnson/hashfs"
	"github.com/davidoram/kratos-selfservice-ui-go/session"
)

// DashboardParams configure the Dashboard http handler
type DashboardParams struct {
	// FS provides access to static files
	FS *hashfs.FS

	session.SessionStore
}

// Dashboard page is accessible to logged in users only, the proptection for that is provide by
// KratoAuthMiddleware middleware
func (p DashboardParams) Dashboard(w http.ResponseWriter, r *http.Request) {
	log.Printf("dashboard")

	dataMap := map[string]interface{}{
		"kratosSession": p.GetKratosSession(r),
		"headers":       []string{},
		"fs":            p.FS,
		"pageHeading":   "Dashboard",
	}
	if err := GetTemplate(dashboardPage).Render("layout", w, r, dataMap); err != nil {
		ErrorHandler(w, r, err)
	}
}

func (p DashboardParams) ResponderFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p.Dashboard(w, r)
	}
}
