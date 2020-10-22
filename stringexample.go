package main

import "fmt"

func main() {
	// Normal String (Can not contain newlines, and can have escape characters like `\n`, `\t` etc)
	var website = "\thttps://www.facebook.com\t\n"

	// Raw String (Can span multiple lines. Escape characters are not interpreted)
	var siteDescription = `\t\tIt is a facebook page.\t\n`

	fmt.Println(website, siteDescription)
}
