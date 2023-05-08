package clouds

// Cloud is the interface that all clouds must implement
type Cloud interface {
	init() error
	// Get the name of the cloud
	close() error
}
