package main

// data stack, i.e. just "stack"
const dataStackDepth uint8 = 255

var dataStack *stack = CreateStack()

// return stack
const returnStackDepth uint8 = 255

var returnStack *stack = CreateStack()

// current mode: 0 - interpreting, 1 - compiling
var state byte = 0

// input buffer
var inputBuffer string

// >in pointer
var pIn int = 0
