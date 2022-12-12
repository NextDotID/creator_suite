package util

// UniqueStrings Deduplication
func UniqueStrings(strs []string) []string {
	var m = make(map[string]bool)
	for _, s := range strs {
		m[s] = true
	}
	var result = make([]string, 0, len(strs))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// StringsIn Determine whether the element is in the string list
func StringsIn(strs []string, target string) bool {
	for _, s := range strs {
		if s == target {
			return true
		}
	}
	return false
}

// MapStrings For each element in strs, perform the f() functon and return
func MapStrings(strs []string, f func(s string) string) []string {
	res := make([]string, len(strs))
	for i, s := range strs {
		res[i] = f(s)
	}
	return res
}

// StringsContains Determine whether strs is completely in targets
func StringsContains(strs, targets []string) bool {
	if len(targets) < len(strs) {
		return false
	}
	targetMap := make(map[string]bool)
	for _, t := range targets {
		targetMap[t] = true
	}
	for _, s := range strs {
		if _, ok := targetMap[s]; !ok {
			return false
		}
	}
	return true
}

// StringsContainsAny Determine whether strs has any elements in targets
func StringsContainsAny(strs, targets []string) bool {
	targetMap := make(map[string]bool)
	for _, t := range targets {
		targetMap[t] = true
	}
	for _, s := range strs {
		if _, ok := targetMap[s]; ok {
			return true
		}
	}
	return false
}
