package models

var RandomValueTypes = []string{"dec", "hex", "str", "stralnum", "uuid"}

type RandomEntity struct {
	GenerationID    string
	RandomValue     string
	RandomValueType string
}
