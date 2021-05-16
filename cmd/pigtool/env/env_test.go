package env

import (
	"fmt"
	"testing"
)

func TestGetGoVersionNum(t *testing.T) {
	num := GetGoVersionNum()
	fmt.Println(num)
}
