package main

import "fmt"

type student struct {
	name  string
	score float32
	id    int
	age   int
}

func newStudent(name string, score float32, id, age int) *student {
	//以下是结构体两种方式的实例化，可对比学习

	return &student{ //键值对的形式初始化结构体
		name:  name,
		score: score,
		id:    id,
		age:   age,
	}

	// s = new(student) //一定要new()先分配内存给结构体指针，才可使用
	// s.name, s.score, s.id, s.age = name, score, id, age
	// return
}

func newMgr() *studentMgr {
	return &studentMgr{
		mgr: make([]*student, 0, 100),
	}
}

type studentMgr struct {
	mgr []*student
}

//添加学生
func (s *studentMgr) addStudent(stu *student) {
	s.mgr = append(s.mgr, stu)
}

//编辑学生信息
func (s *studentMgr) modifyStudent(stu *student) {
	for k, v := range s.mgr {
		if stu.id == v.id {
			s.mgr[k] = stu
			return
		}
	}
	fmt.Printf("不存在学号为%d的学生\n", stu.id)
}

//删除学生信息
func (s *studentMgr) deleteStudent(id int) {
	for k, v := range s.mgr {
		if id == v.id {
			s.mgr = append(s.mgr[:k], s.mgr[k+1:]...)
			return
		}
	}
	fmt.Printf("不存在学号为%d的学生\n", id)
}

//展示所有学生信息
func (s *studentMgr) showStudent() {
	if len(s.mgr) == 0 {
		fmt.Printf("数据库为空，不存在学生信息\n")
	}
	for _, v := range s.mgr {
		fmt.Printf("学号：%d 姓名：%s 年龄：%d 分数：%f\n", v.id, v.name, v.age, v.score)
	}
}
