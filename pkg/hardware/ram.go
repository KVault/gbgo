package hardware

// RAM represents the entire memory bank of the system, which is 0xFFFF long ~64Kb
var RAM = struct {
	Bank [0xFFFF]byte
}{}

//Write to the addr the content specified by the second parameter
func Write(addr uint16, content byte){
	if addr < 0 || addr > uint16(len(RAM.Bank)) {
		return 
	}

	RAM.Bank[addr] = content
}

//Read from the main memory bank the possition specified by addr
func Read(addr uint16) byte{
	if addr < 0 || addr > uint16(len(RAM.Bank)) {
		return 0
	}
	return RAM.Bank[addr]
}