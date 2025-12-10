package testutils

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestUseTempHomeDirSetsTemporaryHome(t *testing.T) {
	originalHome := os.Getenv("HOME")

	cleanup := UseTempHomeDir(t)
	defer cleanup()

	newHome := os.Getenv("HOME")

	if newHome == originalHome {
		t.Error("UseTempHomeDir should set a different HOME directory")
	}

	if newHome == "" {
		t.Error("UseTempHomeDir should set HOME to a non-empty value")
	}

	// Verify the temp directory exists
	if _, err := os.Stat(newHome); err != nil {
		t.Errorf("temp HOME directory should exist: %v", err)
	}
}

func TestUseTempHomeDirRestoresOriginalHome(t *testing.T) {
	originalHome := os.Getenv("HOME")

	cleanup := UseTempHomeDir(t)
	cleanup()

	restoredHome := os.Getenv("HOME")

	if restoredHome != originalHome {
		t.Errorf("HOME = %q, want %q after cleanup", restoredHome, originalHome)
	}
}

func TestUseTempHomeDirHandlesUnsetHome(t *testing.T) {
	// Save and unset HOME
	originalHome, hadHome := os.LookupEnv("HOME")
	os.Unsetenv("HOME")

	// Restore after test
	defer func() {
		if hadHome {
			os.Setenv("HOME", originalHome)
		} else {
			os.Unsetenv("HOME")
		}
	}()

	cleanup := UseTempHomeDir(t)
	defer cleanup()

	// Verify HOME is now set
	newHome := os.Getenv("HOME")
	if newHome == "" {
		t.Error("UseTempHomeDir should set HOME even when originally unset")
	}

	cleanup()

	// Verify HOME is unset again
	if _, exists := os.LookupEnv("HOME"); exists {
		t.Error("HOME should be unset after cleanup when originally unset")
	}
}

func TestUseTempHomeDirResetsViper(t *testing.T) {
	cleanup := UseTempHomeDir(t)
	defer cleanup()

	// Set a viper value
	viper.Set("test.value", "testdata")

	if viper.GetString("test.value") != "testdata" {
		t.Fatal("viper value should be set before cleanup")
	}

	cleanup()

	// Verify viper was reset
	if viper.GetString("test.value") == "testdata" {
		t.Error("viper should be reset after cleanup")
	}
}

func TestUseTempHomeDirCreatesUniqueDirectories(t *testing.T) {
	cleanup1 := UseTempHomeDir(t)
	home1 := os.Getenv("HOME")
	cleanup1()

	cleanup2 := UseTempHomeDir(t)
	home2 := os.Getenv("HOME")
	cleanup2()

	if home1 == home2 {
		t.Error("UseTempHomeDir should create unique temp directories for each call")
	}
}

func TestUseTempHomeDirAllowsFileOperations(t *testing.T) {
	cleanup := UseTempHomeDir(t)
	defer cleanup()

	home := os.Getenv("HOME")
	testFile := home + "/test.txt"

	// Write to temp home directory
	if err := os.WriteFile(testFile, []byte("test content"), 0o600); err != nil {
		t.Fatalf("should be able to write to temp HOME: %v", err)
	}

	// Read from temp home directory
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("should be able to read from temp HOME: %v", err)
	}

	if string(content) != "test content" {
		t.Errorf("content = %q, want %q", string(content), "test content")
	}
}

func TestUseTempHomeDirIsolatesTests(t *testing.T) {
	// First test context
	cleanup1 := UseTempHomeDir(t)
	home1 := os.Getenv("HOME")
	testFile1 := home1 + "/isolation-test.txt"
	if err := os.WriteFile(testFile1, []byte("data1"), 0o600); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}
	cleanup1()

	// Second test context
	cleanup2 := UseTempHomeDir(t)
	defer cleanup2()
	home2 := os.Getenv("HOME")
	testFile2 := home2 + "/isolation-test.txt"

	// File from first context should not exist in second context
	if _, err := os.Stat(testFile2); err == nil {
		t.Error("files from previous temp HOME should not exist in new temp HOME")
	}
}

func TestUseTempHomeDirWithSubdirectories(t *testing.T) {
	cleanup := UseTempHomeDir(t)
	defer cleanup()

	home := os.Getenv("HOME")
	subdir := home + "/.config/test"

	// Create subdirectories
	if err := os.MkdirAll(subdir, 0o700); err != nil {
		t.Fatalf("should be able to create subdirectories in temp HOME: %v", err)
	}

	// Verify subdirectory exists
	if info, err := os.Stat(subdir); err != nil {
		t.Errorf("subdirectory should exist: %v", err)
	} else if !info.IsDir() {
		t.Error("path should be a directory")
	}
}

func TestUseTempHomeDirHelperMarking(t *testing.T) {
	// This test verifies that UseTempHomeDir is properly marked as a test helper
	// by checking that it uses t.Helper() - we can't directly test this,
	// but we can verify it behaves correctly in error scenarios

	cleanup := UseTempHomeDir(t)
	defer cleanup()

	// If UseTempHomeDir is properly marked with t.Helper(),
	// any test failures will point to this test, not to UseTempHomeDir
	home := os.Getenv("HOME")
	if home == "" {
		t.Error("HOME should be set by UseTempHomeDir")
	}
}

func TestUseTempHomeDirAcceptsTestingTB(t *testing.T) {
	// Verify UseTempHomeDir accepts testing.TB interface
	var tb testing.TB = t

	cleanup := UseTempHomeDir(tb)
	defer cleanup()

	home := os.Getenv("HOME")
	if home == "" {
		t.Error("UseTempHomeDir should work with testing.TB interface")
	}
}

func TestUseTempHomeDirCleanupIsIdempotent(t *testing.T) {
	cleanup := UseTempHomeDir(t)

	// Call cleanup multiple times - should not panic
	cleanup()
	cleanup()
	cleanup()

	// If we reach here, multiple cleanup calls didn't cause issues
	t.Log("cleanup is idempotent")
}

func TestUseTempHomeDirWithViperOperations(t *testing.T) {
	cleanup := UseTempHomeDir(t)
	defer cleanup()

	// Perform viper operations
	viper.Set("app.name", "gitgen")
	viper.Set("app.version", "1.0.0")

	if viper.GetString("app.name") != "gitgen" {
		t.Error("viper should work correctly in temp HOME")
	}

	cleanup()

	// Verify viper was reset
	if viper.GetString("app.name") != "" {
		t.Error("viper should be reset after cleanup")
	}
}

func TestUseTempHomeDirPreservesTempDirBehavior(t *testing.T) {
	cleanup := UseTempHomeDir(t)
	defer cleanup()

	home := os.Getenv("HOME")

	// The temp directory should be automatically cleaned up by t.TempDir()
	// We just verify it exists during the test
	if _, err := os.Stat(home); err != nil {
		t.Errorf("temp HOME directory should exist during test: %v", err)
	}
}