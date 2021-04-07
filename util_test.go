package memredis

import (
	"fmt"
	"testing"
)

func TestGET(t *testing.T) {
	fmt.Println(isMatch("x_*", "x_2f23f"))
}
