package main

import "fmt"

func main() {
		//var testArray [3]int                        //数组会初始化为int类型的零值
		//var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
		//var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
		//fmt.Println(testArray)                      //[0 0 0]
		//fmt.Println(numArray)                       //[1 2 0]
	//	//fmt.Println(cityArray)                      //[北京 上海 深圳]
	//var testArray [3]int
	//var numArray = [...]int{1, 2}
	//var cityArray = [...]string{"北京", "上海", "深圳"}
	//fmt.Println(testArray)                          //[0 0 0]
	//fmt.Println(numArray)                           //[1 2]
	//fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	//fmt.Println(cityArray)                          //[北京 上海 深圳]
	//fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string

	//// 声明切片类型
	//var a []string              //声明一个字符串切片
	//var b = []int{}             //声明一个整型切片并初始化
	//var c = []bool{false, true} //声明一个布尔切片并初始化
	////var d = []bool{false, true} //声明一个布尔切片并初始化
	//fmt.Println(a)              //[]
	//fmt.Println(b)              //[]
	//fmt.Println(c)              //[false true]
	//fmt.Println(a == nil)       //true
	//fmt.Println(b == nil)       //false
	////fmt.Println(c == nil)       //false
	//a := make([]int, 2, 10)
	//fmt.Println(a)      //[0 0]
	//fmt.Println(len(a)) //2
	//fmt.Println(cap(a)) //10
	//
	//c := make([]int, 3,200)
	//fmt.Println(c)
	//fmt.Println(len(c))
	//fmt.Println(cap(c))
	//
	//var numSlice []int
	//for i := 0; i < 10; i++ {
	//	numSlice = append(numSlice, i)
	//	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	//}
	//scoreMap := make(map[string]int, 8)
	//scoreMap["张三"] = 90
	//scoreMap["小明"] = 100
	//fmt.Println(scoreMap)
	//fmt.Println(scoreMap["小明"])
	//fmt.Printf("type of a:%T\n", scoreMap)
	//
	//cityMap := make(map[string]string)
	//cityMap["王五"] = "ni"
	//cityMap["zhao6liu"] = "ca"
	//fmt.Println(cityMap)
	//userInfo := map[string]string{
	//	"username": "沙河小王子",
	//	"password": "123456",
	//}
	//fmt.Println(userInfo) //
	//
	//nieInfo := map[string]string{
	//	"ca":"ca",
	//	"caa":"das",
	//}
	//fmt.Println(nieInfo)
	//for k,v  := range nieInfo{
	//	fmt.Println(k,v)
	//}
	//delete(nieInfo,"ca")
	//for k,v  := range nieInfo{
	//	fmt.Println(k,v)
	//}
	//var mapSlice = make([]map[string]string, 3)
	//for index, value := range mapSlice {
	//	fmt.Printf("index:%d value:%v\n", index, value)
	//}
	//fmt.Println("after init")
	//// 对切片中的map元素进行初始化
	//mapSlice[0] = make(map[string]string, 10)
	//mapSlice[0]["name"] = "小王子"
	//mapSlice[0]["password"] = "123456"
	//mapSlice[0]["address"] = "沙河"
	//for index, value := range mapSlice {
	//	fmt.Printf("index:%d value:%v\n", index, value)
	//}
	//a := new(int)
	//b := new(bool)
	//fmt.Printf("%T\n", a) // *int
	//fmt.Printf("%T\n", b) // *bool
	//fmt.Println(*a)       // 0
	//fmt.Println(*b)       // false
	//
	//c := new(string)
	//fmt.Println("%T\n",c)
	//fmt.Println(*c)

	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)

	f := 100
	ca := &f
	fmt.Println("a:%d ca: %p\n",f,&f)
	fmt.Println("b:%p, type:%T\n",ca,ca)


}
