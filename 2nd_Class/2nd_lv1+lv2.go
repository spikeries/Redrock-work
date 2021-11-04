package main

import "fmt"

type Author struct {
	Name        string
	Level       int
	Sex         string
	Fans        int
	Signature   string
	BigVIP      bool
	Goodlooking bool
}
type Videoinfo struct {
	Title        string
	Watch        int
	Coin         int
	Collection   int
	Like         int
	Forward      int
	Introduction string
	Bullet       int
	Time       float32
}
type Re struct {
	Auname    string
	Title string
	Bulletnum int
	Watch     int
}

type Recommend []Re
type Video struct {
	Author
	Videoinfo
	Recommend
}

func main() {
	Video1 := Video{
		Author{
			Name:        "雁巡",
			Level:       6,
			Sex:         "Male",
			Fans:        201000,
			Signature:   "背包装满了家用，路就这样开始走。",
			BigVIP:      false,
			Goodlooking: false,
		},
		Videoinfo{
			Title:"《杀死那个石家庄人》",
			Watch:        14727000,
			Coin:         450000,
			Collection:   369000,
			Like:         621000,
			Forward:      78000,
			Introduction: "-",
			Bullet:       1000,
		},
		Recommend{
			Re{
				Title: "反方向的钟 | 回到当初爱你的时空",
				Auname:    "Yige一格",
				Bulletnum: 228,
				Watch:     286000,
			},
			Re{
				Auname:    "蔡子尤",
				Title:"《宿舍加州旅馆》原先以为是搞笑的，看了直接跪拜了",
				Bulletnum: 54000,
				Watch:     16524000,
			},
			Re{
				Auname:    "雁巡",
				Title:     "《山海》——来对曾经的自己道个别吧。",
				Bulletnum: 3018,
				Watch:     2559000,
			},
			Re{
				Auname:    "哦呼w",
				Title:     "敢 杀 我 的 马？！",
				Bulletnum: 173000,
				Watch:     60887000,
			},
			Re{
				Auname:    "音乐兄弟MTL",
				Title:     "开口跪|《See You Again》油管4亿播放组合红岩谷绝美翻唱",
				Bulletnum: 13000,
				Watch:     8533000,
			},
		},
	}
	Video1.like()
fmt.Println(Video1.Videoinfo.Like)
	v1:=fabu("aaa","bbb")
	fmt.Println(v1)
y(&Video1)
	fmt.Println(Video1)
}
func y(V yjsl)  {
	V.like()
	V.collect()
	V.coin()
}
 type yjsl interface{
	like()
	collect()
	coin()
}
func(v *Video)like(){
	v.Videoinfo.Like++
}
func(v *Video)collect(){
	v.Videoinfo.Collection++
}
func(v *Video)coin(){
	v.Videoinfo.Coin++
}
func(v *Video)yijiansanlian(){
	v.Coin++
	v.Like++
	v.Collection++
}
func fabu(Authorname,Videoname string)Video{
	v:=Video{
		Author:    Author{
			Name: Authorname,
		},
		Videoinfo: Videoinfo{
			Title: Videoname,
		},
	}
	return v
}
