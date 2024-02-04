# omnistore

`omnistore` is designed to store anything and subsequently retrieve it.

## install

`go get -u github.com/JamesChung/omnistore@latest`

## How to use

### Set

```go
omnistore.Set[string]("myname", "James")
omnistore.Set[float64]("pi", 3.14)
```

### StringerSet

```go
myString := MyEnum("examplestringerstring")
omnistore.StringerSet[MyEnum](myString, "myvalue")

myInt := MyEnum("examplestringerint")
omnistore.StringerSet[MyEnum](myInt, 42)

myFloat := MyEnum("examplestringerfloat")
omnistore.StringerSet[MyEnum](myFloat, 3.14)
```

### Get

```go
name := omnistore.Get[string]("myname")
// name == "James"

pi := omnistore.Get[float64]("pi")
// pi == 3.14
```

### StringerGet

```go
myString := MyEnum("examplestringerstring")
omnistore.StringerSet[MyEnum](myString, "myvalue")
fmt.Println(omnistore.StringerGet[MyEnum, string](myString))

myInt := MyEnum("examplestringerint")
omnistore.StringerSet[MyEnum](myInt, 42)
fmt.Println(omnistore.StringerGet[MyEnum, int](myInt))

myFloat := MyEnum("examplestringerfloat")
omnistore.StringerSet[MyEnum](myFloat, 3.14)
fmt.Println(omnistore.StringerGet[MyEnum, float64](myFloat))
```
