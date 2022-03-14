In Go, we can use conditional statements to execute code based on different conditions. We can also use ```if```, ```else if``` and ```if else``` statements to execute code provided if a condition is true or false. If/else statements are commonplace in programming, so if you're already a seasoned programmer you won't spend long here. But as always, we'll take our time and make sure we understand fully what's going on.

In this article we'll cover:

- The ```if``` statement
- The ```if else``` statement
- The ```else if``` statement
- Conditional statements
- ```if``` statements inside ```if``` statements (Ifception).

As always, I have a simple Go file that shows how if statements and conditions work on my [GitHub](https://github.com/willvelida/learn-go-on-azure/blob/main/LearnGo/3-if-else-statements/main.go).

## The If Statement

We can use if statements to execute peices of code if a condition is true. For example:

```go
number1 := 10
number2 := 20

if number1 > number2 {
	fmt.Println("10 is bigger than 20")
}
```

Here, we have two int variables: ```number1``` set to 10 and ```number2``` set to 20. Here we evaluate whether or not ```number1``` is a bigger number than ```number2```. Since 10 isn't bigger than 20, our ```Println``` statement won't be executed.

## The If Else Statement

What if we wanted another peice of code to execute to show the user that the condition we just evaluate was false, rather than doing nothing? This is where we can use the ```else``` statement to do this for us. For example:

```go
number1 := 10
number2 := 20

if number1 > number2 {
	fmt.Println("10 is bigger than 20")
} else {
    fmt.Println("No, 10 is not bigger than 20.")
}
```

Thanks to this else statement, our program will now print out a response telling us that our condition is false.

## The Else If Statement

What if we have multiple conditions we wish to evaluate? If we have several conditions, we should use a ```switch``` statement. But if we have two or three conditions, we can use a ```else if``` statement like so:

```go
if 9%2 == 0 {
	fmt.Println("9 is a even number")
} else if 9 < 10 {
	fmt.Println("9 is smaller than 10!")
} else {
	fmt.Println("No, 9 isn't a even number nor is it smaller than 10.")
}
```

Here, our first condition will fail since the remainder of 9 divided by 2 is not equal 0, so Go will then evaluate our second condition.

Here 9 is smaller than 10, so our Go program will print out that 9 is smaller than 10.

The ```else if``` statement will test a new condition so our first condition evaluated by an ```if``` statement fails.

## Conditional Statements

Let's talk about conditions. A condition in Go can be true or false.

Go supports these mathematical comparison operators:

- Less than ```<```
- Less than or equal ```<=```
- Greater than ```>```
- Greater than or equal ```>=```
- Equal to ```==```
- Not equal to ```!=```

Go also supports these logical operators:

- AND ```&&```
- OR ```||```
- NOT ```!```

So in code, we can write conditions like so:

```go
number1 := 10
number2 := 20

fmt.Println(number1 < number2) // true
fmt.Println(number1 > number2) // false
fmt.Println(number1 <= number2) // true
fmt.Println(number1 >= number2) // false
fmt.Println(number1 == number2) // false
fmt.Println(number1 != number2) // true
```

## Ifception (Ifs inside Ifs)

We can even has ```if``` statements inside ```if``` statements. This is good when we want to evaluate conditions after a previous condition was evaluated. Here's an example:

```go
number1 := 10
number2 := 20

if number1 > 9 {
	if number2 > 19 {
		fmt.Println("Both conditions are true")
	} else {
		fmt.Println("Only the first condition is true")
	}
} else {
	fmt.Println("The first condition is false")
}
```

In this code block, we first evaluate whether or not ```number1``` is bigger than 9. Since this is true, we proceed to our second ```if``` statement and test whether our ```number2``` variable is bigger than 19. Since it is, we'll print out that both conditions are true.

As you can see here, if we had multiple conditions our code would get very messy. 

## Conclusion

In this article, we talked about how we can use if statements and conditionals to control what code gets executed based on certain conditions in Go.

And no, I didn't just write this article to make a crappy pun on Inception. I've never watched the film. Learning control flow and conditional statements are an important milestone in any programming language, so well done you for taking the time to learn how it works in Go!

If you have any questions, feel free to reach out to me on twitter [@willvelida](https://twitter.com/willvelida) and I'll try and help out the best I can.

In my next post, we'll start building Go programs using If, if/else and else/if statements.

Until next time, Happy coding! 🤓🖥️