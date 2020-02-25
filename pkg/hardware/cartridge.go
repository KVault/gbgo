package hardware

import(
	"io/ioutil"
)

// ROM holds all the data of a cartridge 
type ROM struct {
	Dump []byte
}

// LoadROM Loads a rom given the current path, if found, dumps all the contents in a struct.
// Will return an error if the file doesn't exists or the extension isn't right
func LoadROM(absolutePath string) (*ROM, error) {	
	dump, err := ioutil.ReadFile(absolutePath)

	if err != nil {
		//TODO the error
		return nil, err
	}

	rom := &ROM{
		Dump: dump,
	}

	return rom, nil
}
