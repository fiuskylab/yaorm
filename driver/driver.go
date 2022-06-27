package driver

// Driver is a wrap of a Driver methods
type Driver interface {
	Ping() error
	Migrate(models ...any) error
}
