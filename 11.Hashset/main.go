package main

import (
	"errors"
	"fmt"
)

type Set struct {
	data []int
}

func (s *Set) add(val int) {

	for _, v := range s.data {
		if v == val {
			return
		}
	}
	s.data = append(s.data, val)
}
func (s *Set) remove(val int) error {
	index := -1
	for i, v := range s.data {
		if v == val {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("removed item not present")
	}
	s.data = append(s.data[:index], s.data[index+1:]...)
	return nil

}
func main() {
	myHashSet := Set{data: []int{}}
	myHashSet.add(10)
	myHashSet.add(10)
	myHashSet.add(-20)
	myHashSet.add(30)
	myHashSet.add(40)
	myHashSet.add(-50)
	myHashSet.remove(50)
	fmt.Println(myHashSet.data)

}
