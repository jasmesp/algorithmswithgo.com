package module01

import (
	"fmt"
	"reflect"
)

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	for _, i := range list {
		if i == num {
			return true
		} else {
			if reflect.TypeOf(i) != reflect.TypeOf(num) {
				fmt.Println("Type mismatch")
				return false
			}

		}

	}
	return false
}
