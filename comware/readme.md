### Enable netconf
    #
     ssh server enable
    #
     netconf ssh server enable
     netconf log source all
    #
     local-user netconf-user class manage
      password simple ********
      service-type ssh
      authorization-attribute user-role network-admin
      authorization-attribute user-role network-operator
    #
      line vty 0 63
      authentication-mode scheme
      user-role network-admin
      protocol inbound ssh