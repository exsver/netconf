package rawxml

var XMLMessagesHPE = map[string]string{
	"GetDevice": `<rpc message-id="101" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                   <get>
                     <filter type="subtree">
                       <top xmlns="http://www.hp.com/netconf/data:1.0">
                         <Device/>
                       </top>
                     </filter>
                   </get>
                 </rpc>`,
	"GetDeviceBaseHostname": `<rpc message-id="102" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
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
	"SaveForce": `<rpc message-id="103" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0"><save/></rpc>`,
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
}
