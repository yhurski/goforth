package main

// data stack, i.e. just "stack"
const dataStackDepth uint8 = 255

var dataStack *stack = CreateStack()

// return stack
const returnStackDepth uint8 = 255

var returnStack *stack = CreateStack()

// current mode: false - interpreting, true - compiling
var state bool = false
