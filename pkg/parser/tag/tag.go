package tag

import (
	"fmt"
	"strings"

	"github.com/fiuskylab/yaorm/pkg/helpers"
)

func init() {
	for k, v := range tagStrMap {
		strTagMap[v] = k
	}
}

type tag uint

const (
	pk tag = iota
	trigger
	autoincrement
)

var (
	tagStrMap = map[tag]string{
		pk:            "pk",
		trigger:       "trigger",
		autoincrement: "autoincrement",
	}

	strTagMap = map[string]tag{}

	triggerVals = []string{
		"create", "update", "delete",
	}
)

// TODO:
// 	implement some kind of validation for each tag.
func (t tag) validate(val string) error {
	errs := &validationError{
		errorMessages: []string{},
	}

	tagStr := tagStrMap[t]

	// example:
	switch t {
	case trigger:
		// example:
		if !helpers.InSlice(val, triggerVals) {
			errs.append(fmt.Sprintf(invalidTagErr, tagStr, strings.Join(triggerVals, ", ")))
		}
		fallthrough
	default:
	}
	return errs
}
