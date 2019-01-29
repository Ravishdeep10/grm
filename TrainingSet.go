package go_recommend_me

// Training set class that contains the parameyters and the rating matrix
type TrainingSet struct {
	ModelParameters
	sumRatings		float64
	ratingsMatrix	SparseMatrix
}

func (t *TrainingSet) initialize(parameters ModelParameters) {
	t.ModelParameters = parameters
	t.sumRatings = 0
	t.ratingsMatrix.createNew(t.numItems, t.numUsers)

}

// Fill the training set with a known user/item rating
func (t *TrainingSet) setRating(user int, item int, value float64){
	t.ratingsMatrix.set(value, item, user)
	t.sumRatings += value
}

// Get average of vector of ratings by a user
func (t *TrainingSet) averageUserRating(user int) float64 {
	return t.ratingsMatrix.avgColsVals(user)
}

// Get average of vector from ratings of an item
func (t *TrainingSet) averageItemRating(item int) float64 {
	return t.ratingsMatrix.avgRowVals(item)
}

// Add a known rating to the training set
func (t* TrainingSet) addRating(user int, item int, value float64) {
	t.ratingsMatrix.set(value, item, user)
	t.trainingSize++
	t.sumRatings += value
}

// Add a new user to the training set
func (t *TrainingSet) addUser() {
	t.ratingsMatrix.addCol()
}

// Add a new item to the training set
func (t *TrainingSet) addItem() {
	t.ratingsMatrix.addRow()
}
