package nano

import (
	"time"

)

func Measure() string {
	now := time.Now()
	now_rfc3339 := now.Format(time.RFC3339Nano)

	return now_rfc3339
}
