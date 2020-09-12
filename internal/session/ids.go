package session

import (
	"fmt"

	"github.com/hashicorp/boundary/internal/db"
)

const (
	// SessionPrefix for session PK ids
	SessionPrefix = "s"

	// StatePrefix for state PK ids
	StatePrefix = "ss"

	// ConnectionPrefix for connection PK ids
	ConnectionPrefix = "sc"
)

func newId() (string, error) {
	id, err := db.NewPublicId(SessionPrefix)
	if err != nil {
		return "", fmt.Errorf("new session id: %w", err)
	}
	return id, nil
}

func newStateId() (string, error) {
	id, err := db.NewPublicId(StatePrefix)
	if err != nil {
		return "", fmt.Errorf("new session state id: %w", err)
	}
	return id, nil
}

func newConnectionId() (string, error) {
	id, err := db.NewPublicId(ConnectionPrefix)
	if err != nil {
		return "", fmt.Errorf("new session connection id: %w", err)
	}
	return id, nil
}