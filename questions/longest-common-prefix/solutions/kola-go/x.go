package main


func xmain() {

}

func xlongestCommonPrefix(str []string,) string {
	// vertical scan
	// for i := 0; i < len(str[0]); i++ {
	// 	for j := 1; j < len(str); j++ {
	// 		if i >= len(str[j]) || str[j][i] != str[0][i] {
	// 			return str[0][:i]
	// 		}
	// 	}
	// }
	// return str[0]

	// horizontal scan
	if len(str) == 0 {
		return ""
	}
	prefix := str[0]
	for i := 1; i < len(str); i++ {
		for j := 0; j < len(prefix); j++ {
			if j >= len(str[i]) || str[i][j] != prefix[j] {	
				prefix = prefix[:j]
				break
			}
		}
	}
	return prefix


	// divide and conquer
	if len(str) == 0 {
		return ""
	}
	low := 0
	high := len(str[0])
	for low < high {
		mid := (low + high) / 2
		for i := 1; i < len(str); i++ {
			if len(str[i]) < mid || str[i][mid] != str[0][mid] {
				low = mid + 1
				break
			}
			if i == len(str) - 1 {
				high = mid
			}
		}
	}
	return str[0][:low]
}
