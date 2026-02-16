package repo

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	gitssh "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/mitchellh/go-homedir"
)

// ErrDirtyWorktree indicates that the worktree has uncommitted changes.
var ErrDirtyWorktree = errors.New("dirty worktree")

// Pull performs a git pull on the repository at path. It returns
// ErrDirtyWorktree if the worktree has uncommitted changes, and
// git.ErrRepositoryNotExists if path is not a git repository.
func Pull(path string) error {

	// open the repository
	r, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	// get the worktree
	wt, err := r.Worktree()
	if err != nil {
		return err
	}

	// check if the worktree is clean
	status, err := wt.Status()
	if err != nil {
		return err
	}
	if !status.IsClean() {
		return ErrDirtyWorktree
	}

	// build pull options, using SSH auth when the remote is SSH
	opts := &git.PullOptions{}
	if auth, err := sshAuth(r); err == nil && auth != nil {
		opts.Auth = auth
	}

	// pull
	err = wt.Pull(opts)
	if errors.Is(err, git.NoErrAlreadyUpToDate) {
		return nil
	}

	return err
}

// defaultKeyFiles are the SSH key filenames tried in order, matching the
// default behavior of OpenSSH.
var defaultKeyFiles = []string{
	"id_rsa",
	"id_ecdsa",
	"id_ecdsa_sk",
	"id_ed25519",
	"id_ed25519_sk",
	"id_dsa",
}

// sshAuth returns an appropriate SSH auth method if the origin remote uses
// the SSH protocol, or nil if it does not. It tries the SSH agent first, then
// falls back to default key files in ~/.ssh/.
func sshAuth(r *git.Repository) (transport.AuthMethod, error) {
	remote, err := r.Remote("origin")
	if err != nil {
		return nil, err
	}

	urls := remote.Config().URLs
	if len(urls) == 0 {
		return nil, nil
	}

	ep, err := transport.NewEndpoint(urls[0])
	if err != nil {
		return nil, err
	}

	if ep.Protocol != "ssh" {
		return nil, nil
	}

	user := ep.User
	if user == "" {
		user = "git"
	}

	// try default key files first — this is more reliable than the SSH
	// agent, which may report success even when no keys are loaded
	home, err := homedir.Dir()
	if err == nil {
		if auth := findKeyFile(filepath.Join(home, ".ssh"), user); auth != nil {
			return auth, nil
		}
	}

	// fall back to SSH agent
	if auth, err := gitssh.NewSSHAgentAuth(user); err == nil {
		return auth, nil
	}

	return nil, nil
}

// findKeyFile looks for a usable SSH private key in sshDir, trying the
// standard OpenSSH default filenames in order. Returns nil if no usable key
// is found.
func findKeyFile(sshDir, user string) transport.AuthMethod {
	for _, name := range defaultKeyFiles {
		keyPath := filepath.Join(sshDir, name)
		if _, err := os.Stat(keyPath); err != nil {
			continue
		}
		auth, err := gitssh.NewPublicKeysFromFile(user, keyPath, "")
		if err != nil {
			continue
		}
		return auth
	}
	return nil
}
