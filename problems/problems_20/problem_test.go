package problems20

/*

Unit tests for problems 11-20

*/

import (
	"reflect"
	"testing"
)

func Test_Problem11(t *testing.T) {
	test := autoComplete("de")
	expected := []string{"deer", "deal"}
	if !reflect.DeepEqual(test, expected) {
		t.Errorf("problem_11/test01: Expected %v, got %v", expected, test)
	}

}
