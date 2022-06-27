package driver

// DSN is the interface that wraps up DSN builder
// for each Driver type.
type DSN interface {
	String() string
}
