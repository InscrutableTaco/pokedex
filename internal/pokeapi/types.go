package pokeapi

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Response struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}
