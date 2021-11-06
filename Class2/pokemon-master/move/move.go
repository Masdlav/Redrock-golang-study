// @program:     pokemon
// @author:      sunsun

package move

type Move interface {
	MoveName() string
	MoveHurt() int
	MoveCost() int
}

//BaseMove
//招式嘛可以扩展的东西也很多,比如附带属性像麻醉睡眠啊，或则招式cd啊等等等等
//还有这里是比较简单粗暴的直接返回一个招式伤害，那有没有其他的可行流程呢，比如把对面的宝可梦写入我的方法等等
//然后呢，可以看到我这里的字段都是小写的，为什么呢，因为招式不像宝可梦状态，招式不应该在游戏过程中变动，所以我不希望你通过特殊的手段来改变它，这也就是我们上课提到的闭包的概念
//什么？你说招式万一可以互相作用？自己重构吧hh
type BaseMove struct {
	moveName string
	moveHurt int
	moveCost int
}

func NewMove(moveName string, moveHurt, moveCost int) Move {
	return &BaseMove{
		moveName: moveName,
		moveHurt: moveHurt,
		moveCost: moveCost,
	}
}

func (bm *BaseMove) MoveName() string {
	return bm.moveName
}

func (bm *BaseMove) MoveHurt() int {
	return bm.moveHurt
}

func (bm *BaseMove) MoveCost() int {
	return bm.moveCost
}