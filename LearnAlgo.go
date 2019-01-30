package go_recommend_me

import (
	"log"
	"math"
)

// An interface for learning algorithms
type LearnAlgo interface {
	// A stochastic gradient descent
	Learn(set TrainingSet) *LearnedFactors

	// Returns approximates user's rating of an item based on some learned factors
	EstimateRating(userIndex int, itemIndex int, lfactors *LearnedFactors) float64
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
func (m *BasicMF) Learn(tset TrainingSet) *LearnedFactors {
	// initialize learned factors

	itemFactors := randomMatInit(tset.NumItems, tset.Dimensionality)
	userFactors := randomMatInit(tset.Dimensionality, tset.NumUsers)

	for step := 0; step < tset.Steps; step++ {
		for i := 0; i < tset.ratingsMatrix.Rows(); i++ {
			for j := 0; j < tset.ratingsMatrix.Cols(); j++ {
				rating, _ := tset.ratingsMatrix.get(i, j)
				if rating > 0 {
					eij := rating - m.estimateItemRating(itemFactors.getRow(i), userFactors.getCol(j), tset.Dimensionality)
					for z := 0; z < tset.Dimensionality; z++  {
						itemFactors.Set(i, z, itemFactors.Get(i, z) + tset.Alpha*(2*eij*userFactors.Get(z, j) - tset.Beta*itemFactors.Get(i, z)))
						userFactors.Set(z, j, userFactors.Get(z, j) + tset.Alpha*(2*eij*itemFactors.Get(i, z) - tset.Beta*userFactors.Get(z, j)))
					}
				}
			}
		}

		var e = 0.0
		for i := 0; i < tset.ratingsMatrix.Rows(); i++ {
			for j := 0; j < tset.ratingsMatrix.Cols(); j++ {
				rating, _ := tset.ratingsMatrix.get(i, j)
				if rating > 0 {
					e += math.Pow(rating - m.estimateItemRating(itemFactors.getRow(i), userFactors.getCol(j), tset.Dimensionality), 2)
					for z := 0; z < tset.Dimensionality; z++ {
						e += (tset.Beta/2) * (math.Pow(itemFactors.Get(i, z), 2) + math.Pow(userFactors.Get(z, j), 2))
					}
				}
			}
		}

		if e < 0.001 {
			break
		}

	}


	l := &LearnedFactors{
		numItems: tset.NumItems,
		numUser: tset.NumUsers,
		dimensionality: tset.Dimensionality,
		ratingsAvg:0,
		itemFactors: itemFactors,
		userFactors: userFactors,
	}

	return l
}



func (m *BasicMF) EstimateRating(userIndex int, itemIndex int, lfactors *LearnedFactors) float64{
	if userIndex >= lfactors.numUser {
		log.Fatalln("Out of Bounds error")
	}

	if itemIndex >= lfactors.numItems {
		log.Fatalln("Out of Bounds error")
	}


	return m.estimateItemRating(lfactors.userFactors.getCol(userIndex),
		lfactors.itemFactors.getRow(itemIndex), lfactors.dimensionality)
}




