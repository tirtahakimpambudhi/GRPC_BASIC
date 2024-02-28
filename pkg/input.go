package pkg

import (
	"bufio"
	"fmt"
	"grpc_course/pb"
	"os"
)

func Input(question string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s: ", question)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func ChoiceToUnit() pb.Memory_Unit {
	fmt.Print("1.KB\n2.MB\n3.GB\n4.TB\n")
	choice := Input("Unit ")
	switch choice {
	case "1":
		return pb.Memory_KILOBYTE
	case "2":
		return pb.Memory_MEGABYTE
	case "4":
		return pb.Memory_TERABYTE
	default:
		return pb.Memory_GIGABYTE
	}
}
