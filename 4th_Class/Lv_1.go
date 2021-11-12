package main

import "fmt"

func main() {
	skill:=""
	fmt.Println("请输入你要释放的技能！！！！！！")
	fmt.Scanln(&skill)
	moban := 0
	fmt.Println("就决定是你了！模板！（请用模板的序号选择模板）")
	fmt.Println("1.尝尝我的厉害吧！<技能>")
	fmt.Println("2.尝尝这囚禁了一万年的愤怒！<技能>")
	fmt.Println("3.你们这是...自寻死路！<技能>")
	fmt.Println("4.大的要来了！大的要来了！<技能>")
	fmt.Scanln(&moban)
	switch moban {
	case 1 : ReleaseSkill(skill, func(skillName string) {
			fmt.Println("尝尝我的厉害吧！", skillName)
	})
	case 2 : ReleaseSkill(skill, func(skillName string){
		fmt.Println("尝尝这囚禁了一万年的愤怒！",skillName)
	})
	case 3 : ReleaseSkill(skill, func(skillName string){
		fmt.Println("你们这是...自寻死路！",skillName)
	})
	case 4 : ReleaseSkill(skill, func(skillName string) {
		fmt.Println("大的要来了！大的要来了！",skillName)
	})
	default:
		fmt.Println("啊...好像没有这个模板")
	}
}

func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}