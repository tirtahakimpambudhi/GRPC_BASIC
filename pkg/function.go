package pkg

import (
	"fmt"
	"grpc_course/pb"
)

func IsQualified(filter *pb.Filter, laptop *pb.Laptop) bool {
	if laptop.GetPrice().Amount > filter.GetMaxMoney().Amount && laptop.GetPrice().CurrencyCode == filter.GetMaxMoney().CurrencyCode {
		return false
	}

	if laptop.GetCpu().GetNumberCores() < filter.GetMinCpuCores() {
		return false
	}

	if laptop.GetCpu().GetMinGhz() < filter.GetMinCpuGhz() {
		return false
	}

	if toBit(laptop.GetRam()) < toBit(filter.GetMinRam()) {
		return false
	}
	return true
}

func toBit(memory *pb.Memory) uint64 {
	value := memory.GetValue()

	switch memory.GetUnit() {
	case pb.Memory_BIT:
		return value
	case pb.Memory_BYTE:
		return value << 3
	case pb.Memory_KILOBYTE:
		return value << 13
	case pb.Memory_MEGABYTE:
		return value << 23
	case pb.Memory_GIGABYTE:
		return value << 33
	case pb.Memory_TERABYTE:
		return value << 43
	default:
		return 0
	}
}

func PrintLaptopsProtobufToJSON(laptops []*pb.Laptop) error {
	for i, _ := range laptops {
		laptop := laptops[i]
		toJSON, err := ProtoBufToJSON(laptop)
		if err != nil {
			return err
		}
		fmt.Println(toJSON)
	}
	return nil
}
