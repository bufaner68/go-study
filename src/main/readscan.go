// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main
import (
	"bufio"
	"fmt"
	"os"
)

//从输入中读取数据，以EOF或Ctrl+D结束输入
func main() {
	//创建一个存储数据的空map
	counts := make(map[string]int)
	//从标准输入中读取数据
	input := bufio.NewScanner(os.Stdin)
	//循环读取数据，每调用一次input.scan(),读入下一行，并移除行末换行符
	//读取的内容可以调用 input.Text()  得到
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}