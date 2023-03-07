package main


// func main() {

// }

func NumInList(list []int, num int) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

