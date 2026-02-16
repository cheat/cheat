package repo

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-git/go-git/v5"
	gitconfig "github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// testCommitOpts returns a CommitOptions suitable for test commits.
func testCommitOpts() *git.CommitOptions {
	return &git.CommitOptions{
		Author: &object.Signature{
			Name:  "test",
			Email: "test@test",
			When:  time.Now(),
		},
	}
}

// initBareRepoWithCommit creates a bare git repository at dir with an initial
// commit, suitable for use as a remote.
func initBareRepoWithCommit(t *testing.T, dir string) {
	t.Helper()

	// init a non-bare repo to make the commit, then we'll clone it as bare
	tmpWork := t.TempDir()
	r, err := git.PlainInit(tmpWork, false)
	if err != nil {
		t.Fatalf("failed to init repo: %v", err)
	}

	f := filepath.Join(tmpWork, "README")
	if err := os.WriteFile(f, []byte("hello\n"), 0644); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}

	wt, err := r.Worktree()
	if err != nil {
		t.Fatalf("failed to get worktree: %v", err)
	}

	if _, err := wt.Add("README"); err != nil {
		t.Fatalf("failed to stage file: %v", err)
	}

	if _, err = wt.Commit("initial commit", testCommitOpts()); err != nil {
		t.Fatalf("failed to commit: %v", err)
	}

	// clone as bare into the target dir
	if _, err = git.PlainClone(dir, true, &git.CloneOptions{URL: tmpWork}); err != nil {
		t.Fatalf("failed to create bare clone: %v", err)
	}
}

// cloneLocal clones the bare repo at bareDir into a new working directory and
// returns the path.
func cloneLocal(t *testing.T, bareDir string) string {
	t.Helper()

	dir := t.TempDir()
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: bareDir,
	})
	if err != nil {
		t.Fatalf("failed to clone: %v", err)
	}

	return dir
}

// pushNewCommit clones bareDir into a temporary working copy, commits a new
// file, and pushes back to the bare repo.
func pushNewCommit(t *testing.T, bareDir string) {
	t.Helper()

	tmpWork := t.TempDir()
	r, err := git.PlainClone(tmpWork, false, &git.CloneOptions{URL: bareDir})
	if err != nil {
		t.Fatalf("failed to clone for push: %v", err)
	}

	if err := os.WriteFile(filepath.Join(tmpWork, "new.txt"), []byte("new\n"), 0644); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}

	wt, err := r.Worktree()
	if err != nil {
		t.Fatalf("failed to get worktree: %v", err)
	}
	if _, err := wt.Add("new.txt"); err != nil {
		t.Fatalf("failed to stage file: %v", err)
	}
	if _, err := wt.Commit("add new file", testCommitOpts()); err != nil {
		t.Fatalf("failed to commit: %v", err)
	}
	if err := r.Push(&git.PushOptions{}); err != nil {
		t.Fatalf("failed to push: %v", err)
	}
}

// generateTestKey creates an unencrypted ed25519 PEM private key file at path.
func generateTestKey(t *testing.T, path string) {
	t.Helper()

	_, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	der, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		t.Fatalf("failed to marshal key: %v", err)
	}

	pemBytes := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der})
	if err := os.WriteFile(path, pemBytes, 0600); err != nil {
		t.Fatalf("failed to write key file: %v", err)
	}
}

// --- Pull tests ---

func TestPull_NotARepo(t *testing.T) {
	dir := t.TempDir()

	err := Pull(dir)
	if err != git.ErrRepositoryNotExists {
		t.Fatalf("expected ErrRepositoryNotExists, got: %v", err)
	}
}

func TestPull_CleanAlreadyUpToDate(t *testing.T) {
	bare := t.TempDir()
	initBareRepoWithCommit(t, bare)
	clone := cloneLocal(t, bare)

	err := Pull(clone)
	if err != nil {
		t.Fatalf("expected nil (already up-to-date), got: %v", err)
	}
}

func TestPull_NewUpstreamChanges(t *testing.T) {
	bare := t.TempDir()
	initBareRepoWithCommit(t, bare)
	clone := cloneLocal(t, bare)

	// push a new commit to the bare repo after the clone
	pushNewCommit(t, bare)

	err := Pull(clone)
	if err != nil {
		t.Fatalf("expected nil (successful pull), got: %v", err)
	}

	// verify the new file was pulled
	if _, err := os.Stat(filepath.Join(clone, "new.txt")); err != nil {
		t.Fatalf("expected new.txt to exist after pull: %v", err)
	}
}

func TestPull_DirtyWorktree(t *testing.T) {
	bare := t.TempDir()
	initBareRepoWithCommit(t, bare)
	clone := cloneLocal(t, bare)

	// make the worktree dirty with a modified tracked file
	if err := os.WriteFile(filepath.Join(clone, "README"), []byte("changed\n"), 0644); err != nil {
		t.Fatalf("failed to modify file: %v", err)
	}

	err := Pull(clone)
	if err != ErrDirtyWorktree {
		t.Fatalf("expected ErrDirtyWorktree, got: %v", err)
	}
}

func TestPull_DirtyWorktreeUntracked(t *testing.T) {
	bare := t.TempDir()
	initBareRepoWithCommit(t, bare)
	clone := cloneLocal(t, bare)

	// make the worktree dirty with an untracked file
	if err := os.WriteFile(filepath.Join(clone, "untracked.txt"), []byte("new\n"), 0644); err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	err := Pull(clone)
	if err != ErrDirtyWorktree {
		t.Fatalf("expected ErrDirtyWorktree, got: %v", err)
	}
}

// --- sshAuth tests ---

func TestSshAuth_NonSSHRemote(t *testing.T) {
	bare := t.TempDir()
	initBareRepoWithCommit(t, bare)
	clone := cloneLocal(t, bare)

	r, err := git.PlainOpen(clone)
	if err != nil {
		t.Fatalf("failed to open repo: %v", err)
	}

	// the clone's origin is a local file:// path, not SSH
	auth, err := sshAuth(r)
	if err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
	if auth != nil {
		t.Fatalf("expected nil auth for non-SSH remote, got: %v", auth)
	}
}

func TestSshAuth_NoRemote(t *testing.T) {
	dir := t.TempDir()
	r, err := git.PlainInit(dir, false)
	if err != nil {
		t.Fatalf("failed to init repo: %v", err)
	}

	// repo has no remotes
	auth, err := sshAuth(r)
	if err == nil {
		t.Fatalf("expected error for missing remote, got auth: %v", auth)
	}
}

func TestSshAuth_SSHRemote(t *testing.T) {
	dir := t.TempDir()
	r, err := git.PlainInit(dir, false)
	if err != nil {
		t.Fatalf("failed to init repo: %v", err)
	}

	// add an SSH remote
	_, err = r.CreateRemote(&gitconfig.RemoteConfig{
		Name: "origin",
		URLs: []string{"git@github.com:example/repo.git"},
	})
	if err != nil {
		t.Fatalf("failed to create remote: %v", err)
	}

	// sshAuth should not return an error — even if no key is found, it
	// returns (nil, nil) rather than an error
	auth, err := sshAuth(r)
	if err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}

	// we can't predict whether auth is nil or non-nil here because it
	// depends on whether the test runner has SSH keys or an agent; just
	// verify it didn't error
	_ = auth
}

// --- findKeyFile tests ---

func TestFindKeyFile_ValidKey(t *testing.T) {
	sshDir := t.TempDir()
	generateTestKey(t, filepath.Join(sshDir, "id_ed25519"))

	auth := findKeyFile(sshDir, "git")
	if auth == nil {
		t.Fatal("expected non-nil auth for valid key file")
	}
}

func TestFindKeyFile_NoKeys(t *testing.T) {
	sshDir := t.TempDir()

	auth := findKeyFile(sshDir, "git")
	if auth != nil {
		t.Fatalf("expected nil auth for empty directory, got: %v", auth)
	}
}

func TestFindKeyFile_InvalidKey(t *testing.T) {
	sshDir := t.TempDir()
	// write garbage into a file named like a key
	if err := os.WriteFile(filepath.Join(sshDir, "id_ed25519"), []byte("not a key"), 0600); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}

	auth := findKeyFile(sshDir, "git")
	if auth != nil {
		t.Fatalf("expected nil auth for invalid key file, got: %v", auth)
	}
}

func TestFindKeyFile_SkipsInvalidFindsValid(t *testing.T) {
	sshDir := t.TempDir()

	// put garbage in id_rsa (tried first), valid key in id_ed25519 (tried later)
	if err := os.WriteFile(filepath.Join(sshDir, "id_rsa"), []byte("not a key"), 0600); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}
	generateTestKey(t, filepath.Join(sshDir, "id_ed25519"))

	auth := findKeyFile(sshDir, "git")
	if auth == nil {
		t.Fatal("expected non-nil auth; should skip invalid id_rsa and find id_ed25519")
	}
}
