# Go
- Next Gen language
- Compiled = Go tool can run file directly, without VM
- Executables are different for OS
- Go is a cross-platform, open source programming language
- Go can be used to create high-performance applications
- Go is a fast, statically typed, compiled language known for its simplicity and efficiency
- Go was developed at Google by Robert Griesemer, Rob Pike, and Ken Thompson in 2007
- Go's syntax is similar to C++

- What can i build with Golang    
-- already in production  
-- system apps to web apps - Cloud  

- Don't bring baggage  
-- Yeah, I did that earlier.
-- Similarity with lots of languages - C, Java, Pascal

- Object oriented  
-- Yes and No  
-- what you see on the screen is the code

- Missing things  
-- Is it really missing ?  
-- No try catch  
-- lexer does a lot of work

- Memomry Management
-- Memory alocation and deallocation happens automatically  
-- Garbage Collections happens automatically  
-- Out of scope or nil  
-- https://pkg.go.dev/runtime   - low level information available


new()  
-- Allocate memory but no INIT  
-- You will get a memory address  
-- Zeroed storage

make()  
-- Allocate memory and INIT  
-- You will get a memory address  
-- Non - zeroed storage (you can put data)


## What is Go Used For?
- Web development (server-side)
- Developing network-based programs
- Developing cross-platform enterprise applications
- Cloud-native development


## Why Use Go?
- Go is fun and easy to learn
- Go has fast run time and compilation time
- Go supports concurrency
- Go has memory management
- Go works on different platforms (Windows, Mac, Linux, Raspberry Pi, etc.)


## Go Syntax
A Go file consists of the following parts:
- Package declaration
- Import packages
- Functions
- Statements and expressions


## Types
- case sensitive; almost
- variable type should be known in advance
- Everything is a type.
- static langauge


### Primitive types
- String
- Bool
- Integer ==>  uint8   unit64  int8    int64   uintptr     (aliases are involved too)
- Floating  ==>    float32     float64
- Complex

### Advanced types
- Arrays
- Slices
- Maps
- Structs
- Pointers

Almost everything 
- Functions
- Mutex
- Channels....and much more 


## Array
- Array is very less used in the array



<!-- GOOS=windows : The term 'GOOS=windows' is not recognized as the name of 
a cmdlet, function, script file, or operable program. Check the spelling 
of the name, or if a path was included, verify that the path is correct 
and try again.
At line:1 char:1
+ GOOS="windows"  go build
+ ~~~~~~~~~~~~~~
    + CategoryInfo          : ObjectNotFound: (GOOS=windows:String) [],  
   CommandNotFoundException
    + FullyQualifiedErrorId : CommandNotFoundException -->