package pkg

import "testing"

func TestGetAge(t *testing.T) {
	age, err := GetAge("1998-10-09")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("age: %v", age)
}
