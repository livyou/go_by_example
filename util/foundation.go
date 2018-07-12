package util

import (
	"fmt"
	"math"
)
const s string = "constant"

func HelloWorld(){
	fmt.Println("hello world")
}

func Values(){
	fmt.Println("go" + "lang")
	fmt.Println("1+1 = " ,1+1 )
	fmt.Println("7.0/3.0 = " , 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(false || true)
	fmt.Println(!false)
}

func Variables(){
	var a = "initial"
	fmt.Println(a)

	var b,c = 1,2
	fmt.Println(b,c)

	var d = true 
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}

func Constants(){
	fmt.Println(s)

	const n = 500000000
	const d = 3e20/n

	fmt.Println(n)
	fmt.Println(d)
	fmt.Printf("+%T",n)
	fmt.Println("....")
	fmt.Println(int64(d))
	
	fmt.Println(math.Sin(d))
}

