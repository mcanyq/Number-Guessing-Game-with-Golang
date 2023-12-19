package main

import ("fmt"
"math/rand"
"time")


func main(){

	min,max:= 1,100

	rand.Seed(time.Now().UnixNano())
	number:= rand.Intn(max-min)

    attempts:= 10

	for attempts>0{

		fmt.Print("I kept a number between 1 and 100, to guess: ")
        var guess int
        fmt.Scanln(&guess)
	
		if guess>number{
			fmt.Println("Down")
			attempts--
			if attempts==0{
				fmt.Println("trials completed")
			}
		} else if guess<number {
			fmt.Println("Up")
			attempts--
			if attempts==0{
				fmt.Println("trials completed")
			}
		} else{
			fmt.Println("Correct...")
			break
		}
fmt.Println("You have attempts: ", attempts)
	}
}