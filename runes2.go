package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "abcd"
	//fmt.Println(s)

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
		//fmt.Printf("%x ", s[i])
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	//fmt.Println(s)
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
		//fmt.Printf("%#U starts at %d\n")
		//fmt.Println("%U starts at", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		//runeValue, width := utf8.DecodeRuneInString(s)
		fmt.Printf("%#U starts at %d\n", runeValue, width)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
