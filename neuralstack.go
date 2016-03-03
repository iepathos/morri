package main

type NeuralStack struct {
	stackWidth int
	oPrimeDim  int

	V []float64   // stack states
	s [][]float64 // stack strengths
	d []float64   // push strengths
	u []float64   // pop strengths
	r []float64
	o []float64

	VDelta []float64 // stack states
	sDelta []float64 // stack strengths
	dError []float64 // push strengths
	uError []float64 // pop strengths

	t int
}

func (ns *NeuralStack) reset() {
	ns.V = []float64{}
	ns.s = [][]float64{}
	ns.d = []float64{}
	ns.u = []float64{}
	ns.r = []float64{}
	ns.o = []float64{}

	ns.VDelta = []float64{}
	ns.sDelta = []float64{}
	ns.dError = []float64{}
	ns.uError = []float64{}
	ns.t = 0
}

func (ns *NeuralStack) sT(i int) float64 {
	if i >= 0 && i < ns.t {
		innerSum := ns.s[ns.t-1][i+1 : ns.t-0]
		return Relu(ns.s[ns.t-1][i]-Relu(ns.u[ns.t]-Sum(innerSum), false), false)
	} else if i == ns.t {
		return ns.d[ns.t]
	} else {
		panic("An error is happening in NeuralStack.sT")
	}
}

func (ns *NeuralStack) pushAndPopForward(vT, dT, uT float64) {
	ns.d = Append(ns.d, dT)
	ns.dError = Append(ns.dError, 0)

	ns.u = Append(ns.u, uT)
	ns.uError = Append(ns.uError, 0)

	newS := make([]float64, ns.t+1)
	for i := 0; i < ns.t+1; i++ {
		newS[i] = ns.sT(i)
	}
	ns.s = append(ns.s, newS)
	// ns.sDelta = append(np.zeros_like(new_s))
}
