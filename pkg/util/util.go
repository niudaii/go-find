package util

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

func CheckExt(fileWhiteExts []string, ext string) bool {
	for _, fileWhiteExt := range fileWhiteExts {
		if ext == fileWhiteExt {
			return true
		}
	}
	return false
}
