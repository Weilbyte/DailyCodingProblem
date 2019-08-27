//Package problems20 contains the problems from 11-20
package problems20

import (
	"strings"
)

func autoComplete(word string) []string {
	var sList = []string{"dog", "deer", "deal"}
	var result []string
	for _, sLWord := range sList {
		if strings.HasPrefix(sLWord, word) {
			result = append(result, sLWord)
		}
	}
	return result
}

func find(target string, word string) {

}

//f
//furry
//fury

//fu
