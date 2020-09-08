package __kyu

import (
	"reflect"
	"testing"
)

func TestCreatePhoneNumber(t *testing.T) {
	t.Run("should create phone number", func(t *testing.T) {
		got := CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
		want := "(123) 456-7890"
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got: %+v, want: %+v", got, want)
		}
	})
}
