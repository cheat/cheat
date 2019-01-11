# To copy a file from your local machine to a remote server:
scp foo.txt user@example.com:remote/dir

# To copy a file from a remote server to your local machine:
scp user@example.com:remote/dir/foo.txt local/dir

# To scp a file over a SOCKS proxy on localhost and port 9999 (see ssh for tunnel setup):
scp -o "ProxyCommand nc -x 127.0.0.1:9999 -X 4 %h %p" file.txt username@example2.com:/tmp/
