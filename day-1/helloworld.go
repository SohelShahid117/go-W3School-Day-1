//https://www.w3schools.com/go/go_variables.php
//go home,introduction,get started,syntax,comments,variables
package main
import ("fmt")

var a2 int
var b2 int = 22
var c2 = 32
 

func main() {
  fmt.Println("Hello World!")
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)

  var a string
  var b int
  var c bool

  fmt.Println(a)  //a is ""
  fmt.Println(b)  //b is 0
  fmt.Println(c)  //c is false
  
  /*
  By running the code, we can see that they already have the default values of their respective types:

	a is ""
	b is 0
	c is false
  */
  var student3 string
  student3 = "John"
  fmt.Println(student3)

  a2 = 12
  fmt.Println(a2)
  fmt.Println(b2)
  fmt.Println(c2)

  a3 := 1111
  fmt.Println(a3)

  var a5, b5, c5, d5 int = 11, 31, 51, 71

  fmt.Println(a5)
  fmt.Println(b5)
  fmt.Println(c5)
  fmt.Println(d5)


  var (
    a6 int
    b6 int    = 1
    c6 string = "hello"
  )

  fmt.Println(a6)
  fmt.Println(b6)
  fmt.Println(c6)

  /*
  Camel Case:
  myVariableName = "John"

  Pascal Case:
  MyVariableName = "John"

  Snake Case:
  my_variable_name = "John"
  */
}
//Single-line comments start with two forward slashes (//).

/*
Go Multi-line Comments
Multi-line comments start 
with /* and  */