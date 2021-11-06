// @program:     pokemon
// @author:      sunsun

package game

import (
	"fmt"
	"pokemon/pokemon"
)



func Start(){
	fmt.Println("不太行的宝可梦233\n(对应数字对应选项)")
	fmt.Println("总之还是来挑选一只宝可梦吧")
	fmt.Println("现在就只有皮卡丘，这不快去多写点?")
	//如果重复用switch去对宝可梦进行选择，宝可梦多起来了怎么办呢?
	//可以尝试使用map[string]func返回对应的构造器
	pokemon:=pokemon.NewPikachu()
	fmt.Println("出来吧",pokemon.Name())

	x:=0

	for  {
		fmt.Println()
		fmt.Println("1.找个敌人打一架\n2.更换你的宝可梦\n3.等等等\n0.退出游戏")
		fmt.Scanln(&x)
		switch x {
		case 0:
			fmt.Println("游戏结束")
			return
		case 1:
			fight(pokemon)
		default:
			fmt.Println("其他功能当然你来写啊>_<")
		}
	}
}