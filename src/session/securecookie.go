package session

import (
	"github.com/gorilla/securecookie"
	"net/http"
)

type Session struct {
	secureCookie *securecookie.SecureCookie
	r            *http.Request
	w            http.ResponseWriter
}

func New(r *http.Request, w http.ResponseWriter) *Session {
	session := &Session{} // not thread safe
	var hashKey = []byte("osmanthuswine-very-secret")
	// Block keys should be 16 bytes (AES-128) or 32 bytes (AES-256) long.
	// Shorter keys may weaken the encryption used.
	var blockKey = []byte("osmanthuswine-lot-secret")
	session.secureCookie = securecookie.New(hashKey, blockKey)
	session.r = r
	session.w = w
	return session
}

func (session *Session) GetSession() map[string]string {
	value := make(map[string]string)
	if cookie, err := session.r.Cookie("osmseccidhas"); err == nil {
		session.secureCookie.Decode("osmseccidhas", cookie.Value, value)
	}
	return value
}

func (session *Session) SetSession(value map[string]string) {
	if value == nil {
		return
	}
	if encoded, err := session.secureCookie.Encode("osm-sec-cid-has", value); err == nil {
		cookie := &http.Cookie{
			Name:     "osmseccidhas",
			Value:    encoded,
			Path:     "/",
			Secure:   false,
			HttpOnly: false,
		}
		http.SetCookie(session.w, cookie)
	}
}
