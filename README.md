GRM 
=======================


GRM or "Go Recommend Me" is a recommender system library built in Go.

GRM analyzes the feedback of some users and their 
preferences for some items. It learns patterns and predicts the most suitable products 
for a particular user.

Features
--------
 * Latent Factor Approach
 * Collaborative Filtering
 * Basic Matrix Factorization (Looking to add more models in near future)
 * User and Item based recommenders
 * No external dependencies 
 


### Example


```Go
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


fmt.Println("user [1] item[0], rating = ", model.EstimateRating(1, 0, learned))
fmt.Println("user [1] item[1], rating = ", model.EstimateRating(1, 1, learned))
fmt.Println("user [1] item[2], rating = ", model.EstimateRating(1, 2, learned))


```

You can run this example by downloading the repository and running:
```Go
go run tests/test.go
```

References
--------
1. http://en.wikipedia.org/wiki/Recommendation_system
2. https://en.wikipedia.org/wiki/Sparse_matrix
3. https://mathoverflow.net/questions/14412/matrix-factorization-model
4. http://www.quuxlabs.com/blog/2010/09/matrix-factorization-a-simple-tutorial-and-implementation-in-python/

