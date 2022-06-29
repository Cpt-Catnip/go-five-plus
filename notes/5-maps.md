# 49. What's a Map?
- A collection of key/value pairs
  - like object in JS
- Similar to structs but not?
- Both the keys and values of a map are statically typed
  - all keys in map must be the same type
  - all values in a map must be the same type
  - key does not have to be same type as value
- structs are allowed to have fields with different types and I think all keys are strings? Let's find out
- More than one way to declare a map (classic)
```go
colors := map[string]string{
  "red":  "#ff0000",
  "gree": "#00ff00",
}
```
- breaking it down
  - `map`: tells go we're making a map
    - similar to how `[]` tells go we're making a slice
  - `[string]`: __keys__ are of type string
  - `string`: __values__ are of type string
  - `{ ... }`: data goes in here
    - keys and values are separated by a `:`
    - all key value pairs must have a trailing `,`

# 50. Manipulating Maps
- Two other ways of declaring a map
```go
var colors map[string]string

colors := make(map[string]string)
```
- both these new approaches create empty maps
- the second one uses the function `make`, which I hope we get into more later
  - can't find docs on the `make` function
- You can add data to maps using the familiar bracket notation
```go
colors["white"] = "#ffffff"
```
- can you do one of those two initializatino methods and assign data at the same time?
- No, you have to do something gross like
```go
var colors map[string]string = map[string]string{
  "red":  "#ff0000",
  "green": "#00ff00",
}
```
- gross but makes sense
  - left side: make a variable that will store a map that has string keys and string values
  - right side: creating a map itself
- maps do __not__ get the "dot" notation because the keys are typed
- Can delete keys using the `delete` function
  - `delete(map, key)`

# 51. Iterating Over Maps
- My guess is with `range` and an array is sort of just a map with `int` keys (not exactly)
- Yep using range
```go
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for %s is %s\n", color, hex)
	}
}
```
- why isn't this a receiver function???

# 52. Differences Between Maps and Structs
- Map
  - all keys must be the same type
  - all values must be the same type
  - keys are indexed - we can iterate over them
  - use to represent a collection of related props
  - don't need to know all the keys at compile time
    - we can add/remove items over the course of code execution
    - if we know there will always be N keys, maybe use a struct
  - reference type
- Struct
  - values van be of different type
  - keys don't support indexing
    - can't iterate over the fields in a struct (dang)
  - you need to know all fields at compile time
  - use to represent a "thing" with a lot of different properties
  - value type
- structs tend to be used more than maps

# Quiz 9: Test Your Knowledge: Maps
__Q1: Can some of the keys in a single map be of type `int` and others of type `string`?__
- No (!!!!)

__Q2: Can some of the _values_ in a single map be of type `int` and others of type `string`?__
- NO

__Q3: What would the print statement log out?__
```go
package main
import "fmt"
 
func main() {
 m := map[string]string{
   "dog": "bark",
 }
 
 changeMap(m)
 
 fmt.Println(m)
}
 
func changeMap(m map[string]string) {
 m["cat"] = "purr"
}
```
- `map[dog: bark cat: purr]`

__Q4: What would happen if we tried to run the following program?__
```go
package main
import "fmt"
 
func main() {
 m := map[string]string{
   "dog": "bark",
   "cat": "purr",
 }
 
 for key, value := range m {
   fmt.Println(value)
 }
}
```
- Compiler would throw an error because `key` is unused
