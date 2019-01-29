package go_recommend_me

import (
	"log"
	"math"
)

// An interface for learning algorithms
type LearnAlgo interface {
	// A stochastic gradient descent
	learn(set trainingSet, parameters modelParameters) *learnedFactors

	// Returns approximates user's rating of an item based on some learned factors
	estimateRating(userIndex int, itemIndex int, lfactors *learnedFactors) float64
}

// A basic Matrix factorization algorithm

type BasicMF struct {}

func (m *BasicMF) estimateItemRating(a []float64, b[]float64, dimensionality int) float64{
	sum := 0.0
	for i := 0; i < dimensionality; i ++ {
		sum += (a[i] * b[i])
	}
	return sum
}


// A Matrix Factorization algo using Stochastic Gradient Learning
func (m *BasicMF) learn(tset trainingSet, parameters modelParameters) *learnedFactors {
	// initialize learned factors

	itemFactors := randomMatInit(tset.numItems, tset.k)
	userFactors := randomMatInit(tset.k, tset.numUsers)

	for step := 0; step < tset.steps; step++ {
		for i := 0; i < tset.ratingsMatrix.Rows(); i++ {
			for j := 0; j < tset.ratingsMatrix.Cols(); j++ {
				rating, _ := tset.ratingsMatrix.get(i, j)
				if rating > 0 {
					eij := rating - m.estimateItemRating(itemFactors.getRow(i), userFactors.getCol(j), tset.k)
					for z := 0; z < tset.k; z++  {
						itemFactors.Set(i, z, itemFactors.Get(i, z) + tset.alpha*(2*eij*userFactors.Get(z, j) - tset.beta*itemFactors.Get(i, z)))
						userFactors.Set(z, j, userFactors.Get(z, j) + tset.alpha*(2*eij*itemFactors.Get(i, z) - tset.beta*userFactors.Get(z, j)))
					}
				}
			}
		}

		var e = 0.0
		for i := 0; i < tset.ratingsMatrix.Rows(); i++ {
			for j := 0; j < tset.ratingsMatrix.Cols(); j++ {
				rating, _ := tset.ratingsMatrix.get(i, j)
				if rating > 0 {
					e += math.Pow(rating - m.estimateItemRating(itemFactors.getRow(i), userFactors.getCol(j), tset.k), 2)
					for z := 0; z < tset.k; z++ {
						e += (tset.beta/2) * (math.Pow(itemFactors.Get(i, z), 2) + math.Pow(userFactors.Get(z, j), 2))
					}
				}
			}
		}

		if e < 0.001 {
			break
		}

	}


	l := &learnedFactors{
		numItems: tset.numItems,
		numUser: tset.numUsers,
		dimensionality: tset.k,
		ratingsAvg:0,
		itemFactors: itemFactors,
		userFactors: userFactors,
	}

	return l
}



func (m *BasicMF) estimateRating(userIndex int, itemIndex int, lfactors *learnedFactors) float64{
	if userIndex >= lfactors.numUser {
		log.Fatalln("Out of Bounds error")
	}

	if itemIndex >= lfactors.numItems {
		log.Fatalln("Out of Bounds error")
	}


	return m.estimateItemRating(lfactors.userFactors.getCol(userIndex),
		lfactors.itemFactors.getRow(itemIndex), lfactors.dimensionality)
}




