package model

type Classification struct {
	Name string
}

type SubClassification struct {
	Classification string
	Labels         string
	Hint           string
}
