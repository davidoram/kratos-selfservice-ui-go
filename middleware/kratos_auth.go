package middleware

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/davidoram/kratos-selfservice-ui-go/session"
)

// KratosAuthParams configure the KratosAuth http handler
type KratosAuthParams struct {
	session.SessionStore

	// WhoAmIURL is the API endpoint fo the Kratis 'whoami' call that returns the
	// details of an authenticated session
	WhoAmIURL string

	// RedirectUnauthURL is where we will rerirect to if the session is
	// not associated with a valid user
	RedirectUnauthURL string
}

// KratoAuthMiddleware retrieves the user from the session via Kratos WhoAmIURL,
// and if the user is authenticated the request will proceed through the middleware chain.
// If the session is not authenticated, redirects to the RedirectUnauthURL
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
		start := time.Now()
		log.Printf("Calling Kratos %s", p.WhoAmIURL)
		var res *http.Response
		if res, err = client.Do(req); err != nil {
			log.Printf("Error calling %s, error: %v, redirect to %s", p.WhoAmIURL, err, p.RedirectUnauthURL)
			http.Redirect(w, r, p.RedirectUnauthURL, http.StatusPermanentRedirect)
			return
		}
		end := time.Now()
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

		// Save the kratos session in our application session
		if err := p.SaveKratosSession(w, r, string(body)); err != nil {
			log.Printf("Error saving session: %v, redirect to %s", err, p.RedirectUnauthURL)
			http.Redirect(w, r, p.RedirectUnauthURL, http.StatusPermanentRedirect)
			return
		}
		duration := end.Sub(start)
		log.Printf("kratosSession retrieved ok, took %v", duration)

		next.ServeHTTP(w, r)
	})
}
