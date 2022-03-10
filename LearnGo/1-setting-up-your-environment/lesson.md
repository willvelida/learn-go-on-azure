# Setting up our Go Development Environment

I've been wanting to learn Go for a while now. I've tried and failed to keep my learning consistent, so I've decided to force myself to keep on top of my learning through a couple of avenues:

1. This Blog I've set up.
2. [This GitHub repo](https://github.com/willvelida/learn-go-on-azure).

In that GitHub repo, I'll be adding content on both Golang itself and building applications with Go that use Azure. I'll be building this repo up over time, but if you have a particular sample or scenario that you'd like me to create, please feel free to raise an [issue](https://github.com/willvelida/learn-go-on-azure/issues). 

This article is aimed at developers who have never used Go before and want to set up their environment to start playing around with the langugage. In this article, I will cover:

- Why Go?
- Installing Go on our machine
- Verifying Go has been installed correctly
- Installing the Go extension on VS Code
- Building and Running a basic Hello World app!

One thing to note before we start is that I'll be doing my development on a Windows machine, so commands and formatting will be tailored towards the Windows ecosystem. If you have a Mac, I will do my best to try and include instructions for Mac as I go along, but I may miss out from time to time. If you need guidance for Mac systems, please reach out to me on [Twitter](https://twitter.com/willvelida) and I'll do my best to help out 😊

## Why Go?

Go is an open source programming language designed for building simple, fast and reliable softwae. It's expressive, clean, efficient and concise. It compiles quickly to machine code and also has garbage collection.

For cloud applications, Go has a huge advantage in terms of compiling to a static binary and performance. We can use Go to build super-fast server side applications that are scalable and resilient.

Go rules the Cloud Native world (another motivation for me personally to learn Go). Kubernetes, Docker and Prometheus are all written in Go.

## Installing Go on our machine

Before we can write any Go, we'll need to install the tools to do so. We can install these by heading to the [Go website](https://go.dev/dl/). We can install it from a package (*.msi* for Windows, *.pkg* for Mac and *.tar.gz* for Linux) or we can compile it from source.

I went the package route, which isn't as cool I know, but that will automatically install Go in the correct locations, remove any old installations of Go you may have and more importantly, put the Go binary in the default executable path. 

For windows users, you can install Go using Chocolatey like so:

```powershell
choco install golang
```

For Mac users, you can use Homebrew:

```bash
brew install golang
```

## Verifying Go has been installed correctly

We can validate that Go has been installed correctly by opening up a command prompt or terminal and running the following command:

```bash
go version
```

If everything was setup correctly, we should see the following output:

```bash
go version go1.14.4 windows/amd64
```

What this command tells me is that I'm running Go version 1.14.4 on Windows. Your Go version may be different and if you're using a Mac or Linux, you won't see Windows in your output.

Using Go, we will need to install 3rd party tools, which can be installed using the ```go install``` command. By default, these tools will be installed in the *$HOME/go* location. The source code for these tools will be situated in the *$HOME/go/src* directory and the compiled binaries will be stored in the *$HOME/go/bin* directory.

It's a good idea to put our *$HOME/go/bin* directory on our executable path, which makes it clear where our Go workspace is situated and makes it easier to run 3rd party tools. On Windows, we can do this by running the following:

```powershell
setx GOPATH %USERPROFILE%\go
setx path "%path%;%USERPROFILE%\bin"
```

## Installing the Go extension on VS Code

For these tutorials, I'll be using Visual Studio Code for my developer environment. It's free and has a fantastic extension that we can use to write Go code. 

If you haven't got VS Code, you can download it [here](https://code.visualstudio.com/download).

Once you've downloaded VS Code, open it up and install the Go extension by clicking on the **Extensions** icon in the left hand sidebar, and then typing in *Go* in the text box. Pick the Go extension developed by the *Go Team at Google* and click **Install**.

![Go VS Code extension](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/b9wu2x0rpv8u7cjdgru8.png)

Now that this has been installed, let's take everything for a spin and create our first Hello World app in Go!

## Building and Running a basic Hello World app!

In VS Code, create a new folder and in that folder, create a new file called ```main.go```. In that Go file, write the following:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

Let's break this file down:

- Every Go program is made up of packages. We declare our package by writing ```package main```. Programs in Go start running in the man package.
- We then import the *fmt* library in the line ```import "fmt"```. We can import multiple packages into our Go code if we use multiple packages. The *fmt* package implements formatted I/O functions that we can use in our program, without having to write these functions ourselves. We can leverage libraries in our Go code by importing them.
- We then create our entry function by writing ```func main()```. This is our main entry point of our program.
- We then print out our "Hello, World!" statement by writing ```fmt.Println("Hello, World!")```. The ```Println()``` function is a inbuilt function of the *fmt* package which we can use to format our input and write it to standard output.

There are a couple of ways we can run this program. Let's start by running the file like so:

```bash
go run main.go
```

With this command, we'll see the following output:

```bash
Hello, World!
```

What the ```go run``` command does is compile our Go code into a binary, executes it and then deletes the binary. It builds the binary into a temporary directory. This is useful when we want to test out small Go programs or treat our Go code like a script.

If we want to build a binary and keep it, we can run the following.

```bash
go build main.go
```

What this command does is create an executable for us called *main.exe* (or just *main* if you're on Mac) in the diectory where our Go file is. To run this, we just call the executable and it should print out ```Hello, World!``` again. If we want to give our executable a name, we can define that name like so:

```bash
go build -o hello_world main.go
```

## Conclusion

In this article, I showed you how you can install Go onto your machine and how you can setup Visual Studio Code to work with Go. We also wrote our first bit of Go code, built a binary from it and executed it! While this program's not going to make you the next tech unicorn, this was an important first step!

I've kept this first step simple just to get our developer environment going. Once we have this setup, we can start to explore Go a little bit more. If you're confident, I'd recommend checking out the following resources if you want to explore Go on your own:

- [Go by Example](https://gobyexample.com/)
- [The Go Playground](https://go.dev/play/)

If you have any questions or had trouble getting this all set up, feel free to reach out to me on twitter [@willvelida](https://twitter.com/willvelida) and I'll try and help out the best I can.

In my next post, we'll take another small step into the world of Go and discuss variables and types.

Until next time, Happy coding! 🤓🖥️