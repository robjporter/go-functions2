package main

import (
	"fmt"

	"../format/text"
	"github.com/robjporter/go-functions/as"
)

func main() {
	fmt.Println("STRING:                >", "UpperCamelCase")
	fmt.Println("TOSPINAL:              >", text.ToSpinal("UpperCamelCase"))
	fmt.Println("TOTRAIN:               >", text.ToTrain("UpperCamelCase"))
	fmt.Println("TOSNAKE:               >", text.ToSnake("UpperCamelCase"))
	fmt.Println("TOSNAKEUPPER:          >", text.ToSnakeUpper("UpperCamelCase"))
	fmt.Println("TOCAMEL:               >", text.ToCamel("upper_camel_case"))
	fmt.Println("TOCAMELLOWER:          >", text.ToCamelLower("upper_camel_case"))
	fmt.Println("TOREVERSE:             >", text.ToReverse("upper_camel_case"))
	fmt.Println("TOTITLE:               >", text.ToTitle("upper camel case"))
	fmt.Println("GENERATERANDOM:        >", text.GenerateRandomString(40))
	fmt.Println("GENERATERANDOMSPECIAL: >", text.GenerateRandomStringSpecial(40))
	j := text.Jaro("string1", "string2")
	fmt.Println("SIMILARITYJARO:        >", j)
	text.SetBoostThreshold(0.7)
	text.SetPrefixScale(0.1)
	jw := text.JaroWinkler("string1", "string2")
	fmt.Println("SIMILARITYJAROW:       >", jw)
	jw = text.JaroWinkler("string1", "pesty")
	fmt.Println("SIMILARITYJAROW:       >", jw)
	jw = text.JaroWinkler("string1", "string1")
	fmt.Println("SIMILARITYJAROW:       >", jw)
	fmt.Println("FILENAMEEXT:           >", text.FileExtension("/tmp/test/testing.json"))
	fmt.Println("FORMAT:                >", text.Format("The {} says {}", "cow", "MOO!"))
	fmt.Println("FORMAT2:               >", text.Format("I have {} bananas", 871))
	fmt.Println("FORMAT3:               >", text.Format("{} I have bananas", true))
	fmt.Println("FORMAT4:               >", text.Format("It cost £{}", 9.99))
	fmt.Println("FORMAT5:               >", text.Format("This {} part test, is {} and is {} that it is {}% perfect.", 4, true, "trying to prove", 99.99))
	test := []string{"a", "b", "c", "d"}
	test2 := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println("SLICEINCLUDES1:        >", text.Includes(test, "a"))
	fmt.Println("SLICEINCLUDES2:        >", text.Includes(test, "e"))
	fmt.Println("SLICEEQUALS1:          >", text.Equals(test, test2))
	fmt.Println("SLICEEQUALS2:          >", text.Equals(test, test))
	fmt.Println("SLICESUBTRACT:         >", text.Subtract(test2, test))
	fmt.Println("TOMILITARY:            >", text.ToMilitary("TEST"))
	fmt.Println("TOMORSE:               >", text.ToMorse("TEST"))
	fmt.Println("ORDINISE1:             >", text.Ordinise(1))
	fmt.Println("ORDINISE2:             >", text.Ordinise(2))
	fmt.Println("ORDINISE3:             >", text.Ordinise(3))
	fmt.Println("ORDINISE4:             >", text.Ordinise(4))
	fmt.Println("ORDINISE5:             >", text.Ordinise(22))
	fmt.Println("ORDINISE6:             >", text.Ordinise(44))
	fmt.Println("ORDINISE7:             >", text.Ordinise(101))
	fmt.Println("ORDINISE8:             >", text.Ordinise(99))
	fmt.Println("TOORDINISE1:           >", text.ToOrdinise(1))
	fmt.Println("TOORDINISE2:           >", text.ToOrdinise(2))
	fmt.Println("TOORDINISE3:           >", text.ToOrdinise(3))
	fmt.Println("TOORDINISE4:           >", text.ToOrdinise(4))
	fmt.Println("TOORDINISE5:           >", text.ToOrdinise(22))
	fmt.Println("TOORDINISE6:           >", text.ToOrdinise(44))
	fmt.Println("TOORDINISE7:           >", text.ToOrdinise(101))
	a := text.PKCS7PaddingString("Testing", 20)
	fmt.Println("PADDING:               >", a)
	fmt.Println("PADDING STRING:        >", as.ToString(a))
	fmt.Println("UNPADDING:             >", text.PKCS7Unpadding(a))
	fmt.Println("UNPADDING STRING:      >", as.ToString(text.PKCS7Unpadding(a)))
}
