### GO的环境

* GOROOT：

### 语言结构

* package main

* import “fmt”

* func main(){

  ​	i := 1
  }

### 关键字

break、default、func、interface、select

case、defer、go、map、struct

chan、else、goto、package、switch

const、fallthrough、if、range、type

continue、for、import、return、var

  ### 预定义标识符

基本类型：

* bool、error

* byte、int、int8、int16、int32、int64、rune、uint、uint8、uint16、uint32、uint64、uintptr

* string

* complex64、complex128、float32、float64

常量：

  * true、false、iota

函数：

* append、cap、close、complex、copy、delete、imag、len、make、new、print、println、real
* panic、recover （





## 数组、slice、Map、结构体、json、文本和html模板



## 并发

### 通道的使用

1. 创建一个通道，无缓冲区/大小为1

   * 创建int类型

   ```go
   ch := make(chan int, 1)
   ```
   * 创建int类型

   ```go
   ch := make(chan )
   ```

   

## 方法和接口 Methods and interfaces

Go没有类。可以为**类型** 定义**方法**。

方法是一类带特殊的 **接收者**【在func和方法名】 参数的方法.

```go
type Vertex struct{
    X, Y float64
}
func (v Vertex) Abs() float64{
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

Abs方法拥有一个名为`v`，类型为`Vertex`的接收者 

abs:绝对值

方法只是个带接收者参数的函数

**接收者** 的类型定义和方法声明必须在同一包内【内置类型也要在包内再次定义】

```go
type MyFloat float64

func (f MyFloat) Abs() float64{
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}
```

## Strconv包

string与int 的转换 ：

Atoi(ASCII TO INT)、Itoa(int to ascii)

Parse系列函数

ParseBool、ParseInt、ParseUint、ParseFloat

Format函数

FormatBool、FormatInt、FormatUint、FormatFloat
