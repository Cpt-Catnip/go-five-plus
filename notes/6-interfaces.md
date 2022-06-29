# 53. Purpose of Interfaces
- gonna start by looking at code we already wrote and discuss how interfaces could improve it
- recall:
  - every value has a type
  - every function specifies the type of its args
- What if we have multiple functions with identical logic but operates on different types? Think
```go
func addString(s1 string, s2 string) string {
  return s1 + s2
}

func addInts(i1 int, i2 int) int {
  return i1 + i2
}

func addFloats(f1 float64, f2 float64) float64 {
  return f1 + f2
}
```
- and so on and so forth
- Take our shuffle deck function, for example
```go
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
```
- This function could easily applie to _any_ slice
- this is _one of_ the problems interfaces aim to solve
- Let's write a trivial program that will obviously benefit from interfaces, then refactor it
- image two chat bots, `englishBot` and `spanishBot`
- assume bot have `getGreeting` and `printGreeting` functions
  - `func (eb englishBot) getGreeting() string` and `func printGreeting(eb englishBot)`
  - `func (sb spanishBot) getGreeting() string` and `func printGreeting(sb spanishBot)`
- real quick this also reveals a potential assumption about receiver functions: use a receiver function when interacting with the underlying data in the struct (or whatever)
- we're assuming `getGreeting` probably has very diff implementations but `printGreeing` probably are mostly identical

# 54. Problems Without Interfaces
- Leaving structs empty since we only really care about the methods
- Quick note about unused vars
  - if we don't intend to use the instance that a receiver is called on, we can omit the variable name 
  - `func (s string) fn()` vs `func (string) fn()`
  - seems like an analog sort of for static methods but IDK if instances can share static vars or if you can call methods on structs as opposed to only instances
- Go doesn't support function overloading :( but...
  - we _can_ however make two _methods_ with the same name so long as they have different receiver chains (I think that's the proper verbiage)
```go
package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating an spanish greeting
	return "Hola!"
}
```
- look at all this GROSS repeated code

# 55. Interfaces in Practice
- with interface implemented
```go
package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating an spanish greeting
	return "Hola!"
}
```
- Okay first let me try to decipher this
  - we defined an interface that describes a thing `bot`
  - we say that a `bot` is anything that has a method called `getGreeting` that accepts no parameters and returns a `string`
  - we have not specified a receiver type
  - below, we create a function `printGreeting` that accepts this `bot` thing, which, by definition, mush have some method called `getGreeting`
  - so long as the thing we pass in has a method `getGreeting` that matches the specified signature, compiler will be happy
  - Think of an interface like a square hole. Many shapes can fit through a square hole but they _all_ must have some orientation whose 2D projection is a square with the same dimensions as the hole.
  - the `bot` interface is the square hole
  - `englishBot` and `spanishBot` are the two "things" that can fit into the hole
  - nothing has the type `bot`; `bot` is a lense with which to look at variables. Polarized lenses.
- Do interfaces only have methods in them? Can you also use properties to define an interface?

# 56. Rules of Interfaces
gonna eat lunch after this
- receiver functions are _associated_ with some type
- you can also specify parameter types and multiple output types in the interface def
- can list multiple functions
- note no commas, similar to struct def
```go
type bot interface {
  getGreeting(string, int) (string, error)
  getBotVersion() float64
  respondToUser(user) string
}
```
- from [Go By Example](https://gobyexample.com/interfaces), it seems like you only define methods in interfaces. So an interface can be though of something that describes the _behavior_ of some struct
- Hmm lots of different _kinds_ of types
- There are two kinds of types in Go
  - Concrete type: something you can make a value out of. Includes custom types and structs
  - Interface type: is an interface and can't be used to make a value

# 57. Extra Interface Notes
- interfaces are __not__ generic types
  - there is not type that matches all types (like `any` in TS)
- interfaces are "implicit" - we don't have to specify which structs match some interface
- interfaces are a contract to help up manage type
  - I don't know why this needs to be said, but interfaces don't do any tests or validation
- interfaces are tough (allegedly)
- okay cool I guess

# 58. The HTTP Package
- Still much to learn, pupil
- Ooh there are _standard interfaces_
- Let's make a program
  - make an HTTP request to google and print the response
- Going to make use of the [net/http](https://pkg.go.dev/net/http@go1.18.3) package
  - has methods like `http.Get` (cool) [link](https://pkg.go.dev/net/http@go1.18.3#Get)
  - `func Get(url string) (resp *Response, err error)`
  - returns a _pointer_ to a response
```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}
```
- so this print statement isn't really what we want- not the body
- we're going to somehow take advantage of interfaces to get a more gooder log

# 59. Reading the Docs
- readin the docs
- the `resp` thing is a struct with a bunch of fields, one of them being `Body`
- `Body` has the type `io.ReadCloser`
- `io.ReadClose` is actually an interface! What?! Fields can have interface types?
- Furthermore, [the docs](https://pkg.go.dev/io#ReadCloser) for the `ReadClose` interface doesn't have functions in it...
```go
type ReadCloser interface {
	Reader
	Closer
}
```
- again, `Reader` is also an interface!
```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```
- this looks a lot more familiar
- Okay here's the flow of all that in diagram form (thank you lecturer)
  - not sure if this image will show up on gitbub. Might have to commit it in the repo :/

![IMAGE](quiver-image-url/1F1DAB0D8A34F4479979894867F009BF.jpg =932x344)

- Some unanswered questions
  - how can a field in a struct be an interface?
    - I guess that means we're saying "this field is some embedded struct (which we know is possible) but there are many structs that it can be"
  - what's up with that interface shorthand from `ReadCloser`?
    - My guess is that it's similar to the struct shorthand, so `Reader` is probably like a variable that's holding a function or interface or whatever

# 60. More Interface Syntax
- How can we use an interface inside of a struct?
  - yeah what I said
  - The `Body` field can have any value assigned to it so long as it fulfills the interface (contract)
  - essentially `Body` just needs methods `Read` and `Close` (also matching signature)
- What's up with the shorthand syntax?
  - `Reader` and `Closer` are interfaces
  - passing those in to the interface `ReadCloser` _combines_ them
  - meaning for something to be a `ReadCloser`, it has to satisfy _both_ `Reader` __and__ `Closer`
  - Cool!

# 61. Interface Review
- an interface is a contract // polarized lense
- omg that's actually it?
- stoopid

# 62. The Reader Interface
- mans loves repeating himself
- A reader is essentially something that accepts _something_ and outputs it as a byte slice (`[]byte`) or something common that can be passed along
- So it's like the interface is saying "hey I'm gonna need something at somepoint that will accept the output of one black box and spit it out as `[]byte`, otherwise I'm not gonna know what to do".
- mans never heard of editing

# 63. More on the Reader Interface
- rather all input data must have the reader interface so the code can convert it into something it can handle
- recall:
```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```
- There is something somewhere (us, perhaps) that wants to read the body of some response
- we need a way to "translate" the body into something we can understand, so we've sent out an edict that says "if you give us something, it has to have translation instructions"
  - that is, the thing being delivered to us mush subscribe to the `Reader` interface
- the thing that wants to read the data will pass in its own byte slice, the `Read` function will modify the slice and then we can see the data (b/c reference data)
- the int in the return represents how many bytes were read
- So really we're saying "hey here's a slice please fill it with that good good (data)"

# 64. Working with the Read Function
- okay how do we actually get the body now
- initialize the slice using `make`
- "make an empty thing with this type and this many empty spaces
- want to pass this slice into the read function
  - note: the read function is not set up to resize the slice, so we want to make sure it has enough space going in (why use a slice?)
```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// will hold body data
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
```
- wow what a strange design pattern
- Go has some helper functions for dealing with reader functions

# 65. The Writer Interface
- write code first, then break it down
- this is, like, some pretty advanced go stuff
```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}
```
- Mike's guess
  - this is like `sprintf`
  - take something, and put (copy) it somewhere else
  - in this case, we're copying it into the stdout, which is where the terminal reads from 
  - the second arg is probably an interface that implements `Reader`
- recall the `Reader` interface says describes a function that fills up some byte slice with data
- the `Writer` interface is like the opposite
  - instead of handing back the data, it goes and puts it somewhere else, like into an HTTP request or something
- `os.Stdout` probably implements the `Writer` interface

# 66. The io.Copy Function
- [writer interface](https://pkg.go.dev/io#Writer)
```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```
- so, as you can see, where the data gets put is not specified by the interface, that's up to the specific application
- in fact it looks just like the `Reader` interface
  - this harkens back to "an interface is not a test"
- [io.Copy](https://pkg.go.dev/io#Copy)
- `func Copy(dst Writer, src Reader) (written int64, err error)`
- Yep, the first arg must implement the writer interface and second implements reader
- `os.Stdout` has type `*File`, which implements the writer interface

# 67. The Implementation of io.Copy
- recall: [io.Copy docs](https://pkg.go.dev/io#Copy)
- implementation of copy
```go
func Copy(dst Writer, src Reader) (written int64, err error) {
	return copyBuffer(dst, src, nil)
}
```
- this implements `copyBuffer`
```go
func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy.
	// Avoids an allocation and a copy.
	if wt, ok := src.(WriterTo); ok {
		return wt.WriteTo(dst)
	}
	// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	if rt, ok := dst.(ReaderFrom); ok {
		return rt.ReadFrom(src)
	}
	if buf == nil {
		size := 32 * 1024
		if l, ok := src.(*LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					ew = errInvalidWrite
				}
			}
			written += int64(nw)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != EOF {
				err = er
			}
			break
		}
	}
	return written, err
}
```
- this func is making a byte slice with `make` and sending it off into the reader
- passes buffer into read
- then passes populated buffer into write!

# 68. A Custom Writer
- Oh wow lets implement our own writer O_O
- recall: [io.Writer](https://pkg.go.dev/io#Writer)
```go 
type Writer interface {
	Write(p []byte) (n int, err error)
}
```
- Sometimes I see named outputs and sometimes I just see the tyes in the annotation. Why?
- RECALL: nothing is stopping us from making a bad implementation of a `Writer`. The __only__ check that happends is for the correct function signature. Whatever happes inside the code is behind closed doors.
```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Just wrote this many bytes: %d\n", len(bs))
	return len(bs), nil
}
```

# Quiz 10: Test Your Knowledge: Interfaces
__Q1: When we say that interfaces can be satisfied implicitly, we mean that...__
- we don't have to specifically say that some type satisfies an interface

__Q2: To say that a type satisfies an interface means that...__
- the type implements all the funcs mentioned in the interface

__Q3: Does the `square` type satisfy the `shape` interface?__
```go
type shape interface {
    area() int
}
 
type square struct {
    sideLength int
}
 
func (s square) area() int {
    return s.sideLength * s.sideLength
}
 
func printArea(s shape) {
    fmt.Println(s.area())
}
```
- yes, because square implements an `area` function that returns an int

__Q4: Does the `rectangle` type satisfy the `shape` interface?__
```go
type shape interface {
    area() int
}
 
type square struct {
    sideLength int
}
 
type rectangle struct {
    height float64
    width float64
}
 
func (s square) area() int {
    return s.sideLength * s.sideLength
}
 
func (r rectangle) area() float64 {
    return r.height * r.width
}
 
func printArea(s shape) {
    fmt.Println(s.area())
}
```
- no, because rectangle's area returns a float

__Q5: `square` implements the `shape` interface but it always returns `10` instead of the square's true area. Will the `shape` interface do anything to catch this error?__
- No!

__Q6: Types that implement the Reader interface are generally used to...__
- Read information from an outside data source into our application

__Q7: ...__
- this question silly it's just driving home that adhering interfaces doesn't guarantee that the code is doing what it's supposed to

__Q8: Does the `File` type satisfy both the `Reader` and `Writer` interface?__
- yes :(
  - [type File](https://pkg.go.dev/os#File)
  - [func (*File) Read](https://pkg.go.dev/os#File.Read)
  - [func (*File) Write](https://pkg.go.dev/os#File.Write)

# Assignment 2: Interfaces
- Write a program that creates two custom stuct types called `triangle` and `square`
- The `square` type should be a struct with a field called `sideLength` of type `float64`
- The `triangle` type should be a struct with a field called `height` of type `float64` and a field called `base` of type `float64`
- Both types should have functions called `getArea` that returns the calculated area of the square or triangle

$$Area_{triangle} = 0.5 * base * height$$

$$Area_{square} = sideLength * sideLength$$

- Add a `shape` interface that defines a function called `printArea`. This function should calculate the area of the given shape and print it out to the terminal. Design the interface so that the `printArea` function can be called with either a triangle or a square

__Submission:__
```go
package main

import "fmt"

type shape interface {
  getArea() float64
}

type square struct {
  sideLength float64
}

type triangle struct {
  base   float64
  height float64
}

func main() {
  t := triangle{
    base:   4.0,
    height: 3.0,
  }

  s := square{sideLength: 3.5}

  printArea(s)
  printArea(t)
}

func printArea(s shape) {
  fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
  return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
  return 0.5 * t.base * t.height
}
```
- At first I made the `printArea` function a receiver function, but we can't do that since we can't create a value with the type of an interface and therefore can't call functions on them.

__Solution:__
- yeah the same

# Assignment 3: Hard Mode Interfaces
- Create a program that reads the contents of a text file then prints its contents to the terminal
- The file to open should be provided as an argument to the program when it is executed at the terminal. For example `go run main.go myfile.txt` should open up the `myfile.txt` file
- To read in the arguments provided to a program, you can reference the variable `os.Args`, which is a slice of type string
- To open a file, check out the documentation for the `Open` function in the `os` package - https://pkg.go.dev/os#Open
  - `func Open(name string) (*File, error)`
  - >Open opens the named file for reading. If successful, methods on the returned file can be used for reading; the associated file descriptor has mode O_RDONLY. If there is an error, it will be of type *PathError.
- What interfaces does the `File` type implement?
- If the `File` type implements the `Reader` interface, you might be able to reuse that `io.Copy` function!
  - Seems like a big hint...
- the first item in Args is the temporary file that Go compiles to run our code

__Submission:__
```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fn := os.Args[1]
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
```

__Solution:__
- same thing I did
- You can actually read in _any_ file
  - if you want print `main.go` you have to first build then run `./main main.go`
  - oh and also `go mod init`
