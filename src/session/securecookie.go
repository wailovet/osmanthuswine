package session

import (
	"github.com/gorilla/securecookie"
	"net/http"
)

type Session struct {
	secureCookie *securecookie.SecureCookie
}

var instanceSession *Session

func GetInstanceSession() *Session {
	if instanceSession == nil {
		instanceSession = &Session{} // not thread safe

		// Hash keys should be at least 32 bytes long
		var hashKey = []byte("osmanthuswine-very-secret")
		// Block keys should be 16 bytes (AES-128) or 32 bytes (AES-256) long.
		// Shorter keys may weaken the encryption used.
		var blockKey = []byte("osmanthuswine-lot-secret")
		instanceSession.secureCookie = securecookie.New(hashKey, blockKey)
	}
	return instanceSession
}

func GetSession(r *http.Request) map[string]string {
	value := make(map[string]string)
	if cookie, err := r.Cookie("osm-sec-cid-has"); err == nil {
		session := GetInstanceSession()
		session.secureCookie.Decode("osm-sec-cid-has", cookie.Value, value)
	}
	return value
}

func SetSession(w http.ResponseWriter, value map[string]string) {
	session := GetInstanceSession()
	if encoded, err := session.secureCookie.Encode("osm-sec-cid-has", value); err == nil {
		cookie := &http.Cookie{
			Name:     "osm-sec-cid-has",
			Value:    encoded,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
}
