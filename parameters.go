package go_recommend_me


type modelParameters struct{
	numUsers		int
	numItems		int
	dimension		int
	iteration		int
	trainingSize	int

	step			float64
	lambda			float64

	step_bias		float64
	lamba_bias		float64

	algorithmType	int
	binWidth		int
	projFamSize		int

	beta			float64

}
