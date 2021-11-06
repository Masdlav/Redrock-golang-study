package main

import "fmt"

//为视频详情页实现三连方法

type Author struct {
	Name string
	Vip bool
	Icon string
	Signature string
	Focus int
}

type Details struct {
	BV string
	Headline string
	Duration int
	Intro string
	ReleaseTime string
	Label string
}

type Data struct {
	ViewCounts int
	BulletScreenCounts int
	LikeNum int
	CoinNUm int
	CollectionNum int
	ForwardNum int
	CommentNum int
}

type Recommend struct {
	Details
	Author
}

type Video struct {
	VideoAuthor Author
	VideoDetails Details
	VideoData Data
	VideoRecommend Recommend
	IsLiked bool
	IsCollected bool
	IsCoined bool  //布尔值未初始化默认值为false
}

func (a *Video)like()  {
	a.IsLiked = true
}

func (a *Video)collect()  {
	a.IsCollected = true
}

func (a *Video)coin()  {
	a.IsCoined = true
}

func (a *Video)triplet()  {
	a.like()
	a.collect()
	a.coin()
}

func main()  {
	a := Video{
		IsLiked: false,
		IsCoined: false,
		IsCollected: false,
	}
	a.like()
	a.coin()
	a.triplet()
	fmt.Print(a)
}

