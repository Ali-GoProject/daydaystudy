package main
import "fmt"

// 定义个接口，有两个方法Get和Set
type Simpler interface {
	Get() int
	Set(int)
}

// 定义一个Simple的结构体
type Simple struct {
	value int
}


// 这个结构体实现了这个接口
func (value Simple) Get() int{
	return value.value
}

func (value *Simple) Set(v int) {
	value.value = v
}


//测试函数，形参是个接口，只要传递相应的类型就执行相应的方法，一种C++中多态的思想
func TestSimpler(inter Simpler) {
	inter.Set(10)
	fmt.Println(inter.Get())
}


func main() {
	s := Simple{}
	TestSimpler(&s)
}
