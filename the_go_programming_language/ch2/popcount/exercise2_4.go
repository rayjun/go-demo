package popcount

func ShiftPopCount(x uint64) int {
	count := 0
	for i := uint64(0); i < 64; i++ {
		if x&1 == 1 {
			count += 1
		}
		x >>= 1
	}

	return count
}
