package main

import "fmt"

func addRec(n int) int {
	if n == 1 {
		return 1
	}
	addRec(n-1)
	return (n + addRec(n-1))
}

func recHead(n int) {
	if n == 0 {
		return
	}
	recHead(n-1)
	fmt.Println(n)
}

func recTail(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	recTail(n-1)
}

func factorialRec(n int) int {
	if n == 1 {
		return 1
	}
	return (n * factorialRec(n-1))
}

func factorialRecAccumulator(accumulator,n int) int {
	if n == 1 {
		return accumulator
	}
	return factorialRecAccumulator(accumulator*n, n-1)
}

// gcd = greatest common divisor
func gcdRec(n1,n2 int) int {
	if n2 == 0 {
		return n1
	}
	return (gcdRec(n2, n1 % n2))
}

func gcdNonRec(n1,n2 int) int {
	for n2 != 0 {
		temp := n2
		n2 = n1 % n2
		n1 = temp
	}
	return n1
}

// scm = smallest common multiple
func scm(n1,n2 int) int {
	if n1 < 0 {
		n1 = -n1
	}
	if n2 < 0 {
		n2 = -n2
	}
	return (n1*n2/gcdRec(n1,n2))
}

func main(){
	//fmt.Println(addRec(5))
	//recHead(5)
	//recTail(5)
	//fmt.Println(factorialRec(3))
	//fmt.Println(factorialRecAccumulator(1,5))
	//fmt.Println(gcdRec(4,6))
	//fmt.Println(gcdNonRec(4,6))
    //fmt.Println(scm(4,6))
    
}