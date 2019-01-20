package go_recommend_me

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type User struct {
	UserID 	   int
	Age 	   int
	Gender 	   string
	Occupation string
	ZipCode    int
}

type Rating struct {
	UserId		int
	ItemId		int
	Rating		int
	timestamp	int
}

type Movie struct {
	Id int
	Title	string
	Release	string
	IMDbUrl	string
	Genre	[]string


}

func main() {
	getMovieData()
}

func getMovieData() {
	file, err := os.Open("movielens_data/u.item")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringsSlice := strings.FieldsFunc(scanner.Text(), Split)
		movie := Movie{
			getInt(stringsSlice[0]),
			stringsSlice[1],
			stringsSlice[2],
			stringsSlice[3],
			getMovieGenre(stringsSlice[4:]),

		}

		fmt.Println(movie)
	}
}

func getUserData() {
	file, err := os.Open("movielens_data/u.user")
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stringsSlice := strings.FieldsFunc(line, Split)
		uID, _ := strconv.Atoi(stringsSlice[0])
		age, _ := strconv.Atoi(stringsSlice[1])
		zip, _ := strconv.Atoi(stringsSlice[4])
		u := User{uID, age, stringsSlice[2], stringsSlice[3], zip}

		fmt.Println(u)
	}
}

func getRatingData() {
	file, err := os.Open("movielens_data/u.data")
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stringsSlice := strings.Fields(line)
		rate := Rating{getInt(stringsSlice[0]),
			getInt(stringsSlice[1]),
			getInt(stringsSlice[2]),
			getInt(stringsSlice[3]),
		}

		fmt.Println(rate)

	}
}

func getMovieGenre(mapping []string) []string {
	mapGenres := make([]string, 0)

	genres := []string{"unknown", "Action", "Adventure", "Animation", "Children's", "Comedy", "Crime" , "Documentary",
		"Drama", "Fantasy", "Film-Noir", "Horror", "Musical", "Mystery", "Romance", "Sci-Fi", "Thriller", "War", "Western"}

	for i := 0; i < len(mapping); i++ {
		if getInt(mapping[i]) == 1 {
			mapGenres = append(mapGenres, genres[i])
		}
	}

	return mapGenres

}

func Split(r rune) bool {
	return r == '|'
}

func getInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}