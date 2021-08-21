package length

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComparison(t *testing.T) {
	assert := assert.New(t)
	t.Run("Comparing 1 inch to 2.54cm", func(t *testing.T) {
		var lengthInCm  = 2.54* Centimeter
		inch := 1.* In

		assertLength(assert, inch, lengthInCm)
	})
	t.Run("Comparing 11.811 inches is 30cm", func(t *testing.T) {
		var lengthInCm = 30. * Centimeter

		inch := 11.811 * In

		assertLength(assert,inch, lengthInCm)
	})
}

func assertLength(assert *assert.Assertions, expected Length, actual Length) bool {
	return assert.InDelta(float64(expected), float64(actual), 0.0001, "")
}

func TestComparingFeetAndCm(t *testing.T) {
	assert :=assert.New(t)
	t.Run("Comparing 1 feet to 30.48cm", func(t *testing.T) {
		var l = 30.48* Centimeter

		result :=l/ Feet

		assertLength(assert,1., result)
	})

	t.Run("Comparing 10 feet to 304.8cm", func(t *testing.T) {
		var l = 304.8* Centimeter

		result :=l/ Feet

		assertLength(assert,10., result)
	})
}

func TestConversionLength(t *testing.T) {
	assert := assert.New(t)
	t.Run("Convert 100cm to 1meter", func(t *testing.T) {
		hundredCentimeter := 100* Centimeter
		expected := 1.

		convertedValue := hundredCentimeter.Meter()

		assert.InDelta(expected, convertedValue, 0.0001, "")
	})
	t.Run("Convert 200cm to 2meter", func(t *testing.T) {
		twoHundredCentimeter := 200* Centimeter
		expected := 2.

		convertedValue := twoHundredCentimeter.Meter()

		assert.InDelta(expected, convertedValue, 0.0001, "")
	})
}