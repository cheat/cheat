#Exclusive Activation of a Volume Group in a Cluster 
#Link --> https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/7/html/High_Availability_Add-On_Administration/s1-exclusiveactive-HAAA.html 
1> vgs --noheadings -o vg_name
2> volume_list = [ "rhel_root", "rhel_home" ]
3> dracut -H -f /boot/initramfs-$(uname -r).img $(uname -r)
4> Reboot the node
5> uname -r to verify the correct initrd image
