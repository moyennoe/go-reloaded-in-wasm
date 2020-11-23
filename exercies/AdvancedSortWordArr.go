package piscine

//AdvancedSortWordArr will sort a arrays of strings and use a function here Compare.
func AdvancedSortWordArr(a []string, f func(a, b string) int) {

	var len int = 0

	var char string

	for range a {

		len++

	}

	for i := 0; i < len-1; i++ {

		for j := 0; j < len-i-1; j++ {

			if f(a[j], a[j+1]) == 1 { //if f == -1 this will return a reversed sorted tab and if f == 0 this will return the initial arrays we writed

				char = a[j]
				a[j] = a[j+1]
				a[j+1] = char

			}

		}
	}

}

//Compare is a function who will compare a string between them
func Compare(a, b string) int {

	if a == b {

		return 0

	} else if a < b {

		return -1

	} else {

		return 1
	}
}
