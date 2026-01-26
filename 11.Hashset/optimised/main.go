package main

import (
	"errors"
	"fmt"
)

type Set struct {
	data map[int]int
}

func (s *Set) add(val int) {
	s.data[val]++

}
func (s *Set) remove(val int) error {
	_, ok := s.data[val]
	if !ok {
		return errors.New("removed item not present")
	}
	delete(s.data, val)
	return nil
}
func (s *Set) print() {
	for val, _ := range s.data {
		fmt.Println(val)
	}

}
func main() {
	myHashSet := Set{data: make(map[int]int)}
	myHashSet.add(10)
	myHashSet.add(10)
	myHashSet.add(-20)
	myHashSet.add(30)
	myHashSet.add(40)
	myHashSet.add(-50)
	myHashSet.remove(50)
	myHashSet.print()

}
