package branch

import (
	"errors"

	"path/filepath"

	git "gopkg.in/src-d/go-git.v4"
)

var ErrNoGitParent = errors.New("Could not find a parent git directory.")

// Open recursively searches backwards for a git repository root. If it finds one, it returns that repository object.
// If it does not, returns the error ErrNotGitParent
func Open(pathSplit []string) (*git.Repository, error) {
	path := filepath.Join(pathSplit...)
	path = "/" + path
	r, err := git.PlainOpen(path)
	if err != nil {
		if path == "/" {
			return nil, ErrNoGitParent
		}
		pathSplit = pathSplit[:len(pathSplit)-1]
		return Open(pathSplit)
	}
	return r, nil
}
