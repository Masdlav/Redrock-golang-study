// @program:     pokemon
// @author:      sunsun

package pokemon

import (
	"pokemon/move"
)

type Pikachu struct {
	BasePokemon
}

func NewPikachu() Pokemon {
	return &BasePokemon{
		name: "Pikachu",
		hp:   300,
		mp:   100,
		moves: [4]move.Move{
			move.NewMove("十万伏特", 95, 15),
			move.NewMove("电光一闪", 40, 30),
			move.NewMove("打倒", 80, 20),
			move.NewMove("雷电", 120, 10),
		},
	}
}
