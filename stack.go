package main

import "errors"

type stack []int

func CreateStack() *stack {
	var st stack = make(stack, 255)

	return &st
}

func (items *stack) Len() int {
	return len(*items)
}

func (items *stack) Pop() int {

	result := (*items)[items.Len()-1:]
	*items = stack((*items)[:items.Len()-1])

	return result[0]
}

func (items *stack) popn(number int) []int {

	result := (*items)[items.Len()-number:]
	*items = stack((*items)[:items.Len()-number])

	return result
}

func (items *stack) push(item int) (result bool, err error) {
	if items.Len() < 255 {
		*items = append(*items, item)
		result, err = true, nil
	} else {
		result, err = false, errors.New("Stack is full!")
	}

	return
}
