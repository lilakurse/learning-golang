package examples

import
	"fmt"

func varibles(){
	for i:=1;i<=100;i++{
		if i %3==0 {
			fmt.Print("Fizz")
		} else if i%5==0{
			fmt.Print("Buzz")
		}else if i%3==0 && i%5==0{
				fmt.Print("FizzBuzz")
		}
	}
}