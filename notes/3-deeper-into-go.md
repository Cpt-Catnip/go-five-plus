# 13. Poject Overview
* Going to simulate playing around with playing cards, like for poker and stuff
* Package will have a number of functions
  * `newDeck`: Create a list of playing cards. Essentially an array of strings.
  * `print`: Log out the contents of a deck of cards.
  * `shuffle`: Shuffles all the cards in a deck.
  * `deal`: Create a 'hand' of cards.
  * `saveToFile`: Save a list of cards to a file on the local machine.
  * `newDeckFromFile`: Load a list of cards from the local machine.

# 14. New Project Folder
```go
// cards/main.go
package main

func main() {

}
```

# 15. Variable Declarations
* Will also cover some useful VSCode features
* Such as... VSCode automatically added the line `import "fmt"` for us when we wrote `fmt.Println`
* Declare variable like

```go
var card string = "Ace of Spades"
```

* lecturer has an error on this line essentially saying you don't need to assert the type here because it can be inferred from the assignment
  * I don't see that error
* What's going on here?
  * `var`: tell go we're declaring a variable
  * `card`: name of variable
  * `string`: asserting the type- only strings can be assigned to `card`
  * `"Ace of Spades"`: value being assigned to variable
* Go is a "statically typed" language unlike JS or python, which are "dynamically typed" languages
* some "basic" go types (not exhaustive)
  * bool
  * string
  * int
  * float64
* Since we're immediately assigning a string to the var `card` we can use a shorthand syntax that tells go to infer the type

```go
card := "Ace of Spades"
```

* So here we're missing `var` and `string` but does the exact same thing as the original syntax
* The `:=` syntax only gets used when we're declaring a _new variable_

```go
package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	card = "Five of Diamonds"

	fmt.Println(card)
}
```

# Quiz 2: Test Your Knowledge: Variable Assignment
__Q1: Is the following a valid way of initializing and assigning a value to a varialbe?__
```go
var bookTitle string = "Harry Potter"
```
* yes

__Q2: Is the following a valid way of initializing and assigning a value to a variable:__
```go
fruitCount := 5
```
* yes

__Q3: After running the following code, Go will assume that the variable `quizQuestionCount` is of what type?__
```go
quizQuestionCount := 10
```
* integer

__Q4: Will the following code compile? WHy or why not?__
```go
paperColor := "Green"
paperColor := "Blue"
```
* No, because the var is being initialized twice

__Q5: Are the two lines following ways of initializing the variable `pi` equivalent?__
```go
pi := 3.14
var pi float64 = 3.14
```
* yes

__Q6: Is the following code valid?__
```go
package main
 
import "fmt"
 
deckSize := 20
 
func main() {
  fmt.Println(deckSize)
}
```
* No
  * This is an example of something that wasn't covered in the slides!
  * I think it's invalid because code can't be run outside of the _main_ function since that's the main driver of the code itself

__Q7: We _might_ be able to initialize a variable and then later assign a variable to it. Is the following code valid?__
```go
package main
 
import "fmt"
 
func main() {
  var deckSize int
  deckSize = 52
  fmt.Println(deckSize)
}
```
* yes

__Q8: Is the following code valid?__
```go
package main
 
import "fmt"
 
var deckSize int
 
func main() {
  deckSize = 50
  fmt.Println(deckSize)
}
```
* yes (got this wrong)
  * we can initialize a var outside of a function, we just can't assign it a value (okay...)

__Q9: Is the following code valid? Why or why not?___
```go
package main
 
import "fmt"
 
func main() {
  deckSize = 52
  fmt.Println(deckSize)
}
```
* No, `deckSize` is never initialized

# 16. Functions and Return Types
* Want a new function the returns the value of a card we're trying to create

```go
func newCard() string {
	return "Five of Diamonds"
}
```

* the `string` after `newCard()` is the _return type_ of the function
* without a return type, go will encounter a compilation error essentially saying that the compiler was expecting a function with no return type but instead got one that return a string
* We can still use the shorthand declaration for variable when assigning the return value of a function
* go typically has very good error messages

# Quiz 3: Test Your Knowledge: Functions
__Q1: If a function returns a value, do we have to annotate, or mark, the function with the type of value that is being returned?__
* yes

__Q2: The Go compiler is presenting an error message about the following function. What should we do to fix the error?__
```go
func getSize() {
    return "30 meters"
}
```
* `func getSize string() {}`

__Q3: Is the following function valid?__
```go
func estimatePi() float64 {
    return 3.14
}
```
* yes

__Q4: Is the following code valid? Why or why not?__
```go
package main
 
import "fmt"
 
func main() {
    fmt.Println(getTitle())
}
 
func getTitle() string {
    return "Harry Potter"
}
```
* Yes, it's valid

__Q5: Imagine we have two files in the same folder. Will the following code successfully compile?__
```go
// main.go
package main
 
func main() {
    printState()
}

// state.go
package main
 
import "fmt"
 
func printState() {
    fmt.Println("California")
}
```
* yes, because both files are part of the same package

# 17. Slices and For Loops
* Arrays in go!
* Two basic structures for handling lists of records
  * `array`: Fixed length list of things
  * `slice`: An array that can grow or shrink
    * This is more like what I'm familiar with in JS (I think)
* Silces and arrays must be defined with a data type
* this means that all records inside must be of the __same type__
  * this differs from JS, where something like `["hello", 5, false, "world"]` is a valid array
* declaring slices

```go
cards := []string{"Ace of Diamonds", newCard()}
```

* breaking it down
  * `[]`: signifies that we're making a slice
  * `string`: specifies the type of the records
  * `{ el_0, el_1, ..., el_n }`: the elements to put in the slice
* to add a new element

```go
cards = append(cards, "Six of Spades")
```

* note that `append` **returns a new slice**
* to iterate over a slice, use a for loop

```go
for index, card := range cards {
	fmt.Println(i, card)
}
```

* breaking it down (I actually love these)
  * `range cards`: make an iterable out of given slice (cards)
  * `:=`: using this to pass the output of `range` to...
  * `index`/`card`: loop variables to be updated on each iteration
  * `{ ... }`: code to run on each loop
* we lose the loop vars on each iteration, which is why we use the `:=` on each step
* Code up until now

```go
package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
```

# Quiz 4: Test Your Knowledge: Slices and For Loops
__Q1: Which of the following represents a slice where each element in it is of type int?__
* `[]int{}`

__Q2: Is the following code valid?__
```go
colors := []strings{"Red", "Yellow", "Blue"}
```
*  No, because `strings` isn't a valid type

__Q3: How do we iterate through each element in a slice and print out its value?__
```go
colors := []string{"red", "yellow", "blue"}
for index, color := range colors {
  fmt.Println(index, color)
}
```

__Q4: Would the following code compile successfully?__
```go
for index, card := range cards {
  fmt.Println(card)
}
```
* no, because we're not using `index`

__Q5: Can a slice have __both__ values of type `int` and of type `string` in it?__
* no

# 18. OO Approach vs Go Approach
* To be clear, go is __not__ an object oriented language
* If we _were_ using an object oriented language, we might have a deck "class" with various props and methods (print, shuffle, etc.)
* In go, we'll have a deck "type" and functions with deck as a "receiver"
  * `type deck []string`
  * receiver functions are like analogs for instance methods- functions attached to types
* Desired proj structure
  * `main.go`
  * `deck.go`
  * `deck_test.go`

# 19. Custom Type Declarations
* Gonna define the `deck` type!
* This won't differ in any significant way from a slice of strings, but it allows us to create receiver functions for it
* Once creating the new type, we can replace the string slice examples prior into a deck!

```go
// deck.go
package main

// Create a new type of "deck"
// which is a slice of strings
type deck []string

// main.go
package main

import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
```

* I'm getting an error when trying to initialize `cards`
  * undeclared name: deck compiler (UndeclaredName)
  * It might be resolved in a bit- hold tight
* To run the script now, we have to tell go about both files, otherwise `main.go` won't know about the deck type

```bash
$ go run main.go deck.go
```

* Now that we have a deck type, we can make a function that _belongs_ to that type that does our loop/print logic for us

```go
// deck.go
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// main.go
func main() {
  // ...
  deck.print()
}
```

* the `(d deck)` is the "receiver" part of this function

# 20. Receiver Functions
* receiver goes _before_ the function name
* the receiver consists of a name (`d`) and a type (`deck`)
* the name is just essentially a stand in for "the thing with `type` that will be calling this function"
  * e.g. `d` represents `cards` in `main.go`
* all variables with the type specified in the receiver def will automatically get access to the receiver function
* `d` is the actual copy of the desk we're working with- it's a reference to the "instance", kind of list `this`
* The convention in go is to use a one or two letter abbrev to match the type of the receiver
  * `d deck`
  * `f fart`
  * `th thursday`
* nothing actually forcing you to do this
* like classes, instances, and instance methods, but those terms are forbidden in go
  * go doesn't acutally have a class declaration syntax, just the type declaration

# Quiz 5: Test Your Knowledge: Functions with Receivers
__Q1: What would the following code print out?__
```go
package main
 
import "fmt"
 
type book string
 
func (b book) printTitle() {
    fmt.Println(b)
}
 
func main() {
    var b book = "Harry Potter"
    b.printTitle()
}
```
* "Harry Potter"

__Q2: Complete the sentence! By creating a new type with a function that has a receiver, we...__
* Are adding a 'method' to any value of that type

__Q3: In the following snippet, what does the variable `ls` represent?__
```go
type laptopSize float
 
func (ls laptopSize) getSizeOfLaptop() {
    return ls
}
```
* A value of the type `laptopSize`

__Q4: Is the following code valid?__

```go
type laptopSize float64
 
func (this laptopSize) getSizeOfLaptop() laptopSize {
    return this
}
```
* Yes, but it breaks convention by using `this`

# 21. Creating a New Deck

# 22. Slice Range Syntax

# 23. Multiple Return Values

# Quiz 6: Test Your Knowledge: Multiple Return Values

# 24. Byte Slices

# 25. Deck to String

# 26. Joining a Slice of Strings

# 27. Saving Data to the Hard Drive

# 28. Reading From the Hard Drive

# 29. Error Handling

# 30. Shuffling a Deck

# 31. Random Number Generation

# 32. Testing With Go

# 33. Writing Useful Tests

# 34. Asserting Elements in a Slice

# 35. Testing File IO

# 36. Project Review

# Assignment 1: Even and Odd

