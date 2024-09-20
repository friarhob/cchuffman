package adt

import (
	"testing"
)

func TestNewDefaultMap(t *testing.T) {
	defaultVal := "default"
	dm := NewDefaultMap(defaultVal)

	if dm == nil {
		t.Error("Expected new DefaultMap to be initialized")
	}

	if dm.Get("missing_key") != defaultVal {
		t.Errorf("Expected default value '%v', got '%v'", defaultVal, dm.Get("missing_key"))
	}
}

func TestSetAndGet(t *testing.T) {
	defaultVal := "default"
	dm := NewDefaultMap(defaultVal)

	key := "test_key"
	val := "test_value"

	dm.Set(key, val)

	if dm.Get(key) != val {
		t.Errorf("Expected value '%v', got '%v'", val, dm.Get(key))
	}

	if dm.Get("missing_key") != defaultVal {
		t.Errorf("Expected default value '%v' for missing key, got '%v'", defaultVal, dm.Get("missing_key"))
	}
}

func TestSetDefaultValue(t *testing.T) {
	initialDefaultVal := "initial_default"
	newDefaultVal := "new_default"
	dm := NewDefaultMap(initialDefaultVal)

	dm.SetDefaultValue(newDefaultVal)

	if dm.Get("missing_key") != newDefaultVal {
		t.Errorf("Expected default value '%v' after updating, got '%v'", newDefaultVal, dm.Get("missing_key"))
	}
}

func TestOverwriteValue(t *testing.T) {
	defaultVal := "default"
	dm := NewDefaultMap(defaultVal)

	key := "test_key"
	firstVal := "first_value"
	secondVal := "second_value"

	dm.Set(key, firstVal)
	if dm.Get(key) != firstVal {
		t.Errorf("Expected value '%v', got '%v'", firstVal, dm.Get(key))
	}

	dm.Set(key, secondVal)
	if dm.Get(key) != secondVal {
		t.Errorf("Expected value '%v', got '%v'", secondVal, dm.Get(key))
	}
}
