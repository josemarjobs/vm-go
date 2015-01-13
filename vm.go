package main

import "fmt"
	
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
	v.stack = make([]int, 100)
	v.sp = -1

	v.code = c
	v.pc = 0

	for {
		// Fetch
		op := v.code[v.pc]
		v.pc++

		// Decode
		switch op {
		case PUSH:
			val := v.code[v.pc]
			v.pc++

			v.sp++
			v.stack[v.sp] = val
		case ADD:
			a := v.stack[v.sp]
			v.sp--;
			b := v.stack[v.sp]
			v.sp--

			v.sp++
			v.stack[v.sp] = a+b
		case PRINT:
			val := v.stack[v.sp]
			v.sp--
			fmt.Println(val)
		case HALT:
			return;
		}
	}
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