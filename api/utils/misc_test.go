package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestNanoTime(t *testing.T) {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)
	millis := nanos / 1000000

	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
	if 1 == 2 {
		t.Errorf("Error")
	}
	fmt.Println(time.Now().UnixNano() / 1000000000)
}
