# go_array

Go Slice operates like a JavaScript Array

# How to use?

You have to declare an Array object:
```cassandraql
    var sliceData = []int{1, 2, 3}
    array, err := Array(sliceData)
    if err != nil {
        fmt.Errorf("%v", err.Error())
    }
    str := array.ToString()
    fmt.Println(str) // 1,2,3
```

# Methods

- GetData: Return interface array data, you can use assertions, for example:
```cassandraql
    array.GetData().([]int)
```
- Len: Return array length
- ForEach
- Concat
- CopyWithin ........

#### There are detailed examples in the code