package factory

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"grpc_course/pb"
)

func NewKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout:   RandomKeyboardLayout(),
		Backlist: RandomBool(),
	}
}

func NewCPU() *pb.CPU {
	brand := RandomBrandCPU()
	name := RandomCPUName(brand)
	cores := RandomInt(2, 8)
	threads := RandomInt(cores, 12)
	minGhz := RandomFloat64(2.0, 3.5)
	maxGhz := RandomFloat64(minGhz, 5.0)
	return &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(cores),
		NumberThreads: uint32(threads),
		MaxGhz:        maxGhz,
		MinGhz:        minGhz,
	}
}

func NewGPU() *pb.GPU {
	brand := RandomBrandGPU()
	name := RandomGPUName(brand)

	minGHz := RandomFloat64(1.0, 1.5)
	maxGHz := RandomFloat64(minGHz, 2.0)

	memory := &pb.Memory{
		Value: uint64(RandomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}
	return &pb.GPU{
		Brand:  brand,
		Name:   name,
		MaxGhz: maxGHz,
		MinGhz: minGHz,
		Memory: memory,
	}
}

func NewRAM() *pb.Memory {
	return &pb.Memory{
		Value: uint64(RandomInt(2, 8)),
		Unit:  pb.Memory_GIGABYTE,
	}
}

func NewSSD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(RandomInt(128, 1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
}

func NewHDD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(RandomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
}

func NewScreen() *pb.Screen {
	height := RandomInt(1080, 4320)
	width := height * 16 / 9

	return &pb.Screen{
		SizeInch: RandomFloat32(13, 17),
		Resolution: &pb.Screen_Resulotion{
			Width:  uint32(height),
			Height: uint32(width),
		},
		Panel:      RandomPanel(),
		MultiTouch: RandomBool(),
	}
}
func NewLaptop() *pb.Laptop {
	brand := RandomBrandLaptop()
	name := RandomModelLaptop(brand)
	currencyCode := RandomCurrenyCode()
	amount := RandomAmount(currencyCode)
	return &pb.Laptop{
		Id:       uuid.New().String(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*pb.GPU{NewGPU(), NewGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight:   &pb.Laptop_WeightKg{WeightKg: RandomFloat64(2.0, 5.0)},
		Price: &pb.Money{
			Amount:       amount,
			CurrencyCode: currencyCode,
		},
		Os:          RandomOS(),
		ReleaseYear: uint32(RandomInt(2000, 2024)),
		UpdatedAt:   ptypes.TimestampNow(),
	}
}
