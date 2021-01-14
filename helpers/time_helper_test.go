package helpers

import "testing"

func TestString2Time(t *testing.T) {

	in := "2020-12-15T17:44:08+08:00"
	result := String2Time(in)

	t.Log("result => ", result)
}
