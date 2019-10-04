package internal

import (
	"fmt"
	"reflect"

	log "github.com/sirupsen/logrus"
)

// CheckNil - Checks if given value is nil as expected, panics otherwise
func CheckNil(eid string, val interface{}) {
	if !isNil(val) {
		message := fmt.Sprintf(
			"[%s] checkNil: Received a non nil value \"%v\", but not expecting that. A bug propably!",
			eid, val,
		)
		log.Panic(message)
		panic(message)
	}
}

func isNil(val interface{}) bool {
	return val == nil || (reflect.ValueOf(val).Kind() == reflect.Ptr && reflect.ValueOf(val).IsNil())
}
