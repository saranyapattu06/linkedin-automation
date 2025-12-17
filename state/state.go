package state

import (
	"encoding/json"
	"os"
	"sync"
)

type State struct {
	SentConnections map[string]bool `json:"sent_connections"`
	SentMessages    map[string]bool `json:"sent_messages"`
	mu              sync.Mutex
}

func LoadState(path string) (*State, error) {
	state := &State{
		SentConnections: make(map[string]bool),
		SentMessages:    make(map[string]bool),
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return state, nil // fresh state
	}

	_ = json.Unmarshal(file, state)
	return state, nil
}

func (s *State) Save(path string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, _ := json.MarshalIndent(s, "", "  ")
	return os.WriteFile(path, data, 0644)
}

func (s *State) MarkConnectionSent(profile string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.SentConnections[profile] = true
}

func (s *State) MarkMessageSent(profile string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.SentMessages[profile] = true
}
