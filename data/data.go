package data

// Game Model
type Game struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
	Rating    int    `json:"rating"`
	Sales     int    `json:"sales"`
}
