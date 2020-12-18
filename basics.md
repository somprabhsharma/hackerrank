- data types
```
bool, string, rune, byte, float32, float64, complex64, complex128
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
```

- var initialization
```
c, python, java := true, false, "no!"
```

- Arrays
```
var a [10]int
primes := [6]int{2, 3, 5, 7, 11, 13}
s := []int{2, 3, 5, 7, 11, 13}
a := make([]int, 5)
a[low : high], includes low and excludes high, its a slice
board := [][]string{
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
}
```

Slices are like references to arrays
A slice does not store any data, it just describes a section of an underlying array.
Changing the elements of a slice modifies the corresponding elements of its underlying array.
Other slices that share the same underlying array will see those changes.


- type conversion
```
i := 42
f := float64(i)
```


- normal loop
```
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

- for range
```
for i, v := range pow {
	fmt.Printf("2**%d = %d\n", i, v)
}
for _, value := range pow
for i := range pow
```

- While loop
```
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```

- infinite for loop
```
func main() {
	for {
	}
}
```

- special if case, expression before condition
```
if v := math.Pow(x, n); v < lim {
	return v
}
```

- switch case
```
today := time.Now().Weekday()
switch time.Saturday {
case today + 0:
	fmt.Println("Today.")
case today + 1:
	fmt.Println("Tomorrow.")
default:
	fmt.Println("Too far away.")
}
```

- Maps
```
m := make(map[string]int)
m[key] = elem
elem = m[key]
delete(m, key)
elem, ok = m[key]

```

- Strings package
```
p("Contains:  ", strings.Contains("test", "es"))
p("Count:     ", strings.Count("test", "t"))
p("HasPrefix: ", strings.HasPrefix("test", "te"))
p("HasSuffix: ", strings.HasSuffix("test", "st"))
p("Index:     ", strings.Index("test", "e"))
p("Join:      ", strings.Join([]string{"a", "b"}, "-"))
p("Repeat:    ", strings.Repeat("a", 5))
p("Replace:   ", strings.Replace("foo", "o", "0", -1))
p("Replace:   ", strings.Replace("foo", "o", "0", 1))
p("Split:     ", strings.Split("a-b-c-d-e", "-"))
p("ToLower:   ", strings.ToLower("TEST"))
p("ToUpper:   ", strings.ToUpper("test"))
```

- Sort package
```
sort.Strings(strs)
s := sort.IntsAreSorted(ints)

type byLength []string

func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}
```

- strconv Package
```
package main

import (
    "fmt"
    "strconv"
)

func main() {

    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

    k, _ := strconv.Atoi("135")
    fmt.Println(k)

    _, e := strconv.Atoi("wat")
    fmt.Println(e)
}
```

- time package
```
package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

    now := time.Now()
    p(now)

    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    p(then.Weekday())

    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    diff := now.Sub(then)
    p(diff)

    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    p(then.Add(diff))
    p(then.Add(-diff))
}
```
