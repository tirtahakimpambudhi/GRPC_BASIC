package factory

import (
	"fmt"
	"grpc_course/pb"
	"math/rand"
	"strconv"
)

func RandomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTZ
	case 2:
		return pb.Keyboard_AZERTY
	default:
		return pb.Keyboard_QWERTY
	}
}

func RandomBool() bool {
	return rand.Intn(2) == 1
}

func RandomBrandCPU() string {
	return randomFromStringSet("INTEL", "AMD", "ARM", "RYZEN", "APPLE SILICON")
}

func RandomCPUName(brand string) string {
	switch brand {
	case "AMD":
		return randomFromStringSet("AMD Ryzen 9 5950X", "AMD Ryzen 7 5800X", "AMD Ryzen 5 5600X", "AMD EPYC 7763", "AMD Athlon 3000G")
	case "ARM":
		return randomFromStringSet("ARM Cortex-A77", "ARM Cortex-A78", "ARM Cortex-A55")
	case "RYZEN":
		return randomFromStringSet("AMD Ryzen 9 5900X", "AMD Ryzen 7 5800X", "AMD Ryzen 5 5600X")
	case "APPLE SILICON":
		return randomFromStringSet("Apple M1", "Apple M1 Pro", "Apple M1 Max")
	default:
		return randomFromStringSet("Intel Core i9-11900K", "Intel Core i7-11700K", "Intel Core i5-11600K", "Intel Xeon E5-2699 v4", "Intel Pentium Gold G5600")
	}
}

func randomFromStringSet(a ...string) string {
	if m := len(a); m == 0 {
		return ""
	}
	return a[rand.Intn(len(a))]
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func RandomBrandGPU() string {
	return randomFromStringSet("NVIDIA", "AMD", "INTEL", "APPLE SILICON", "RADEON")
}

func RandomGPUName(brand string) string {
	switch brand {
	case "AMD":
		return randomFromStringSet("Radeon RX 6900 XT", "Radeon RX 6800 XT", "Radeon RX 6700 XT")
	case "NVIDIA":
		return randomFromStringSet("GeForce RTX 3090", "GeForce RTX 3080", "GeForce RTX 3070")
	case "RADEON":
		return randomFromStringSet("RX 6900 XT", "RX 6800 XT", "RX 6700 XT")
	case "APPLE SILICON":
		return randomFromStringSet("M1 GPU", "M1 Pro GPU", "M1 Max GPU")
	default:
		return randomFromStringSet("Xe Graphics", "Iris Xe MAX", "DG1")
	}
}

func RandomPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func RandomBrandLaptop() string {
	return randomFromStringSet("DELL", "ASUS", "APPLE", "LENOVO", "ACER", "MICROSOFT", "SAMSUNG")
}

func RandomModelLaptop(brand string) string {
	switch brand {
	case "DELL":
		return randomFromStringSet("XPS", "Inspiron", "Latitude", "Alienware", "Precision")
	case "ASUS":
		return randomFromStringSet("ZenBook", "VivoBook", "ROG (Republic of Gamers)", "TUF")
	case "APPLE":
		return randomFromStringSet("MacBook Air", "MacBook Pro", "iMac", "Mac mini")
	case "LENOVO":
		return randomFromStringSet("ThinkPad", "IdeaPad", "Yoga", "Legion")
	case "ACER":
		return randomFromStringSet("Aspire", "Predator", "Swift", "TravelMate")
	case "MICROSOFT":
		return randomFromStringSet("Surface Laptop", "Surface Pro", "Surface Book", "Surface Studio")
	default:
		return randomFromStringSet("Galaxy Book", "Notebook", "Odyssey")
	}
}

func RandomCurrenyCode() string {
	return randomFromStringSet("USD", "EUR", "JPY", "IDR")
}

func RandomAmount(currencyCode string) int64 {
	switch currencyCode {
	case "USD":
		return int64(RandomInt(300, 5000))
	case "EUR":
		return int64(RandomInt(300, 5000))
	case "JPY":
		return int64(RandomInt(30000, 500000))
	default:
		return int64(RandomInt(5000000, 80000000))
	}
}

func RandomOS() *pb.OS {
	os := new(pb.OS)
	switch rand.Intn(3) {
	case 1:
		versions := randomWindowsVersion()
		edition := randomWindowsEdition(versions)
		os.TypeOs = &pb.OS_Windows{Windows: &pb.Windows{
			Version: versions,
			Edition: edition,
		}}
	case 2:
		os.TypeOs = &pb.OS_MacOs{MacOs: &pb.MAC{
			Version:       randowMACVersion(),
			Build:         randomFromStringSet("20G71", "19G2021", "18G8022"),
			KernelVersion: strconv.FormatFloat(RandomFloat64(17.0, 20.5), 'f', 2, 64),
		}}
	default:
		os.TypeOs = &pb.OS_Linux{Linux: &pb.Linux{
			Distribution:  randomDistro(),
			Version:       strconv.FormatFloat(RandomFloat64(10.0, 23.5), 'f', 2, 64),
			KernelVersion: fmt.Sprintf("Linux %v", strconv.FormatFloat(RandomFloat64(1.0, 5.0), 'f', 2, 64)),
		}}
	}
	return os
}

func randomDistro() pb.Linux_Distro {
	switch rand.Intn(6) {
	case 1:
		return pb.Linux_UBUNTU
	case 2:
		return pb.Linux_FEDORA
	case 3:
		return pb.Linux_CENTOS
	case 4:
		return pb.Linux_DEBIAN
	case 5:
		return pb.Linux_ARCH
	case 6:
		return pb.Linux_ALPINE
	default:
		return pb.Linux_Mint
	}
}

func randomWindowsVersion() string {
	return randomFromStringSet("Windows 10", "Windows 8.1", "Windows 8", "Windows 7", "Windows XP")
}
func randowMACVersion() string {
	return randomFromStringSet("MacOS BIG SUR", "MacOS Catalina", "MacOS Mojave")
}

func randomWindowsEdition(windowsVersion string) string {
	switch windowsVersion {
	case "Windows 10":
		return randomFromStringSet("Home", "Pro", "Enterprise", "Education", "Pro Education", "Pro Worksatations")
	case "Windows XP":
		return randomFromStringSet("Home", "Professional", "Media Center", "Tablet PC")
	case "Windows 7":
		return randomFromStringSet("Starter", "Home Basic", "Home Premium", "Professional", "Enterprise", "Ultimate")
	default:
		return randomFromStringSet("Pro", "Enterprise")
	}
}
