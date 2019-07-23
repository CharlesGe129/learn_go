package archive

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() []int {
	seq := []int{1, 1}
	return func() []int {
		l := len(seq)
		seq = append(seq, seq[l-1] + seq[l-2])
		fmt.Println(seq)
		return seq
	}
}

func main() {
	//pos := adder()
	fib := fibonacci()
	for i := 0; i < 10; i ++ {
		//fmt.Println(pos(i))
		fib()
	}
}
