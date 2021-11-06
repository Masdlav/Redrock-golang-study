package main

//视频详情页结构体

import (
	"fmt"
)

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

//又是一个要填的坑....要用切片实现这个推荐结构体
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
func main(){
	WhyiPhone := Video{
		VideoAuthor: Author{
			Name: "我的宇宙_Yoga",
			Vip: true,
			Icon: "Yoga's universe",
			Signature: "请用理科生的脑子，造出文科生的产品。 欢迎来到，我的宇宙。YouTube: Yoga's Universe 商务合作V：Yoga_Works",
			Focus: 122000,
		},
		VideoDetails: Details{
			BV: "BV15r4y1176q",
			Headline: "【杜比】「我的宇宙」你为什么还在用 iPhone？",
			Duration: 1230,
			Intro: "这是「我的宇宙」的第四十个视频，希望大家喜欢。\nYouTube ：Yoga's Universe",
			ReleaseTime: "2021-10-30 10:00:07",
			Label: "iPhone科技数码\n科技猎手\n数码\n苹果\n生态\n安卓\n电影效果\n13\n统一\n120Hz\n杜比视界\n",
		},
		VideoData: Data{
			ViewCounts: 44000,
			BulletScreenCounts: 924,
			LikeNum: 3867,
			CoinNUm: 2911,
			CollectionNum: 1115,
			ForwardNum: 250,
			CommentNum: 802,
		},
	}
	fmt.Print(WhyiPhone)
}