package statistic

import "testing"

type testpair struct {
	name    int
	values  []float64
	average float64
	sum     float64
}

var tests = []testpair{
	{
		name:    1,
		values:  []float64{3, 3, 3},
		average: 3,
		sum:     9,
	},
	{
		name:    2,
		values:  []float64{1, 2},
		average: 1.5,
		sum:     3,
	},
	{
		name:    3,
		values:  []float64{},
		average: 0,
		sum:     0,
	},
}

func TestAverage(t *testing.T) {

	for _, pair := range tests {

		v := Average(pair.values)
		if v[1] != pair.average {
			t.Error(
				"For TestName =", pair.name,
				"expected", pair.average,
				"got", v[1],
			)
		}
		if v[0] != pair.sum {
			t.Error(
				"For TestName =", pair.name,
				"expected", pair.sum,
				"got", v[0],
			)
		}
	}

}
