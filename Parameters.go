package go_recommend_me

// Parmeters for the algorithm
type ModelParameters struct{
	NumUsers		int
	NumItems		int

	// k or the dimensionality of the joint latent factor space
	// ie kind of determining the space size for the latent factors
	Dimensionality	int

	//number of known ratings
	TrainingSize	int

	// Step size in the sdg algo
	Steps			int

	//constant that controls the extent of regularization
	Alpha			float64
	Beta			float64

	algorithmType	int
	binWidth		int
	projFamSize		int


}
