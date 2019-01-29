package go_recommend_me

type trainingSet struct {
	modelParameters
	sumRatings		float64
	ratingsMatrix	SparseMatrix
}

func (t *trainingSet) initialize(parameters modelParameters) {
	t.modelParameters = parameters
	t.sumRatings = 0
	t.ratingsMatrix.createNew(t.numItems, t.numUsers)

}

// Fill the training set with a known user/item rating
func (t *trainingSet) setRating(user int, item int, value float64){
	t.ratingsMatrix.set(value, item, user)
	t.sumRatings += value
}

// Get average of vector of ratings by a user
func (t *trainingSet) averageUserRating(user int) float64 {
	return t.ratingsMatrix.avgColsVals(user)
}

// Get average of vector from ratings of an item
func (t *trainingSet) averageItemRating(item int) float64 {
	return t.ratingsMatrix.avgRowVals(item)
}

// Add a known rating to the training set
func (t* trainingSet) addRating(user int, item int, value float64) {
	t.ratingsMatrix.set(value, item, user)
	t.trainingSize++
	t.sumRatings += value
}

// Add a new user to the training set
func (t *trainingSet) addUser() {
	t.ratingsMatrix.addCol()
}

// Add a new item to the training set
func (t *trainingSet) addItem() {
	t.ratingsMatrix.addRow()
}
