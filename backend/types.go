package main

type Response struct {
	Data any `json:"data"`
}

type UpdateCounterRequest struct {
	NewVal int `json:"newVal"`
}
