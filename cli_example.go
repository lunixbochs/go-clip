package clip

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func Run(cmd ...string) {
	switch path.Base(cmd[0]) {
	case "copy":
		fallthrough
	case "pbcopy":
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		Copy(string(data))
	case "paste":
		fallthrough
	case "pbpaste":
		os.Stdout.Write([]byte(Paste()))
	default:
		if len(cmd) > 1 {
			Run(cmd[1])
		} else {
			fmt.Printf("Usage: %s <command>\n", cmd[0])
			fmt.Println("Commands:")
			fmt.Println("  copy: place stdin on the clipboard")
			fmt.Println("  paste: print the clipboard contents")
			fmt.Println()
			fmt.Println("Note: can be (hard|sym)linked to (pb)?(copy|paste) to perform commands directly")
			fmt.Println()
		}
	}
}

func main() {
	Run(os.Args...)
}
