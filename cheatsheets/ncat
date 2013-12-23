# Connect mode (ncat is client) | default port is 31337
ncat <host> [<port>]

# Listen mode (ncat is server) | default port is 31337
ncat -l [<host>] [<port>]

# Transfer file (closes after one transfer)
ncat -l [<host>] [<port>] < file

# Transfer file (stays open for multiple transfers)
ncat -l --keep-open [<host>] [<port>] < file

# Receive file
ncat [<host>] [<port>] > file

# Brokering | allows for multiple clients to connect
ncat -l --broker [<host>] [<port>]

# Listen with SSL | many options, use ncat --help for full list
ncat -l --ssl [<host>] [<port>]

# Access control
ncat -l --allow <ip>
ncat -l --deny <ip>

# Proxying
ncat --proxy <proxyhost>[:<proxyport>] --proxy-type {http | socks4} <host>[<port>]

# Chat server | can use brokering for multi-user chat
ncat -l --chat [<host>] [<port>]
