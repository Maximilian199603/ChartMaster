package typechart

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDeserialize(t *testing.T) {
	// Define the path to the CSV file
	filepath := filepath.Join("testData", "shortTypeChart.csv")

	// Call the Deserialize function
	typeChart, err := DeserializeFile(filepath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if the typeChart is not nil
	if typeChart == nil {
		t.Fatal("Expected typeChart to be non-nil")
	}

	// Construct the expected TypeChart
	expectedChart := NewTypeChart()
	expectedChart.
		AddInteraction("Normal", "Normal", 1.0).
		AddInteraction("Normal", "Fire", 1.0).
		AddInteraction("Normal", "Water", 1.0).
		AddInteraction("Normal", "Electric", 1.0).
		AddInteraction("Fire", "Normal", 1.0).
		AddInteraction("Fire", "Fire", 0.5).
		AddInteraction("Fire", "Water", 0.5).
		AddInteraction("Fire", "Electric", 1.0).
		AddInteraction("Water", "Normal", 1.0).
		AddInteraction("Water", "Fire", 2.0).
		AddInteraction("Water", "Water", 0.5).
		AddInteraction("Water", "Electric", 1.0).
		AddInteraction("Electric", "Normal", 1.0).
		AddInteraction("Electric", "Fire", 1.0).
		AddInteraction("Electric", "Water", 2.0).
		AddInteraction("Electric", "Electric", 0.5)

	// Check if the deserialized chart equals the expected chart
	if !typeChart.Equals(expectedChart) {
		t.Error("Deserialized TypeChart does not match the expected TypeChart")
	}
}

func TestSerialize(t *testing.T) {
	// Create a TypeChart for testing
	typeChart := NewTypeChart()
	typeChart.AddInteraction("Normal", "Normal", 1.0).
		AddInteraction("Normal", "Fire", 1.0).
		AddInteraction("Normal", "Water", 1.0).
		AddInteraction("Normal", "Electric", 1.0).
		AddInteraction("Fire", "Normal", 1.0).
		AddInteraction("Fire", "Fire", 0.5).
		AddInteraction("Fire", "Water", 0.5).
		AddInteraction("Fire", "Electric", 1.0).
		AddInteraction("Water", "Normal", 1.0).
		AddInteraction("Water", "Fire", 2.0).
		AddInteraction("Water", "Water", 0.5).
		AddInteraction("Water", "Electric", 1.0).
		AddInteraction("Electric", "Normal", 1.0).
		AddInteraction("Electric", "Fire", 1.0).
		AddInteraction("Electric", "Water", 2.0).
		AddInteraction("Electric", "Electric", 0.5)

	// Define the file path for the output CSV
	outputPath := filepath.Join("testData", "outputTypeChart.csv")

	// Serialize the TypeChart
	err := SerializeToFile(typeChart, outputPath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if the file was created
	if _, err := os.Stat(outputPath); err != nil {
		t.Fatalf("Expected file %s to be created, but it was not", outputPath)
	}

	// Clean up the test file after the test completes
	defer os.Remove(outputPath)

	// Deserialize the file to verify its content matches the original TypeChart
	deserializedChart, err := DeserializeFile(outputPath)
	if err != nil {
		t.Fatalf("Expected no error during deserialization, got %v", err)
	}

	// Verify that the deserialized chart equals the original TypeChart
	if !typeChart.Equals(deserializedChart) {
		t.Error("Deserialized TypeChart does not match the original TypeChart")
	}
}
