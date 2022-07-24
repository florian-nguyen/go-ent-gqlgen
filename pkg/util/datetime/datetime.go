package datetime

import (
	"fmt"
	"time"
)

func FormatDate(date time.Time) string {
	return fmt.Sprintf("%v", date.UTC())
}
