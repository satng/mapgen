package example

import "sync"

/*
   GENERATED FILE, DO NOT EDIT

   This file was generated by https://github.com/ammario/mapgen
*/

// ChannelIntMap is a generated thread safe map
// with key Channel and value int
type ChannelIntMap struct {
	mu sync.Mutex
	m  map[Channel]int
}

// NewChannelIntMap returns an instantiated thread safe map
// with key Channel and value int
func NewChannelIntMap() *ChannelIntMap {
	return &ChannelIntMap{
		m: make(map[Channel]int),
	}
}

// Lock locks the map and returns it.
func (m *ChannelIntMap) Lock() map[Channel]int {
	m.mu.Lock()
	return m.m
}

// Unlock unlocks the map
func (m *ChannelIntMap) Unlock() {
	m.mu.Unlock()
}

// Set sets a key on the map
func (m *ChannelIntMap) Set(key Channel, val int) {
	m.mu.Lock()
	m.m[key] = val
	m.mu.Unlock()
}

// Delete removes a key from the map
func (m *ChannelIntMap) Delete(key Channel) {
	m.mu.Lock()
	delete(m.m, key)
	m.mu.Unlock()
}

// Get retrieves a key from the map
func (m *ChannelIntMap) Get(key Channel) int {
	m.mu.Lock()
	v := m.m[key]
	m.mu.Unlock()
	return v
}

// Len returns the length of the map
func (m *ChannelIntMap) Len() int {
	m.mu.Lock()
	n := len(m.m)
	m.mu.Unlock()
	return n
}

// GetEx retrieves a key from the map
// and whether it exists
func (m *ChannelIntMap) GetEx(key Channel) (int, bool) {
	m.mu.Lock()
	v, exists := m.m[key]
	m.mu.Unlock()
	return v, exists
}

// Exists returns if a key exists
func (m *ChannelIntMap) Exists(key Channel) bool {
	m.mu.Lock()
	_, exists := m.m[key]
	m.mu.Unlock()
	return exists
}
