package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// Distance 函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance 方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) //function call 包级别的函数调用
	fmt.Println(p.Distance(q))  // method call 调用的是Point类下声明的Point.Distance方法

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}

type Path []Point // slice类型
// Distance 给slice定义方法，事实上能够给任意类型定义方法
// 只要不是一个指针或者一个interface
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

//以上可以看出不同的类型可以有相同的方法名，且相同的类型必须有唯一的方法名
