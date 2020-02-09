### Arrays

A **fixed** data structure that accepts a set of values of the **same data type**. Arrays are pased by value, meaning that its instance is copied when is passed to a function. Therefore, within the function, we cannot modify the existing values of the initial data structure

```
words := [4]string{"the","quick","brown","fox"}
```

### Slice

A **dynamically-sized** data structure built on top of an array. A slice is a reference to an array and it's declared similar to an array without the size of the slice. That means that a **slice is passed by reference** to a function. A slice contains three components:

- **Pointer:** the pointer is used to point to the first element of the array that is accessible through the slice. Here, it is not necessary that the pointed element is the first element of the array.
- **Length:** the length is the total number of elements present in the array.
- **Capacity:** the capacity represents the maximum size upto which it can expand.

```
words := []string{"the","quick","brown","fox"}
```

#### Methods
- append
- copy

### Maps

It's a reference to a hash table containung key-value pairs. The keys within a map do not **have a specific order** and it does not allow duplicate keys. It only accepts key-value pairs of the same data type and has a **constant time complexity O(k)**.

```
	dayMonths := make(map[string]int)
	dayMonths["January"] = 31
	dayMonths["February"] = 28
	dayMonths["March"] = 31
	dayMonths["April"] = 30
	dayMonths["May"] = 31
	dayMonths["June"] = 30
	dayMonths["July"] = 31
	dayMonths["August"] = 31
	dayMonths["September"] = 30
	dayMonths["October"] = 31
	dayMonths["November"] = 30
	dayMonths["December"] = 31

  // or

  dayMonths = map[string]int{
    "January": 31,
    "February": 28,
    ...
  }
```

### Method

- delete

## Goroutines and channels

Goroutines are *functions* or *methods* in go that run concurrently in different **threads**. Goroutines can communicate with each other using a **channel**, which is a specific data structure that allow bi-directional communication between then goroutines. Any type of communication needs to happen by sending and receiving data of the same data type. Goroutines and channels are handled by go scheduler which is part of go language. 
In Go language, a channel is created using chan keyword and it can only transfer data of the same type, different types of data are not allowed to transport from the same channel.

```
// creates a channel
var ch = make(chan int)

// sends data to a channel
ch <- 1

// receives data from a channel
a := <- ch
```


## Useful links

[geekforgeeks.org](https://www.geeksforgeeks.org/slices-in-golang/)