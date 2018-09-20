# psql is the PostgreSQL terminal interface. The following commands were tested on version 9.5.
# Connection options:
# -U username (if not specified current OS user is used).
# -p port.
# -h server hostname/address.

# Connect to a specific database:
psql -U postgres -h serverAddress -d dbName

# Get databases on a server:
psql -U postgres -h serverAddress --list

# Execute sql query and save output to file:
psql -U postgres -d dbName -c 'select * from tableName;' -o fileName

# Execute query and get tabular html output:
psql -U postgres -d dbName -H -c 'select * from tableName;'

# Execute query and save resulting rows to csv file:
psql -U postgres -d dbName -t -A -P fieldsep=',' -c 'select * from tableName;' -o fileName.csv

# Read commands from file:
psql -f fileName

# Restore databases from file:
psql -f fileName.backup postgres
