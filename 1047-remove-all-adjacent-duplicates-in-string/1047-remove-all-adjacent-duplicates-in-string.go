package main

func removeDuplicates(input string) string {

	buf := make([]byte, len(input))

	checkPoint := 0
	recordPoint := 0

	for i := 0; i < len(input); i++ {

		if buf[0] == 0 && i < len(input)-1 {
			buf[0] = input[i]
			recordPoint++
			i++
		}

		if input[i] != buf[checkPoint] {
			buf[recordPoint] = input[i]
			checkPoint++
			recordPoint++
		} else {
			buf[checkPoint] = 0

			if checkPoint != 0 {
				checkPoint--
				recordPoint--
			} else {
				recordPoint = 0
			}

		}

	}

	buf = buf[:recordPoint]

	return string(buf)
}
