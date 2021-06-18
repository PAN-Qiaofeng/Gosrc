package mysort

func InsertSort(elements []int) {
	n := len(elements)
	for i := 1; i < n; i++ {
		temp := elements[i]
		var j int
		for j = i - 1; temp < elements[j]; j-- {
			elements[j+1] = elements[j]
			if j == 0 {
				j -= 1
				break
			}
		}
		elements[j+1] = temp
	}
}
