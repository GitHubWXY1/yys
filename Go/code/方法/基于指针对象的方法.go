package main

import "fmt"

type Point struct {
	X, Y float64
}

// ScaleBy 使用指针。当需要更新一个变量，或者其中的参数太大希望避免这种拷贝时
// 指针类型 的方法，则该类型的方法都应是指针类型，即使其他方法可能不会更改参数值
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
func main() {
	//使用指针类型的方法
	// 01
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)
	// 02
	p := Point{1, 2}
	(&p).ScaleBy(2)
	//03 编译器隐式调用&p
	p.ScaleBy(2)
	//04 临时变量的地址无法获取到 报错
	Point{1, 2}.ScaleBy(2)

}
