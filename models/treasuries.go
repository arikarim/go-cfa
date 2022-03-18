package models

// type Status int

// // create status enum
// const (
// 	Active Status = iota
// 	Inactive
// )

type Treasury struct {
	Id     string `json:"id"`
	NameEn string `json:"name_en"`
	Status string `json:"status"`
}

func (t *Treasury) GetTreasury() (Treasury, error) {
	return Treasury{}, nil
}