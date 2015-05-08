package juicage

import "testing"
import "flag"

var _ = &testing.T{}

func ExampleUsage() {
	flagSet := flag.NewFlagSet("testola", flag.ContinueOnError)
	var _ = flagSet.String("o", "default", "o has a default value")

	usage := NewUsage("testola", flagSet)

	usage.Mandatory("hello", "A hello argument")
	usage.Rest("files", "files to use")
	usage.Optional("maybe", "Perhaps you'd like to specify this one, too")
	usage.Optional("maybe2", "More maybe arg!")
	usage.Mandatory("really1", "Set this one though")
	usage.Optional("maybe3", "Maybe the third")

	usage.Print()

	// Output:
	// usage: testola hello really1 [maybe] [maybe2] [maybe3] [files]*
	//
	//    hello: A hello argument
	//  really1: Set this one though
	//    maybe: Perhaps you'd like to specify this one, too
	//   maybe2: More maybe arg!
	//   maybe3: Maybe the third
	//    files: files to use
	//
	//   -o="default": o has a default value

}
