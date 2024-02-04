package omnistore

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	testCases := []struct {
		name          string
		key           string
		inputValue    any
		expectedValue any
	}{
		{
			name:          "set an int value",
			key:           "myint",
			inputValue:    42,
			expectedValue: 42,
		},
		{
			name:          "set a float value",
			key:           "myfloat",
			inputValue:    3.14,
			expectedValue: 3.14,
		},
		{
			name:          "set a struct value",
			key:           "mystruct",
			inputValue:    struct{ value string }{"test"},
			expectedValue: struct{ value string }{"test"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			Set(testCase.key, testCase.inputValue)
			if !reflect.DeepEqual(internalStore[testCase.key], testCase.expectedValue) {
				t.Errorf("expected key: %s to have value: %#v but had value: %#v\n",
					testCase.key, testCase.expectedValue, internalStore[testCase.key])
			}
		})
	}
}

type MyEnum string

func (m MyEnum) String() string {
	return string(m)
}

func TestStringerSet(t *testing.T) {
	testCases := []struct {
		name          string
		key           MyEnum
		inputValue    any
		expectedValue any
	}{
		{
			name:          "set an int value",
			key:           MyEnum("myint"),
			inputValue:    42,
			expectedValue: 42,
		},
		{
			name:          "set a float value",
			key:           MyEnum("myfloat"),
			inputValue:    3.14,
			expectedValue: 3.14,
		},
		{
			name:          "set a struct value",
			key:           MyEnum("mystruct"),
			inputValue:    struct{ value string }{"test"},
			expectedValue: struct{ value string }{"test"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			StringerSet(testCase.key, testCase.inputValue)
			if !reflect.DeepEqual(internalStore[testCase.key.String()], testCase.expectedValue) {
				t.Errorf("expected key: %s to have value: %#v but had value: %#v\n",
					testCase.key, testCase.expectedValue, internalStore[testCase.key.String()])
			}
		})
	}
}

func TestGet(t *testing.T) {
	t.Run("get a zero value", func(t *testing.T) {
		key := "myzero"
		input := 0
		got := Get[int](key)
		if !reflect.DeepEqual(input, got) {
			t.Errorf("expected key: %s to have value: %#v but had value: %#v\n",
				key, input, got)
		}
	})
	t.Run("get an int value", func(t *testing.T) {
		key := "myint"
		input := 42
		Set(key, input)
		got := Get[int](key)
		if !reflect.DeepEqual(input, got) {
			t.Errorf("expected key: %s to have value: %#v but had value: %#v\n",
				key, input, got)
		}
	})
	t.Run("get a float value", func(t *testing.T) {
		key := "myfloat"
		input := 3.14
		Set(key, input)
		got := Get[float64](key)
		if !reflect.DeepEqual(input, got) {
			t.Errorf("expected key: %s to have value: %#v but had value: %#v\n",
				key, input, got)
		}
	})
	t.Run("get a string value", func(t *testing.T) {
		key := "mystring"
		input := "hello, world!"
		Set(key, input)
		got := Get[string](key)
		if !reflect.DeepEqual(input, got) {
			t.Errorf("expected key: %s to have value: %#v but had value: %#v\n",
				key, input, got)
		}
	})
	t.Run("get a struct value", func(t *testing.T) {
		key := "mystruct"
		input := struct {
			svalue string
			ivalue int
		}{"hello, world!", 42}
		Set(key, input)
		got := Get[struct {
			svalue string
			ivalue int
		}](key)
		if !reflect.DeepEqual(input, got) {
			t.Errorf("expected key: %s to have value: %#v but had value: %#v\n",
				key, input, got)
		}
	})
}

func TestStringerGet(t *testing.T) {
	t.Run("get a zero value", func(t *testing.T) {
		key := MyEnum("mystringerzero")
		input := 0
		got := StringerGet[MyEnum, int](key)
		if !reflect.DeepEqual(input, got) {
			t.Errorf("expected key: %s to have value: %#v but had value: %#v\n",
				key, input, got)
		}
	})
	t.Run("get an int value", func(t *testing.T) {
		key := MyEnum("mystringerint")
		input := 42
		StringerSet(key, input)
		got := StringerGet[MyEnum, int](key)
		if !reflect.DeepEqual(input, got) {
			t.Errorf("expected key: %s to have value: %#v but had value: %#v\n",
				key, input, got)
		}
	})
	t.Run("get a float value", func(t *testing.T) {
		key := MyEnum("mystringerfloat")
		input := 3.14
		StringerSet(key, input)
		got := StringerGet[MyEnum, float64](key)
		if !reflect.DeepEqual(input, got) {
			t.Errorf("expected key: %s to have value: %#v but had value: %#v\n",
				key, input, got)
		}
	})
	t.Run("get a string value", func(t *testing.T) {
		key := MyEnum("mystringerstring")
		input := "hello, world!"
		StringerSet(key, input)
		got := StringerGet[MyEnum, string](key)
		if !reflect.DeepEqual(input, got) {
			t.Errorf("expected key: %s to have value: %#v but had value: %#v\n",
				key, input, got)
		}
	})
	t.Run("get a struct value", func(t *testing.T) {
		key := MyEnum("mystringerstruct")
		input := struct {
			svalue string
			ivalue int
		}{"hello, world!", 42}
		StringerSet(key, input)
		got := StringerGet[MyEnum, struct {
			svalue string
			ivalue int
		}](key)
		if !reflect.DeepEqual(input, got) {
			t.Errorf("expected key: %s to have value: %#v but had value: %#v\n",
				key, input, got)
		}
	})
}
