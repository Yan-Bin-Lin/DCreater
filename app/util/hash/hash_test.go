package hash

import "testing"

func TestGetHashString(t *testing.T) {
	if GetHashString("any_mos/project/blog") != "-ue6MLWCROIuILAxZzMNFFAX5kjx8tHmmUUL3xfUyk0" {
		t.Errorf("hash result wrong")
	}
}
