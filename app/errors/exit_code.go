package errors

// List of exit code available
const (
	_ = iota // 0 means success
	MissingConfigDatabase
	InitDatabase
)
