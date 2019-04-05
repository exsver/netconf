package rawxml

var XMLMessagesHPE = map[string]string{
	"GetSubtree": `<rpc message-id="101" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                              <get>
                                <filter type="subtree">
                                  <top xmlns="http://www.hp.com/netconf/data:1.0">
                                    <ACL/>
                                  </top>
                                </filter>
                              </get>
                            </rpc>`,
	/* Subtree list:
		<ACL/>                           5130 + |
		<ARP/>                           5130 + |
		<ATK/>							 5130 - |
	    <CAR/>
	    <Configuration/>                 5130 + |
		<Device/>
	    <Device><ExtPhysicalEntities><Entity><PhysicalIndex>228</PhysicalIndex><AdminState/><OperState/><StandbyState/><CpuUsage/><CpuUsageThreshold/><MemUsage/><MemAvgUsage/><MemSize/><PhyMemSize/></Entity></ExtPhysicalEntities></Device>
	    <Device><Boards><Board/></Boards><PhysicalEntities><Entity><PhysicalIndex/><Chassis/><Slot/><SubSlot/><Description/><VendorType/><ContainedIn/><Class/><ParentRelPos/><Name/></Entity></PhysicalEntities><ExtPhysicalEntities><Entity><PhysicalIndex/><AdminState/><OperState/><StandbyState/><CpuUsage/><CpuUsageThreshold/><MemUsage/><MemAvgUsage/><MemSize/><PhyMemSize/></Entity></ExtPhysicalEntities></Device><Ifmgr><Interfaces><Interface><IfIndex/><Name/><AbbreviatedName/><PortIndex/><PhysicalIndex/><ifTypeExt/></Interface></Interfaces><EthInterfaces><Interface><IfIndex/><Combo/></Interface></EthInterfaces><EthInterfaceCapabilities><Interface><IfIndex/><Combo/></Interface></EthInterfaceCapabilities></Ifmgr>
	    <DHCP/>
		<DHCPSP/>
		<DHCPSP6/>
		<Domain/>                        5130 + |
	    <FileSystem/>                    5130 + |
	    <FTP/>
	    <Fundamentals/>
	    <HardwareQueue/>
		<Ifmgr/>                         5130 + |
		<IPCIM/>
		<IPV4ADDRESS/>                   5130 + |
	    <IRF/>							 5130 + | 5940 +
		<L2VPN/>								| 5940 +
		<LAGG/>                          5130 + |
		<Login/>
	    <License/>                       5130 - |
		<MAC/>							 5130 + | 5940 +
	    <MGROUP/>                        5130 + |
	    <MLDSnooping/>
	    <MQC/>
		<NTP/>                           5130 + |
	    <ND/>
	    <Package/>
	    <PasswordControl/>
	    <PBR/>                           5130 + |
	    <PIM/>
	    <Radius/>                        5130 + |
		<RBAC/>                          5130 + |
	    <ResourceMonitor/>               5130 - | 5940 +
	    <RMON/>
	    <Route/>
	    <SFLOW/>
	    <Super/>                         5130 - |
		<SNMP/>                          5130 + |
		<SSH/>                           5130 + |
		<StaticRoute/>                   5130 + |
		<STP/>                           5130 + |
		<Syslog/>
	    <URPF4/>
		<UserAccounts/>                  5130 + |
		<VCF/>                           5130 + |
		<VLAN/>                          5130 + |
		<VXLAN/>								| 5940 +
	*/
	"GetDevice": `<rpc message-id="102" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                   <get>
                     <filter type="subtree">
                       <top xmlns="http://www.hp.com/netconf/data:1.0">
                         <Device/>
                       </top>
                     </filter>
                   </get>
                 </rpc>`,
	"GetDeviceBaseHostname": `<rpc message-id="103" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                                <get-config>
                                  <source>
                                    <running/>
                                  </source>
                                <filter type="subtree">
                                  <top xmlns="http://www.hp.com/netconf/config:1.0">
                                    <Device>
                                      <Base>
                                        <HostName/>
                                      </Base>
                                    </Device>
                                  </top>
                                </filter>
                              </get-config>
                             </rpc> `,
	"GetTransceivers": `<rpc message-id="104" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                          <get>
                            <filter type="subtree">
                              <top xmlns="http://www.hp.com/netconf/data:1.0">
                                <Device>
                                  <Transceivers/>
                                </Device>
                              </top>
                            </filter>
                          </get>
                        </rpc> `,
	"GetSessions": `<rpc message-id="105" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0"><get-sessions/></rpc>`,
	"GetRunningConfig": `<rpc xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="103">
                           <get-config>
                             <source>
                               <running/>
                             </source>
                           </get-config>
                         </rpc>`,
	//https://support.hpe.com/hpsc/doc/public/display?docId=emr_na-c04966589
	"LockRunningConfig": `<rpc message-id="301" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                            <lock>
                              <target>
                                <running/>
                              </target>
                            </lock>
                          </rpc>`,
	"UnLockRunningConfig": `<rpc message-id="302" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                              <unlock>
                                <target>
                                  <running/>
                                </target>
                              </unlock>
                            </rpc>`,
	"SaveConfig": `<rpc message-id="303" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                     <save>
                       <file>my_config.cfg</file>
                     </save>
                   </rpc>`,
	"SaveForce": `<rpc message-id="304" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0"><save/></rpc>`,
	"GetIfConfig": `<rpc message-id="104" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                      <get-config>
                        <source>
                          <running/>
                        </source>
                        <filter type="subtree">
                          <top xmlns="http://www.hp.com/netconf/config:1.0">
                            <Ifmgr>
                              <Interfaces/>
                            </Ifmgr>
                          </top>
                        </filter>
                      </get-config>
                    </rpc> `,
	"GetIf": `<rpc message-id="111" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                <get>
                  <filter type="subtree">
                    <top xmlns="http://www.hp.com/netconf/data:1.0">
                      <Ifmgr>
                        <Interfaces/>
                      </Ifmgr>
                    </top>
                  </filter>
                </get>
              </rpc> `,
	"GetEthIfConfig": `<rpc message-id="105" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                         <get-config>
                           <source>
                             <running/>
                           </source>
                           <filter type="subtree">
                             <top xmlns="http://www.hp.com/netconf/config:1.0">
                               <Ifmgr>
                                 <EthInterfaces/>
                               </Ifmgr>
                             </top>
                           </filter>
                         </get-config>
                       </rpc> `,
	"GetIfConfigFilter": `<rpc message-id="105" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                            <get>
                              <filter type="subtree">
                                <top xmlns="http://www.hp.com/netconf/data:1.0">
                                  <Ifmgr>
                                  </Ifmgr>
                                </top>
                              </filter>
                            </get>
                          </rpc> `,
	"GetEthIf": `<rpc message-id="111" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                   <get>
                     <filter type="subtree">
                       <top xmlns="http://www.hp.com/netconf/data:1.0">
                         <Ifmgr>
                           <EthInterfaces/>
                         </Ifmgr>
                       </top>
                     </filter>
                   </get>
                 </rpc>`,
	"GetPorts": `<rpc message-id="400" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                  <get>
                    <filter type="subtree">
                      <top xmlns="http://www.hp.com/netconf/data:1.0">
                        <Ifmgr>
                          <Ports/>
                        </Ifmgr>
                      </top>
                    </filter>
                  </get>
                </rpc>`,
	"GetInterfaceStatistics": `<rpc message-id="401" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                                 <get>
                                   <filter type="subtree">
                                     <top xmlns="http://www.hp.com/netconf/data:1.0">
                                       <Ifmgr>
                                         <Statistics/>
                                       </Ifmgr>
                                     </top>
                                   </filter>
                                 </get>
                               </rpc>`,
	"GetTrafficStatistics": `<rpc message-id="401" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                               <get>
                                 <filter type="subtree">
                                   <top xmlns="http://www.hp.com/netconf/data:1.0">
                                     <Ifmgr>
                                       <TrafficStatistics/>
                                     </Ifmgr>
                                   </top>
                                 </filter>
                               </get>
                             </rpc>`,
	"GetBindings": `<rpc message-id="106" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                      <get-config>
                        <source>
                          <running/>
                        </source>
                        <filter type="subtree">
                          <top xmlns="http://www.hp.com/netconf/config:1.0">
                            <IPCIM>
                              <IpSourceBindingInterface/>
                            </IPCIM>
                          </top>
                        </filter>
                      </get-config>
                    </rpc> `,
	"GetARPDetection": `<rpc message-id="107" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                          <get-config>
                            <source>
                              <running/>
                            </source>
                            <filter type="subtree">
                              <top xmlns="http://www.hp.com/netconf/config:1.0">
                                <ARP/>
                              </top>
                            </filter>
                          </get-config>
                        </rpc> `,
	"GetARPTable": `<rpc message-id="108" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                      <get>
                        <filter type="subtree">
                          <top xmlns="http://www.hp.com/netconf/config:1.0">
                            <ARP/>
                          </top>
                        </filter>
                      </get>
                    </rpc> `,
	"GetMACTable": `<rpc message-id="111" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                      <get>
                        <filter type="subtree">
                          <top xmlns="http://www.hp.com/netconf/data:1.0">
                            <MAC>
                              <MacUnicastTable/>
                            </MAC>
                          </top>
                        </filter>
                      </get>
                    </rpc>`,
	"CLI": `<rpc message-id="109" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
              <CLI>
                <Execution>
                  display vlan
                </Execution>
              </CLI>
            </rpc> `,
	"SetLogBuffer": `<rpc message-id="200" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                       <edit-config>
                         <target>
                           <running/>
                         </target>
                         <config>
                           <top xmlns="http://www.hp.com/netconf/config:1.0" xmlns:name="urn:ietf:params:xml:ns:netconf:base:1.0" name:operation="merge">
                             <Syslog>
                               <LogBuffer>
                                 <BufferSize>512</BufferSize>
                               </LogBuffer>
                             </Syslog>
                           </top>
                         </config>
	                   </edit-config>
	</rpc>`,
	"SetBinding": `<rpc message-id="201" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                     <edit-config>
                       <target>
                         <running/>
                       </target>
                       <config>
                         <top xmlns="http://www.hp.com/netconf/config:1.0" xmlns:name="urn:ietf:params:xml:ns:netconf:base:1.0" name:operation="merge">
                           <IPCIM>
                             <IpSourceBindingInterface>
                               <SourceBinding>
                                 <IfIndex>8</IfIndex>
                                 <Ipv4Address>10.10.0.10</Ipv4Address>
                                 <MacAddress>52-54-00-9A-C9-20</MacAddress>
                                 <VLANID>220</VLANID>
                               </SourceBinding>
                             </IpSourceBindingInterface>
                           </IPCIM>
                         </top>
                       </config>
	                 </edit-config>
                   </rpc>`,
	"Get-bulk": `<rpc message-id="101" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:web="urn:ietf:params:xml:ns:netconf:base:1.0">
                  <get-bulk>
                    <filter type="subtree">
                      <top xmlns="http://www.hp.com/netconf/data:1.0" xmlns:web="http://www.hp.com/netconf/base:1.0">
                        <Ifmgr>
                          <Interfaces web:count="5"></Interfaces>
                        </Ifmgr>
                      </top>
                    </filter>
                  </get-bulk>
                 </rpc>`,
	"Get-mon": `<rpc message-id='101' xmlns='urn:ietf:params:xml:ns:netconf:base:1.0' xmlns:web='urn:ietf:params:xml:ns:netconf:base:1.0'>
                  <get>
                    <filter type='subtree'>
                      <top xmlns='http://www.hp.com/netconf/data:1.0' xmlns:web='http://www.hp.com/netconf/base:1.0' xmlns:data='http://www.hp.com/netconf/data:1.0' web:count='1'>
                        <Device>
                          <ExtPhysicalEntities>
                            <Entity>
                              <PhysicalIndex>192</PhysicalIndex>
                              <AdminState/>
                              <OperState/>
                              <StandbyState/>
                              <CpuUsage/>
                              <CpuUsageThreshold/>
                              <MemUsage/>
                              <MemAvgUsage/>
                              <MemSize/>
                              <PhyMemSize/>
                            </Entity>
                          </ExtPhysicalEntities>
                        </Device>
                      </top>
                    </filter>
                  </get>
                </rpc>`,
	"regexp": `<rpc xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1">
      <get>
        <filter type="subtree">
            <top xmlns="http://www.hp.com/netconf/data:1.0">
                <Ifmgr>
                    <Ports>
                        <Port>
							<IfIndex xmlns:re="http://www.hp.com/netconf/base:1.0" re:regExp="2"/>
                        </Port>
                    </Ports>
                </Ifmgr>
            </top>
        </filter>
      </get>
    </rpc>`,
	"yang": `<rpc message-id="101" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
			<get-schema xmlns="urn:ietf:params:xml:ns:yang:ietf-netconf-monitoring">
				<identifier>ietf-netconf</identifier>
				<version>2014-10-12</version>
				<format>yang</format>
				</get-schema>
			</rpc>`,
	"ConditionalMatch": `<rpc xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1">
<get>
	<filter type="subtree">
		<top xmlns="http://www.hp.com/netconf/data:1.0">
			<MAC>
				<MacUnicastTable>
					<Unicast>
						<VLANID>99</VLANID>
						<MacAddress/>
						<PortIndex xmlns:re="http://www.hp.com/netconf/base:1.0" re:match="notEqual:719"/>
						<Status/>
						<Aging/>
					</Unicast>
				</MacUnicastTable>
			</MAC>
		</top>
	</filter>
</get>
</rpc>`,
}
