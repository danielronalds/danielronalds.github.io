package parsing

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/danielronalds/danielronalds.github.io/internal/site"
)

func ParseRoutes(logger *slog.Logger, srcDir, destDir string) ([]site.Route, error) {
	sourceInfo, err := os.Lstat(srcDir)
	if err != nil {
		return nil, fmt.Errorf("inspect source directory %q: %w", srcDir, err)
	}

	if !sourceInfo.IsDir() {
		return nil, fmt.Errorf("source path %q is not a directory", srcDir)
	}

	routes := make([]site.Route, 0)
	sourcesByDestination := make(map[string]string)

	err = filepath.WalkDir(srcDir, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if path == srcDir {
			return nil
		}

		if entry.Type()&os.ModeSymlink != 0 {
			logger.Warn("skipping symbolic link", "path", path, "reason", "symbolic links are not supported")
			return nil
		}

		if strings.HasPrefix(entry.Name(), ".") {
			logger.Debug("skipping hidden path", "path", path, "reason", "hidden path")
			if entry.IsDir() {
				return fs.SkipDir
			}
			return nil
		}

		if entry.IsDir() {
			return nil
		}

		if filepath.Ext(entry.Name()) != ".md" {
			logger.Debug("skipping file", "path", path, "reason", "unsupported file type")
			return nil
		}

		relativePath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return fmt.Errorf("calculate path relative to source directory for %q: %w", path, err)
		}

		var destinationPath string
		if entry.Name() == "index.md" {
			destinationPath = filepath.Join(destDir, filepath.Dir(relativePath), "index.html")
		} else {
			pagePath := strings.TrimSuffix(relativePath, ".md")
			destinationPath = filepath.Join(destDir, pagePath, "index.html")
		}

		if existingSource, exists := sourcesByDestination[destinationPath]; exists {
			return fmt.Errorf(
				"route destination collision: %q and %q both map to %q",
				existingSource,
				path,
				destinationPath,
			)
		}

		sourcesByDestination[destinationPath] = path
		routes = append(routes, site.Route{Src: path, Dest: destinationPath})
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse routes from %q: %w", srcDir, err)
	}

	sort.Slice(routes, func(leftIndex, rightIndex int) bool {
		return routes[leftIndex].Src < routes[rightIndex].Src
	})

	return routes, nil
}
