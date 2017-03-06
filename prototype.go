package learn_pattern

import (
	"fmt"
)

type WorkExperience struct {
	timeArea string
	company  string
}

func (w *WorkExperience) getWorkDate() string {
	return w.timeArea
}
func (w *WorkExperience) setWorkDate(timeArea string) {
	w.timeArea = timeArea
}
func (w *WorkExperience) getCompany() string {
	return w.company
}
func (w *WorkExperience) setCompany(company string) {
	w.company = company
}

type Resume struct {
	name string
	sex  string
	age  string
	WorkExperience
}

func (r *Resume) setPersonalInfo(name, sex, age string) {
	r.name = name
	r.age = age
	r.sex = sex
}
func (r *Resume) setWorkExperience(timeArea, company string) {
	r.company = company
	r.timeArea = timeArea
}
func (r *Resume) display() {
	fmt.Println(r.name, r.sex, r.age)
	fmt.Println("工作经历：", r.timeArea, r.company)
}
func (r *Resume) clone() *Resume {
	new_obj := (*r)
	return &new_obj
}
