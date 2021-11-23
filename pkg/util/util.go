package util

func ToSet(list1, list2 []string) []string {
	list1 = append(list1, list2...)
	var set []string
	hashSet := make(map[string]struct{})
	for _, v := range list1 {
		if v != "" {
			hashSet[v] = struct{}{}
		}
	}
	for k, _ := range hashSet {
		set = append(set, k)
	}
	return set
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
