# All the following command work as well with multiple package names

# To search for a package
pacman -Ss <package name>

# To update the local package base and upgrade all out of date packages
pacman -Suy

# To install a package
pacman -S <package name>

# To uninstall a package
pacman -R <package name>

# To uninstall a package and his depedencies, removing all new orphans
pacman -Rcs <package name>

# To get informations about a package
pacman -Si <package name>

# To install a package from builded package file (.tar.xz)
pacman -U <file name/file url>

# To list the commands provided by an installed package
pacman -Ql <package name> | sed -n -e 's/.*\/bin\///p' | tail -n +2

# To list explicitly installed packages
pacman -Qe

# To list orphan packages (installed as dependencies and not required anymore)
pacman -Qdt


# You can't directly install packages from the Arch User Database (AUR) with pacman.
# You need yaourt to perform that. But considering yaourt itself is in the AUR, here is how to build a package from its tarball.
# Installing a package from AUR is a relatively simple process:
# - Retrieve the archive corresponding to your package from AUR website
# - Extract the archive (preferably in a folder for this purpose)
# - Run makepkg in the extracted directory. (makepkg-s allows you to install any dependencies automatically from deposits.)
# - Install the package created using pacman
# Assuming $pkgname contains the package name.
wget "https://aur.archlinux.org/packages/${pkgname::2}/$pkgname/$pkgname.tar.gz"
tar zxvf "$pkgname.tar.gz"
cd "$pkgname"
# Build the package
makepkg -s
# Install
sudo pacman -U <package file (.pkg.tar.xz)>
