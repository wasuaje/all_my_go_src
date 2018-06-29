package main

import (
"fmt"
)

type ErrSqrt  struct {
	errNum  int
	errMes  string
}


func (e *ErrSqrt) Error() string {
	return fmt.Sprintf("For value= %d - %s", e.errNum, e.errMes)
}


func Sqrt(x float64) (float64, error) {
	if x > 0 {
		z := float64(x)
		for i:=0; i<100; i++ {
			z -= (z*z - x) / (2*z)
		}
		return z, &ErrSqrt{int(x), "Your result is: "}
	}else if x < 0{
		return -1, &ErrSqrt{int(x), "Can't do this with negative numbers" }
	}else{
		return -1, &ErrSqrt{int(x), "Oops i don't work with 0" }
	}
}

func main() {

	for i :=-10; i<11; i++ {
		if result, message := Sqrt(float64(i)); result > 0 {
			fmt.Println("Success: ", message, result)
		} else {
			fmt.Println("Failed: ", message)
		}

	}
}
