// LML - LiveMediaList
// Service for choice LiveMedia
// Roadmap
// 1 - message at start programm
// 1 - Pattern
// 2 - Read start data from file and write to DB
// 3 - Read config file
// 4 - Read file variable
package main

import (
	"fmt"
	//"gopkg.in/yaml.v2"
)

type listMediaLive struct {
	number       int
	name         string
	homepage     string
	download     string
	wikipedia    string
	distrowatch  string
	sizeMin      int
	sizeMax      int
	stableVer    string
	lastRelease  string
	target       []int
	os           byte
	state        byte
	media        []int
	architecture []int
	note         string
}

var (
	lmlDB        []listMediaLive
	State        = map[string]byte{}
	Target       = map[string]byte{}
	OS           = map[string]byte{}
	Media        = map[string]byte{}
	Architecture = map[string]byte{}
)

//	State: Target: OS: Media: Architecture:

func init() {
	fmt.Println(" ")
	fmt.Println("Start Init func: ")
	fmt.Println(" ")
	readValueDb()
	readStartDb()

}

func main() {

	fmt.Println(" ")
	fmt.Println("Start Main func: ")
	fmt.Println(" ")

	// State = make(map[string]byte)
	//Target = make(map[string]byte)
	//OS = make(map[string]byte)
	//Media = make(map[string]byte)
	//Architecture = make(map[string]byte)

	// fmt.Println(State)
	// fmt.Println(Target)
	// fmt.Println(OS)
	// fmt.Println(Media)
	// fmt.Println(Architecture)

	//fmt.Println(lmlDB[0].name)
	fmt.Println(" ")
}
