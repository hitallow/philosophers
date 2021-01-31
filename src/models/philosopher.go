package models

// Philosophers : List a philosophers
type Philosophers struct {
	Philosophers []Philosopher `json:"philosophers"`
}

// Philosopher : a philosopher entity
type Philosopher struct {
	Name  string `json:"name"`
	Phrases []string
}