package go_recommend_me


type modelParameters struct{
	numUsers		int
	numItems		int

	// k or the dimensionality of the joint latent factor space
	// ie kind of determining the space size for the latent factors
	k				int

	//number of known ratings
	trainingSize	int

	// Step size in the sdg algo
	steps			int

	//constant that controls the extent of regularization
	alpha			float64
	beta			float64

	step_bias		float64
	lamba_bias		float64

	algorithmType	int
	binWidth		int
	projFamSize		int


}
