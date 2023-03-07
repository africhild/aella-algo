package main

// func SumAll(n []int ) int {
// 	sum := 0
// 	// for i := 0; i < len(n); i++ {
// 	// 	sum += n[i]
// 	// }

// 	for _, v := range n {
// 		sum += v
// 	}
// 	return sum
// }

func SumAll(n []int ) int {
    if len(n) == 0 {
		return 0
	}
	return n[0] + SumAll(n[1:]);
}
// func Reverse(s string) string {
// 	runes := []rune(s)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }

//func Reverse(s string) string {
//	var res string
//	for _, r := range s {
//		res = string(r) + res
//	}
//	return res
//}

//func Reverse(s string) string {
//	var sb strings.Builder
//	for i := len(s) - 1; i >= 0; i-- {
//		sb.WriteByte(s[i])
//	}
//	return sb.String()
//}

func Reverse(s string) string {
	runes := []rune(s)
	j := len(runes) - 1
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[j-i] = runes[j - i], runes[i]
	}
	return string(runes)
}

