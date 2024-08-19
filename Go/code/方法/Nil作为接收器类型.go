package main

// IntList 是一个int的链表
// nil List 代表空表
type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
func main() {

}
