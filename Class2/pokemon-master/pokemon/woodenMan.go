// @program:     pokemon
// @author:      sunsun

package pokemon

import (
	"pokemon/move"
)

//喜闻乐见的不会还手的木头人
//怎么让他动起来?
type WoodenMan struct {
	BasePokemon
}

func NewWoodenMan()Pokemon{
	return &BasePokemon{
		name:  "wooden man",
		hp:    300,
		mp:    0,
		moves: [4]move.Move{
			move.NewMove("只需要1招",20,0),
		},
	}
}