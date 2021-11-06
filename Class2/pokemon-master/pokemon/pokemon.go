// @program:     pokemon
// @author:      sunsun

package pokemon

import (
	"fmt"

	"pokemon/move"
)

type Pokemon interface {
	Attack(other Pokemon)

	Name() string
	IsDead()bool
	ShowState()

	ReduceHP(num int)
	ReduceMP(num int)
}

//为了大家都能轻松看懂，毕竟我们是考察接口不是考察宝可梦，这里弱化了相关属性，
//你想让它更“原滋原味”比如搞搞什么种族值啊,异常状态啊之类的就自己重构吧
type BasePokemon struct {
	name  string
	hp    int          //血量
	mp    int          //蓝量
	moves [4]move.Move //一个宝可梦可以记住四个招式
}

func (bp *BasePokemon) Attack(other Pokemon) {
	fmt.Println("请选择你的宝可梦想要发动的招式")
	fmt.Println("招式列表 hurt cost")
	fmt.Println("0.",bp.moves[0].MoveName(),"\n","1.",bp.moves[1].MoveName(),"\n","2.",bp.moves[2].MoveName(),"\n","3.",bp.moves[3].MoveName())

	x := 0
	fmt.Scanln(&x)
	move := bp.moves[x]
	fmt.Println(bp.Name(),"对",other.Name(),"使用了", move.MoveName())
	other.ReduceHP(move.MoveHurt())
	bp.ReduceMP(move.MoveCost())
}

func (bp *BasePokemon) Name() string {
	return bp.name
}

func (bp *BasePokemon) IsDead()bool  {
	return bp.hp<=0
}

func (bp *BasePokemon)ShowState(){
	fmt.Println(bp.name+" HP: ",bp.hp," mp: ",bp.mp)
}

func (bp *BasePokemon) ReduceHP(num int) {
	bp.hp -= num
}

func (bp *BasePokemon) ReduceMP(num int) {
	bp.mp -= num
}
