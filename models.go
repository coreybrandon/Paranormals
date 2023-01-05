package main

type Product struct {
	Name      string `json:"name"`
	Mind      string `json:"mind,omitempty"`
	Character string `json:"character,omitempty"`
	Thoughts  string `json:"thoughts,omitempty"`
	Feelings  string `json:"feelings,omitempty"`
}

type Name struct {
	GivenName string
	Surname   string
}
