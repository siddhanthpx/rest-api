// Package classification Games API.
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package data

import "github.com/go-playground/validator"

// A list of games returned in the response
// swagger:response gamesResponse
type gamesResponse struct {
	// All products in the database
	// in: body
	Body []Game
}

// Game Model
type Game struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" validate:"required"`
	Publisher string `json:"publisher" validate:"required"`
	Rating    int    `json:"rating" validate:"required,gte=0,lte=10,number"`
	Sales     int    `json:"sales" validate:"number"`
}

// Validate validates the game object
// using the struct tags associated with it
func (g *Game) Validate() error {

	validate := validator.New()
	return validate.Struct(g)

}
