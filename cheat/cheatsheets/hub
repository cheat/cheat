As a contributor to open-source
-------------------------------

# clone your own project
$ git clone dotfiles
→ git clone git://github.com/YOUR_USER/dotfiles.git

# clone another project
$ git clone github/hub
→ git clone git://github.com/github/hub.git

# see the current project's issues
$ git browse -- issues
→ open https://github.com/github/hub/issues

# open another project's wiki
$ git browse mojombo/jekyll wiki
→ open https://github.com/mojombo/jekyll/wiki

## Example workflow for contributing to a project:
$ git clone github/hub
$ cd hub
# create a topic branch
$ git checkout -b feature
→ ( making changes ... )
$ git commit -m "done with feature"
# It's time to fork the repo!
$ git fork
→ (forking repo on GitHub...)
→ git remote add YOUR_USER git://github.com/YOUR_USER/hub.git
# push the changes to your new remote
$ git push YOUR_USER feature
# open a pull request for the topic branch you've just pushed
$ git pull-request
→ (opens a text editor for your pull request message)


As an open-source maintainer
----------------------------

# fetch from multiple trusted forks, even if they don't yet exist as remotes
$ git fetch mislav,cehoffman
→ git remote add mislav git://github.com/mislav/hub.git
→ git remote add cehoffman git://github.com/cehoffman/hub.git
→ git fetch --multiple mislav cehoffman

# check out a pull request for review
$ git checkout https://github.com/github/hub/pull/134
→ (creates a new branch with the contents of the pull request)

# directly apply all commits from a pull request to the current branch
$ git am -3 https://github.com/github/hub/pull/134

# cherry-pick a GitHub URL
$ git cherry-pick https://github.com/xoebus/hub/commit/177eeb8
→ git remote add xoebus git://github.com/xoebus/hub.git
→ git fetch xoebus
→ git cherry-pick 177eeb8

# `am` can be better than cherry-pick since it doesn't create a remote
$ git am https://github.com/xoebus/hub/commit/177eeb8

# open the GitHub compare view between two releases
$ git compare v0.9..v1.0

# put compare URL for a topic branch to clipboard
$ git compare -u feature | pbcopy

# create a repo for a new project
$ git init
$ git add . && git commit -m "It begins."
$ git create -d "My new thing"
→ (creates a new project on GitHub with the name of current directory)
$ git push origin master
