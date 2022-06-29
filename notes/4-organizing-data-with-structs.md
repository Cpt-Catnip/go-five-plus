# 37. Structs in Go
- Going back to cards project! Want to talk about something that was maybe a little awkward
- Each card was just a plain string, which would have made it hard to easily pull out a card's suit or value
  - string manipulation is possible but clunky
- What if we could use a special data structure to describe a card?
- Introducing... `struct`!!!
- short for structure
- A collection of related data props
- A card struct could have
  - suit
  - value
- The "struct" is the "object" itself. The thing that describes structs is the "struct field definion"

# 38. Defining Structs
- gonna make a "person" struct
- a `struct` is an extension of a `type`
- no separators or anything, just new lines
- This code is _outside_ of `func main`

```go
type person struct {
	firstName string
	lastName  string
}
```

# 39. Declaring Structs
- There are many ways to define (create?) a struct
```go
// method 1
alex := person{"Alex", "Anderson"}
```
- Here we're kind of treating the type keyword like a function and passing the arguments into the curly braces
  - lecturer doesn't like this approach because what if someone swaps the order of the prop defs?
    - what if someone does that in a function? Like that's such a stupid argument.
```go
// method 2
alex := person{firstName: "Alex", lastName: "Anderson"}
```
- Pretty much the same but we're passing more as key:value pairs- this is similar to how python function calls work
- There's a third way that we'll look at later
- printing a struct looks like `{Alex Anderson}`

# 40. Updating Struct Values
- The third way! Part one:
```go
var alex person
```
- Since we're only initializing a person variable and not adding any fields, `firstName` and `lastName` get the __zero values__ for their given types
- Each type gets their own zero value

| type | zero value |
|------|------------|
| `string` | `""` |
| `int` | `0` |
| `float` | `0` |
| `bool` | `false` |

- printing alex now yields `{ }`
- There's a special way to print structs to show their field names
```go
fmt.Printf("%+v", alex)
```
- `%v` is the verb but adding the `+` to make it `%+v` says "if `v` is a struct, print the field names"
- We can update struct fields using dot-notation

```go
package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Printf("%+v\n", alex)
}
```

# 41. Embedding Structs
- struct-ception
- what if a person has contact information that itself has a number of properties?
- we can make a contact info struct and have one of the fields in the person struct house that struct (yo dawg)
- Struct fields can have _any_ data type we please!
```go
package main

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	
	fmt.Printf("%+v\n", jim
}
```
- go requires trailing commas

# 42. Structs with Receiver Functions
- One additional way to embed a struct
- Similar to JS, if we want to make a field in the struct def with the same name as the name of the struct, you can simply just put down the name of the struct type, for example
```go
type person struct {
	firstName string
	lastName  string
	contactInfo
}
```
- now the field on person is called `contactInfo` instead of `contact` but still has type `contactInfo`
  - still have to specify it twice when we're assigning a value to the field
- This will apparently become useful in the future
- Since structs are just type, you can make receiver functions with them
```go
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
```
- Okay _one last thing_. Consider this function
```go
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
```
- this function will _not_ update the `firstName` field on `jim`. More on that (and pointers) next.

# 43. Pass By Value
- When we create a new person struct, go will create data and store it somewhere in local memory (ram)
- `jim` now points to that location in memory
- go is a "pass by value" language
- this means that when data is passed to a function, the data gets passed and then assigned to a new variable, putting into a __new location in memory__
  - this is __unlike javascript__ which will sometimes pass by value and sometimes pass by reference; go always passes by value
- although the receiver function is operating on an "instance", it gets sent a copy of the variable
- When we update the `firstName` of the person in the receiver function, we're updating the `firstName` of the _copy_ that the receiver gets, not the variable that the function was called on

# 44. Structs with Pointers
- update `updateName` to use pointers and have the desired outcome
```go
func main() {
  // ...
  jimPointer := &jim
  jimPointer.updateName("jimmy")
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
```
- Okay so let's go over what's going on here... in the next section

# 45. Pointer Operations
- what are `&` and `*` here?
- `&variable` creates a pointer (memory address) of `variable`
  - think about it like a function that accepts some variable and returns that variables memory address
- `*pointer` returns the data that lives at the address referenced by `pointer`
- So let's understand our new code
  - `jimPointer` is `jim`'s memor address
  - The receiver type is now `*person`, which means "a pointer that points to a `person`"
  - this may seem confusing since we literally just said that `*pointer` returns the data at memory, but now as a type def it means "specifically a pointer and not data"
    - weird!
    - `*` in front of a _type_ means "a type that is itself a pointer to the given type" - it's a type description
    - `*` in front of a pointer means "the data that lives at this address - it's an operation
    - it's up to the engineer to notice which scenario we're in!
  - now when we do `*pointerToPerson` we're getting _the original `person` that me made_ and updating that
    - I think the paren is so we're cleare about "use `*` on the pointer" and not `pointer.field`, which would probably throw an error or something
- To recap, there are two types of variables
  - pointers: memory address where data lives
  - variables: the data
- Turn `address` into `value` with `*address`
- Turn `value` into `address` with `&value`

# 46. Pointer Shortcut
- recall we wrote
```go
jimPointer := &jim
jimPointer.updateName("jimmy")
```
- Go provides a shortcut for doing this very kind of operation
- You actually don't have to call pointer receiver functions on pointers. You can simply call it on the data type itself and go will work out the details
- The following code is perfectly valid
```go
func main() {
  //...
  jim.updateName("jimmy")
  jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
  (*pointerToPerson).firstName = newFirstName
}
```
- the receiver function _must_ be for a pointer though if you want the data addess to be passed to the function

# 47. Gotchas With Pointers
- recall: go is a pass by value language
- doing this in the go playground (maybe because memory leaks?) 
```go
package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "there", "how", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
```
- okay this code _will_ update the first element in `mySlice` even though we're not passing a pointer. More on that later...
  - oh I think this is because a slice is just a reference to an array

# Quiz 7: Test Your Knowledge: Pointers
__Q1: Whenever you pass an integer, float, string, or struct into a function, what does Go do with that argument?__
- It creates a copy of each argument, and these copies are used inside of the function
  - this is a hint at the next lecture!

__Q2: What will the following program print out?__
```go
package main
import "fmt"
 
func main() {
   name := "Bill"
 
   fmt.Println(&name)
}
```
- The memory address that "Bill" is stored at

__Q3: What is the `&` operator used for?__
- Turning a value into a pointer

__Q4: When you see a `*` operator in front of a pointer, what will it turn the pointer into?__
- A value

__Q5: When the following program runs, the `fmt.Println` call reports that the `latitude` field of `newYork` is still equal to `40.73`. What changes should we make to get the `latitude` of `newYork` to update to `41.0`?__
```go
package main
import "fmt"
 
type location struct {
 longitude float64
 latitude float64
}
 
func main() {
 newYork := location{
   latitude: 40.73,
   longitude: -73.93,
 }
 
 newYork.changeLatitude()
 
 fmt.Println(newYork)
}
 
func (lo location) changeLatitude() {
 lo.latitude = 41.0
}
```
- change the receiver type of the function to `*location` and replace `lo` with `(*lo` in the function body

__Q6: in the `changeLatitude` function, what is the `*location` in the receiver list (after the word 'func') communicating to us?__
```go
package main
 
import "fmt"
 
type location struct {
 longitude float64
 latitude float64
}
 
func main() {
 newYork := location{
 latitude: 40.73,
 longitude: -73.93,
 }
 
 newYork.changeLatitude()
 
 fmt.Println(newYork)
}
 
func (lo *location) changeLatitude() {
 (*lo).latitude = 41.0
}

```
- it specifies the type of the receiver that the function expects

__Q7: What will the `Println` function in the `main` function print out?__
```go
package main
 
import "fmt"
 
func main() {
    name := "Bill"
    updateValue(name)
    fmt.Println(name)
}
 
func updateValue(n string) {
    n = "Alex"
}
```
- "Bill"

__Q8: What will happed when this code is executed?__
```go
package main
 
import "fmt"
 
type location struct {
 longitude float64
 latitude float64
}
 
func main() {
 newYork := location{
   latitude: 40.73,
   longitude: -73.93,
 }
 
 newYork.changeLatitude()
 
 fmt.Println(newYork)
}
 
func (lo *location) changeLatitude() {
 (*lo).latitude = 41.0
}
```
- The program uses the shortcut!

__Q9: Here's a tricky one! What will the following program print out?__
```go
package main
 
import "fmt"
 
func main() {
    name := "Bill"
 
    fmt.Println(*&name)
}
```
- The string "Bill"
  - tricky my ass

# 48. Reference vs Value Types
- recall:
```go
package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "there", "how", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
```
- this snipped updated the original slice `mySlice` even though we didn't pass the function `updateSlice` a pointer
- So what really _is_ the difference between an array and a struct
- Array:
  - primitive data structure
  - can't be resized
  - rarely used directly
- Slices:
  - can grow and shrink
  - used 99% of the time for lists of elements
- I'm sure there's more than this
- a slice is actually two diff data structures
  - the "slice" is a pointer to the head of an array, a capacity number, and a length
  - the "array" is where the data is actually currently being held!
- So this makes sense! A copy of the slice _was_ being passed to the function but the actualy data inside was itself a reference to a memory address
  - similar to how objects are passed in JS
- a "slice" itself is not the list of elements

| Value Types | Reference Types |
|-------------|-----------------|
| int         | slices          |
| float       | maps            |
| string      | channels        |
| bool        | pointers        |
| structs     | functions       |

- for "value types", you need to pass pointers
- for "reference types", you can pass the data structure without worry (depending on what you want the outcome to be)
- again, like JS, go will _always_ copy the varialbe passed to a function, it's just a matter of whether or not the variable holds data or a memory address

# Quiz 8: Test Your Knowledge: Value vs Reference Types
__Q1: When we create a slice, Go will automatically create which two data structures?__
- An array and a structure of metadata (full answer omitted for brevity)

__Q2: is the `mySlice` avlue being copied before being passed into the function?__
```go
package main
 
import "fmt"
 
func main() {
 mySlice := []string{"Hi", "There", "how", "are", "you?"}
 
 updateSlice(mySlice)
 
 fmt.Println(mySlice)
}
 
func updateSlice(s []string) {
 s[0] = "Bye"
}
```
- yes

__Q3: With 'value types' in Go, do we have to worry about pointers if we want to pass a value to a function and modify the original value inside the function?__
- Yes

__Q4: With 'reference types' in Go, Do we have to worry about pointers if we want to pass a value to a function and modify the original value inside the function?__
- no

__Q5: Is a slice a 'value type' or a 'reference type'?__
- reference type

__Q6: Here's a tough one. Will the memory address printed by both `Println` calls be the same?__
```go
package main
 
import "fmt"
 
func main() {
 name := "bill"
 
 namePointer := &name
 
 fmt.Println(&namePointer)
 printPointer(namePointer)
}
 
func printPointer(namePointer *string) {
 fmt.Println(&namePointer)
}
```
- the addresses will be different
  - each address will themselves point to a common memory address but the two pointers storing that address are stored in different locations
