package model

// Log represents log payload.
type Log struct {
	Type    int32
	Message string
	Caller  string
}
