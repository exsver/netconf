package rawxml

var XMLMessagesJunOS = map[string]string{
	"Get":                      `<rpc><get/></rpc>`,
	"CLI":                      `<rpc><command format=\"text\">show configuration vlans</command></rpc>`,
	"GetChassisInventory":      `<rpc><get-chassis-inventory/></rpc>`,
	"GetSoftwareInformation":   `<rpc><get-software-information/></rpc>`,
	"GetSystemInformation":     `<rpc><get-system-information/></rpc>`,
	"GetChassisMacAddresses":   `<rpc><get-chassis-mac-addresses/></rpc>`,
	"GetInterfacesInformation": `<rpc><get-interface-information><terse/></get-interface-information></rpc>`,
	"GetVRRPInformation":       `<rpc><get-vrrp-information/></rpc>`,
	"GetARP":                   `<rpc><get-arp-table-information/></rpc>`,
	"ClearARP":                 `<rpc><clear-arp-table><hostname>10.0.0.100</hostname></clear-arp-table></rpc>`, // run clear arp hostname 10.0.0.100
	"Commit":                   `<rpc><commit/></rpc>`,
	"Reboot":                   `<rpc><request-reboot/></rpc>`,
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
