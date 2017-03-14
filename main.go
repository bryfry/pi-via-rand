package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"math"
	"math/big"
	"math/rand"
	"time"
)

type Trial struct {
	A       *big.Int
	B       *big.Int
	Coprime bool
}

func main() {
	var num_trials = 500
	var i uint64

	fmt.Printf("n trials, max rand, variance, min, q1, q2, q3, max\n")
	for i = 1; i < 100; i++ {
		var max_rand = big.NewInt(2 << i)
		var exp_pis []float64
		for i := 1; i <= 4096; i++ {
			exp_pis = append(exp_pis, Experiment(max_rand, num_trials))
		}
		max, _ := stats.Max(exp_pis)
		min, _ := stats.Min(exp_pis)
		mean, _ := stats.Mean(exp_pis)
		variance, _ := stats.Variance(exp_pis)
		q, _ := stats.Quartile(exp_pis)
		fmt.Printf("%d, %d, %f, %f, %f, %f, %f, %f\n", num_trials, max_rand, variance, min, q.Q1, mean, q.Q3, max)
	}
}

func Experiment(max_rand *big.Int, num_trials int) (exp_pi float64) {
	trials := make([]Trial, num_trials)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	cc := 0 // coprime count
	for i, t := range trials {
		var C *big.Int
		t.A, t.B, C = big.NewInt(0), big.NewInt(0), big.NewInt(0)
		trials[i].A = t.A.Rand(r, max_rand)
		trials[i].B = t.B.Rand(r, max_rand)
		C = C.GCD(nil, nil, t.A, t.B)
		if C.Cmp(big.NewInt(1)) == 0 {
			trials[i].Coprime = true
			cc += 1
		}
	}
	//for _, t := range trials {
	//	fmt.Printf("%s, %s, %s, %t\n", t.A.String(), t.B.String(), t.C.String(), t.Coprime)
	//}
	return math.Sqrt(6 / (float64(cc) / float64(num_trials)))
}
