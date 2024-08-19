package main

import (
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

// Distance 方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type ColoredPoint struct {
	Point //简写形式，并使其包含Point类型所具有的一切字段
	Color color.RGBA
}
type ColoredPoints struct {
	*Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	//当作自身的字段使用，完全不需要在调用时指出Point
	cp.X = 1
	cp.Y = 2
	red := color.RGBA{R: 255, A: 255}
	blue := color.RGBA{B: 255, A: 255}
	// 也可以直接把ColoredPoint类型当作接收器来调用Point里的方法
	cp.Distance(Point{1, 1})
	var p = ColoredPoint{Point{1, 1}, color.RGBA{R: 255, G: 255, B: 255}}
	var Q = ColoredPoint{Point{5, 4}, color.RGBA{R: 255, G: 255, B: 255}}
	// Color Point当作接收器调用 Point里的方法，里面传入的必须是Point
	p.Distance(Q.Point)
	//传入ColorPoint 则报错
	p.Distance(Q)

	// 匿名字段是指针,字段和方法会被间接地引入到类型中
	redP := ColoredPoints{&Point{1, 1}, red}
	blueP := ColoredPoints{&Point{5, 4}, blue}
	redP.Distance(*blueP.Point)
}

// 一个struct类型可能会有多个匿名字段，该类型会拥有Point和RGBA的所有方法，以及直接定义到ColoredPoint中的方法。
type ColoredPointss struct {
	Point
	color.RGBA
}

//p.ScaleBy -》编译器 先找CPss，然后找Point和RGBA,然后一直递归找下去。若有二义性，编译器会报错
