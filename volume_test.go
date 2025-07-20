package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	kubus 				Kubus 	= Kubus{4}
	volumeSeharusnya	float64 = 64
	luasSeharusnya		float64 = 96
	kelilingSeharusnya	float64 = 48
)

func TestHitungVol(t *testing.T) {
	t.Logf("Volume : %.2f", kubus.Volume())

	//if kubus.Volume() != volumeSeharusnya {
	//	t.Errorf("Salah, seharusnya %.2f", volumeSeharusnya)
	//}
	assert.Equal(t, kubus.Volume(), volumeSeharusnya, "perhitungan volume salah")
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %.2f", kubus.Luas())

	//if kubus.Luas() != luasSeharusnya {
	//	t.Errorf("Salah, seharusnya %.2f", luasSeharusnya)
	//}

	assert.Equal(t, kubus.Luas(), luasSeharusnya, "perhitungan luas salah")
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling : %2.f", kubus.Keliling())

	//if kubus.Keliling() != kelilingSeharusnya {
	//	t.Errorf("Salah, seharusnya %.2f", kelilingSeharusnya)
	//}

	assert.Equal(t, kubus.Keliling(), kelilingSeharusnya, "perhitungan keliling salah")
}