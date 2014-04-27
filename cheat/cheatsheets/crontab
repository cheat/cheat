# set a shell
SHELL=/bin/bash

# crontab format
* * * * *  command_to_execute
- - - - -
| | | | |
| | | | +- day of week (0 - 7) (where sunday is 0 and 7)
| | | +--- month (1 - 12)
| | +----- day (1 - 31)
| +------- hour (0 - 23)
+--------- minute (0 - 59)

# example entries
# every 15 min
*/15 * * * * /home/user/command.sh
# every midnight
0 * * * * /home/user/command.sh
# every Saturday at 8:05 AM
5 8 * * 6 /home/user/command.sh
