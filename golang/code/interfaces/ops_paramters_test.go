package interfaces

import (
	"fmt"
	"testing"
)

func TestOptionsParameter(t *testing.T) {
	s1 := newStudent("test", 20)
	s2 := newStudent("test", 20, setAddress("cd"), setPhone(123))
	if s1 == s2 {
		t.Errorf("s1 is  %v, s2 is %v", s1, s2)
	}
	fmt.Println(s1, s2)
}
