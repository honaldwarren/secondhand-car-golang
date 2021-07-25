package dto

type CarFilterDto struct {
	Make     string
	Model    string
	MinYear  int
	MaxYear  int
	MinPrice int
	MaxPrice int
}
