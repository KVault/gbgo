package cpu

// RAM represents the entire memory bank of the system, which is 0xFFFF long ~64Kb
var RAM = struct {
	bank [0xFFFF]byte
}{}
