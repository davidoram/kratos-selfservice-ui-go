// session package provides typesafe access to session data
package session

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

// SessionStore holds a connection to the application Session store
type SessionStore struct {
	// Session store
	Store *sessions.CookieStore
}

const (
	// The cookie we use to store this applications session in
	sessionCookieName = "kgc-sess"

	// Keys we store in our application session
	keyKratosSession = "kratosSession"
)

// SaveKratosSession stores a string as the KratosSession value
func (s SessionStore) SaveKratosSession(w http.ResponseWriter, r *http.Request, ks string) error {

	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, err := s.Store.Get(r, sessionCookieName)
	if err != nil {
		log.Printf("Error decoding session, %v", err)
		return err
	}

	// Add the value into the session store
	session.Values[keyKratosSession] = ks

	// Save it before we write to the response/return from the handler.
	return session.Save(r, w)
}

// GetKratosSession returns the KratosSession or nil
func (s SessionStore) GetKratosSession(r *http.Request) *KratosSession {

	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, err := s.Store.Get(r, sessionCookieName)
	if err != nil {
		log.Printf("Error decoding session, %v", err)
		return nil
	}
	if v, exists := session.Values[keyKratosSession]; exists {
		ks := NewKratosSession(v.(string))
		return &ks
	}
	return nil

}
