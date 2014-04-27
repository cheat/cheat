# Create a key

 gpg --gen-key


# Show keys

  To list a summary of all keys

    gpg --list-keys

  To show your public key

    gpg --armor --export

  To show the fingerprint for a key

    gpg --fingerprint KEY_ID

# Search for keys

  gpg --search-keys 'user@emailaddress.com'


# To Encrypt a File

  gpg --encrypt --recipient 'user@emailaddress.com' example.txt


# To Decrypt a File

  gpg --output example.txt --decrypt example.txt.gpg


# Export keys

  gpg --output ~/public_key.txt --armor --export KEY_ID
  gpg --output ~/private_key.txt --armor --export-secret-key KEY_ID

  Where KEY_ID is the 8 character GPG key ID.

  Store these files to a safe location, such as a USB drive, then
  remove the private key file.

    shred -zu ~/private_key.txt

# Import keys

  Retrieve the key files which you previously exported.

    gpg --import ~/public_key.txt
    gpg --allow-secret-key-import --import ~/private_key.txt

  Then delete the private key file.

    shred -zu ~/private_key.txt

# Revoke a key

  Create a revocation certificate.

    gpg --output ~/revoke.asc --gen-revoke KEY_ID

  Where KEY_ID is the 8 character GPG key ID.

  After creating the certificate import it.

    gpg --import ~/revoke.asc

  Then ensure that key servers know about the revokation.

    gpg --send-keys KEY_ID

# Signing and Verifying files

  If you're uploading files to launchpad you may also want to include
  a GPG signature file.

    gpg -ba filename

  or if you need to specify a particular key:

    gpg --default-key <key ID> -ba filename

  This then produces a file with a .asc extension which can be uploaded.
  If you need to set the default key more permanently then edit the
  file ~/.gnupg/gpg.conf and set the default-key parameter.

  To verify a downloaded file using its signature file.

  gpg --verify filename.asc

# Signing Public Keys

  Import the public key or retrieve it from a server.

    gpg --keyserver <keyserver> --recv-keys <Key_ID>

  Check its fingerprint against any previously stated value.

    gpg --fingerprint <Key_ID>

  Sign the key.

    gpg --sign-key <Key_ID>

  Upload the signed key to a server.

    gpg --keyserver <keyserver> --send-key <Key_ID>

# Change the email address associated with a GPG key

  gpg --edit-key <key ID>
  adduid

  Enter the new name and email address. You can then list the addresses with:

    list

  If you want to delete a previous email address first select it:

    uid <list number>

  Then delete it with:

    deluid

  To finish type:

    save

  Publish the key to a server:

    gpg --send-keys <key ID>

# Creating Subkeys

  Subkeys can be useful if you don't wish to have your main GPG key
  installed on multiple machines. In this way you can keep your
  master key safe and have subkeys with expiry periods or which may be
  separately revoked installed on various machines. This avoids
  generating entirely separate keys and so breaking any web of trust
  which has been established.

    gpg --edit-key <key ID>

  At the prompt type:

    addkey

  Choose RSA (sign only), 4096 bits and select an expiry period.
  Entropy will be gathered.

  At the prompt type:

    save

  You can also repeat the procedure, but selecting RSA (encrypt only).
  To remove the master key, leaving only the subkey/s in place:

    gpg --export-secret-subkeys <subkey ID> > subkeys
    gpg --export <key ID> > pubkeys
    gpg --delete-secret-key <key ID>

  Import the keys back.

    gpg --import pubkeys subkeys

  Verify the import.

    gpg -K

  Should show sec# instead of just sec.
