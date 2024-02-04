package omnistore_test

import (
	"fmt"

	"github.com/JamesChung/omnistore"
)

func ExampleSet() {
	omnistore.Set[string]("examplestring", "myvalue")
	omnistore.Set[int]("exampleint", 42)
	omnistore.Set[float64]("examplefloat", 3.14)
}

type MyEnum string

func (m MyEnum) String() string {
	return string(m)
}

func ExampleStringerSet() {
	myString := MyEnum("examplestringerstring")
	omnistore.StringerSet[MyEnum](myString, "myvalue")

	myInt := MyEnum("examplestringerint")
	omnistore.StringerSet[MyEnum](myInt, 42)

	myFloat := MyEnum("examplestringerfloat")
	omnistore.StringerSet[MyEnum](myFloat, 3.14)
}

func ExampleGet() {
	omnistore.Set[string]("examplestring", "myvalue")
	fmt.Println(omnistore.Get[string]("examplestring"))

	omnistore.Set[int]("exampleint", 42)
	fmt.Println(omnistore.Get[int]("exampleint"))

	omnistore.Set[float64]("examplefloat", 3.14)
	fmt.Println(omnistore.Get[float64]("examplefloat"))
	// Output:
	// myvalue
	// 42
	// 3.14
}

func ExampleGetE() {
	omnistore.Set[string]("examplestring", "myvalue")
	fmt.Println(omnistore.GetE[string]("examplestring"))

	omnistore.Set[int]("exampleint", 42)
	fmt.Println(omnistore.GetE[int]("exampleint"))

	omnistore.Set[float64]("examplefloat", 3.14)
	fmt.Println(omnistore.GetE[float64]("examplefloat"))
	// Output:
	// myvalue <nil>
	// 42 <nil>
	// 3.14 <nil>
}

func ExampleStringerGet() {
	myString := MyEnum("examplestringerstring")
	omnistore.StringerSet[MyEnum](myString, "myvalue")
	fmt.Println(omnistore.StringerGet[MyEnum, string](myString))

	myInt := MyEnum("examplestringerint")
	omnistore.StringerSet[MyEnum](myInt, 42)
	fmt.Println(omnistore.StringerGet[MyEnum, int](myInt))

	myFloat := MyEnum("examplestringerfloat")
	omnistore.StringerSet[MyEnum](myFloat, 3.14)
	fmt.Println(omnistore.StringerGet[MyEnum, float64](myFloat))
	// Output:
	// myvalue
	// 42
	// 3.14
}

func ExampleStringerGetE() {
	myString := MyEnum("examplestringerstring")
	omnistore.StringerSet[MyEnum](myString, "myvalue")
	fmt.Println(omnistore.StringerGetE[MyEnum, string](myString))

	myInt := MyEnum("examplestringerint")
	omnistore.StringerSet[MyEnum](myInt, 42)
	fmt.Println(omnistore.StringerGetE[MyEnum, int](myInt))

	myFloat := MyEnum("examplestringerfloat")
	omnistore.StringerSet[MyEnum](myFloat, 3.14)
	fmt.Println(omnistore.StringerGetE[MyEnum, float64](myFloat))
	// Output:
	// myvalue <nil>
	// 42 <nil>
	// 3.14 <nil>
}
