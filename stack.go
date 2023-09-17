package main

// TODO: add err handling

type Stack []string

func (s *Stack) Push(ele string) {
	*s = append(*s, ele)
}

func (s *Stack) Peak() string {
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	temp := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return temp
}

func (s *Stack) Equals(s2 *Stack) bool {
	if len(*s) != len(*s2) {
		return false
	}
	for i := range *s {
		if (*s)[i] != (*s2)[i] {
			return false
		}
	}
	return true
}
