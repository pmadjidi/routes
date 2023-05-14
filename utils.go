package main

func uniqueDsts(dsts []string) []string {
	dict := make(map[string]bool)
	uDsts := make([]string, 0)
	for _, dst := range dsts {
		_, ok := dict[dst]
		if !ok {
			uDsts = append(uDsts, dst)
			dict[dst] = true
		}
	}
	return uDsts
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
