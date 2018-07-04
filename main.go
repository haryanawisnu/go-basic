package main

import "fmt"
import "strconv"

func main() {
	fmt.Println("==========================")
	fmt.Printf("Hello, %s\n", "senpai")
	
	fmt.Println("==========================")
	fmt.Printf("1 + 1 = %s\n", strconv.Itoa(1+1))
	
	fmt.Println("==========================")
	for i := 0; i < 10; i++ {
		fmt.Printf("Bilangan ke - %s\n",strconv.Itoa(i+1))
	}
	
	fmt.Println("==========================")
	var empty_array []int
	var message string
	have_array := []int{1,2,3,4} 
	fmt.Println(empty_array)
	empty_array, message = arrayAppendarray(empty_array, have_array)
	fmt.Println(have_array)
	fmt.Println(message)
}

func arrayAppendarray(array1 ,array2 []int) ([]int, string) {
	var message string
	message = fmt.Sprintf("Panjang array ke 1 : %s\nPanjang array ke 2 : %s\n", strconv.Itoa(len(array1)), strconv.Itoa(len(array2)))
	for i := 0; i < len(array2); i++ {
		array1 = append(array1, array2[i])
	}
	return array1, message
}