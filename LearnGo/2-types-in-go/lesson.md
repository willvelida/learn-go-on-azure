So the Golang learning adventure continues! We have our [developer environment](https://www.willvelida.com/posts/setting-up-dev-environment-golang/) setup so we're ready to dive into more Golang code.

In this post, I'll be talking about types and variables in Go. If you've used other languages, you may notice as we go through this that Go has some differences in how it uses types compared to other languages. In this article, we'll be covering:

- What types are available to us?
- Assigning values in Go
- Constants in Go

I've built a very simple Go file that works with different types in Go which you can view [here](https://github.com/willvelida/learn-go-on-azure/blob/main/LearnGo/2-types-in-go/main.go). Please feel free to clone it, copy/paste it and play around with it.

## What types are available to us?

Go has many built-in types that we would find in other languages. These are:

- Booleans (true or false values)
- Integers (numbers like 1, 2, 3)
- Floats (like 1.0, 2.0, 3.0)
- Strings (like "Hello!")

Of course there's a bit more detail to cover here. Starting with number types, Go has support for 12 different types of numerics that can be grouped into three categories. Go provides support for the following integer types:

| Type | Range of Values |
| ---- | --------------- |
| int8 | -128 to 127 |
| int16 | –32768 to 32767 |
| int32 | –2147483648 to 2147483647 |
| int64 | –9223372036854775808 to 9223372036854775807 |
| uint8 | 0 to 255 |
| uint16 | 0 to 65536 |
| uint32 | 0 to 4294967295 |
| uint64 | 0 to 18446744073709551615 |

The zero value (or default value) for all integer types is 0. 

With integers, we have the usual operators that we'd find in other languages: ```+```, ```*```, ```/```, ```-``` and ```%``` for modulus operations. We can also combine any of these operators with ```=``` to modify the value of a variable. For example:

```go
var myNumber int = 10
myNumber *= 2
// myNumber is now 20
```

We can also perform comparison operations on integers using ```==```, ```!=```, ```>```, ```>=```, ```<```,```<=```.

Now let's turn our attention to floats. Floats in Go are similar to floats in other programming languages and Go supports two different types of floats: ```float32``` and ```float64```. Floating point literals have a default type of ```float64```. To use floats, we can do so like this:

```go
var number3 float64 = 5.0
number4 := 3.0
```

Like integers, floats have a default value of 0 unless initialized. We can use comparison operators to work with floats, but the exception here is that we would generally avoid using ```===``` or ```!=``` for floats. Float values have a very **wide** range, so we might not get the result we expect when using these operators on floats.

Finally, let's take a look at Strings. The zero value for strings is just an empty string. Go supports Unicode, so we can put any Unicode character we like into our strings. They are immutable in Go, so we can reassign the value, we just can't change the value.

In a future post, I'll talk about how we can work with strings in Go, but for now, here's a small snippet on how we can declare a string, assign a value to it and then print it out.

```go
var greeting string = "This is a string!"
fmt.Println(greeting)
// This prints out "This is a string!"
```

Like integers and floats, we can perform operations on our strings. We can compare them using ```==``` or ```!=``` and we can order them using ```>, >=, <, <=```. We can concantenate strings together like so:

```go
fmt.Println("Hello, " + "World")
// Prints out Hello, World
```

## Assigning values in Go

In Go, we explicitly declare our variables which are used by the compiler to check the type-correctness of function calls. We can use the ```var``` keyword to declare a single variable like so:

```go
var greeting string = "Hello!"
```

we can even declare multiple variables at once like so:

```go
var number1, number2 int = 10, 20
```

We can also declare variables without specifying the type. Go will infer what type the variable is when we initialize it. So in the following snippet, the ```isTypeSpecified``` variable is infered to be of boolean type:

```go
var isTypeSpecified = false
```

We also dont have to write ```var``` everytime we declare a variable. We can use shorthand syntax like so:

```go
number1 := 5
```

The great thing about using ```:=``` notation is that we can use this to assign values to existing variables. However, we can only use this notation if we are declaring our variables inside a function. If we do it at the package level (outside our functions), this will throw an error.

We can also declare our variables using something called a **declaration list** like so:

```go
var (
    number1 int
    number2 = 20
    number3 int = 30
    number4, number5 = 40, 50
    number6, number7 int
)
```

## Constants in Go

Go also has support for constants. These are constant values that do not change. We can declare constants like so

```go
constant statement string = "This is a constant statement"
```

Constants can appear anywhere we put our ```var statements```. They can also be declared outside our functions, so for example, the following Go code is valid:

```go
const number1 int = 10

func main() {
    const number2 int = 20
}
```

The important thing to note about constants is that they are **immutable**. This means that once we have initialized them with a value we can't change the value later in our program (Making it constant...). So for example, the following code would throw an error:

```go
const myNumber int = 10
myNumber = 10 + 1
```

## Conclusion

In this post, we took a very small and simple step into types in Go. The purpose of this was to keep it simple and not overwhelm ourselves with information that we won't need right now. As we dive into more complex Go, we'll need to dig a little deeper and understand how to keep our Go code efficient by ensuring we declare our variables with the correct type. But for now, we know the basics to start working with Go.

If you have any questions, feel free to reach out to me on twitter [@willvelida](https://twitter.com/willvelida) and I'll try and help out the best I can.

In my next post, we'll start building Go programs using If, if/else and else/if statements.

Until next time, Happy coding! 🤓🖥️