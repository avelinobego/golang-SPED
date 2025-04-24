package tipos

import (
	"time"
)

type Data time.Time

func (p Data) MarshalText() ([]byte, error) {
	result := time.Time(p).Format("2006-01")
	return []byte(result), nil
}
