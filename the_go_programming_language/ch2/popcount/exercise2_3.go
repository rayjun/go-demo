package popcount

func LoopPopCount(x uint64) int {
	count := 0

	for i := uint64(0); i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}

	return count
}
