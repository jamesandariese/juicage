package main

import "github.com/jamesandariese/juicage"
import "flag"

var _ = flag.String("o", "default", "o has a default value")

func init() {
	juicage.Mandatory("hello", "A hello argument")
	juicage.Rest("files", "files to use")
	juicage.Optional("maybe", "Perhaps you'd like to specify this one, too")
	juicage.Optional("maybe2", "More maybe arg!")
	juicage.Mandatory("really1", "Set this one though")
	juicage.Optional("maybe3", "Maybe the third")
}

func main() {
	juicage.Print()
}
