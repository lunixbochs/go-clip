package clip

import (
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	Copy(string(data))
}
