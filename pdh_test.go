package win

import (
	"fmt"
	"testing"
)

func TestNameById(t *testing.T) {
	// Expect On German System for number 4 -> Arbeitsspeicher
	PdhLookupPerfNameByIndex(30)
	t.Fail()
}

func TestIdByName(t *testing.T) {
	ret := PdhLookupPerfIndexByName("hanau", "Arbeitsspeicher")
	fmt.Printf("ret: %v\n", ret)
	t.Fail()

}
