package branch

import (
	"errors"
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

var ErrNoGitParent = errors.New("Could not find a parent git directory.")

// Open recursively searches backwards for a git repository root. If it finds one, it returns that repository object.
// If it does not, returns the error ErrNotGitParent
func Open(path string) (*git.Repository, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		err = os.Chdir("..")
		if err != nil {
			return nil, err
		}
		d, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		if d == "/" {
			return nil, ErrNoGitParent
		}
		return Open(d)
	}
	return r, nil
}
