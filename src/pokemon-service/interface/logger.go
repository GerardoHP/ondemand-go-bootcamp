package interfaces

// Methods available to log either error or access
type Logger interface {
	LogError(string, ...interface{})
	LogAccess(string, ...interface{})
}
