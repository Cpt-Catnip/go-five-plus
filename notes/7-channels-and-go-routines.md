# 69. Website Status Checker
- again, approaching idea by writing a program that will benefir from concurrency and then refactor
- program will iterate over list of websites and check that we can get an OK response from them
  - I think you can see where this is going
- omg we did so little
```go
package main

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, l := range links {
		
	}
}
```

# 70. Printing Site Status
- we don't care about the response itself, just whether or not we get an error
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, l := range links {
		checkLink(l)
	}
}

func checkLink(l string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Printf("%s might be down!\n", l)
		return
	}

	fmt.Printf("%s is up!\n", l)
}
```
- So we're making a request for each link
- status messages are appearing in the order that we list them in the links slice

# 71. Serial Link Checking
- before moving on to the next link/loop iteration, we have to wait to get the status for the current link
  - i.e. we're checking statuses one at a time
- this is bad because
  - this program would take too much time to execute if the list gets too long
  - the result of one check doesn't depend on the result of any other; there's no reason why one status should have to wait for the other statuses
- CaN wE dO tHiS iS pArAlLeL?
- I dare this man to make longer videos

# 72. Go Routines
- When we run a program, we create a "Go Routine" that runs code line-by-line
- Code that takes time to run is called a _blocking call_
  - nothing can run until this line finishes and we have no control over how long that will be
- we can run any arbitrary line in their own go routing by prepending it with the `go` keyword
- so `checkLink(l)` can be changed to `go checkLink(l)` and now that check will happen in its own go routine
  - the http request is still blocking within its routine, but now the main routine can notice that and move on to the next item it its routine

# 73. Theory of Go Routines
- behind the scenes there's a "go scheduler"
- go will always try to use one CPU core, so one routine is being run at a given time
- the schedule monitors the code being run in a given routine
- if the scheduler notices that a routine is finished or has made a blocking call, it will move on to the next routine
- _only one routine is being run at a time_
- We can ask Go to use more than one core, now we can run multiple routines at a time
- The scheduler will assign different routines to different cores
  - concurrent vs parallel
- concurrency is the one-core example
  - concurrent code is one that can spin up multiple threads (routines) and run each one independently - we're able to stop/wait for one thread and pick up another
- parallelism is when different threads are literally being run at the same time (in english you might say that they're being run concurrently <_<)
  - don't get hung up on semantics - language is meant to express ideas and different industries establish different conventions
  - if two routines can be run _concurrently_, then they are independent of each other and can also be run in parallel.
- when we run a program from the command line, we have one main routine
  - routines spawned from within the code are _child routines_
  - child routines do not have the same level of priority as the main routine
- we'll see a relevant bug soon

# 74. Channels
- oh, __we only use the `go` keyword in front of function calls__
- that implies some honus is on us, the engineers, to write code that can be run concurrently
- new for-loop
```go
for _, l := range links {
	go checkLink(l)
}
```
- when we run this code we see that _nothing gets logged_! Gasp!
- This is because the main routine finishes before any of the chile routines!
- Program completion is limited by the _main routine_ not any of the _child routines_
  - so once the scheduler gets to the bottom of main, we're done
- We need to tell the main routine to wait for the child routines
- __Channels__ are used to communicate between go routines
  - we can use channels to tell the main channel to wait for the child routines
- channles are _typed_, meaning you specify _what kind of data_ gets piped through the channel- the channel itself doesn't have a type

# 75. Channel Implementation
- okay so how do we use these so called _"channels"_ to block the main routine from finishing?
- to make a string channel, we do
```go
c := make(chan string)
```
- do we _have_ to use the `make` keyword?
  - [it would](https://go.dev/tour/concurrency/2) [appear so](https://gobyexample.com/channels)
  - ok nbd
- channels are treated like any other value in our app
  - channels are a reference type
- for a func to use a channel, it needs to be passe as an argument
- There are a number of channel operations used for sending data
  - `channel <- 5`: send the value `5` in the this channel
  - `myNumber <- channel`: _wait_ for a value to be sent into the channel. when we get one, assign the value to `myNumber`
  - `fmt.Println(<- channel)`: wait for a value to be sent into the channel. When we get one, log it out immediately
    - same operation as above but value isn't being captured
  - seems like the communication always goes _right to left_
- now with channel implementation
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
	}

	fmt.Println(<-c)
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Printf("%s might be down!\n", l)
		c <- "Might be down I think"
		return
	}

	fmt.Printf("%s is up!\n", l)
	c <- "Yep it's up"
}
```
- channels have been places very intentionally
- Okay _another_ fun bug!
- We are only getting _one response_ each time we execute the program
  - lecturer always get's back google, which is the first link, but I sometimes get SO and sometimes google - probably depends on whichever response we get back first

# 76. Blocking Channels
- first, Mike's educated guess/attempt at remembering from previous readings
  - I think the `<- c` syntax only means "wait until you get _something_ and then move on"
  - even if more things get piped into the channel, the main routine stops paying attention (notice me, senpai)
- `<- c` is a blocking call just like `http.Get` is
- once the channel receives a message, `<- c` sees is, wakes up the main channel, executes that line, and finally _moves on_
- if _moving on_ means reaching the end of `func main()`, the main routine exits
- We need a way to say "wait for any additional messages and _only move on if there are none left_"
- We can copy the blocking call to receive a message for each item in the slice but that's __very_ unwieldy
- this is also very prone to error since if we have _more_ receives than sends, the program will just hang up waiting for a message that isn't coming

# 77. Receiving Messages
- Just gonna do a for-loop that iterates for each element in the link slice
- except we don't actually want to iterate _over_ the slice, so we'll use a more familiar syntax ;)
```go
for i := 0; i < len(links); i++ {
	fmt.Println(<-c)
}
```
- still, imo, not a good implementation. This only works in the very specific scenario where we know exactly how many channel pipes we should be waiting for
- I'm pretty sure I remember reading a way to loop (while) until there are no more messages
- each iteration is still run serially
- man is advocating for this approach

# 78. Repeating Routines

a quick thought from mike: I like this approach to concurrency much more than JS's. In JS, you need to just know which function are async and which aren't. The only thing that's clear is when you're waiting for an async function. In Go, everything is run _synchronously_ by default and you explicitly say "GO run this function asynchronously/in its own go routine". You still also get told when you are awaiting something "asynchronous" with the `<- c` syntax. Much nicer.

- Okay now what if we want to repeatedly check the status of a given website?
  - ha ha now we can't just loop over the length of the slice
- now, when a given check is done, we can pipe back the link and spin up another status check routine back on the other end of the channel
```go
func main() {
  // ...
  
  for {
		go checkLink(<-c, c)
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Printf("%s might be down!\n", l)
		c <- l
		return
	}

	fmt.Printf("%s is up!\n", l)
	c <- l
}
```
- this new `for { ... }` syntax is just a loop that runs forever, in this case limited in rate by receving a message in the channel `c`
- okay let's put a delay between each status check

# 79. Alternative Loop Syntax
- gonna make an opinionated stylistic change to this loop
```go
for l := range c{
	go checkLink(l, c)
}
```
- this is exactly equivalent to the previous loop but is a Go shortcut for doing this very common operation
  - I need to look into `range` more deeply. I bet it's argument is an interface and what I was calling "iterables" (python syntax) is something that satisfies that interface
- `range c` will, behind the scenes, wait for the message
- Seems a _little_ silly to have these two loops right next to each other, but I guess the first one can be though of as "spinning up" the communication channel
```go
for _, l := range links {
	go checkLink(l, c)
}

for l := range c {
	go checkLink(l, c)
}
```

# 80. Sleeping a Routine
- Note that _as soon_ as the the channel receives a message, our code is designed to immediately spin up another status check routine- we want to wait between status checks!
  - implementing the pause is straightforward but _where_ to put it won't be ;)
- We're going to use [time.Sleep](https://pkg.go.dev/time@go1.18.3#Sleep)
  - `func Sleep(d Duration)`
  - looks like `d` is in ~~milliseconds~~ nanoseconds and that __Go has unit conversions built in!!!__ (`time.Millisecond`)
- Sleep puases the __current goroutine__
- You can do something like `time.Pause(5 * time.Second)` to pause for 5 seconds
- If we want pauses between status checks _for a given link_, that goroutine needs to pause and not the main routine

# 81. Function Literals
- we're going to use something called a _function literal_
- similar to an anonymous function in JS
- We're going to use the func literal on the receving end of the channel and
  - first sleep for a duration
  - then call `checkLink`
- the function literal here gets called like an IFFE
```go
for l := range c {
	go func() {
		time.Sleep(5 * time.Second)
		checkLink(l, c)
	}()
}
```
- Go doesn't like us using the loop variable in the func literal
  - `loop variable l captured by func literal loopclosure`
  - guess: `l` is a reference value and its value might change by the time the goroutine uses it

# 82. Channels Gotcha!
- the issue above is a warning and not an actual error, so code will still compile and run... let's see what happens
```bash
$ go run main.go
http://stackoverflow.com is up!
http://google.com is up!
http://amazon.com is up!
http://facebook.com is up!
http://golang.org is up!
http://golang.org is up!
http://golang.org is up!
http://golang.org is up!
http://golang.org is up!
http://golang.org is up!
http://golang.org is up!
http://golang.org is up!
http://golang.org is up!
http://golang.org is up!
http://golang.org is up!
^Csignal: interrupt
```
- We get the first five off the bat but then keep getting repeats. Weird!
- Seems like `l` is always getting the golang link
- in the main routine, the var `l` is ponting at some location in memory that holds a link string
- the child routine is referencing `l` in the parent scope, so it's pointing at the same location in memory
- note that at each loop iteration, `l` changes its value but lives in the same location in memory
- so by the time the sleep is finished, `l` will be something other than what it started out as when we first spun up that goroutine
  - `l` is `google.com`
  - spin up goroutine
  - sleep for 5 seconds
  - `l` is `golang.org`
  - call `checkLink(l, c)`
- or maybe we even make the http request with the right link, but the time _that_ finishes, we pipe back the wrong link
- the subroutine is using a variable that's managed by a different routine, the main routine
- we _never_ want to reference the same variable in two separate routines
  - that's bad for concurrency!!
- we can take advantage of Go's pass-by-value nature
- pass the link into the function literal as an argument
```go
for l := range c {
	go func(link string) {
		time.Sleep(5 * time.Second)
		checkLink(link, c)
	}(l)
}
```
- we can still call the parameter `l` but that would be confusing for the engineer

```bash
$ go run main.go
http://stackoverflow.com is up!
http://google.com is up!
http://amazon.com is up!
http://facebook.com is up!
http://golang.org is up!
http://stackoverflow.com is up!
http://google.com is up!
http://amazon.com is up!
http://facebook.com is up!
http://golang.org is up!
http://stackoverflow.com is up!
http://google.com is up!
http://amazon.com is up!
http://facebook.com is up!
http://golang.org is up!
^Csignal: interrupt
```

# Quiz 11: Channels and Go Routines
__Q1: Which of the following best describes what a go routine is?__
* A separate line of code execution that can be used to handle blocking code

__Q2: What's the purpose of a channel?__
* for communicating between go routines

__Q3: Are there any issues with this program?__
```go
package main
 
import (
 "fmt"
)
 
func main() {
 greeting := "Hi There!"
 
 go (func() {
     fmt.Println(greeting) 
 })()
}
```
* the greeting variable is referenced in two separate goroutines
* the program will exit before printing

__Q4: Is there any issue with the following code?__
```go
package main
 
func main() {
 c := make(chan string)
 c <- []byte("Hi there!")
}
```
* chan wants `string` but we're feeding it `[]byte` (picky eater)

__Q5: Is there any issue with the following code?__
```go
package main
 
func main() {
     c := make(chan string)
     c <- "Hi there!"
}
```
* syntax is okay but the program will get stuck waiting for something to receive the message in the channel
  * okay this wasn't explicitly mentioned in the lectures. I guess the program won't "finish" if that routine has a non-empty channel?
* I'm so confused as to how to even make this work
* The comments reveal the true answer! It's not that channels _can_ communicate between routines, it's that they __must__ communicate between routines.
* That is, some other goroutine needs to receive the message to unblock the main routine

__Q6: Ignoring whether or not the program will exit correctly, are the following two code snippets equivalent?__
__snippet 1__
```go
package main
 
import "fmt"
 
func main() {
 c := make(chan string)
 for i := 0; i < 4; i++ {
     go printString("Hello there!", c)
 }
 
 for s := range c {
     fmt.Println(s)
 }
}
 
func printString(s string, c chan string) {
 fmt.Println(s)
 c <- "Done printing." 
}
```
__snippet 2__
```go
package main
 
import "fmt"
 
func main() {
 c := make(chan string)
 
 for i := 0; i < 4; i++ {
     go printString("Hello there!", c)
 }
 
 for {
     fmt.Println(<- c)
 }
}
 
func printString(s string, c chan string) {
 fmt.Println(s)
 c <- "Done printing." 
}
```
* they are the same!

