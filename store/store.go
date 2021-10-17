package store

import (
	"encoding/json"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

// Store key-value db and have key holder
type Store struct {
	db map[string]string
	sync.RWMutex
	Ops *int32
}

// NewStore create new in-memory key-value db and return interval finisher chan
func NewStore() (*Store, chan struct{}) {
	var t int32 = 0
	store := &Store{db: map[string]string{}, Ops: &t}
	if _, err := os.Stat("./test.tmp"); os.IsExist(err) {
		store.UnMarshall()
	} else {
		store.Marshall()
	}
	qChan := store.intervalSaveStart(1)
	return store, qChan
}

// Get item with given key
func (k Store) Get(key string) string {
	k.RLock()
	defer k.RUnlock()
	return k.db[key]
}

// Put can item in store with given value , Put create own storeKey
func (k Store) Put(value string) {
	k.Lock()
	defer k.Unlock()
	k.db[string(*k.Ops)] = value
	k.increment()
}

// Marshall crete key-value store file for persistence mod
func (k Store) Marshall() error {
	k.Lock()
	defer k.Unlock()
	v, _ := json.Marshal(k.db)
	tempfile := string("test") + ".tmp"
	if err := os.WriteFile(tempfile, v, 0644); err != nil {
		return err
	}
	return nil
}

// UnMarshall is load up key-values in test.tmp dir.
func (k Store) UnMarshall() error {
	k.Lock()
	defer k.Unlock()

	tempfile := string("test") + ".tmp"
	contents, err := os.ReadFile(tempfile)
	if err != nil {
		return err
	}

	json.Unmarshal(contents, &k.db)

	return nil
}

// Flush delete all key-value object in to db
func (k Store) Flush() {
	k.Lock()
	defer k.Unlock()
	for v := range k.db {
		delete(k.db, v)
	}
	atomic.SwapInt32(k.Ops, 0)
}

func (k Store) intervalSaveStart(minute int) chan struct{} {
	ticker := time.NewTicker((time.Duration(minute) * time.Minute))
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				k.Marshall()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return quit

}
func (k Store) increment() int {
	atomic.AddInt32(k.Ops, 1)
	return int(*k.Ops)
}
func (k Store) decrement() int {
	atomic.AddInt32(k.Ops, -1)
	return int(*k.Ops)
}
func (k Store) zero() int {
	atomic.SwapInt32(k.Ops, 0)
	return int(*k.Ops)
}
