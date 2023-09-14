package main

type Stack []string

func (s *Stack) Push(ele string){
	*s = append(*s, ele)
}

func (s *Stack) Peak() string{
	return (*s)[len(*s) - 1]
}

func (s *Stack) Pop() string{
	if len(*s) == 0 {
		return  ""
	}
	temp := (*s)[len(*s) - 1]
	*s = (*s)[0:len(*s) - 1]
	return temp
}
