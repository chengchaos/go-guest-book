package paths

import (
	"log"
	"testing"
)

func Test_PrintPath(t *testing.T) {
	ap := GetPath(".")
	log.Println("ap => ", ap)

}
