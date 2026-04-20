func convert(s string, numRows int) string {
    r := 0
	dir := true
    arr := make([]*strings.Builder, numRows)
	for i := range numRows {
		arr[i] = &strings.Builder{}
	}

	updateDir := func() {
        if numRows == 1 {
            return
        }

		if dir {
			if r < numRows - 1 {
				r++
			} else {
				r--
				dir = !dir
			}
		} else {
			if r > 0 {
				r--
			} else {
				r++
				dir = !dir
			}
		}
	}

    for _, v := range s {
		arr[r].WriteRune(v)
		updateDir()
	}

	ans := &strings.Builder{}
	for i := range numRows {
		ans.WriteString(arr[i].String())
	}

	return ans.String()
}
