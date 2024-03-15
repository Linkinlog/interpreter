package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "hello"}
	hello2 := &String{Value: "hello"}
	diff1 := &String{Value: "diff"}
	diff2 := &String{Value: "diff"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
