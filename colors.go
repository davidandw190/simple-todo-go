package todo

import (
	"fmt"
	"io"
	"os"
)

const (
	ColorDefault = "\x1b[39m"

	ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
	ColorBlue  = "\x1b[94m"
	ColorGrey  = "\x1b[90m"
)

func Red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func Green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func Blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

func Grey(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGrey, s, ColorDefault)
}

func PrintRed(w io.Writer, s string) {
	fmt.Fprintf(w, "%s%s%s", ColorRed, s, ColorDefault)
}

func PrintGreen(w io.Writer, s string) {
	fmt.Fprintf(w, "%s%s%s", ColorGreen, s, ColorDefault)
}

func PrintBlue(w io.Writer, s string) {
	fmt.Fprintf(w, "%s%s%s", ColorBlue, s, ColorDefault)
}

func PrintGrey(w io.Writer, s string) {
	fmt.Fprintf(w, "%s%s%s", ColorGrey, s, ColorDefault)
}

// PrintRedStderr prints red text to os.Stderr
func PrintRedStderr(s string) {
	PrintRed(os.Stderr, s)
}
