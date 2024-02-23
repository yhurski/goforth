package forth

// data stack, i.e. just "stack"
const dataStackDepth uint8 = 255

var dataStack *stack = CreateStack()

// return stack
const returnStackDepth uint8 = 255

var returnStack *stack = CreateStack()

// instruction pointer
var ip uint32 = 0

// current mode: 0 - interpreting, 1 - compiling
var state int = 0

// input buffer
var inputBuffer string

// >in pointer
var pIn int = 0

// store error code of an instruciton
var errorCode byte = 0
