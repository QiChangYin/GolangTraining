package main

type People interface {
	ReturnName() string
}

type Student struct {
	Name string
}

type Teacher struct {
	Name string
}


func (s Student) ReturnName() string{
	return s.Name
}
func (t *Teacher) ReturnName() string{
	return t.Name
}

func main(){
	cbs := Student{Name:"学生的咖啡色"}
	cbt := Teacher{Name:"老师的咖啡色"}
	var a People
	a = cbs
	println(a.ReturnName())
	a = &cbt
	println(a.ReturnName())
}