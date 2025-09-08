package outbound

import (
	"fmt"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	res, err := GetExchangeRate(time.Now(), USD)
	print(err)
	fmt.Printf("%v", res)
}
