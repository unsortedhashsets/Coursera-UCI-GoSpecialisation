/* 
Example:
[...@... 2.w2]$ ./displacement 
Enter value of Acceleration(A): 10
Enter value of Initial Velocity(V0): 2
Enter value of Initial Displacement(S0): 1
Enter value of Time(t): 5
Displacement after 5 Seconds: 136
*/

package main

import "fmt"

func PrepareArgs(A *float64, V0 *float64, S0 *float64, t *float64){
	fmt.Printf("Enter value of Acceleration(A): ")
	fmt.Scanln(A)

	fmt.Printf("Enter value of Initial Velocity(V0): ")
	fmt.Scanln(V0)

	fmt.Printf("Enter value of Initial Displacement(S0): ")
	fmt.Scanln(S0)

	fmt.Printf("Enter value of Time(t): ")
	fmt.Scanln(t)
}

func GenDisplaceFn(acc, iniVelocity, iniDisplacement float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*acc*(t*t) + iniVelocity*t + iniDisplacement
	}
}

func main() {
	var A, V0, S0, t float64

	PrepareArgs(&A, &V0, &S0, &t)
	fn := GenDisplaceFn(A, V0, S0)

	fmt.Printf("Displacement after %v Seconds: %v\n", t, fn(t))
}

