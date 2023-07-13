package util

func FindLargestNumberOfNodeId(arr []string) int {
	if len(arr) == 0 {
		return 0
	}
	cur := StringToInt(arr[0])
	for _, v := range arr {
		if StringToInt(v) > cur {
			cur = StringToInt(v)
		}
	}
	return cur

}
