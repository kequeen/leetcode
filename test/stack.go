package leetcode

//用切片来模拟栈
type Stack struct {
	arr []int
}

func (s *Stack) Push(value int) {
	s.arr = append(s.arr, value)
}

func (s *Stack) Pop() int {
	value := s.arr[len(s.arr)-1]
	s.arr = s.arr[0 : len(s.arr)-1]
	return value
}

func (s *Stack) GetLen() int {
	return len(s.arr)
}
