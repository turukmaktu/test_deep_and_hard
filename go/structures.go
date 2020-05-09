package main

type resultApi struct {
	Status bool `json:"status"`
	Message string `json:"message"`
}

type eventApi struct {
	Id string `json:"id"`
	Label string `json:"label"`
}
