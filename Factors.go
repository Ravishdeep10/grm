package go_recommend_me

// Output of a learning algorithm's stochastic gradient descent process
type LearnedFactors struct{
	// Passed in from tset parameters
	numUser			int
	numItems		int
	dimensionality	int


	// average of the ratings
	ratingsAvg		float64

	// Two matrices based on latent factors whose product would
	//	reult in the matrix passed in through the tset parameters
	userFactors		*DenseMatrix
	itemFactors		*DenseMatrix

}

