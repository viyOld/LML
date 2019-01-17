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
	"log"
	//"gopkg.in/yaml.v2"
)

type listMediaLive struct {
	number       byte
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

var lmlDB []listMediaLive

func init() {

	log.Println("Start Init func: ")
	readValueDb()
	readStartDb()

}

func main() {
	fmt.Println(" ")
	fmt.Println("Start Main func: ")
	fmt.Println(" ")
	//fmt.Println(lmlDB[0].name)
	fmt.Println(" ")
}
