package middleware

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

// KratosAuthParams configure the KratosAuth http handler
type KratosAuthParams struct {
	// WhoAmIURL is the API endpoint fo the Kratis 'whoami' call that returns the
	// details of an authenticated session
	WhoAmIURL string

	// RedirectUnauthURL is where we will rerirect to if the session is
	// not associated with a valid user
	RedirectUnauthURL string

	// Session store
	Store *sessions.CookieStore
}

// KratoAuthMiddleware retrieves the user from the session via Kratos WhoAmIURL,
// proceed through the middleware chain.
// If the session is not authenticated, will redirects to the RedirectUnauthURL
func (p KratosAuthParams) KratoAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		csrfCookie, csrfErr := r.Cookie("csrf_token")
		sessionCookie, sessionErr := r.Cookie("ory_kratos_session")

		if !(csrfErr == nil && sessionErr == nil) {
			log.Printf("Error retrieving cookies: csrf_token: %v, ory_kratos_session: %v, redirect to %s", csrfErr, sessionErr, p.RedirectUnauthURL)
			http.Redirect(w, r, p.RedirectUnauthURL, http.StatusPermanentRedirect)
			return
		}

		client := &http.Client{}
		var req *http.Request
		var err error
		if req, err = http.NewRequest("POST", p.WhoAmIURL, nil); err != nil {
			log.Printf("Error creating request: %v, redirect to %s", err, p.RedirectUnauthURL)
			http.Redirect(w, r, p.RedirectUnauthURL, http.StatusPermanentRedirect)
			return
		}

		// Add Cookies onto the next call
		req.AddCookie(csrfCookie)
		req.AddCookie(sessionCookie)

		// Setup Headers
		req.Header.Add("Accept", "application/json")

		// Make the API call
		log.Printf("Calling Kratos %s", p.WhoAmIURL)
		var res *http.Response
		if res, err = client.Do(req); err != nil {
			log.Printf("Error calling %s, error: %v, redirect to %s", p.WhoAmIURL, err, p.RedirectUnauthURL)
			http.Redirect(w, r, p.RedirectUnauthURL, http.StatusPermanentRedirect)
			return
		}
		if res.StatusCode != 200 {
			log.Printf("Error status code != 200: %d, redirect to %s", res.StatusCode, p.RedirectUnauthURL)
			http.Redirect(w, r, p.RedirectUnauthURL, http.StatusPermanentRedirect)
			return
		}
		var body []byte
		if body, err = ioutil.ReadAll(res.Body); err != nil {
			log.Printf("Error reading body: %v, redirect to %s", err, p.RedirectUnauthURL)
			http.Redirect(w, r, p.RedirectUnauthURL, http.StatusPermanentRedirect)
			return
		}
		defer res.Body.Close()

		s := string(body)

		// Get a session. We're ignoring the error resulted from decoding an
		// existing session: Get() always returns a session, even if empty.
		session, _ := p.Store.Get(r, "my-app-session")
		session.Values["kratosSession"] = s
		// Save it before we write to the response/return from the handler.
		if err := session.Save(r, w); err != nil {
			log.Printf("Error saving session: %v, redirect to %s", err, p.RedirectUnauthURL)
			http.Redirect(w, r, p.RedirectUnauthURL, http.StatusPermanentRedirect)
			return
		}

		log.Printf("kratosSession retrieved ok")

		next.ServeHTTP(w, r)
	})
}
