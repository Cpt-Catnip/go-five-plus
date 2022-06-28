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
  * __update__: VSCode is confused because I have the entore folder open, not just the cards module. Run `code ~/code/go-five-plus/cards` and there's no error
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
* Create and return a list of playing cards (deck)
* This will _not_ be a receiver function; receiver functions act on "instances" of a type but we don't have a deck to operate on yet
* Going to build this deck using a nested loop
  * have two slices, one for sutis and one for values
  * for each suit, iterate over each value
  * add "value of suit" to deck
* We're making these slices string slices since it's _just_ the suit or value and a card is a suit _and_ value.
* the `range` keyword creates an iterable (not sure go syntax here) that returns the index and value- what if we don't need the index?
* You can replace an output with `_` (underscore) to tell go that we're not using this return value. This also suppresses compiler errors for unused vars.
* Code now looks like

```go
// main.go
package main

func main() {
	cards := newDeck()

	cards.print()
}

// deck.go
package main

import "fmt"

// Create a new type of "deck"
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Fout"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
```

# 22. Slice Range Syntax
* Gonna make the `deal` function now
  * accept a deck and a number of cards to deal, return new deck of specified size
* Essentially transforming a deck into two decks
* If a deck has $N$ cards and we're dealing a hand with $x$ cards, we should expect
  * a deck with $x$ cards (the dealt hand)
  * a deck with $N - x$ cards (the remaining deck)
* To do this, we need to use some slice methods
* Slices are zero-indexed
* Slices use square bracket notation (`slc[n]`)
* We can use the range notation for getting values from a slice
* `slc[i:j]` returns the values on the half-open interval $[i, j)$
  * i.e. inclusive of `i` and exclusive of `j`
* ex:

```go
fruits := []string{"apple", "banana", "grape", "orange"}

fmt.Println(frutis[1])    // banana
fmt.Println(frutis[0:2])  // ["apple", "banana"]
```

* Like python, you can omit _either_ `i` or `j`
  * `slc[:j]`: means from the beginning and until, but not including, `j`
  * `slc[i:]`: means from `i` until the end, inclusive
* We can utilize this to easily split a deck into a hand and deck
  * `cards[:handSize]` and `cards[handSize:]`, respectively

# 23. Multiple Return Values
* Lets write!
* This function has arguments! Args need to be typed and you can't pass in args with types other than what's specified in the signature
* syntax is `func fnName(arg1 type1, arg2 type2) { ... }`
* Go supports returning multiple values
* Instead of just adding the type of the return before the `{`, put all the types in a parenthesis separated by commas
  * `func fn() (rt1, rt2) { ... }`
* Here we have _not_ made a receiver function. I think this is because we're not doing an operation _on_ a deck, rather we're creating two new decks
  * there's an implication here that receiver functions shouldn't create new things or modify the thing it's operating on. Keep an eye on that assumption, Mike.
* To capture both return values, use a comma (already saw this with `range`)

```go
// main.go
hand, remainingCards := deal(cards, 5)

// deck.go
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
```

* we're using the shorthand assignment operator for _both_ vars here. We can infer from what we say in `range` that the types don't have to be the same for this to work
  * we use `i, val := range slc` where `i` is _always_ an integer and `val` can be any arbitrary type

# Quiz 6: Test Your Knowledge: Multiple Return Values
__Q1: In the following code snippet, what will the value and type of `title` and `pages` be?
```go
func getBookInfo() (string, int) {
    return "War and Peace", 1000
}
 
title, pages := getBookInfo()
```
* `string` and `int` respectively
  * lol I picked the wrong answer; I don't think my meds have fully kicked in yet

__Q2: What will the following program log out?__
```go
package main
 
import "fmt"
 
func main() {
    color1, color2, color3 := colors()
 
    fmt.Println(color1, color2, color3)
}
 
func colors() (string, string, string) {
    return "red", "yellow", "blue"
}
```
* "red yellow blue`

__Q3: What will the following program log out?__
```go
package main
 
import "fmt"
 
func main() {
   c := color("Red")
 
   fmt.Println(c.describe("is an awesome color"))
}
 
type color string
 
func (c color) describe(description string) (string) {
   return string(c) + " " + description
}
```
* "Red is an awesome color"

__Q4: Which of the following best explains the `describe` function listed below?__
```go
package main
 
import "fmt"
 
func main() {
   c := color("Red")
 
   fmt.Println(c.describe("is an awesome color"))
}
 
type color string
 
func (c color) describe(description string) (string) {
   return string(c) + " " + description
}
```
* color receiver func, accepts string, returns string

__Q5: After calling "deal" and passing in "cards", does the list of strings that the "cards" variable point at change?  In other words, did we modify the 'cards' slice by calling 'deal'?__
```go
func main() {
    cards := newDeck()
 
    hand, remainingCards := deal(cards, 5)
 
    hand.print()
    remainingCards.print()
}
```
* Great question. Kind of gets at whether or not slice range returns new slices or references to the original slice. Like what's the first index of `remainingCards`, `0` or `5`?
* I think the answer it no, cards will be the same but we can modify `cards` by modifying `hand` or `remainingCards`
> Answer: You got it. We created two new references that point at subsections of the 'cards' slice. We never directly modified the slice that 'cards' is pointing at.
* I'm so smart (and have taken this part of the course before)

# 24. Byte Slices
* covering `saveToFile`; how do we save the deck to our harddrive?
* can use [ioutil.WriteFile](https://pkg.go.dev/io/ioutil@go1.18.3#WriteFile) to write to disk
  * Write to new or existing file
  * signature: `func WriteFile(filename string, data []byte, perm fs.FileMode) error`
* the data has a type `[]byte`, a byte slice.
* the third arg is permisions
* How do we convert our deck instance (`[]string`) into a byte slice?
* A byte slice is a slice of ascii character codes (decimal), essentially
  * represents a string in a computer-friendly way

# 25. Deck to String
* _actually_ covering `saveToFile`, lol
* We can use "type conversion" to switch between types
* syntax is something like `desiredType(varWithOtherType)`
  * e.g. `[]byte("Hi there!")` becomes `[72 105 32 116 104 101 114 101 33]`
* want `deck -> []string -> string -> []byte`
* Gonna start with a function that turns a deck into a string
* note we _are_ making this a receiver function. I guess it's not explicitly modifying the original `deck` so maybe it's kosher? Lecturer says we'll talk more about when to use a recevier later on

# 26. Joining a Slice of Strings
* to go from `[]string` to `string`, we can joing all strings in the slice on `,` using [strings.Join](https://pkg.go.dev/strings@go1.18.3#Join)
* `func Join(elems []string, sep string) string`
  * need to import `strings` to use this
* When using multiple imports, wrap all imports in parenthesis and separate each package with a newline

* Where we're at

```go
// main.go
package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
}

// deck.go
package main

import (
	"fmt"
	"strings"
)

// Create a new type of "deck"
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Fout"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
```

* yields 

```bash
Ace of Spades,Two of Spades,Three of Spades,Fout of Spades,Ace of Diamonds,Two of Diamonds,Three of Diamonds,Fout of Diamonds,Ace of Hearts,Two of Hearts,Three of Hearts,Fout of Hearts,Ace of Clubs,Two of Clubs,Three of Clubs,Fout of Clubs
```

# 27. Saving Data to the Hard Drive
* NOW we can write to a file
* Will utilize the `error` type
* Wowww when autocompleting `d.toString` in the function call, VSCode automatically cast it to `[]byte`
* The permissions argument is in case the file doesn't already exist

```go
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
```

* note that ioutil gets imported as `io/ioutil` in the import statement
* Looks like `ioutil.WriteFile` writes to current directory if no path is specified

# 28. Reading From the Hard Drive
* gonna use [ioutil.ReadFile](https://pkg.go.dev/io/ioutil@go1.18.3#ReadFile)
* `func ReadFile(filename string) ([]byte, error)`
  * oop! this returns an error! Gonna learn fun error handling
* What _is_ this `error` thing?
* If everything goes according to plan, `error` will be `nil`
  * `nil` is just a type of data in go that means "nothing", so "there is no error"
* otherwise, `error` will be non-nil
* to check if there was an error, we want to check that `error` is NOT `nil`
* How to handle an error in go depends strongly on _how_ the code failed
* To end execution, use [os.Exit](https://pkg.go.dev/os@go1.18.3#Exit)
* `func Exit(code int)`
  * any non-zero exit code indicates an error
* function so far

```go
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
```

# 29. Error Handling
* this `if err != nil` is a __very__ common patter (common complaint)
* Now we want `[]byte -> string -> []string -> deck`
* to splig a string, use [strings.Split](https://pkg.go.dev/strings@go1.18.3#Split)
* `func Split(s string, sep string) []string`
* lecturer keeps driving home that the reason we can do something like `deck(s)`, where `s` has type `[]string` is because "a deck is really just a slice of string"
  * makes me think converting to a custom type is usually more hands on and _maybe_ there's a way to define how that conversion happens so you can still type `customType(someVar)`

```go
// main.go
func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
}

// deck.go
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}
```

* if we try to open a file called `my`, we get the following error message
```bash
$ go run main.go deck.go
Error: open my: no such file or directory
exit status 1
```

# 30. Shuffling a Deck
* how to be shuffling deck?
* Want to randomize the order of all the cards in the deck
  * same elements and number of elements
* logic we're going to use

```plaintext
for each idx, card in cards:
  generate random number btwn 1 and len(cards) - 1
  swap cards[idx] and cards[randomIdx]
```

* Pretty sure this isn't a true shuffling but oh well, not necessary to be rigorous here
* how to be gen random number? Use [Intn](https://pkg.go.dev/math/rand@go1.18.3#Intn)
  * `func Intn(n int) int`
  * generates a pseudo-random integer on the half-open interval $[0, n)$
* making this a receiver func
* using `range` but _not_ using the second return value. In this case we don't need to replace it with `_`, we can just say `for i := range d`
* use `Intn` like `rand.Intn`
  * why not `math.rand.Intn`? What's the logic here? Got imported as "math/rand"
* get length of a slice (and probably array?) with `len(slc)`
* We can use the fun one-line swap like in JS just with out the surrounding brackets- nice
```go
func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
```
* This will actually return the same randomization _each time_ the code is run

# 31. Random Number Generation
* the random number generator depends on a _seed value_, which has a default value
  * need to generate a random seed to generate a random number
* random number generator depends on the `Rand` type, which has the seed
* We need to make our own `Rand` object, which depends on a `Source` type
  * this accepts an `int64`, which we need to generate randomly
* okay `rand.New(s Source)` actually returns something that looks like `*Rand`
* This relates to pointers but we haven't gotten there yet so hold tight, but now that it's not returning `Rand`
* For our random `int64`, we're gonna use nanoseconds since epoch, which __really should__ always be different
  * use [time.Now()](https://pkg.go.dev/time@go1.18.3#Now)
  * `func Now() Time`
* Shuffling is always different now!
* All caught up!!

```go
// main.go
package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// deck.go
package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of "deck"
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
```

# 32. Testing With Go

# 33. Writing Useful Tests

# 34. Asserting Elements in a Slice

# 35. Testing File IO

# 36. Project Review

# Assignment 1: Even and Odd

