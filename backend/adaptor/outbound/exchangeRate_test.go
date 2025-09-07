package outbound

import (
	"fmt"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	res, err := GetExchangeRate(time.Now().Add(-24 * time.Hour))
	print(err)
	fmt.Printf("%v", res)
}
