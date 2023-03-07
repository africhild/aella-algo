package main


func main() {

}
// func twoSums(num []int, target int) []int{
// 	for i := 0; i < len(num) - 1; i++ {
// 		for j := i +1; j <  len(num); j++ {
// 			if num[i] + num[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{-1, -1}
// }

func twoSums(num []int, target int) []int {
	indexMap := make(map[int]int)
	// for i := 0; i < len(num); i++ {
	// 	if ind, val := indexMap[target - num[i]]; val {
	// 		return []int{ind, i}
	// 	}
	// 	indexMap[num[i]] = i
	// }

	for i, n :=  range num {
		if idx, state := indexMap[target - n]; state {
			return []int{idx, i}
		}
		indexMap[n] = i
	}
	return []int{}
}

