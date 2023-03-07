package main


func main() {

}

func romanToInt(s string) int {
	dataMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000}
	var result = dataMap[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if dataMap[s[i]] < dataMap[s[i+1]] {
			result -= dataMap[s[i]]
		} else {
			result += dataMap[s[i]]
		}
	}
	return result
}
