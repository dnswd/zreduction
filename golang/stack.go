package main

// Adapted from https://stackoverflow.com/a/28542256
type stack struct {
     s []int
}

func NewStack() *stack {
    return &stack{make([]int,0)}
}

func (s *stack) Push(v int) {
    s.s = append(s.s, v)
}

func (s *stack) Pop() int {
    l := len(s.s)
    if l == 0 {
        return -1
    }

    res := s.s[l-1]
    s.s = s.s[:l-1]
    return res
}

func (s *stack) Size() int {
    return len(s.s)
}

func (s *stack) Clear() {
    s.s = nil
}

func (s *stack) Top() int {
    l := len(s.s)
    return s.s[l-1]
}