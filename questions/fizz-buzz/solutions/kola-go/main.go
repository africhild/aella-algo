package main

import (
    "strconv"
)

// type hash struct {
// 	key   int
// 	value string
// }
// type hashMap struct {
// 	m map[int]string
// }

func main() {

}

// func FizzBuzz(n int) []string {
// 	var result []string
// 	for i := 1; i <= n; i++ {
// 		if i%3 == 0 && i%5 == 0 {
// 			result = append(result, "FizzBuzz")
// 		} else if i%3 == 0{
// 			result = append(result, "Fizz")
// 		}else if i%5 == 0 {
// 			result = append(result, "Buzz")	
// 		} else{
// 			result = append(result, strconv.Itoa(i))
// 		}
// 	}
// 	return result
// }


func FizzBuzz(n int) [] string {
	var result [] string
	for i := 1; i <= n; i++ {
		num_ans_str := ""
		if i%3 == 0 {
			num_ans_str += "Fizz"
		}
		if i%5 == 0 {
			num_ans_str += "Buzz"
		}
		if num_ans_str == "" {
			num_ans_str = strconv.Itoa(i)
		}
		result = append(result, num_ans_str)
	}
	return result
}

