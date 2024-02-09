package main

const DataStackDepth uint8 = 255

var DataStack [DataStackDepth]int = [DataStackDepth]int{}

var Sp uint8

// current mode: false - interpreting, true - compiling
var state bool = false
