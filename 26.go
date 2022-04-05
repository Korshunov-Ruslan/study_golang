package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var (
		str, substr string
	)
	flag.StringVar(&str, "str", "", "setstr")
	flag.StringVar(&substr, "substr", "", "set substr")
	flag.Parse()
	isSubtstr(
		&str,
		&substr,
	)
}
func isSubtstr(str *string, substr *string) error {
	if str == nil || substr == nil {
		return fmt.Errorf("nil error")
	}
	fmt.Println(strings.Contains(*str, *substr))
	return nil
}
