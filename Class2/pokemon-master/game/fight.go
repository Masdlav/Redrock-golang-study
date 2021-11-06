// @program:     pokemon
// @author:      sunsun

package game

import (
	"fmt"

	"pokemon/pokemon"
)

func fight(me pokemon.Pokemon){
	other:=pokemon.NewWoodenMan()
	for {
		me.Attack(other)
		me.ShowState()
		other.ShowState()
		if other.IsDead(){
			fmt.Println("赢了好欸")
			return
		}
		//该怎么调度使木头人也会攻击我?
	}
}