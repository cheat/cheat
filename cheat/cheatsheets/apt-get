# Desc: Allows to update the operating system

# To fetch package list
apt-get update

# To download and install updates without installing new package.
apt-get upgrade

# To download and install the updates AND install new necessary packages
apt-get dist-upgrade

# Full command:
apt-get update && apt-get dist-upgrade

# To install a new package(s)
apt-get install package(s)

# Download a package without installing it. (The package will be downloaded in your current working dir)
apt-get download modsecurity-crs

# Change Cache dir and archive dir (where .deb are stored).
apt-get -o Dir::Cache="/path/to/destination/dir/" -o Dir::Cache::archives="./" install ...

# Show apt-get installed packages.
grep 'install ' /var/log/dpkg.log
