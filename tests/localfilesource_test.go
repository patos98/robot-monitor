package tests

import (
	"robot-monitor/filesource"
	"strings"
	"testing"
)

func TestLocalFileSource(t *testing.T) {
	testFilePath := "testdata/local_file_source_test.txt"
	expectedContent := "test content"

	// replace this block with file source object
	content, err := filesource.Local(testFilePath).GetContent()
	if err != nil {
		t.Fatal(err)
	}

	if expectedContent != strings.TrimSpace(string(content)) {
		t.Fatalf("Content does not match expected!\nexpected: %s\nactual: %s", expectedContent, content)
	}
}
