package engine

type key int

const (
	// ContextOriginalPath holds the original requrest URL
	ContextOriginalPath key = iota
	// ContextRequestStart holds the time when the request was received
	ContextRequestStart
	// ContextUserID holds the user id
	ContextUserID
)
