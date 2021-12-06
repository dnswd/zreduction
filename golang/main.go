package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	T    string
	t    int
	st   *stack
	Pals []int
	w    strings.Builder
)

func deletion(j int, i int) {
	temp := w.String()[:w.Len()-i+j-1]
	w.Reset()
	w.WriteString(temp)
	Pals = Pals[:len(Pals)-i+j-1]
}

func slowExtend(c int) bool {
	r := w.Len() - 2 - c
	for w.String()[c+r+1] == w.String()[c-r] {
		r++
		if Pals[c-r] >= r {
			deletion(c-r+1, c+r)
			return true
		}
		Pals[c+r] = Pals[c-r]
		t++
		w.WriteByte(T[t])
		Pals = append(Pals, -1)
	}
	Pals[c] = r
	return false
}

func fastExtend(c int) bool {
	for st.Size() != 0 {
		r := st.Top() - c
		if Pals[c-r] >= Pals[c+r] {
			st.Pop()
		} else {
			Pals[c] = r + Pals[c-r]
			return false
		}
	}
	return true
}

func stabilize(c int) bool {
	b := w.Len() - 1
	unstable := true
	for unstable {
		unstable = false
		if slowExtend(c) {
			return true
		}
		for d := c + Pals[c]; d >= b; d-- {
			if d+Pals[d] >= c+Pals[c] && fastExtend(d) && stabilize(d) {
				if c == (w.Len() - 1) {
					return true
				}
				if d == w.Len()-1 {
					Pals[d] = Pals[2*c-d]
				}
				t++
				w.WriteByte(T[t])
				Pals = append(Pals, -1)
				unstable = true
				break
			}
			st.Push(d)
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	T, _ = reader.ReadString('\n')
	T = "$" + T + "#"
	t = 0
	st = NewStack()
	Pals = make([]int, 0)

	w.WriteByte(T[t])
	Pals = append(Pals, 0)
	for t < len(T)-1 {
		t++
		w.WriteByte(T[t])
		Pals = append(Pals, -1)
		stabilize(w.Len() - 2)
		for st.Size() != 0 {
			st.Pop()
		}
	}
	fmt.Println(w.String()[1 : w.Len()-1])
}
