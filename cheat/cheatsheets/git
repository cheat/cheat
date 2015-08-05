# To set your identity:
git config --global user.name "John Doe"
git config --global user.email johndoe@example.com

# To set your editor:
git config --global core.editor emacs

# To enable color:
git config --global color.ui true

# To stage all changes for commit:
git add --all

# To commit staged changes
git commit -m "Your commit message"

# To edit previous commit message
git commit --amend

# Git commit in the past
git commit --date="`date --date='2 day ago'`"
git commit --date="Jun 13 18:30:25 IST 2015"
# more recent versions of Git also support --date="2 days ago" directly

# To change the date of an existing commit
git filter-branch --env-filter \
    'if [ $GIT_COMMIT = 119f9ecf58069b265ab22f1f97d2b648faf932e0 ]
     then
         export GIT_AUTHOR_DATE="Fri Jan 2 21:38:53 2009 -0800"
         export GIT_COMMITTER_DATE="Sat May 19 01:01:01 2007 -0700"
     fi'

# To removed staged and working directory changes
git reset --hard

# To go 2 commits back
git reset --hard HEAD~2

# To remove untracked files
git clean -f -d

# To remove untracked and ignored files
git clean -f -d -x

# To push to the tracked master branch:
git push origin master

# To push to a specified repository:
git push git@github.com:username/project.git

# To delete the branch "branch_name"
git branch -D branch_name

# To make an exisiting branch track a remote branch
git branch -u upstream/foo

# To see who commited which line in a file
git blame filename

# To sync a fork with the master repo:
git remote add upstream git@github.com:name/repo.git    # Set a new repo
git remote -v                                           # Confirm new remote repo
git fetch upstream                                      # Get branches
git branch -va                                          # List local - remote branches
git checkout master                                     # Checkout local master branch
git checkout -b new_branch                              # Create and checkout a new branch
git merge upstream/master                               # Merge remote into local repo
git show 83fb499                                        # Show what a commit did.
git show 83fb499:path/fo/file.ext                       # Shows the file as it appeared at 83fb499.
git diff branch_1 branch_2                              # Check difference between branches
git log                                                 # Show all the commits
git status                                              # Show the changes from last commit

# Commit history of a set of files
git log --pretty=email --patch-with-stat --reverse --full-index -- Admin\*.py > Sripts.patch

# Import commits from another repo
git --git-dir=../some_other_repo/.git format-patch -k -1 --stdout <commit SHA> | git am -3 -k

# View commits that will be pushed
git log @{u}..

# View changes that are new on a feature branch
git log -p feature --not master
git diff master...feature

# Interactive rebase for the last 7 commits
git rebase -i @~7

# Diff files WITHOUT considering them a part of git
# This can be used to diff files that are not in a git repo!
git diff --no-index path/to/file/A path/to/file/B

# To pull changes while overwriting any local commits
git fetch --all
git reset --hard origin/master

# Update all your submodules
git submodule update --init --recursive

# Perform a shallow clone to only get latest commits
# (helps save data when cloning large repos)
git clone --depth 1 <remote-url>

# To unshallow a clone
git pull --unshallow
