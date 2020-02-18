## JunOS
### Enable netconf
    set system login user netconf-user class super-user
    set system login user netconf-user authentication encrypted-password "$1$rJFgTFU4$VQoCpIWVxMsLgb0hrYr1C/"
    set system services netconf ssh
    
### Enable netconf logging
    set system services netconf traceoptions file netconf.log
    set system services netconf traceoptions file size 5m
    set system services netconf traceoptions file files 20
    set system services netconf traceoptions file world-readable
    set system services netconf traceoptions flag all
