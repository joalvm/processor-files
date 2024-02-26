package utils

import (
	"math"

	"github.com/joalvm/processor-medias/pkg/models"
)

func GetAspectRatio(width int, height int) models.AspectRatio {
	ratio := float64(width) / float64(height)
	minDiff := math.MaxFloat64
	closest := models.AspectRatio{X: 0, Y: 0}

	standardRatios := []models.AspectRatio{
		{X: 1, Y: 1},               // Cuadrado
		{X: 2, Y: 1}, {X: 1, Y: 2}, // Panorámico
		{X: 3, Y: 2}, {X: 2, Y: 3}, // Fotografía
		{X: 4, Y: 3}, {X: 3, Y: 4}, // Estándar de televisión en la era analógica
		{X: 16, Y: 9}, {X: 9, Y: 16}, // Estándar de televisión en la era digital, Full HD, 2K, 4K, 5K y 8K
		{X: 18, Y: 9}, {X: 9, Y: 18}, // Estándar de televisión en la era digital, Full HD, 2K, 4K, 5K y 8K
		{X: 21, Y: 9}, {X: 9, Y: 21}, // Ultra panorámico
	}

	for _, standardRatio := range standardRatios {
		diff := math.Abs(ratio - float64(standardRatio.X)/float64(standardRatio.Y))
		if diff < minDiff {
			minDiff = diff
			closest = standardRatio
		}
	}

	return closest
}
