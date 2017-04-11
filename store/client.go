package store

import (
	"errors"
	"sync"

	"github.com/haoyixin/oauth2"
)

func NewClientStore() *ClientStore {
	return &ClientStore{
		data: make(map[string]oauth2.ClientInfo),
	}
}

// ClientStore client information store
type ClientStore struct {
	sync.RWMutex
	data map[string]oauth2.ClientInfo
}

// RemoveByID deleting the client information by the ID
func (cs *ClientStore) RemoveByID(id string) (err error) {
	cs.Lock()
	defer cs.Unlock()
	delete(cs.data, id)
	return
}

// GetByID according to the ID for the client information
func (cs *ClientStore) GetByID(id string) (cli oauth2.ClientInfo, err error) {
	cs.RLock()
	defer cs.RUnlock()
	if c, ok := cs.data[id]; ok {
		cli = c
		return
	}
	err = errors.New("not found")
	return
}

// Create set client information
func (cs *ClientStore) Create(cli oauth2.ClientInfo) (err error) {
	cs.Lock()
	defer cs.Unlock()
	cs.data[cli.GetID()] = cli
	return
}
