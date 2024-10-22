package typechart

import "strings"

type TypeChart struct {
	chart map[string]map[string]float64
}

func NewTypeChart() *TypeChart {
	return &TypeChart{
		chart: make(map[string]map[string]float64),
	}
}

func (self *TypeChart) AddInteraction(attackingType, defendingType string, effectiveness float64) *TypeChart {
	attackingType, defendingType = normalizeInput(attackingType, defendingType)
	if self.chart[attackingType] == nil {
		self.chart[attackingType] = make(map[string]float64)
	}
	self.chart[attackingType][defendingType] = effectiveness
	return self
}

func (self *TypeChart) RemoveInteraction(attackingType, defendingType string) *TypeChart {
	attackingType, defendingType = normalizeInput(attackingType, defendingType)
	defendingMap, attackingTypeExists := self.chart[attackingType]
	if attackingTypeExists {
		_, defendingTypeExists := defendingMap[defendingType]
		if defendingTypeExists {
			delete(defendingMap, defendingType)
		}

		if len(defendingMap) == 0 {
			delete(self.chart, attackingType)
		}
	}
	return self
}

func (self *TypeChart) Effectiveness(attackingType, defendingType string) (float64, bool) {
	attackingType, defendingType = normalizeInput(attackingType, defendingType)
	effectiveness, exists := self.chart[attackingType][defendingType]
	if exists {
		return effectiveness, true
	}
	return -1.0, false
}

func normalizeInput(attacking, defending string) (string, string) {
	return strings.ToUpper(attacking), strings.ToUpper(defending)
}

func (self *TypeChart) Equals(other *TypeChart) bool {
	if self == other {
		return true // Check for self-comparison
	}

	if other == nil || len(self.chart) != len(other.chart) {
		return false
	}

	for attackingType, defendingMap := range self.chart {
		otherDefendingMap, exists := other.chart[attackingType]
		if !exists || len(defendingMap) != len(otherDefendingMap) {
			return false
		}

		for defendingType, effectiveness := range defendingMap {
			if otherEffectiveness, exists := otherDefendingMap[defendingType]; !exists || effectiveness != otherEffectiveness {
				return false
			}
		}
	}

	return true
}

func (self *TypeChart) AttackingTypes() map[string]struct{} {
	result := make(map[string]struct{})
	for attackingType := range self.chart {
		result[attackingType] = struct{}{}
	}
	return result
}

func (tc *TypeChart) DefendingTypes() map[string]struct{} {
	result := make(map[string]struct{})

	for _, defendingMap := range tc.chart {
		for defendingType := range defendingMap {
			result[defendingType] = struct{}{} // Use empty struct for set behavior
		}
	}
	return result
}
