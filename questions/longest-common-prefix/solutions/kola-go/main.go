package main


func main() {

}

func longestCommonPrefix(str []string) string {

	// horizontal scan
	// if len(str) == 0 {
	// 	return ""
	// }
	// prefix := str[0]
	// // 01 - sample
	// // for i := 1; i < len(str); i++ {
	// // 	for chr := 0; chr < len(prefix); chr++ {
	// // 		if chr >= len(str[i]) || prefix[chr] != str[i][chr] {
	// // 			prefix = prefix[:chr]
	// // 			break
	// // 		}
	// // 	}
	// // }
	// // 02 - sample
	// fmt.Println(str[1:])
	// for _, s := range str[1:] {
	// 	for chr := 0; chr < len(s); chr++ {
	// 		if chr >= len(prefix) || prefix[chr] != s[chr] {
	// 			prefix = prefix[:chr]
	// 			break
	// 		}
	// 	}
	// }
	// return prefix

	// vertical scan
	for chr := 0; chr < len(str[0]); chr++ {
		for s := 1; s < len(str); s++ {
			if chr >= len(str[s]) || str[0][chr] != str[s][chr] {
				return str[0][:chr]
			}
		}
	}
	return str[0]
}
