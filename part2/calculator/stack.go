package calculator

type Stack struct {
	data []int32
}

type IStack interface {
	Push(s int32)
	Top() int32
	Pop()
	Length() int
}

func NewStack() IStack {
	return &Stack{}
}

func (st *Stack) Push(s int32) {
	st.data = append(st.data, s)
}

func (st *Stack) Top() int32 {
	if len(st.data) == 0 {
		return -1
	}
	return st.data[len(st.data) - 1]
}

func (st *Stack) Pop() {
	st.data = st.data[:len(st.data) - 1]
}

func (st Stack) Length() int {
	return len(st.data)
}
