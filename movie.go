package main

type Movie struct {
	Title       string   `json:"title"`
	Poster      string   `json:"poster"`
	ReleaseYear int      `json:"release_year"`
	Actors      []string `json:"actors"`
	SimilarIDs  []string `json:"similar_ids"`
}

type Movies []Movie
