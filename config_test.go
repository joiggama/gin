package gin

import (
	"testing"
)

func TestConfigGetDefinedKey(t *testing.T) {
	r := New()

	r.Config.Set("foo", "bar")

	if c := r.Config.Get("foo"); c != "bar" {
		t.Errorf("Expected foo to be nil but was %s", c)
	}
}

func TestConfigGetUndefinedKey(t *testing.T) {
	r := New()

	if c := r.Config.Get("foo"); c != nil {
		t.Errorf("Expected foo to be nil but was %s", c)
	}
}

func TestConfigSetDefinedKey(t *testing.T) {
	r := New()

	r.Config.Set("foo", "bar")

	if c := r.Config.Get("foo"); c != "bar" {
		t.Errorf("Expected config foo to be bar but was %s", c)
	}

	r.Config.Set("foo", "baz")

	if c := r.Config.Get("foo"); c != "baz" {
		t.Errorf("Expected config foo to be baz but was %s", c)
	}
}

func TestConfigSetUndefinedKey(t *testing.T) {
	r := New()

	r.Config.Set("foo", "bar")

	if c := r.Config.Get("foo"); c != "bar" {
		t.Errorf("Expected config foo to be bar but was %s", c)
	}
}
