package solution

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func GetMessage(hello string) string {
	return emoji.Sprint(hello)
}

func main() {
	a := GetMessage("Hello :world_map:!")
	fmt.Printf(a)
}
