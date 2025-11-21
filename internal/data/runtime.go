package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJOSNValue := strconv.Quote(jsonValue)
	return []byte(quotedJOSNValue), nil
}
