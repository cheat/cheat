# Read from {/dev/urandom} 2*512 Bytes and put it into {/tmp/test.txt}
# Note: At the first iteration, we read 512 Bytes.
# Note: At the second iteration, we read 512 Bytes.
dd if=/dev/urandom of=/tmp/test.txt count=2 bs=512

# Watch the progress of 'dd'
dd if=/dev/zero of=/dev/null bs=4KB &; export dd_pid=`pgrep '^dd'`; while [[ -d /proc/$dd_pid ]]; do kill -USR1 $dd_pid && sleep 1 && clear; done

# Watch the progress of 'dd' with `pv` and `dialog` (apt-get install pv dialog)
(pv -n /dev/zero | dd of=/dev/null bs=128M conv=notrunc,noerror) 2>&1 | dialog --gauge "Running dd command (cloning), please wait..." 10 70 0

# Watch the progress of 'dd' with `pv` and `zenity` (apt-get install pv zenity)
(pv -n /dev/zero | dd of=/dev/null bs=128M conv=notrunc,noerror) 2>&1 | zenity --title 'Running dd command (cloning), please wait...' --progress

# Watch the progress of 'dd' with the built-in `progress` functionality (introduced in coreutils v8.24)
dd if=/dev/zero of=/dev/null bs=128M status=progress

# DD with "graphical" return
dcfldd if=/dev/zero of=/dev/null bs=500K

# This will output the sound from your microphone port to the ssh target computer's speaker port. The sound quality is very bad, so you will hear a lot of hissing.
dd if=/dev/dsp | ssh -c arcfour -C username@host dd of=/dev/dsp
