# Download a single file
curl http://path.to.the/file

# Download a file and specify a new filename
curl http://example.com/file.zip -o new_file.zip

# Download multiple files
curl -O URLOfFirstFile -O URLOfSecondFile

# Download all sequentially numbered files (1-24)
curl http://example.com/pic[1-24].jpg

# Download a file and pass HTTP Authentication
curl -u username:password URL 

# Download a file with a Proxy
curl -x proxysever.server.com:PORT http://addressiwantto.access

# Download a file from FTP
curl -u username:password -O ftp://example.com/pub/file.zip

# Get an FTP directory listing
curl ftp://username:password@example.com

# Resume a previously failed download
curl -C - -o partial_file.zip http://example.com/file.zip

# Fetch only the HTTP headers from a response
curl -I http://example.com

# Fetch your external IP and network info as JSON
curl http://ifconfig.me/all/json

# Limit the rate of a download
curl --limit-rate 1000B -O http://path.to.the/file
