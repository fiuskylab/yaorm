package pgsql

import (
	"fmt"
	"strings"
)

// DSN is the DSN struct for storing data to
// set up a connection with PostgreSQL DB.
type DSN struct {
	User     string
	Password string
	Host     string
	Port     uint
	DBName   string
	Params   map[string]string
}

const (
	dsn = `postgresql://%s@%s/%s%s`
)

func (p *DSN) String() string {
	userSpec := strings.Join([]string{p.User, p.Password}, ":")
	hostSpec := strings.Join([]string{p.Host, fmt.Sprintf("%d", p.Port)}, ":")
	params := ""
	if len(p.Params) >= 0 {
		paramSlice := []string{}
		for k, v := range p.Params {
			paramSlice = append(paramSlice, strings.Join([]string{k, v}, "="))
		}
		if len(paramSlice) > 0 {
			params = "?" + strings.Join(paramSlice, "&")
		}
	}

	return fmt.Sprintf(dsn, userSpec, hostSpec, p.DBName, params)
}
