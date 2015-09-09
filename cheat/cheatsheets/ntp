# Verify ntpd running:
service ntp status

# Start ntpd if not running:
service ntp start

# Display current hardware clock value:
sudo hwclock -r

# Apply system time to hardware time:
sudo hwclock --systohc

# Apply hardware time to system time:
sudo hwclock --hctosys

# Set hwclock to local time:
sudo hwclock --localtime

# Set hwclock to UTC:
sudo hwclock --utc

# Set hwclock manually:
sudo hwclock --set --date="8/10/15 13:10:05"

# Query surrounding stratum time servers
ntpq -pn

# Config file:
/etc/ntp.conf

# Driftfile:
location of "drift" of your system clock compared to ntp servers
/var/lib/ntp/ntp.drift
