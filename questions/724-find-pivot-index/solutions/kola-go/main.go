package main

// func main() {
// 	nums := []int{3, 5, 2, 6, 7}
// 	fmt.Println(nums[:1], nums[2:])
// 	// for i, n := range nums {
// 	// 	fmt.Println(i, n,nums[:i], sum(nums[:i]))
// 	// }
// }

func sum(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}
func pivotIndex(nums []int) int {
	for i, _ := range nums {
		if sum(nums[:i]) == sum(nums[i+1:]) {
			return i
		}
	}
	return -1
}
