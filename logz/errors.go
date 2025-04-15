package logz

import "errors"

// Aggregat all common errors here
var (
	InitDefaultErr error = errors.New("unable to define default logz")
)

// Aggregat all certain errors here
var (
	OtelNoSpanId  error = errors.New("unavailable span-id")
	OtelNoTraceId error = errors.New("unavailable trace-id")
)
