package sample

import (
	"github.com/google/uuid"
	"gitlab.com/techschool/pcbook/pb"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomCPUBrand() string {
	return randomStringFormSet("Intel", "AMD")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFormSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9950H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	}
	return randomStringFormSet(
		"Ryzen 7 PRO 2780U",
		"Ryzen 5 PRO 3500U",
		"Ryzen 3 PRO 3200GE",
	)
}

func randomGPUBrand() string {
	return randomStringFormSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFormSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"RTX 1070",
		)
	}
	return randomStringFormSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-56",
	)
}

func randomLaptopBrand() string {
	return randomStringFormSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFormSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randomStringFormSet("Latitude", "Vostro", "XPS", "Alienware")
	default:
		return randomStringFormSet("Thinkpad X1", "Thinkpad P1", "Thinkpad P53")
	}
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1000, 4320)
	weight := height * 16 / 9

	resolution := &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(weight),
	}
	return resolution
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomStringFormSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}
func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomID() string {
	return uuid.New().String()
}
