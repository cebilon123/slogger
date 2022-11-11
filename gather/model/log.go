package model

// Log represents structured log as
// the domain object.
type Log struct {
	Type    int32
	Message string
	Caller  string
}
