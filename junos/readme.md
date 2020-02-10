#### Enable netconf
    set system login user netconf-user class super-user
    set system login user netconf-user authentication encrypted-password "$1$rJFgTFU4$VQoCpIWVxMsLgb0hrYr1C/"
    set system services netconf ssh