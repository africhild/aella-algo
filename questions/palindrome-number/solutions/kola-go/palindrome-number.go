package main

func main() {

}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	temp := x
	var reversedResult int
	for temp > 0 {
		reversedResult = (reversedResult * 10) + temp % 10
		temp /= 10
	}
	return reversedResult == x
}
