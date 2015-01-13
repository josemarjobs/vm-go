package main
	
const (
	PUSH = iota
	ADD
	PRINT
	HALT
)
type VM struct {
	code []int
	pc int

	stack []int
	sp int
}

func (v *VM) run(c []int) {
	
}
func main() {
	code := []int{
		PUSH, 2,
		PUSH, 3,
		ADD,
		PRINT,
		HALT,
	}

	v := &VM{}
	v.run(code)

}