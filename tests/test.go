package main

import (
	"fmt"
	"go_recommend_me"
)

func main() {

	// A Basic Matrix Factorization learning model which includes the algorithm and estimator
	var model go_recommend_me.BasicMF

	//  Setup the model parameters
	var params go_recommend_me.ModelParameters
	params.Dimensionality = 2
	params.NumItems = 3
	params.NumUsers = 2
	params.Steps = 5000
	params.Alpha = 0.0002
	params.Beta = 0.02
	params.TrainingSize = 5

	var tset go_recommend_me.TrainingSet

	// Initialize the training set with parameters and fill with known user/item ratings
	tset.Initialize(params)

	tset.SetRating(0, 0, 4)
	tset.SetRating(0, 1, 1)
	tset.SetRating(0, 2, 6)

	tset.SetRating(1, 1, 1)
	tset.SetRating(1, 0, 2)

	// Factors measure the extent to the extent that an item possess some feature
	learned := model.Learn(tset)

	// Rating estimation
	fmt.Println("user [0] item[0], rating = ", model.EstimateRating(0, 0, learned))
	fmt.Println("user [0] item[1], rating = ", model.EstimateRating(0, 1, learned))
	fmt.Println("user [0] item[2], rating = ", model.EstimateRating(0, 2, learned))

	fmt.Println("user [1] item[0], rating = ", model.EstimateRating(1, 0, learned))
	fmt.Println("user [1] item[1], rating = ", model.EstimateRating(1, 1, learned))
	fmt.Println("user [1] item[2], rating = ", model.EstimateRating(1, 2, learned))





}