package auth

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type SessionStore struct {
	sessions map[string]int64
	mu       sync.RWMutex
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		sessions: make(map[string]int64),
	}
}

func (s *SessionStore) Create(userID int64) (string, error) {
	bytes := make([]byte, 32)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	sessionID := hex.EncodeToString(bytes)

	s.mu.Lock()
	s.sessions[sessionID] = userID
	s.mu.Unlock()

	return sessionID, nil
}

func (s *SessionStore) Get(sessionID string) (int64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	userID, ok := s.sessions[sessionID]
	return userID, ok
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

		userID, ok := s.Get(sessionID)
		if !ok {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
