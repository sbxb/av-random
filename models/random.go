package models

type RandomEntity struct {
	GenerationID string
	RandomValue  int64 // TODO Let us use a simple built-in type for now
}

func (re RandomEntity) IsEmpty() bool {
	return re.GenerationID == "" && re.RandomValue == 0
}
