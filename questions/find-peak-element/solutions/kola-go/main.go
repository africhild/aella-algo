package main


func main() {

}

func findPeakElement(n []int) int {
	if len(n) == 1 {
		return 0
	}
	left, right := 0, len(n) - 1
	for left < right {
		mid := (left + right) /2
		if n[mid] > n[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

