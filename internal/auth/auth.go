package auth

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type sessionValue struct {
	id   int64
	name string
}

type sessionData map[string]sessionValue

type SessionStore struct {
	sessions sessionData
	mu       sync.RWMutex
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		sessions: make(sessionData),
	}
}

func (s *SessionStore) Create(userID int64, name string) (string, error) {
	bytes := make([]byte, 32)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	sessionID := hex.EncodeToString(bytes)

	s.mu.Lock()
	s.sessions[sessionID] = sessionValue{
		id:   userID,
		name: name,
	}
	s.mu.Unlock()

	return sessionID, nil
}

func (s *SessionStore) Get(sessionID string) (sessionValue, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.sessions[sessionID]
	return v, ok
}

func (s *SessionStore) Delete(sessionID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.sessions, sessionID)
}

func AuthGuard(s *SessionStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("sesid")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		sessionValue, ok := s.Get(sessionID)
		if !ok {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		c.Set("userID", sessionValue.id)
		c.Set("name", sessionValue.name)
		c.Next()
	}
}
