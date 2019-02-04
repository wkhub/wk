package backends

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4"

	"github.com/wkhub/wk/fs"
	"github.com/wkhub/wk/user"
)

type GitBackend struct {
}

var gitPrefixes = []string{}

func (b GitBackend) Name() string {
	return "Git"
}

func (b GitBackend) Match(source string) bool {
	return strings.HasSuffix(source, ".git")
}

func (b GitBackend) Fetch(source string) (string, error) {
	currentUser := user.Current()
	basename := strings.TrimSuffix(path.Base(source), ".git")
	target := filepath.Join(currentUser.Home.TemplatesDir(), basename)

	if fs.Exists(target) {
		repo, err := git.PlainOpen(target)
		if err != nil {
			return "", err
		}
		// Get the working directory for the repository
		worktree, err := repo.Worktree()
		if err != nil {
			return "", err
		}
		err = worktree.Pull(&git.PullOptions{RemoteName: "origin", Progress: os.Stdout})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			return "", err
		}
	} else {
		_, err := git.PlainClone(target, false, &git.CloneOptions{
			URL:      source,
			Progress: os.Stdout,
		})
		if err != nil {
			return "", err
		}
	}
	return target, nil
}

func init() {
	Register(GitBackend{})
}
