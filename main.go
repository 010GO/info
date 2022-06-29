//学生信息管理系统
//需求：
//1.添加学生信息
//2.编辑学生信息
//3.删除学生信息
//4.展示学生信息

package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("*********************************")
	fmt.Println("欢迎进入学生信息管理系统")
	fmt.Println("1.添加学生")
	fmt.Println("2.编辑学生信息")
	fmt.Println("3.删除学生")
	fmt.Println("4.展示所有学生信息")
	fmt.Println("5.退出系统")
	fmt.Println("*********************************")
}

func getStudentInfo() (stu *student) {
	var (
		name  string
		score float32
		id    int
		age   int
	)

	fmt.Println("按以下提示输入用户信息")
	// fmt.Println("请输入学生姓名：")
	// fmt.Scanf("%s\n", &name)
	// fmt.Println("请输入学生分数：")
	// fmt.Scanf("%f\n", &score)
	// fmt.Println("请输入学生学号：")
	// fmt.Scanf("%d\n", &id)
	// fmt.Println("请输入学生年龄：")
	// fmt.Scanf("%d\n", &age)
	fmt.Println("请依次输入学生姓名、学号、年龄、分数：")
	fmt.Scanf("%s %d %d %f\n", &name, &id, &age, &score)
	stu = newStudent(name, score, id, age)
	return
}

func main() {
	var s = newMgr()
	for {
		//供用户选择的菜单栏
		showMenu()
		//等待用户要选择执行的选项
		var input int
		fmt.Println("请输入您要选择的序号:")
		fmt.Scanf("%d\n", &input)
		fmt.Println("用户输入的是", input)

		//执行用户选择的动作
		switch input {
		//1.添加学生信息
		case 1:
			stu := getStudentInfo()
			s.addStudent(stu)
		//2.编辑学生信息
		case 2:
			stu := getStudentInfo()
			s.modifyStudent(stu)
		//3.删除学生信息
		case 3:
			var id int
			fmt.Println("请输入删除学生学号")
			fmt.Scanf("%d\n", &id)
			s.deleteStudent(id)
		//4.展示学生信息
		case 4:
			s.showStudent()
		//退出系统
		case 5:
			os.Exit(0)
		}
	}
}
