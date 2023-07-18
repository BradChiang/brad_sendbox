package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	p := fmt.Println

	p("command-line subcommand :")

	foocmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooenable := foocmd.Bool("enable", false, "enable")
	fooname := foocmd.String("name", "", "name")

	barcmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barlevel := barcmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		p("excepted foo or bar subcommands")
	}

	switch os.Args[1] {
	case "foo":
		foocmd.Parse(os.Args[2:])
		p("subcommand 'foo'")
		p("		enable : ", *fooenable)
		p("		name   : ", *fooname)
		p("		tail   : ", foocmd.Args())
	case "bar":
		barcmd.Parse(os.Args[2:])
		p("subcommand 'bar'")
		p("		level  : ", *barlevel)
		p("		tail   : ", barcmd.Args())
	default:
		p("excepted foo or bar subcommands")
		os.Exit(1)
	}

}
