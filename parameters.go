package go_recommend_me


type modelParameters struct{
	numUsers		int
	numItems		int
	dimensionality	int
	iteration		int
	trainingSize	int

	step			float64
	lambda			float64

	step_bias		float64
	lamba_bias		float64

	algorithmType	int
	seed			int
	binWidth		int
	projFamSize		int

	beta			float64

}

type learnedFactors struct{
	numUser			int
	numItems		int
	dimensionality	int

	userBias		[]float64
	itemBias		[]float64

	ratingsAvg		float64

	userFactors		[]float64
	itemFactors		[]float64

}

type ratingEstimatorParams struct{
	userIndex	int
	itemIndex	int
	tset		trainingSet
	lfactors    learnedFactors
}

type learningAlgoParams struct {
	tset		trainingSet
	params  	modelParameters
	socialMat	SparseMatrix
}
