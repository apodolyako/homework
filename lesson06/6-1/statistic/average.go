package statistic

func Average(xs []float64) []float64 {
	total := float64(0)

	if len(xs) == 0 {
		return []float64{0, 0}
	}

	for _, x := range xs {
		total += x
	}
	return []float64{total, (total / float64(len(xs)))}
}
