package main

import (
// "fmt"
)

type NeuralStack struct {
	stackWidth  int
	o_prime_dim int

	V []float64 // stack states
	s []float64 // stack strengths
	d []float64 // push strengths
	u []float64 // pop strengths
	r []float64
	o []float64

	V_delta []float64 // stack states
	s_delta []float64 // stack strengths
	d_error []float64 // push strengths
	u_error []float64 // pop strengths

	t int
}

func (ns *NeuralStack) reset() {
	ns.V = []float64{0.0}
	ns.s = []float64{0.0}
	ns.d = []float64{0.0}
	ns.u = []float64{0.0}
	ns.r = []float64{0.0}
	ns.o = []float64{0.0}

	ns.V_delta = []float64{0.0}
	ns.s_delta = []float64{0.0}
	ns.d_error = []float64{0.0}
	ns.u_error = []float64{0.0}
	ns.t = 0
}

// func (ns *NeuralStack) s_t(i int) {
// 	if i >= 0 && i < ns.t {
// 		inner_sum := ns.s[ns.t-1][i+1 : ns.t-0]
// 		return Relu(ns.s[ns.t-1][i] - Relu(ns.u[ns.t]-Sum(inner_sum)))
// 	} else if i == ns.t {
// 		return ns.d[ns.t]
// 	} else {
// 		fmt.Println("Problem")
// 	}

// }

// func (ns *NeuralStack) pushAndPopForward(v_t, d_t, u_t float64) {
// 	ns.d = Append(ns.d, d_t)
// 	ns.d_error = Append(ns.d_error, 0)

// 	ns.u = Append(ns.u, u_t)
// 	ns.u_error = Append(ns.u_error, 0)

//     new_s = [ns.t+1]float64
//     for i := 0; i < ns.t+1; i++ {
//         new_s[i] = ns.s_t
//     }
// }
