package serve

import (
	"testing"
)

func TestServe_SplitWork(t *testing.T) {
	type expect struct {
		project []string
		blog    string
	}

	// build test table
	var splitTest = []struct {
		url     string // input
		expect_ expect // expected result
	}{
		{"/foo", expect{[]string{""}, "foo"}},
		{"/foo/bar", expect{[]string{"", "foo"}, "bar"}},
		{"/foo/bar/blog", expect{[]string{"", "foo", "bar"}, "blog"}},
	}

	for _, value := range splitTest {
		proj, blog := splitWork(value.url)
		for i, v := range proj {
			if v != value.expect_.project[i] {
				t.Errorf("splitWork(%s).project[%d] = %s; expected %s\n", value.url, i, v, value.expect_.project[i])
			}
		}
		// test blog
		if blog != value.expect_.blog {
			t.Errorf("splitWork(%s).blog = %s; expected %s\n", value.url, blog, value.expect_.blog)
		}
	}
}
