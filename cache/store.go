package cache

import (
	"encoding/json"
	"errors"
	"maps"
	"sync"
)

// RWMutex guards concurrent access — reads don't block each other; writes are exclusive.
type Store struct {
	store   map[string]json.RawMessage
	rwMutex sync.RWMutex
}

func New() *Store {
	storeMap := make(map[string]json.RawMessage)

	return &Store{store: storeMap}
}

func (s *Store) Get(id string) (json.RawMessage, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	val, ok := s.store[id]
	if !ok {
		return nil, errors.New("key does not exist")
	}

	return val, nil
}

func (s *Store) Set(key string, val json.RawMessage) {
	//Lock to allow concurrent eads and defer unlock
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	s.store[key] = val
}

func (s *Store) Delete(id string) error {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	_, ok := s.store[id]
	if !ok {
		return errors.New("key does not exist")
	}

	delete(s.store, id)
	return nil
}

func (s *Store) All() map[string]json.RawMessage {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	clone := maps.Clone(s.store)

	return clone
}
