package parsing

import (
	"bytes"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestParseRoutesProducesRoutesForExampleTree(t *testing.T) {
	tempDir := t.TempDir()
	srcDir := filepath.Join(tempDir, "web")
	destDir := filepath.Join(tempDir, "dist")

	writeTestFile(t, filepath.Join(srcDir, "recommended-readings.md"))
	writeTestFile(t, filepath.Join(srcDir, "posts", "index.md"))
	writeTestFile(t, filepath.Join(srcDir, "projects.md"))
	writeTestFile(t, filepath.Join(srcDir, "index.md"))
	writeTestFile(t, filepath.Join(srcDir, "posts", "building-wade.md"))

	routes, err := ParseRoutes(slog.New(slog.NewTextHandler(io.Discard, nil)), srcDir, destDir)
	if err != nil {
		t.Fatalf("ParseRoutes() error = %v", err)
	}

	expectedRoutes := []Route{
		{Src: filepath.Join(srcDir, "index.md"), Dest: filepath.Join(destDir, "index.html")},
		{Src: filepath.Join(srcDir, "posts", "building-wade.md"), Dest: filepath.Join(destDir, "posts", "building-wade", "index.html")},
		{Src: filepath.Join(srcDir, "posts", "index.md"), Dest: filepath.Join(destDir, "posts", "index.html")},
		{Src: filepath.Join(srcDir, "projects.md"), Dest: filepath.Join(destDir, "projects", "index.html")},
		{Src: filepath.Join(srcDir, "recommended-readings.md"), Dest: filepath.Join(destDir, "recommended-readings", "index.html")},
	}
	if !reflect.DeepEqual(routes, expectedRoutes) {
		t.Errorf("ParseRoutes() = %#v, want %#v", routes, expectedRoutes)
	}

	if _, err := os.Stat(destDir); !os.IsNotExist(err) {
		t.Errorf("ParseRoutes() created destination directory, os.Stat() error = %v", err)
	}
}

func TestParseRoutesMapsNestedMarkdownPages(t *testing.T) {
	tempDir := t.TempDir()
	srcDir := filepath.Join(tempDir, "web")
	destDir := filepath.Join(tempDir, "dist")
	pagePath := filepath.Join(srcDir, "posts", "2026", "hello.md")
	writeTestFile(t, pagePath)

	routes, err := ParseRoutes(slog.New(slog.NewTextHandler(io.Discard, nil)), srcDir, destDir)
	if err != nil {
		t.Fatalf("ParseRoutes() error = %v", err)
	}

	expectedRoutes := []Route{{
		Src:  pagePath,
		Dest: filepath.Join(destDir, "posts", "2026", "hello", "index.html"),
	}}
	if !reflect.DeepEqual(routes, expectedRoutes) {
		t.Errorf("ParseRoutes() = %#v, want %#v", routes, expectedRoutes)
	}
}

func TestParseRoutesOnlyTreatsExactIndexFilenameAsIndexPage(t *testing.T) {
	tempDir := t.TempDir()
	srcDir := filepath.Join(tempDir, "web")
	destDir := filepath.Join(tempDir, "dist")

	writeTestFile(t, filepath.Join(srcDir, "index.md"))
	writeTestFile(t, filepath.Join(srcDir, "Index.md"))
	writeTestFile(t, filepath.Join(srcDir, "posts", "index.md"))

	routes, err := ParseRoutes(slog.New(slog.NewTextHandler(io.Discard, nil)), srcDir, destDir)
	if err != nil {
		t.Fatalf("ParseRoutes() error = %v", err)
	}

	expectedRoutes := []Route{
		{Src: filepath.Join(srcDir, "Index.md"), Dest: filepath.Join(destDir, "Index", "index.html")},
		{Src: filepath.Join(srcDir, "index.md"), Dest: filepath.Join(destDir, "index.html")},
		{Src: filepath.Join(srcDir, "posts", "index.md"), Dest: filepath.Join(destDir, "posts", "index.html")},
	}
	if !reflect.DeepEqual(routes, expectedRoutes) {
		t.Errorf("ParseRoutes() = %#v, want %#v", routes, expectedRoutes)
	}
}

func TestParseRoutesSkipsUnsupportedFiles(t *testing.T) {
	tempDir := t.TempDir()
	srcDir := filepath.Join(tempDir, "web")
	destDir := filepath.Join(tempDir, "dist")
	unsupportedPaths := []string{
		filepath.Join(srcDir, "document.html"),
		filepath.Join(srcDir, "image.png"),
		filepath.Join(srcDir, "LICENSE"),
		filepath.Join(srcDir, "page.MD"),
	}
	for _, path := range unsupportedPaths {
		writeTestFile(t, path)
	}
	writeTestFile(t, filepath.Join(srcDir, "page.md"))

	var logs bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&logs, &slog.HandlerOptions{Level: slog.LevelDebug}))
	routes, err := ParseRoutes(logger, srcDir, destDir)
	if err != nil {
		t.Fatalf("ParseRoutes() error = %v", err)
	}

	expectedRoutes := []Route{{
		Src:  filepath.Join(srcDir, "page.md"),
		Dest: filepath.Join(destDir, "page", "index.html"),
	}}
	if !reflect.DeepEqual(routes, expectedRoutes) {
		t.Errorf("ParseRoutes() = %#v, want %#v", routes, expectedRoutes)
	}

	logOutput := logs.String()
	if debugLogCount := strings.Count(logOutput, "level=DEBUG"); debugLogCount != len(unsupportedPaths) {
		t.Errorf("ParseRoutes() wrote %d DEBUG logs, want %d; logs = %s", debugLogCount, len(unsupportedPaths), logOutput)
	}
	for _, path := range unsupportedPaths {
		if !strings.Contains(logOutput, path) {
			t.Errorf("ParseRoutes() logs do not contain skipped path %q; logs = %s", path, logOutput)
		}
	}
	if !strings.Contains(logOutput, `reason="unsupported file type"`) {
		t.Errorf("ParseRoutes() logs do not contain unsupported file reason; logs = %s", logOutput)
	}
}

func TestParseRoutesSkipsHiddenFiles(t *testing.T) {
	tempDir := t.TempDir()
	srcDir := filepath.Join(tempDir, "web")
	destDir := filepath.Join(tempDir, "dist")
	hiddenPath := filepath.Join(srcDir, ".draft.md")
	writeTestFile(t, hiddenPath)
	writeTestFile(t, filepath.Join(srcDir, "published.md"))

	var logs bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&logs, &slog.HandlerOptions{Level: slog.LevelDebug}))
	routes, err := ParseRoutes(logger, srcDir, destDir)
	if err != nil {
		t.Fatalf("ParseRoutes() error = %v", err)
	}

	expectedRoutes := []Route{{
		Src:  filepath.Join(srcDir, "published.md"),
		Dest: filepath.Join(destDir, "published", "index.html"),
	}}
	if !reflect.DeepEqual(routes, expectedRoutes) {
		t.Errorf("ParseRoutes() = %#v, want %#v", routes, expectedRoutes)
	}

	logOutput := logs.String()
	if !strings.Contains(logOutput, "level=DEBUG") || !strings.Contains(logOutput, hiddenPath) || !strings.Contains(logOutput, `reason="hidden path"`) {
		t.Errorf("ParseRoutes() did not log hidden file at DEBUG; logs = %s", logOutput)
	}
}

func TestParseRoutesSkipsHiddenDirectoriesWithoutTraversingThem(t *testing.T) {
	tempDir := t.TempDir()
	srcDir := filepath.Join(tempDir, "web")
	destDir := filepath.Join(tempDir, "dist")
	hiddenDir := filepath.Join(srcDir, ".drafts")
	writeTestFile(t, filepath.Join(hiddenDir, "index.md"))
	writeTestFile(t, filepath.Join(hiddenDir, "nested.md"))

	var logs bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&logs, &slog.HandlerOptions{Level: slog.LevelDebug}))
	routes, err := ParseRoutes(logger, srcDir, destDir)
	if err != nil {
		t.Fatalf("ParseRoutes() error = %v", err)
	}
	if len(routes) != 0 {
		t.Errorf("ParseRoutes() = %#v, want no routes", routes)
	}

	logOutput := logs.String()
	if debugLogCount := strings.Count(logOutput, "level=DEBUG"); debugLogCount != 1 {
		t.Errorf("ParseRoutes() wrote %d DEBUG logs, want 1; logs = %s", debugLogCount, logOutput)
	}
	if !strings.Contains(logOutput, hiddenDir) {
		t.Errorf("ParseRoutes() logs do not contain hidden directory %q; logs = %s", hiddenDir, logOutput)
	}
	if strings.Contains(logOutput, "index.md") || strings.Contains(logOutput, "nested.md") {
		t.Errorf("ParseRoutes() logged contents of hidden directory; logs = %s", logOutput)
	}
}

func TestParseRoutesSkipsSymbolicLinks(t *testing.T) {
	tempDir := t.TempDir()
	srcDir := filepath.Join(tempDir, "web")
	destDir := filepath.Join(tempDir, "dist")
	targetFile := filepath.Join(tempDir, "targets", "page.md")
	targetDir := filepath.Join(tempDir, "targets", "posts")
	writeTestFile(t, targetFile)
	writeTestFile(t, filepath.Join(targetDir, "index.md"))
	if err := os.MkdirAll(srcDir, 0o755); err != nil {
		t.Fatalf("os.MkdirAll() error = %v", err)
	}

	linkedFile := filepath.Join(srcDir, "linked.md")
	if err := os.Symlink(targetFile, linkedFile); err != nil {
		t.Skipf("os.Symlink() unavailable: %v", err)
	}
	linkedDir := filepath.Join(srcDir, "linked-posts")
	if err := os.Symlink(targetDir, linkedDir); err != nil {
		t.Skipf("os.Symlink() unavailable: %v", err)
	}

	var logs bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&logs, &slog.HandlerOptions{Level: slog.LevelDebug}))
	routes, err := ParseRoutes(logger, srcDir, destDir)
	if err != nil {
		t.Fatalf("ParseRoutes() error = %v", err)
	}
	if len(routes) != 0 {
		t.Errorf("ParseRoutes() = %#v, want no routes", routes)
	}

	logOutput := logs.String()
	if warningCount := strings.Count(logOutput, "level=WARN"); warningCount != 2 {
		t.Errorf("ParseRoutes() wrote %d WARN logs, want 2; logs = %s", warningCount, logOutput)
	}
	for _, path := range []string{linkedFile, linkedDir} {
		if !strings.Contains(logOutput, path) {
			t.Errorf("ParseRoutes() logs do not contain symlink path %q; logs = %s", path, logOutput)
		}
	}
	if !strings.Contains(logOutput, `reason="symbolic links are not supported"`) {
		t.Errorf("ParseRoutes() logs do not contain symlink reason; logs = %s", logOutput)
	}
}

func TestParseRoutesRejectsDestinationCollisions(t *testing.T) {
	tempDir := t.TempDir()
	srcDir := filepath.Join(tempDir, "web")
	destDir := filepath.Join(tempDir, "dist")
	pagePath := filepath.Join(srcDir, "posts.md")
	indexPath := filepath.Join(srcDir, "posts", "index.md")
	writeTestFile(t, pagePath)
	writeTestFile(t, indexPath)

	routes, err := ParseRoutes(slog.New(slog.NewTextHandler(io.Discard, nil)), srcDir, destDir)
	if err == nil {
		t.Fatal("ParseRoutes() error = nil, want destination collision error")
	}
	if routes != nil {
		t.Errorf("ParseRoutes() routes = %#v, want nil", routes)
	}

	errorMessage := err.Error()
	for _, expectedText := range []string{pagePath, indexPath, filepath.Join(destDir, "posts", "index.html")} {
		if !strings.Contains(errorMessage, expectedText) {
			t.Errorf("ParseRoutes() error = %q, want it to contain %q", errorMessage, expectedText)
		}
	}
}

func TestParseRoutesRejectsMissingSourceDirectory(t *testing.T) {
	tempDir := t.TempDir()
	srcDir := filepath.Join(tempDir, "missing")

	routes, err := ParseRoutes(slog.New(slog.NewTextHandler(io.Discard, nil)), srcDir, filepath.Join(tempDir, "dist"))
	if err == nil {
		t.Fatal("ParseRoutes() error = nil, want missing source directory error")
	}
	if routes != nil {
		t.Errorf("ParseRoutes() routes = %#v, want nil", routes)
	}
	if !strings.Contains(err.Error(), srcDir) {
		t.Errorf("ParseRoutes() error = %q, want it to contain %q", err.Error(), srcDir)
	}
}

func TestParseRoutesRejectsSourcePathThatIsAFile(t *testing.T) {
	tempDir := t.TempDir()
	srcPath := filepath.Join(tempDir, "web.md")
	writeTestFile(t, srcPath)

	routes, err := ParseRoutes(slog.New(slog.NewTextHandler(io.Discard, nil)), srcPath, filepath.Join(tempDir, "dist"))
	if err == nil {
		t.Fatal("ParseRoutes() error = nil, want source path type error")
	}
	if routes != nil {
		t.Errorf("ParseRoutes() routes = %#v, want nil", routes)
	}
	if !strings.Contains(err.Error(), "is not a directory") {
		t.Errorf("ParseRoutes() error = %q, want not-a-directory error", err.Error())
	}
}

func TestParseRoutesReturnsRoutesInSourcePathOrder(t *testing.T) {
	tempDir := t.TempDir()
	srcDir := filepath.Join(tempDir, "web")
	destDir := filepath.Join(tempDir, "dist")
	writeTestFile(t, filepath.Join(srcDir, "zebra.md"))
	writeTestFile(t, filepath.Join(srcDir, "middle", "page.md"))
	writeTestFile(t, filepath.Join(srcDir, "alpha.md"))

	routes, err := ParseRoutes(slog.New(slog.NewTextHandler(io.Discard, nil)), srcDir, destDir)
	if err != nil {
		t.Fatalf("ParseRoutes() error = %v", err)
	}

	expectedSources := []string{
		filepath.Join(srcDir, "alpha.md"),
		filepath.Join(srcDir, "middle", "page.md"),
		filepath.Join(srcDir, "zebra.md"),
	}
	for index, expectedSource := range expectedSources {
		if routes[index].Src != expectedSource {
			t.Errorf("ParseRoutes() route %d source = %q, want %q", index, routes[index].Src, expectedSource)
		}
	}
}

func writeTestFile(t *testing.T, path string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("os.MkdirAll() error = %v", err)
	}
	if err := os.WriteFile(path, []byte("content"), 0o644); err != nil {
		t.Fatalf("os.WriteFile() error = %v", err)
	}
}
