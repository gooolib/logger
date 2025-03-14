package logger

import (
	"io"
	"os"
	"strings"
	"testing"
)

func captureErrorOutput(f func()) string {
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	f()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stderr = old

	return string(out)
}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = old

	return string(out)
}

func TestDefaultLogger__Debugf(t *testing.T) {
	out := captureOutput(func() {
		DefaultLogger.SetLevel(LogLevelDebug)
		DefaultLogger.Debugf("Hello, world, %s!", "foo")
		DefaultLogger.SetLevel(LogLevelInfo)
	})

	if !strings.Contains(string(out), "Hello, world, foo!") {
		t.Errorf("Output should contain '%s' but got %s", "\"Hello, world, foo!\"", out)
	}
	if !strings.HasPrefix(string(out), "[DEBUG]") {
		t.Errorf("Output should starts with '%s' but got %s", "[DEBUG]", out)
	}
}

func TestDefaultLogger__Infof(t *testing.T) {
	out := captureOutput(func() {
		DefaultLogger.Infof("Hello, world, %s!", "foo")
	})

	if !strings.Contains(string(out), "Hello, world, foo!") {
		t.Errorf("Output should contain '%s' but got %s", "\"Hello, world, foo!\"", out)
	}
	if !strings.HasPrefix(string(out), ColorCyan+"[INFO]"+ColorReset) {
		t.Errorf("Output should starts with '%s' but got %s", ColorCyan+"[INFO]"+ColorReset, out)
	}
}

func TestDefaultLogger__Warnf(t *testing.T) {
	out := captureOutput(func() {
		DefaultLogger.Warnf("Hello, world, %s!", "foo")
	})

	if !strings.Contains(string(out), "Hello, world, foo!") {
		t.Errorf("Output should contain '%s' but got %s", "\"Hello, world, foo!\"", out)
	}
	if !strings.HasPrefix(string(out), ColorYellow+"[WARN]"+ColorReset) {
		t.Errorf("Output should starts with '%s' but got %s", ColorYellow+"[INFO]"+ColorReset, out)
	}
}

func TestDefaultLogger__Errorf(t *testing.T) {
	out := captureOutput(func() {
		DefaultLogger.Errorf("Hello, world, %s!", "foo")
	})

	if !strings.Contains(string(out), "Hello, world, foo!") {
		t.Errorf("Output should contain '%s' but got %s", "\"Hello, world, foo!\"", out)
	}
	if !strings.HasPrefix(string(out), ColorRed+"[ERROR]"+ColorReset) {
		t.Errorf("Output should starts with '%s' but got %s", ColorRed+"[ERROR]"+ColorReset, out)
	}
}

func TestInfoLabel(t *testing.T) {
	if InfoLabel() != ColorCyan+"[INFO]"+ColorReset {
		t.Errorf("Expected InfoLabel() to return '%s[INFO]%s', got: %s", ColorCyan, ColorReset, InfoLabel())
	}
}
