func isIsomorphic(s string, t string) bool {
	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	for i := range len(s) {
		charS := s[i]
		charT := t[i]

		if v, ok := sMap[charS]; ok && v != charT {
			return false
		}

		if v, ok := tMap[charT]; ok && v != charS {
			return false
		}

		sMap[charS] = charT
		tMap[charT] = charS
	}

	return true
}