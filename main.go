package main

//TODO  打印字符串时 补全高位的0
import (
	"flag"
	"fmt"
	"io"
	"os"
)

type BinaryFile struct {
	data string
	//length    int64
	printType string
}

// func (bf *BinaryFile) toHex() (res string) {

// }

func (bf *BinaryFile) PrintTo(des io.Writer) {
	formatting := ""
	switch bf.printType {
	case "x":
		formatting = fmt.Sprintf("%%02%s ", bf.printType)
	case "b":
		formatting = fmt.Sprintf("%%08%s ", bf.printType)
	case "o":
		formatting = fmt.Sprintf("%%04%s ", bf.printType)

	}

	for i := 0; i < len(bf.data); i++ {
		if i%8 == 0 {
			fmt.Fprintf(des, "0x%08x    ", i)

		} else {
			if (i+1)%8 == 0 {
				fmt.Fprintf(des, formatting, bf.data[i])
				fmt.Fprintln(des)
			} else {
				fmt.Fprintf(des, formatting, bf.data[i])
			}
		}

	}

}

func main() {
	path := flag.String("path", "", "file path")
	mode := flag.String("mode", "b", " x o b")

	flag.Parse()
	var print BinaryFile
	print.printType = *mode
	file, err := os.Open(*path)
	if err != nil {
		os.Exit(0)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		os.Exit(0)
	}
	print.data = string(data)

	if err != nil {
		fmt.Print(err)
	}
	print.PrintTo(os.Stdout)

}
