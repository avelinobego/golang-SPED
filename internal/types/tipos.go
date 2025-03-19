package types

import (
	"fmt"
	"time"
)

func Periodo(de time.Time, para time.Time) string {
	return fmt.Sprintf("%s,%s", de.Format("2006-01"), para.Format("2006-01"))
}
