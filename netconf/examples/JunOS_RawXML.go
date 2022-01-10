package main

import (
	"fmt"
	"time"

	"github.com/exsver/netconf/netconf"
	"golang.org/x/crypto/ssh"
)

func main() {
	netconf.LogLevel.Verbose()

	targetDevice := &netconf.TargetDevice{
		IP:   "10.10.10.10",
		Port: 830,
		SSHConfig: ssh.ClientConfig{
			Config: ssh.Config{
				Ciphers: []string{"aes128-ctr", "hmac-sha1"}, // aes128-cbc for HP5940  aes128-ctr for juniper QFX5100 or juniper MX
			},
			User:            "netconf-user",
			Auth:            []ssh.AuthMethod{ssh.Password("netconf-password")},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         120 * time.Second,
		},
	}

	err := targetDevice.Connect(300 * time.Second)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	defer targetDevice.Disconnect()

	_, _ = targetDevice.NetconfSession.SendAndReceive([]byte(xmlMessagesJunOS["GetChassisInventory"]))

}

// xmlMessagesJunOS provides a few raw xml examples (for JunOS)
var xmlMessagesJunOS = map[string]string{
	"Get":                          `<rpc><get/></rpc>`,
	"CLI":                          `<rpc><command format=\"text\">show configuration vlans</command></rpc>`,
	"GetChassisInventory":          `<rpc><get-chassis-inventory/></rpc>`,
	"GetSoftwareInformation":       `<rpc><get-software-information/></rpc>`,
	"GetSystemInformation":         `<rpc><get-system-information/></rpc>`,
	"GetSystemStatistics":          `<rpc><get-statistics-information/></rpc>`,   // run show system statistics
	"GetSystemUsers":               `<rpc><get-system-users-information/></rpc>`, // run show system users
	"GetAlarmInformation":          `<rpc><get-alarm-information/></rpc>`,        // run show chassis alarm
	"GetChassisMacAddresses":       `<rpc><get-chassis-mac-addresses/></rpc>`,
	"GetInterfacesInformation":     `<rpc><get-interface-information><terse/></get-interface-information></rpc>`,
	"GetVRRPInformation":           `<rpc><get-vrrp-information/></rpc>`,
	"GetARPTable":                  `<rpc><get-arp-table-information/></rpc>`,
	"ClearARP":                     `<rpc><clear-arp-table><hostname>10.0.0.100</hostname></clear-arp-table></rpc>`, // run clear arp hostname 10.0.0.100
	"Commit":                       `<rpc><commit/></rpc>`,
	"ShowCompare":                  `<rpc><get-configuration compare="rollback" rollback="0" format="text"/></rpc>`,
	"Reboot":                       `<rpc><request-reboot/></rpc>`,
	"OpenConfigurationPrivate":     `<rpc><open-configuration><private/></open-configuration></rpc>`,   // configure private
	"OpenConfigurationExclusive":   `<rpc><open-configuration><exclusive/></open-configuration></rpc>`, // configure exclusive
	"GetDatabaseStatusInformation": `<rpc><get-database-status-information/></rpc>`,
	"GetRunningConfig": `<rpc>
                            <get-config>
                              <source>
                                <running/>
                              </source>
                            </get-config>
                           </rpc>`,
	"GetRunningConfigFilter": `<rpc>
                                <get-config>
                                  <source>
                                    <running/>
                                  </source>
                                  <filter type="subtree">
                                    <configuration>
                                      <system/>
                                    </configuration>
                                  </filter>
                                </get-config>
                               </rpc>`,
	"LockConfig": `<rpc>
                     <lock>
                       <target>
                         <candidate/>
                       </target>
                     </lock>
                   </rpc>`,
	"UnLockConfig": `<rpc>
                      <unlock>
                        <target>
                          <candidate/>
                        </target>
                      </unlock>
                     </rpc>`,
	"SetHostname": `<rpc>
                      <edit-config>
                        <target>
                          <candidate/>
                        </target>
                        <config>
                          <configuration>
                            <system>
                              <host-name>test</host-name>
                            </system>
	                      </configuration>
                        </config>
                      </edit-config>
                     </rpc>`,
	"CommitCheck": `<rpc>
                      <validate>
                        <source>
                          <candidate/>
                        </source>
                      </validate>
                     </rpc>`,
	"CommitConfirmed": `<rpc>
                           <commit>
                             <confirmed/>
                             <confirm-timeout>" 60 "</confirm-timeout>
                           </commit>
                         </rpc>`, //time in seconds
	"RollbackCompare": `<rpc>
                          <get-rollback-information>
                            <rollback>0</rollback>
                            <compare>1</compare>
                          </get-rollback-information>
                        </rpc>`,
	"LoadConfigurationRolback": `<rpc><load-configuration rollback="0"/></rpc>`,
	"GetConfig": `<rpc>
                   <get-config>
                     <source>
                       <candidate/>
                     </source>
                   </get-config>
                 </rpc>`,
	"GetConfigSubtree": `<rpc>
                          <get-config>
                            <source>
                              <candidate/>
                            </source>
                            <filter type="subtree">
                              <configuration>
                                <system>
                                  <login/>
                                </system>
                              </configuration>
                            </filter>
                          </get-config>
                         </rpc>`,
	"GetConfiguration": `<rpc>
                           <get-configuration format="set" inherit="defaults"></get-configuration>
                         </rpc>`, //format="( json | set | text | xml )"
	"LoadConfigurationText": `<rpc>
                           <load-configuration action="merge" format="text">
                             <configuration-text>
vlans {
    v101 {
        description test-vlan101;
        vlan-id 101;
    }
}
                             </configuration-text>
                           </load-configuration></rpc>`,
	"LoadConfigurationXML":    `<rpc></rpc>`,
	"GetInterfaceInformation": `<rpc><get-interface-information/></rpc>`,
	"GetInterfaceDescription": `<rpc>
                                  <get-interface-information>
	                                <descriptions/>
	                              </get-interface-information>
	                            </rpc>`,
	"GetInterfaceInformationText": `<rpc><get-interface-information format="text"></get-interface-information></rpc>`,
	"GetInterfaceInformationFilter": `<rpc>
                                        <get-interface-information format="xml">
                                           <interface-name>em0</interface-name>
                                        </get-interface-information>
                                      </rpc>`, //format="(ascii | json| text | xml)"
	"GetInterfaceInformationDetail": `<rpc>
                                        <get-interface-information>
                                          <detail/>
                                        </get-interface-information>
                                      </rpc>`,
	"GetInterfaceInformationTerse": `<rpc>
                                       <get-interface-information>
                                         <terse/>
                                       </get-interface-information>
                                     </rpc>`,
	"GetInterfaceBrief": `<rpc>
                             <get-interface-information>
                               <brief/>
                               <interface-name>et-0/0/0</interface-name>
                             </get-interface-information>
                           </rpc>`,
	"GetInterfaceOptics": `<rpc>
                             <get-interface-optics-diagnostics-information>
                               <interface-name>et-0/0/0</interface-name>
                             </get-interface-optics-diagnostics-information>
                           </rpc>`,
	"GetFirewallCounter": `<rpc>
		<get-firewall-counter-information>
			<countername>udp-drop-counter</countername>
			<filter>testFilter</filter>
		</get-firewall-counter-information></rpc>`, // run show firewall counter udp-drop-counter filter testFilter
	"GetFirewallFilterInformation": `<rpc>
		<get-firewall-filter-information>
			<filtername>testFilter</filtername>
		</get-firewall-filter-information></rpc>`, // run show firewall filter testFilter
	"Ping": `<rpc>
               <ping>
                 <host>8.8.8.8</host>
                 <count>1</count>
               </ping>
              </rpc>`,
	"GetXNM": `<rpc>
                 <get-xnm-information>
                   <type>xml-schema</type>
                   <namespace>junos-configuration</namespace>
                 </get-xnm-information>
               </rpc>`,
	"AddVlan": `<rpc xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="2">
<edit-config>
	<target><candidate/></target>
	<default-operation>merge</default-operation>
	<config>
		<configuration>
			<vlans>
				<vlan>
					<name>add-test-vlan</name>
					<vlan-id>36</vlan-id>
				</vlan>
			</vlans>
		</configuration>
	</config>
</edit-config>
</rpc>`,
}
