func isIsomorphic(s string, t string) bool {
    sMap := make(map[string]string)
	tMap := make(map[string]string)

    for i := range len(s) {
        v, ok := sMap[string(s[i])]
		if ok && v != string(t[i]) {
			return false
		}
		v, ok = tMap[string(t[i])]
		if ok && v != string(s[i]) {
			return false
		}
		sMap[string(s[i])] = string(t[i])
		tMap[string(t[i])] = string(s[i])
    }

	return true
}