package singleton

import "sync"

// Singleton holds the single instance data.
type Singleton struct {
	// Add fields as needed
	Data string
}

var (
	instance *Singleton
	once     sync.Once
)

// GetInstance returns the single instance of Singleton, creating it if necessary.
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{Data: "Initialized Singleton"}
	})
	return instance
}
