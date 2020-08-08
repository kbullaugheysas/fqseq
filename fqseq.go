package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Args struct {
	Names bool
	Limit int
}

var args = Args{}

func init() {
	log.SetFlags(0)
	flag.BoolVar(&args.Names, "names", false, "include read names in tab-separated output")
	flag.IntVar(&args.Limit, "limit", 0, "output only the first LIMIT reads")

	flag.Usage = func() {
		log.Println("usage: fqseq [options]")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	lineNum := 0
	printed := 0
	var readName string
	for scanner.Scan() {
		line := scanner.Text()
		if lineNum%4 == 0 {
			readName = strings.TrimSpace(line[1:])
		}
		if lineNum%4 == 1 {
			line = strings.ToUpper(line)
			if args.Names {
				fmt.Printf("%s\t%s\n", readName, line)
			} else {
				fmt.Println(line)
			}
			printed += 1
		}
		lineNum++
		if args.Limit > 0 && printed == args.Limit {
			break
		}
	}
}
