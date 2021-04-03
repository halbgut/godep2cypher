package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		line := scanner.Text()
		split := strings.Split(line, " ")
		if len(split) == 0 {
			continue
		}
		if len(split) != 2 {
			fmt.Printf("%#v", split)
			return fmt.Errorf("Unexpected input format")
		}
		a := tidy(split[0])
		b := tidy(split[1])
		write(`MERGE (a:Dependency { name: "%v" })`, a)
		write(`MERGE (b:Dependency { name: "%v" })`, b)
		write("MERGE (a) -[:DEPENDS_ON]-> (b)")
		write(";")
	}
	return nil
}

func write(input string, args ...interface{}) {
	os.Stdout.Write([]byte(fmt.Sprintf(input+"\n", args...)))
}

func tidy(dep string) string {
	for _, abbr := range [][2]string{
		{"github.com", "gh"},
		{"gopkg.in", "gp"},
		{"golang.org/x", "x"},
	} {
		dep = strings.Replace(dep, abbr[0], abbr[1], 1)
	}
	return dep
}
