# WARNING:
#   In order to avoid headaches when moving ZFS physical devices around,
#   one will be much better served to reference devices by their *immutable*
#   ID - as in /dev/disk/by-id/* - rather than their block device name -
#   as in /dev/{sd,nvme}* - which is bound to change as per PCI enumeration
#   order.
# For the sake of briefness, we'll use the following variables:
#   ${device}    device (/dev/disk/by-id/${device})
#   ${part}      partition (/dev/disk/by-id/${part=${device}-part${N}})
#   ${pool}      ZFS pool (name)
#   ${fs_vol}    ZFS file system or volume (name)
#   ${snapshot}  ZFS snapshot (name)


## Pools

# Create a new "RAID-5" (raidz1) pool
# Recommended: use entire devices rather than partitions
zpool create ${pool} raidz1 ${device} ${device} ${device} [...]

# Add 2nd-level "RAID-1" (mirror) ZFS Intent Log (ZIL; synchronous write cache)
# Recommended: use separate, fast, low-latency devices (e.g. NVMe)
zpool add ${pool} log mirror ${part} ${part}

# Add 2nd-level "RAID-0" Adaptive Replacement Cache (ARC; read cache)
# Recommended: use separate, fast, low-latency devices (e.g. NVMe)
zpool add ${pool} cache ${part} ${part} [...]

# Remove log or cache components
zpool remove zfs ${part} [...]

# Import (enable) existing pool from newly connected devices
# Note: this will create the /etc/zfs/zpool.cache devices cache
zpool import -d /dev/disk/by-id -aN

# Import (enable) existing pool using the devices cache
zpool import -c /etc/zfs/zpool.cache -aN

# Export (disable) pool (e.g. before shutdown)
zpool export -a

# List all (imported) pools
zpool list

# See pool status
zpool status ${pool}

# See detailed pool I/O statistics
zpool iostat ${pool} -v

# Verify pool integrity (data checksums)
# (watch progress with 'zpool status')
zpool scrub ${pool}

# Remove a failing device from a pool
# Note: redundant pools (mirror, raidz) will continue working in degraded state
zpool detach ${pool} ${device}

# Replace a failed device in a pool
# Note: new device will be "resilvered" automatically (parity reconstruction)
#       (watch progress with 'zpool status')
zpool replace ${pool} ${failed-device} ${new-device}

# Erase zpool labels ("superblock") from a device/partition
# WARNING: MUST do before reusing a device/partition for other purposes
zpool labelclear ${device}

# Query pool configuration (properties)
zpool get all ${pool}

# Change pool configuration (property)
zpool set <property>=<value> ${pool}

# Dump the entire pool (commands) history
zpool history ${pool}

# More...
man zpool


## File systems / Volumes

# Create a new file system
zfs create ${pool}/${fs_vol}

# Create a new volume ("block device")
# Note: look for it in /dev/zvol/${pool}/${fs_vol}
zfs create -V <size> ${pool}/${fs_vol}

# List all file systems / volumes
zfs list

# Mount all file systems
# Note: see 'zfs get mountpoint ${pool}' for mountpoint root path
zfs mount -a

# Create a snapshot
zfs snapshot ${pool}/${fs_vol}@${snapshot}

# Delete a snapshot
zfs destroy ${pool}/${fs_vol}@${snapshot}

# Full backup
# Note: pipe (|) source to destination through netcat, SSH, etc.
# ... on source:
zfs send -p -R ${pool}/${fs_vol}@${snapshot}
# ... on destination:
zfs receive -F ${pool}/${fs_vol}

# Incremental backup
# Note: pipe (|) source to destination through netcat, SSH, etc.
# ... on source:
zfs send -p -R -i ${pool}/${fs_vol}@${snapshot-previous} ${pool}/${fs_vol}@${snapshot}
# ... on destination:
zfs receive -F ${pool}/${fs_vol}

# Query file system / volume configuration (properties)
zfs get all ${pool}
zfs get all ${pool}/${fs_vol}

# Change file system / volume configuration (property)
zfs set <property>=<value> ${pool}/${fs_vol}

# More...
man zfs

