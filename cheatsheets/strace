# Basic stracing
strace <command>

# save the trace to a file
strace -o strace.out <other switches> <command>

# follow only the open() system call
strace -e trace=open <command>

# follow all the system calls which open a file
strace -e trace=file <command>

# follow all the system calls associated with process
# management
strace -e trace=process <command>

# follow child processes as they are created
strace -f <command>

# count time, calls and errors for each system call
strace -c <command>

# trace a running process (multiple PIDs can be specified)
strace -p <pid>