package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(".com")
	fmt.Println(re.FindString("google.com"))                                   //FindString()
	fmt.Println(re.FindString("magic.com.main.com"))          //return the string having the text from the leftmost match.
	fmt.Println(re.FindString("ne.ctfolixm.org"))

	fmt.Println(re.FindStringIndex("google.com"))                                   //FindStringIndex()
	fmt.Println(re.FindStringIndex("magic.com.main.com"))          					//return the index 
	fmt.Println(re.FindStringIndex("netflix.org"))


	re1:= regexp.MustCompile("f([a-z]+)ing")
	fmt.Println(re1.FindStringSubmatch("flying"))							//FindStringSubmatch()
	fmt.Println(re1.FindStringSubmatch("fling"))		// return a slice of strings having the text leftmost match and the matches.
	fmt.Println(re1.FindStringSubmatch("abcfloatingroll"))


	fmt.Println(re1.MatchString("flng"))											//MatchString
	fmt.Println(re1.MatchString("fling"))					//reports whether a string contains any match of regular expression.
	fmt.Println(re1.MatchString("flinge"))

	fmt.Println(re.ReplaceAllString("google.com", "X"))
	fmt.Println(re.ReplaceAllString("magic.com.main.com", "X"))						//ReplaceAllString
	fmt.Println(re.ReplaceAllString("ne.ctfolixm.org", "X"))	//returns a copy of string, replacing matches of RE with replacement
	fmt.Println(re1.ReplaceAllString("fling", "X"))
	fmt.Println(re1.ReplaceAllString("abcfloatingroll", "X"))


	
}