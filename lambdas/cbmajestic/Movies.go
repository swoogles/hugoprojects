package cbmajestic

type Rating struct {
	Source string
	Value string
}

type Movie struct {
	Title string
	Rated string
	Runtime string
	Genre string
	Director string
	Plot string
	Poster string
	Ratings []Rating
}

