package helpers

import (
	"reflect"
	"sort"
)

// InArray will search element inside array with any type.
// Will return boolean and index for matched element.
// True and index more than 0 if element is exist.
// needle is element to search, haystack is slice of value to be search.
// Thanks to https://github.com/SimonWaldherr/golang-examples/blob/master/advanced/in_array.go
func InArray(needle interface{}, haystack interface{}) (index int, exists bool) {
	exists = false
	index = -1

	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

// InStringArray is an optimized version of InArray for string arrays
// It is useful only for big arrays and repeted searches
func InStringArray(needle string, haystack []string) (index int, exists bool) {
	exists = false
	index = -1

	sort.Strings(haystack)
	i := sort.SearchStrings(haystack, needle)

	if i < len(haystack) && haystack[i] == needle {
		exists = true
		index = i
		return
	}

	return
}