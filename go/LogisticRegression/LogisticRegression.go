package LogisticRegression

import (
	"math"
)

type LogisticRegression struct {
	N int
	N_in int
	N_out int
	W [][]float64
	B []float64
}


func LogisticRegression__construct(this *LogisticRegression, N int, n_in int, n_out int) {
	this.N = N
	this.N_in = n_in
	this.N_out = n_out

	this.W = make([][]float64, n_out)
	for i := 0; i < n_out; i++ { this.W[i] = make([]float64, n_in) }
	
	this.B = make([]float64, n_out)
}

func LogisticRegression_train(this *LogisticRegression, x []int, y []int, lr float64) {
	p_y_given_x := make([]float64, this.N_out)
	dy := make([]float64, this.N_out)
	
	for i := 0; i < this.N_out; i++ {
		p_y_given_x[i] = 0
		for j := 0; j < this.N_in; j++ {
			p_y_given_x[i] += this.W[i][j] * float64(x[j])
		}
		p_y_given_x[i] += this.B[i]
	}
	LogisticRegression_softmax(this, p_y_given_x)
	
	for i := 0; i < this.N_out; i++ {
		dy[i] = float64(y[i]) - p_y_given_x[i]
		
		for j := 0; j < this.N_in; j++ {
			this.W[i][j] += lr * dy[i] * float64(x[j]) / float64(this.N)
		}

		this.B[i] += lr * dy[i] / float64(this.N)
	}
	
}

func LogisticRegression_softmax(this *LogisticRegression, x []float64) {
	var (
		max float64
		sum float64
	)

	for i := 0; i < this.N_out; i++ { if max < x[i] {max = x[i]} }
	for i := 0; i < this.N_out; i++ {
		x[i] = math.Exp(x[i] - max)
		sum += x[i]
	}

	for i := 0; i < this.N_out; i++ { x[i] /= sum }
}

func LogisticRegression_predict(this *LogisticRegression, x []int, y []float64) {
	for i := 0; i < this.N_out; i++ {
		y[i] = 0
		for j := 0; j < this.N_in; j++ {
			y[i] += this.W[i][j] * float64(x[j])
		}
		y[i] += this.B[i]
	}

	LogisticRegression_softmax(this, y)
}
