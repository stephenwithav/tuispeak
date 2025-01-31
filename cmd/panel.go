package main

type Board struct {
	Title     string   `yaml:"title"`
	Questions []string `yaml:"questions"`
}

type Boards struct {
	Boards []Board `yaml:"boards"`
}
