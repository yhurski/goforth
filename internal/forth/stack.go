package forth

import "errors"

type stack []int

func CreateStack() *stack {
	var st stack = make(stack, 0, 255)

	return &st
}

func (items *stack) Len() int {
	return len(*items)
}

func (items *stack) Pop() (int, error) {
	if items.Len() == 0 {
		return -1, errors.New("Stack is empty")
	}

	result := (*items)[items.Len()-1:]
	*items = stack((*items)[:items.Len()-1])

	return result[0], nil
}

func (items *stack) Popn(number int) ([]int, error) {
	if items.Len() < number {
		return *items, errors.New("Stack doesn't have enough elements")
	}

	result := (*items)[items.Len()-number:]
	*items = stack((*items)[:items.Len()-number])

	return result, nil
}

func (items *stack) Push(item int) (result bool, err error) {
	if items.Len() < 255 {
		*items = append(*items, item)
		result, err = true, nil
	} else {
		result, err = false, errors.New("Stack is full!")
	}

	return
}

func (items *stack) Get(index int) int {
	return (*items)[index]
}
