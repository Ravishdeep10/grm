package go_recommend_me


type learnedFactors struct{
	numUser			int
	numItems		int
	dimensionality	int

	userBias		[]float64
	itemBias		[]float64

	ratingsAvg		float64

	userFactors		*DenseMatrix
	itemFactors		*DenseMatrix

}

