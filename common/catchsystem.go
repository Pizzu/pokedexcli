package common

import (
	"fmt"
	"math/rand"
	"time"
)

func AttempCatch(baseExperience int) bool {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	maxValue := 500
	catchChance := 1.0 - (float64(baseExperience) / float64(maxValue))

	if catchChance < 0 {
		catchChance = 0
	}

	if catchChance > 1 {
		catchChance = 1
	}

	randomValue := generator.Float64()

	fmt.Println(randomValue, catchChance)

	return randomValue < catchChance
}
