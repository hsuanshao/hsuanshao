package main

func main() {}

//Runtime: 4 ms
// Memory Usage: 6.3 MB

func convert(s string, numRows int) string {
	cov := make([]string, numRows)
	idx := 0

	direction := true
	for _, c := range s {
		if idx < 0 {
			idx = 0
		}
		cov[idx] += string(c)

		if !direction {
			idx--
			if idx < 0 {
				direction = true
				idx = 0
			}
		}

		if direction {
			idx++
			if idx == numRows {
				direction = false
				idx = numRows - 2
			}

		}
	}

	str := ""
	for _, s := range cov {
		str += s
	}

	return str
}
