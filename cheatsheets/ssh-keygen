# To generate an SSH key:
ssh-keygen -t rsa

# To generate a 4096-bit SSH key:
ssh-keygen -t rsa -b 4096

# To update a passphrase on a key
ssh-keygen -p -P old_passphrase -N new_passphrase -f /path/to/keyfile

# To remove a passphrase on a key
ssh-keygen -p -P old_passphrase -N '' -f /path/to/keyfile

# To generate a 4096 bit RSA key with a passphase and comment containing the user and hostname
ssh-keygen -t rsa -b 4096 -C "$USER@$HOSTNAME" -P passphrase
