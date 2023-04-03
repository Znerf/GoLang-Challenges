package kata

import (
   "strings"
)

func GetCount(str string) (count int) {
  // Enter solution here
  return strings.Count(str, "a")+strings.Count(str, "e")+strings.Count(str, "i")+strings.Count(str, "o")+strings.Count(str, "u")
}

// count vowel

func GetCount(str string) (count int) {
	r := regexp.MustCompile("[aeiou]")
	vowels := r.FindAllString(str, -1)
	return len(vowels)
}

//test
