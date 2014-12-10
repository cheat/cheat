# Single target scan:
nmap [target]

# Scan from a list of targets:
nmap -iL [list.txt]

# iPv6:
nmap -6 [target]

# OS detection:
nmap -O --osscan_guess [target]

# Save output to text file:
nmap -oN [output.txt] [target]

# Save output to xml file:
nmap -oX [output.xml] [target]

# Scan a specific port:
nmap -source-port [port] [target]

# Do an aggressive scan:
nmap -A [target]

# Speedup your scan:
# -n => disable ReverseDNS
# --min-rate=X => min X packets / sec
nmap -T5 --min-parallelism=50 -n --min-rate=300 [target]

# Traceroute:
nmap -traceroute [target]

# Ping scan only: -sP
# Don't ping:     -PN <- Use full if a host don't reply to a ping.
# TCP SYN ping:   -PS
# TCP ACK ping:   -PA
# UDP ping:       -PU
# ARP ping:       -PR

# Example: Ping scan all machines on a class C network
nmap -sP 192.168.0.0/24

# Force TCP scan: -sT
# Force UDP scan: -sU

# Use some script:
nmap --script default,safe

# Loads the script in the default category, the banner script, and all .nse files in the directory /home/user/customscripts.
nmap --script default,banner,/home/user/customscripts

# Loads all scripts whose name starts with http-, such as http-auth and http-open-proxy.
nmap --script 'http-*'

# Loads every script except for those in the intrusive category.
nmap --script "not intrusive"

# Loads those scripts that are in both the default and safe categories.
nmap --script "default and safe"

# Loads scripts in the default, safe, or intrusive categories, except for those whose names start with http-.
nmap --script "(default or safe or intrusive) and not http-*"

# Scan for the heartbleed
# -pT:443 => Scan only port 443 with TCP (T:)
nmap -T5 --min-parallelism=50 -n --script "ssl-heartbleed" -pT:443 127.0.0.1

# Show all informations (debug mode)
nmap -d ...
