// +build !example
package hello

const TestVersion = 1

func HelloWorld(s string) string {
	if s == "" {
		s = "World"
	}
	return "Hello, " + s + "!";
}
