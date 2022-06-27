# 7. Boring Ol' Hello World
* This is the literal title of the subsection don't @ me
* Small program with big potential

```go
package main

import "fmt"

func main() {
	fmt.Println("Hi there!")
}
```

* okay not getting syntax highlighting. Something is amiss.
  * Restarting VSCode again got rid of the `[Unsupported]` tag in the top and triggered the helper tool installation. All is well in the digi-world.

# 8. Five Important Questions
* How do we run the code in our project?
  * That was easy! (see below)
  * the `go` command exposes many cmds in the CLI
    * `go run`: Compiles and executes one or two (or more?) files
      * does not save compiled code
    * `go build`: Compiles a bunch of go source code files
      * differs from `run` in that it does __not__ execute code
    * `go fmt`: Formats all the code in each file in the current directory
    * `go install`: Compiles and "installs" a package
    * `go get`: Downloads the raw source code of someone else's package
    * `go test`: Runs any tests associated with the current project

```bash
$ go run main.go
> Hi there!
```

* What does `package main` mean?
* What does `import "fmt"` mean?
* What's that `func` thing?
* How is the `main.go` file organized?

# 9. Go Packages
> What does `package main` mean?

* A package is a collection of common source code files
* Each file in a package _must_ declare which package it belongs to in the first line (e.g. `package main`)
* But why call it `main`?
* Two types of packages in Go
  * executable: generates a file that we can run 
    * will output an executable file when you run `go build`
  * reusable: like dependencies or libraries. Helpers.
    * hmm...
* _Only_ packages called `main` will result in an executable type package
* changing the name of the package in `main.go` and rtying to run in via CLI results in an error
  * you can still build though

```bash
$ go run main.go
> package command-line-arguments is not a main package
```

* An executable package also must have a function inside of it called `main`
  * that's what the `func main` bit was all about

# 10. Import Statements
> What does `import "fmt"` mean?

* `import` is used to give our package to code from another package
* `fmt` is part of the _standard library_, which comes with all installations of go
* packages have no default to standard library packages so you need to explicitly import them
  * This is unlike JavaScript, where you can run something like `Math.random()` without any overhead
* We can also import packages made by other engineers not already included
* Documentation for all standard library (stdlib) packages can be found at [pkg.go.dev/std](https://pkg.go.dev/std)

# 11. File Organization
> What's that `func` thing?

```go
func main() {
  // ...
}
```

* `func` is the function keyword!
  * same as `function` in JavaScript
* parts of a function
  * 
  * `func`: Tells go we're about to declare a function
  * `main`: Sets the name of the function
  * `()`: List of arguments to pass the function
  * `{ ... }`: Function body. Calling the function runs this code

> How is the `main.go` file organized?

* Always the same exact pattern no matter what file we create
* In order:
  * package declaration (`package main`)
  * imports (`import "fmt"`)
  * Code body/login (`func main() {...}`)

# Quiz 1: Test Your Knowledge: Packages
_editor's note: sometimes there is information tested in the quizes that weren't covered in the slides, so it's worth paying close attention to them_

__Q1: What is the purpose of a package in Go?__
To group together code with a similar purpose

__Q2: What is the special name we use for a package to tell Go that we want it to be turned into a file that can be executed?__
main

__Q3: The one requirement of packages named "main" is that we...__
Define a function named "main", which is ran automatically when the program runs

__Q4: Why do we use "import" statements?__
To give our package access to code written in another package

# 12. How to Access Course Diagrams
All of the diagrams in this course can be downloaded and marked up by you!  Here's how:

* Go to https://github.com/StephenGrider/GoCasts/tree/master/diagrams
* Open the folder containing the set of diagrams you want to edit
* Click on the ‘.xml’ file
* Click the ‘raw’ button
* Copy the URL
* Go to https://www.draw.io/
* On the ‘Save Diagrams To…’ window click ‘Decide later’ at the bottom
* Click ‘File’ -> ‘Import From’ -> ‘URL’
* Paste the link to the XML file
* Tada!
