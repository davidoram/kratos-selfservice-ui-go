package handlers

import (
	"log"
	"net/http"

	"github.com/davidoram/kratos-selfservice-ui-go/middleware"
	"github.com/gorilla/sessions"
)

// DashboardParams configure the Dashboard http handler
type DashboardParams struct {
	// Session store
	Store *sessions.CookieStore
}

// Dashboard page is accessible to logged in users only
func (p DashboardParams) Dashboard(w http.ResponseWriter, r *http.Request) {
	log.Printf("dashboard")

	session, _ := p.Store.Get(r, "my-app-session")
	kratosSession := middleware.NewKratosSession(session.Values["kratosSession"].(string))

	log.Print(kratosSession.JsonPretty())
	dataMap := map[string]interface{}{
		"kratosSession": kratosSession,
		"headers":       []string{},
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
