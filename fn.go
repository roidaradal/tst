package tst

import (
	"errors"
	"maps"
	"reflect"
	"slices"
	"testing"
)

type assertFn[T any] = func(*testing.T, string, T, T)

// Convert maps the list of items to a new list using the conversion function
func Convert[T any](items []T, convert func(T) T) []T {
	items2 := make([]T, len(items))
	for i, item := range items {
		items2[i] = convert(item)
	}
	return items2
}

// AssertTrue asserts that the given condition is true
func AssertTrue(t *testing.T, name string, condition bool) {
	if condition != true {
		t.Errorf("%s = condition failed", name)
	}
}

// AssertFalse asserts that the given condition is false
func AssertFalse(t *testing.T, name string, condition bool) {
	if condition != false {
		t.Errorf("%s = condition failed", name)
	}
}

// AssertDeepEqual asserts that the two `any` items are deeply equal
func AssertDeepEqual(t *testing.T, name string, a, b any) {
	if reflect.DeepEqual(a, b) == false {
		t.Errorf("%s = %v, want %v", name, a, b)
	}
}

// AssertDeepEqualAnd asserts that the two `any` items are deeply equal and the boolean flags are equal
func AssertDeepEqualAnd(t *testing.T, name string, a, b any, flag1, flag2 bool) {
	if flag1 != flag2 || reflect.DeepEqual(a, b) == false {
		t.Errorf("%s = %v, %t, want %v, %t", name, a, flag1, b, flag2)
	}
}

// AssertDeepEqualError asserts that the two `any` items are deeply equal and that the errors are equal
func AssertDeepEqualError(t *testing.T, name string, a, b any, err1, err2 error) {
	if !errors.Is(err1, err2) || reflect.DeepEqual(a, b) == false {
		t.Errorf("%s = %v, %v, want %v, %v", name, a, err1, b, err2)
	}
}

// AssertEqual asserts that the two given values are equal
func AssertEqual[T comparable](t *testing.T, name string, a, b T) {
	if a != b {
		t.Errorf("%s = %v, want %v", name, a, b)
	}
}

// AssertEqualAnd asserts that the two given values are equal and the boolean flags are equal
func AssertEqualAnd[T comparable](t *testing.T, name string, a, b T, flag1, flag2 bool) {
	if flag1 != flag2 || a != b {
		t.Errorf("%s = %v, %t, want %v, %t", name, a, flag1, b, flag2)
	}
}

// AssertEqualError asserts that the two given values are equal and the errors are equal
func AssertEqualError[T comparable](t *testing.T, name string, a, b T, err1, err2 error) {
	if !errors.Is(err1, err2) || a != b {
		t.Errorf("%s = %v, %v, want %v, %v", name, a, err1, b, err2)
	}
}

// AssertEqualAny checks if two `any` items are equal
func AssertEqualAny(t *testing.T, name string, a, b any) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("%s panicked", name)
		}
	}()
	if a != b {
		t.Errorf("%s = %v, want %v", name, a, b)
	}
}

// AssertEqualAnyAnd checks if two `any` items are equal and the boolean flags are equal
func AssertEqualAnyAnd(t *testing.T, name string, a, b any, flag1, flag2 bool) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("%s panicked", name)
		}
	}()
	if flag1 != flag2 || a != b {
		t.Errorf("%s = %v, %t, want %v, %t", name, a, flag1, b, flag2)
	}
}

// AssertEqualAnyError checks if two `any` items are equal and the errors are equal
func AssertEqualAnyError(t *testing.T, name string, a, b any, err1, err2 error) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("%s panicked", name)
		}
	}()
	if !errors.Is(err1, err2) || a != b {
		t.Errorf("%s = %v, %v, want %v, %v", name, a, err1, b, err2)
	}
}

// AssertListEqual asserts that the two given lists are equal
func AssertListEqual[S ~[]T, T comparable](t *testing.T, name string, a, b S) {
	if slices.Equal(a, b) == false {
		t.Errorf("%s = %v, want %v", name, a, b)
	}
}

// AssertListEqualAnd asserts that the two given lists are equal and the boolean flags are equal
func AssertListEqualAnd[S ~[]T, T comparable](t *testing.T, name string, a, b S, flag1, flag2 bool) {
	if flag1 != flag2 || slices.Equal(a, b) == false {
		t.Errorf("%s = %v, %t, want %v, %t", name, a, flag1, b, flag2)
	}
}

// AssertListEqualError asserts that the two given lists are equal and the errors are equal
func AssertListEqualError[S ~[]T, T comparable](t *testing.T, name string, a, b S, err1, err2 error) {
	if !errors.Is(err1, err2) || slices.Equal(a, b) == false {
		t.Errorf("%s = %v, %v, want %v, %v", name, a, err1, b, err2)
	}
}

// AssertMapEqual asserts that the two given maps are equal
func AssertMapEqual[M1, M2 ~map[K]V, K, V comparable](t *testing.T, name string, a M1, b M2) {
	if maps.Equal(a, b) == false {
		t.Errorf("%s = %v, want %v", name, a, b)
	}
}

// AssertMapEqualAnd asserts that the two given maps are equal and the boolean flags are equal
func AssertMapEqualAnd[M1, M2 ~map[K]V, K, V comparable](t *testing.T, name string, a M1, b M2, flag1, flag2 bool) {
	if flag1 != flag2 || maps.Equal(a, b) == false {
		t.Errorf("%s = %v, %t, want %v, %t", name, a, flag1, b, flag2)
	}
}

// AssertMapEqualError asserts that the two given maps are equal and the errors are equal
func AssertMapEqualError[M1, M2 ~map[K]V, K, V comparable](t *testing.T, name string, a M1, b M2, err1, err2 error) {
	if !errors.Is(err1, err2) || maps.Equal(a, b) == false {
		t.Errorf("%s = %v, %v, want %v, %v", name, a, err1, b, err2)
	}
}

// AssertPanic asserts that the end of the function will panic
// Usage: defer AssertPanic(t, name)
func AssertPanic(t *testing.T, name string) {
	if err := recover(); err == nil {
		t.Errorf("%s did not panic", name)
	}
}

// assertTest calls the test function
func assertTest(t *testing.T, name string, test func() bool) {
	if test != nil && test() == false {
		t.Errorf("%s post test failed", name)
	}
}
