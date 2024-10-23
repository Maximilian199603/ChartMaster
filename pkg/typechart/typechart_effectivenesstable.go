package typechart

import "math"

type WeaknessTable struct {
	table map[float64]map[string]struct{}
}

func NewWeaknesstable(chart TypeChart, types ...string) *WeaknessTable {
	var result = WeaknessTable{table: make(map[float64]map[string]struct{})}
	attackingSet := chart.AttackingTypes()
	for attackingType := range attackingSet {
		effect := combiEffectiveness(attackingType, types, chart)
		result.add(effect, attackingType)
	}
	return &result
}

func combiEffectiveness(attacking string, types []string, chart TypeChart) float64 {
	var activeAttackEffect float64 = 1.0
	for _, defendingType := range types {
		current, existance := chart.Effectiveness(attacking, defendingType)
		if !existance {
			activeAttackEffect = math.NaN()
			// found an type interaction that is not mapped break this loop
			break
		} else {
			activeAttackEffect = activeAttackEffect * current
		}
	}
	return activeAttackEffect
}

func (w *WeaknessTable) add(effectiveness float64, attackingType string) *WeaknessTable {
	if w.table[effectiveness] == nil {
		w.table[effectiveness] = make(map[string]struct{})
	}
	w.table[effectiveness][attackingType] = struct{}{}
	return w
}

func (w *WeaknessTable) AsMap() map[float64][]string {
	result := make(map[float64][]string)
	for k, v := range w.table {
		result[k] = convert(v)
	}
	return result
}

func convert(input map[string]struct{}) []string {
	var result []string
	for k := range input {
		result = append(result, k)
	}
	return result
}
