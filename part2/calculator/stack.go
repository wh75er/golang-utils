package calculator

type Stack struct {
	data []interface{}
}

type IStack interface {
	Push(s interface{})
	Top() interface{}
	Pop()
	Length() int
}

func NewStack() IStack {
	return &Stack{}
}

func (st *Stack) Push(s interface{}) {
	st.data = append(st.data, s)
}

func (st *Stack) Top() interface{} {
	if len(st.data) == 0 {
		return -1
	}
	return st.data[len(st.data)-1]
}

func (st *Stack) Pop() {
	st.data = st.data[:len(st.data)-1]
}

func (st Stack) Length() int {
	return len(st.data)
}
