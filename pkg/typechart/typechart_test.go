package typechart

import (
	"fmt"
	"testing"
)

func TestAddInteraction(t *testing.T) {
	tc := NewTypeChart()

	tests := []struct {
		attackingType  string
		defendingType  string
		effectiveness  float64
		expectedExists bool
		expectedEff    float64
	}{
		{"FIRE", "GRASS", 2.0, true, 2.0},
		{"WATER", "GRASS", 0.5, true, 0.5},
		{"NORMAl", "NORMAL", 1.0, true, 1.0},
		{"GHOST", "NORMAL", 0.0, true, 0.0},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Add %s -> %s", tt.attackingType, tt.defendingType)
		t.Run(testname, func(t *testing.T) {
			tc.AddInteraction(tt.attackingType, tt.defendingType, tt.effectiveness)
			effectiveness, exists := tc.Effectiveness(tt.attackingType, tt.defendingType)

			if exists != tt.expectedExists {
				t.Errorf("got existence %v, want %v", exists, tt.expectedExists)
			}
			if effectiveness != tt.expectedEff {
				t.Errorf("got effectiveness %.1f, want %.1f", effectiveness, tt.expectedEff)
			}
		})
	}
}

func TestRemoveInteraction(t *testing.T) {
	tc := NewTypeChart()

	// Initial setup
	tc.AddInteraction("WATER", "FIRE", 2.0)
	tc.AddInteraction("WATER", "GROUND", 2.0)

	tests := []struct {
		attackingType  string
		defendingType  string
		expectedExists bool
		remainingEntry bool // Whether we expect the attackingType to remain after removal
	}{
		{"WATER", "FIRE", false, true},    // Removing "FIRE" should leave "WATER -> GROUND"
		{"WATER", "GROUND", false, false}, // Removing "GROUND" should remove "WATER" entirely
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Remove %s -> %s", tt.attackingType, tt.defendingType)
		t.Run(testname, func(t *testing.T) {
			tc.RemoveInteraction(tt.attackingType, tt.defendingType)
			_, exists := tc.Effectiveness(tt.attackingType, tt.defendingType)

			if exists != tt.expectedExists {
				t.Errorf("got existence %v, want %v", exists, tt.expectedExists)
			}

			_, attackingTypeExists := tc.chart[tt.attackingType]
			if attackingTypeExists != tt.remainingEntry {
				t.Errorf("got attacking type remaining %v, want %v", attackingTypeExists, tt.remainingEntry)
			}
		})
	}
}

func TestEffectiveness(t *testing.T) {
	tc := NewTypeChart()

	// Add some interactions for testing
	tc.AddInteraction("FIRE", "GRASS", 2.0)
	tc.AddInteraction("WATER", "GRASS", 0.5)
	tc.AddInteraction("GHOST", "NORMAL", 0.0)
	tc.AddInteraction("NORMAL", "FIRE", 1.0)

	tests := []struct {
		attackingType  string
		defendingType  string
		expectedEff    float64
		expectedExists bool
	}{
		{"FIRE", "GRASS", 2.0, true},
		{"WATER", "GRASS", 0.5, true},
		{"GHOST", "NORMAL", 0.0, true},
		{"NORMAL", "FIRE", 1.0, true},
		{"ELECTRIC", "GROUND", -1.0, false}, // No interaction present
		{"FAIRY", "GROUND", -1.0, false},    // No interaction present
		{"DRAGON", "GROUND", -1.0, false},   // No interaction present
		{"ROCK", "GROUND", -1.0, false},     // No interaction present
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Get %s -> %s", tt.attackingType, tt.defendingType)
		t.Run(testname, func(t *testing.T) {
			effectiveness, exists := tc.Effectiveness(tt.attackingType, tt.defendingType)

			if exists != tt.expectedExists {
				t.Errorf("got existence %v, want %v", exists, tt.expectedExists)
			}

			if effectiveness != tt.expectedEff {
				t.Errorf("got effectiveness %.1f, want %.1f", effectiveness, tt.expectedEff)
			}
		})
	}
}

func TestTypeChartEqual(t *testing.T) {
	chart1 := NewTypeChart()
	chart1.AddInteraction("FIRE", "GRASS", 2)
	chart1.AddInteraction("WATER", "FIRE", 2)

	chart2 := NewTypeChart()
	chart2.AddInteraction("FIRE", "GRASS", 2)
	chart2.AddInteraction("WATER", "FIRE", 2)

	chart3 := NewTypeChart()
	chart3.AddInteraction("WATER", "GRASS", 2)

	tests := []struct {
		name     string
		chartA   *TypeChart
		chartB   *TypeChart
		expected bool
	}{
		{"Equal charts", chart1, chart2, true},
		{"Different charts", chart1, chart3, false},
		{"Nil comparison", chart1, nil, false},
		{"Self comparison", chart1, chart1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.chartA.Equals(tt.chartB)
			if got != tt.expected {
				t.Errorf("Equal() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAttackingTypeSet(t *testing.T) {
	// Setup
	chart := NewTypeChart()
	chart.AddInteraction("NORMAL", "NORMAL", 1)
	chart.AddInteraction("NORMAL", "FIRE", 1)
	chart.AddInteraction("FIRE", "NORMAL", 1)
	chart.AddInteraction("FIRE", "FIRE", 0.5)

	expected := make(map[string]struct{})
	expected["NORMAL"] = struct{}{}
	expected["FIRE"] = struct{}{}

	// Act
	actual := chart.AttackingTypes()

	// Assert
	if len(expected) != len(actual) {
		t.Errorf("Expected length %d but got %d for attacking types set.", len(expected), len(actual))
	}

	for key := range expected {
		if _, exists := actual[key]; !exists {
			t.Errorf("Expected attacking type '%s' not found in actual set.", key)
		}
	}
}

func TestDefendingTypeSet(t *testing.T) {
	// Setup
	chart := NewTypeChart()
	chart.AddInteraction("NORMAL", "NORMAL", 1)
	chart.AddInteraction("NORMAL", "FIRE", 1)
	chart.AddInteraction("FIRE", "WATER", 1)
	chart.AddInteraction("FIRE", "GROUND", 0.5)

	expected := make(map[string]struct{})
	expected["NORMAL"] = struct{}{}
	expected["FIRE"] = struct{}{}
	expected["WATER"] = struct{}{}
	expected["GROUND"] = struct{}{}

	// Act
	actual := chart.DefendingTypes()

	// Assert
	if len(expected) != len(actual) {
		t.Errorf("Expected length %d but got %d for attacking types set.", len(expected), len(actual))
	}

	for key := range expected {
		if _, exists := actual[key]; !exists {
			t.Errorf("Expected attacking type '%s' not found in actual set.", key)
		}
	}
}
