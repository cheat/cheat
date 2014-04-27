# All pacman commands are working the same way with yaourt.
# Just check the pacman cheatsheet.
# For instance, to install a package : 
pacman -S <package name>
yaourt -S <package name>
# The difference is that yaourt will also query the Arch User Repository,
# and if appropriate, donwload the source and build the package requested.

# Here are the commands yaourt provides while pacman doesn't :

# To search for a package and install it
yaourt <package name>

# To update the local package base and upgrade all out of date package, including the ones from 
AUR and the packages based on development repos (git, svn, hg...)
yaourt -Suya --devel

# For all of the above commands, if you want yaourt to stop asking constantly for confirmations, 
use the option --noconfirm

# To build a package from source
yaourt -Sb <package name>

