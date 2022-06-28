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

# 47. Gotchas With Pointers

# Quiz 7: Test Your Knowledge: Pointers

# 48. Reference vs Value Types

# Quiz 8: Test Your Knowledge: Value vs Reference Types
