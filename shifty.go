// shifty prints the number of characters in files that require pressing the
// shift key, and the ratio of unshifted to shifted characters.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var yes, no int
	for _, file := range os.Args[1:] {
		y, n := shiftedChars(file)
		yes += y
		no += n
	}
	fmt.Println(yes, no, float64(no)/float64(yes))
}

func shiftedChars(file string) (yes, no int) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range data {
		if shiftedChar(c) {
			yes++
		} else {
			no++
		}
	}
	return
}

func shiftedChar(c byte) bool {
	if 'A' <= c && c <= 'Z' {
		return true
	}
	switch c {
	case '~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+':
		return true
	case '{', '}', '|', ':', '"', '<', '>', '?':
		return true
	}
	return false
}
