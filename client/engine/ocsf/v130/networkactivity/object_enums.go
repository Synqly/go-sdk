package networkactivity

// AccountTypeId Values
// Type ID. The normalized account type identifier.

// Unknown. The account type is unknown.
const Account_Type_Unknown AccountTypeId = 0
// LDAPAccount
const Account_Type_LDAPAccount AccountTypeId = 1
// WindowsAccount
const Account_Type_WindowsAccount AccountTypeId = 2
// AWSIAMUser
const Account_Type_AWSIAMUser AccountTypeId = 3
// AWSIAMRole
const Account_Type_AWSIAMRole AccountTypeId = 4
// GCPAccount
const Account_Type_GCPAccount AccountTypeId = 5
// AzureADAccount
const Account_Type_AzureADAccount AccountTypeId = 6
// MacOSAccount
const Account_Type_MacOSAccount AccountTypeId = 7
// AppleAccount
const Account_Type_AppleAccount AccountTypeId = 8
// LinuxAccount
const Account_Type_LinuxAccount AccountTypeId = 9
// AWSAccount
const Account_Type_AWSAccount AccountTypeId = 10
// Other. The account type is not mapped.
const Account_Type_Other AccountTypeId = 99

// AffectedPackageTypeId Values
// Type ID. The type of software package.

// Unknown. The type is unknown.
const AffectedPackage_Type_Unknown AffectedPackageTypeId = 0
// Application. An application software package.
const AffectedPackage_Type_Application AffectedPackageTypeId = 1
// OperatingSystem. An operating system software package.
const AffectedPackage_Type_OperatingSystem AffectedPackageTypeId = 2
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const AffectedPackage_Type_Other AffectedPackageTypeId = 99

// AgentTypeId Values
// Type ID. The normalized representation of an agent or sensor. E.g., EDR, vulnerability management, APM, backup & recovery, etc.

// Unknown. The type is unknown.
const Agent_Type_Unknown AgentTypeId = 0
// EndpointDetectionandResponse. Any EDR sensor or agent. Or any tool that provides similar threat detection, anti-malware, anti-ransomware, or similar capabilities. E.g., Crowdstrike Falcon, Microsoft Defender for Endpoint, Wazuh.
const Agent_Type_EndpointDetectionandResponse AgentTypeId = 1
// DataLossPrevention. Any DLP sensor or agent. Or any tool that provides similar data classification, data loss detection, and/or data loss prevention capabilities. E.g., Forcepoint DLP, Microsoft Purview, Symantec DLP.
const Agent_Type_DataLossPrevention AgentTypeId = 2
// Backup_Recovery. Any agent or sensor that provides backups, archival, or recovery capabilities. E.g., Azure Backup, AWS Backint Agent.
const Agent_Type_Backup_Recovery AgentTypeId = 3
// PerformanceMonitoring_Observability. Any agent or sensor that provides Application Performance Monitoring (APM), active tracing, profiling, or other observability use cases and optionally forwards the logs. E.g., New Relic Agent, Datadog Agent, Azure Monitor Agent.
const Agent_Type_PerformanceMonitoring_Observability AgentTypeId = 4
// VulnerabilityManagement. Any agent or sensor that provides vulnerability management or scanning capabilities. E.g., Qualys VMDR, Microsoft Defender for Endpoint, Crowdstrike Spotlight, Amazon Inspector Agent.
const Agent_Type_VulnerabilityManagement AgentTypeId = 5
// LogForwarding. Any agent or sensor that forwards logs to a 3rd party storage system such as a data lake or SIEM. E.g., Splunk Universal Forwarder, Tenzir, FluentBit, Amazon CloudWatch Agent, Amazon Kinesis Agent.
const Agent_Type_LogForwarding AgentTypeId = 6
// MobileDeviceManagement. Any agent or sensor responsible for providing Mobile Device Management (MDM) or Mobile Enterprise Management (MEM) capabilities. E.g., JumpCloud Agent, Esper Agent, Jamf Pro binary.
const Agent_Type_MobileDeviceManagement AgentTypeId = 7
// ConfigurationManagement. Any agent or sensor that provides configuration management of a device, such as scanning for software, license management, or applying configurations. E.g., AWS Systems Manager Agent, Flexera, ServiceNow MID Server.
const Agent_Type_ConfigurationManagement AgentTypeId = 8
// RemoteAccess. Any agent or sensor that provides remote access capabilities to a device. E.g., BeyondTrust, Amazon Systems Manager Agent, Verkada Agent.
const Agent_Type_RemoteAccess AgentTypeId = 9
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const Agent_Type_Other AgentTypeId = 99

// CvssDepth Values
// CVSS Depth. The CVSS depth represents a depth of the equation used to calculate CVSS score.

// Base
const Cvss_Depth_Base CvssDepth = "Base"
// Environmental
const Cvss_Depth_Environmental CvssDepth = "Environmental"
// Temporal
const Cvss_Depth_Temporal CvssDepth = "Temporal"

// DeviceRiskLevelId Values
// Risk Level ID. The normalized risk level id.

// Info
const Device_RiskLevel_Info DeviceRiskLevelId = 0
// Low
const Device_RiskLevel_Low DeviceRiskLevelId = 1
// Medium
const Device_RiskLevel_Medium DeviceRiskLevelId = 2
// High
const Device_RiskLevel_High DeviceRiskLevelId = 3
// Critical
const Device_RiskLevel_Critical DeviceRiskLevelId = 4
// Other. The risk level is not mapped. See the <code>risk_level</code> attribute, which contains a data source specific value.
const Device_RiskLevel_Other DeviceRiskLevelId = 99

// DeviceTypeId Values
// Type ID. The device type ID.

// Unknown. The type is unknown.
const Device_Type_Unknown DeviceTypeId = 0
// Server. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Server/'>server</a>.
const Device_Type_Server DeviceTypeId = 1
// Desktop. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:DesktopComputer/'>desktop computer</a>.
const Device_Type_Desktop DeviceTypeId = 2
// Laptop. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:LaptopComputer/'>laptop computer</a>.
const Device_Type_Laptop DeviceTypeId = 3
// Tablet. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:TabletComputer/'>tablet computer</a>.
const Device_Type_Tablet DeviceTypeId = 4
// Mobile. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:MobilePhone/'>mobile phone</a>.
const Device_Type_Mobile DeviceTypeId = 5
// Virtual. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:VirtualizationSoftware/'>virtual machine</a>.
const Device_Type_Virtual DeviceTypeId = 6
// IOT. An <a target='_blank' href='https://www.techtarget.com/iotagenda/definition/IoT-device'>IOT (Internet of Things) device</a>.
const Device_Type_IOT DeviceTypeId = 7
// Browser. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Browser/'>web browser</a>.
const Device_Type_Browser DeviceTypeId = 8
// Firewall. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Firewall/'>networking firewall</a>.
const Device_Type_Firewall DeviceTypeId = 9
// Switch. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Switch/'>networking switch</a>.
const Device_Type_Switch DeviceTypeId = 10
// Hub. A <a target='_blank' href='https://en.wikipedia.org/wiki/Ethernet_hub'>networking hub</a>.
const Device_Type_Hub DeviceTypeId = 11
// Router. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Router/'>networking router</a>.
const Device_Type_Router DeviceTypeId = 12
// IDS. An <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:IntrusionDetectionSystem/'>intrusion detection system</a>.
const Device_Type_IDS DeviceTypeId = 13
// IPS. An <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:IntrusionPreventionSystem/'>intrusion prevention system</a>.
const Device_Type_IPS DeviceTypeId = 14
// LoadBalancer. A <a target='_blank' href='https://en.wikipedia.org/wiki/Load_balancing_(computing)'> Load Balancer device.
const Device_Type_LoadBalancer DeviceTypeId = 15
// ImagingEquipment. Equipment for processing optical data, such as a camera.
const Device_Type_ImagingEquipment DeviceTypeId = 89
// PLC. A Programmable logic controller.
const Device_Type_PLC DeviceTypeId = 90
// SCADA. A supervisory control and data acquisition system.
const Device_Type_SCADA DeviceTypeId = 91
// DCS. A distributed control system.
const Device_Type_DCS DeviceTypeId = 92
// CNC. A computer numerical control system, including computerized machine tools.
const Device_Type_CNC DeviceTypeId = 93
// ScientificEquipment. A piece of scientific equipment such as an oscilloscope or spectrometer.
const Device_Type_ScientificEquipment DeviceTypeId = 94
// MedicalDevice. A medical device such as an MRI machine or infusion pump.
const Device_Type_MedicalDevice DeviceTypeId = 95
// LightingControls. A lighting control for internal or external applications.
const Device_Type_LightingControls DeviceTypeId = 96
// EnergyMonitoringSystem. An energy monitoring, security or safety system.
const Device_Type_EnergyMonitoringSystem DeviceTypeId = 97
// TransportationDevice. A transportation device or transportation supporting device.
const Device_Type_TransportationDevice DeviceTypeId = 98
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const Device_Type_Other DeviceTypeId = 99

// DigitalSignatureAlgorithmId Values
// Algorithm ID. The identifier of the normalized digital signature algorithm.

// Unknown. The algorithm is unknown.
const DigitalSignature_Algorithm_Unknown DigitalSignatureAlgorithmId = 0
// DSA. Digital Signature Algorithm (DSA).
const DigitalSignature_Algorithm_DSA DigitalSignatureAlgorithmId = 1
// RSA. Rivest-Shamir-Adleman (RSA) Algorithm.
const DigitalSignature_Algorithm_RSA DigitalSignatureAlgorithmId = 2
// ECDSA. Elliptic Curve Digital Signature Algorithm.
const DigitalSignature_Algorithm_ECDSA DigitalSignatureAlgorithmId = 3
// Authenticode. Microsoft Authenticode Digital Signature Algorithm.
const DigitalSignature_Algorithm_Authenticode DigitalSignatureAlgorithmId = 4
// Other. The algorithm is not mapped. See the <code>algorithm</code> attribute, which contains a data source specific value.
const DigitalSignature_Algorithm_Other DigitalSignatureAlgorithmId = 99

// DigitalSignatureStateId Values
// State ID. The normalized identifier of the signature state.

// Unknown. The state is unknown.
const DigitalSignature_State_Unknown DigitalSignatureStateId = 0
// Valid. The digital signature is valid.
const DigitalSignature_State_Valid DigitalSignatureStateId = 1
// Expired. The digital signature is not valid due to expiration of certificate.
const DigitalSignature_State_Expired DigitalSignatureStateId = 2
// Revoked. The digital signature is invalid due to certificate revocation.
const DigitalSignature_State_Revoked DigitalSignatureStateId = 3
// Suspended. The digital signature is invalid due to certificate suspension.
const DigitalSignature_State_Suspended DigitalSignatureStateId = 4
// Pending. The digital signature state is pending.
const DigitalSignature_State_Pending DigitalSignatureStateId = 5
// Other. The state is not mapped. See the <code>state</code> attribute, which contains a data source specific value.
const DigitalSignature_State_Other DigitalSignatureStateId = 99

// DnsAnswerFlagIds Values
// DNS Header Flags. The list of DNS answer header flag IDs.

// Unknown. The flag is unknown.
const DnsAnswer_FlagIds_Unknown DnsAnswerFlagIds = 0
// AuthoritativeAnswer
const DnsAnswer_FlagIds_AuthoritativeAnswer DnsAnswerFlagIds = 1
// TruncatedResponse
const DnsAnswer_FlagIds_TruncatedResponse DnsAnswerFlagIds = 2
// RecursionDesired
const DnsAnswer_FlagIds_RecursionDesired DnsAnswerFlagIds = 3
// RecursionAvailable
const DnsAnswer_FlagIds_RecursionAvailable DnsAnswerFlagIds = 4
// AuthenticData
const DnsAnswer_FlagIds_AuthenticData DnsAnswerFlagIds = 5
// CheckingDisabled
const DnsAnswer_FlagIds_CheckingDisabled DnsAnswerFlagIds = 6
// Other. The flag is not mapped. See the <code>flags</code> attribute, which contains a data source specific value.
const DnsAnswer_FlagIds_Other DnsAnswerFlagIds = 99

// DomainContactTypeId Values
// Domain Contact Type ID. The normalized domain contact type ID.

// Unknown. The type is unknown.
const DomainContact_Type_Unknown DomainContactTypeId = 0
// Registrant. The contact information provided is for the domain registrant.
const DomainContact_Type_Registrant DomainContactTypeId = 1
// Administrative. The contact information provided is for the domain administrator.
const DomainContact_Type_Administrative DomainContactTypeId = 2
// Technical. The contact information provided is for the domain technical lead.
const DomainContact_Type_Technical DomainContactTypeId = 3
// Billing. The contact information provided is for the domain billing lead.
const DomainContact_Type_Billing DomainContactTypeId = 4
// Abuse. The contact information provided is for the domain abuse contact.
const DomainContact_Type_Abuse DomainContactTypeId = 5
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const DomainContact_Type_Other DomainContactTypeId = 99

// FileConfidentialityId Values
// Confidentiality ID. The normalized identifier of the file content confidentiality indicator.

// Unknown. The confidentiality is unknown.
const File_Confidentiality_Unknown FileConfidentialityId = 0
// NotConfidential
const File_Confidentiality_NotConfidential FileConfidentialityId = 1
// Confidential
const File_Confidentiality_Confidential FileConfidentialityId = 2
// Secret
const File_Confidentiality_Secret FileConfidentialityId = 3
// TopSecret
const File_Confidentiality_TopSecret FileConfidentialityId = 4
// Private
const File_Confidentiality_Private FileConfidentialityId = 5
// Restricted
const File_Confidentiality_Restricted FileConfidentialityId = 6
// Other. The confidentiality is not mapped. See the <code>confidentiality</code> attribute, which contains a data source specific value.
const File_Confidentiality_Other FileConfidentialityId = 99

// FileTypeId Values
// Type ID. The file type ID.

// Unknown. The type is unknown.
const File_Type_Unknown FileTypeId = 0
// RegularFile
const File_Type_RegularFile FileTypeId = 1
// Folder
const File_Type_Folder FileTypeId = 2
// CharacterDevice
const File_Type_CharacterDevice FileTypeId = 3
// BlockDevice
const File_Type_BlockDevice FileTypeId = 4
// LocalSocket
const File_Type_LocalSocket FileTypeId = 5
// NamedPipe
const File_Type_NamedPipe FileTypeId = 6
// SymbolicLink
const File_Type_SymbolicLink FileTypeId = 7
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const File_Type_Other FileTypeId = 99

// FingerprintAlgorithmId Values
// Algorithm ID. The identifier of the normalized hash algorithm, which was used to create the digital fingerprint.

// Unknown. The algorithm is unknown.
const Fingerprint_Algorithm_Unknown FingerprintAlgorithmId = 0
// MD5. MD5 message-digest algorithm producing a 128-bit (16-byte) hash value.
const Fingerprint_Algorithm_MD5 FingerprintAlgorithmId = 1
// SHA_1. Secure Hash Algorithm 1 producing a 160-bit (20-byte) hash value.
const Fingerprint_Algorithm_SHA_1 FingerprintAlgorithmId = 2
// SHA_256. Secure Hash Algorithm 2 producing a 256-bit (32-byte) hash value.
const Fingerprint_Algorithm_SHA_256 FingerprintAlgorithmId = 3
// SHA_512. Secure Hash Algorithm 2 producing a 512-bit (64-byte) hash value.
const Fingerprint_Algorithm_SHA_512 FingerprintAlgorithmId = 4
// CTPH. The ssdeep generated fuzzy checksum. Also known as Context Triggered Piecewise Hash (CTPH).
const Fingerprint_Algorithm_CTPH FingerprintAlgorithmId = 5
// TLSH. The TLSH fuzzy hashing algorithm.
const Fingerprint_Algorithm_TLSH FingerprintAlgorithmId = 6
// quickXorHash. Microsoft simple non-cryptographic hash algorithm that works by XORing the bytes in a circular-shifting fashion.
const Fingerprint_Algorithm_quickXorHash FingerprintAlgorithmId = 7
// Other. The algorithm is not mapped. See the <code>algorithm</code> attribute, which contains a data source specific value.
const Fingerprint_Algorithm_Other FingerprintAlgorithmId = 99

// HttpRequestHttpMethod Values
// HTTP Method. The <a target='_blank' href='https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods'>HTTP request method</a> indicates the desired action to be performed for a given resource.

// Connect. The CONNECT method establishes a tunnel to the server identified by the target resource.
const HttpRequest_HttpMethod_Connect HttpRequestHttpMethod = "CONNECT"
// Delete. The DELETE method deletes the specified resource.
const HttpRequest_HttpMethod_Delete HttpRequestHttpMethod = "DELETE"
// Get. The GET method requests a representation of the specified resource. Requests using GET should only retrieve data.
const HttpRequest_HttpMethod_Get HttpRequestHttpMethod = "GET"
// Head. The HEAD method asks for a response identical to a GET request, but without the response body.
const HttpRequest_HttpMethod_Head HttpRequestHttpMethod = "HEAD"
// Options. The OPTIONS method describes the communication options for the target resource.
const HttpRequest_HttpMethod_Options HttpRequestHttpMethod = "OPTIONS"
// Post. The POST method submits an entity to the specified resource, often causing a change in state or side effects on the server.
const HttpRequest_HttpMethod_Post HttpRequestHttpMethod = "POST"
// Put. The PUT method replaces all current representations of the target resource with the request payload.
const HttpRequest_HttpMethod_Put HttpRequestHttpMethod = "PUT"
// Trace. The TRACE method performs a message loop-back test along the path to the target resource.
const HttpRequest_HttpMethod_Trace HttpRequestHttpMethod = "TRACE"

// Ja4FingerprintTypeId Values
// Type ID. The identifier of the JA4+ fingerprint type.

// Unknown. The type is unknown.
const Ja4Fingerprint_Type_Unknown Ja4FingerprintTypeId = 0
// JA4. TLS Client Fingerprint.
const Ja4Fingerprint_Type_JA4 Ja4FingerprintTypeId = 1
// JA4Server. TLS Server Response/Session Fingerprint.
const Ja4Fingerprint_Type_JA4Server Ja4FingerprintTypeId = 2
// JA4HTTP. HTTP Client Fingerprint.
const Ja4Fingerprint_Type_JA4HTTP Ja4FingerprintTypeId = 3
// JA4Latency. Latency Measurement/Light Distance Fingerprint.
const Ja4Fingerprint_Type_JA4Latency Ja4FingerprintTypeId = 4
// JA4X509. X509 TLS Certificate Fingerprint.
const Ja4Fingerprint_Type_JA4X509 Ja4FingerprintTypeId = 5
// JA4SSH. SSH Traffic Fingerprint.
const Ja4Fingerprint_Type_JA4SSH Ja4FingerprintTypeId = 6
// JA4TCP. Passive TCP Client Fingerprint.
const Ja4Fingerprint_Type_JA4TCP Ja4FingerprintTypeId = 7
// JA4TCPServer. Passive TCP Server Fingerprint.
const Ja4Fingerprint_Type_JA4TCPServer Ja4FingerprintTypeId = 8
// JA4TCPScan. Active TCP Server Fingerprint.
const Ja4Fingerprint_Type_JA4TCPScan Ja4FingerprintTypeId = 9
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const Ja4Fingerprint_Type_Other Ja4FingerprintTypeId = 99

// KbArticleInstallStateId Values
// Install State ID. The normalized install state ID of the kb article.

// Unknown. The normalized install state is unknown.
const KbArticle_InstallState_Unknown KbArticleInstallStateId = 0
// Installed. The item is installed.
const KbArticle_InstallState_Installed KbArticleInstallStateId = 1
// NotInstalled. The item is not installed.
const KbArticle_InstallState_NotInstalled KbArticleInstallStateId = 2
// InstalledPendingReboot. The item is installed pending reboot operation.
const KbArticle_InstallState_InstalledPendingReboot KbArticleInstallStateId = 3
// Other. The install state is not mapped. See the <code>install_state</code> attribute, which contains a data source specific value.
const KbArticle_InstallState_Other KbArticleInstallStateId = 99

// KillChainPhasePhaseId Values
// Kill Chain Phase ID. The cyber kill chain phase identifier.

// Unknown. The kill chain phase is unknown.
const KillChainPhase_Phase_Unknown KillChainPhasePhaseId = 0
// Reconnaissance. The attackers pick a target and perform a detailed analysis, start collecting information (email addresses, conferences information, etc.) and evaluate the victim’s vulnerabilities to determine how to exploit them.
const KillChainPhase_Phase_Reconnaissance KillChainPhasePhaseId = 1
// Weaponization. The attackers develop a malware weapon and aim to exploit the discovered vulnerabilities.
const KillChainPhase_Phase_Weaponization KillChainPhasePhaseId = 2
// Delivery. The intruders will use various tactics, such as phishing, infected USB drives, etc.
const KillChainPhase_Phase_Delivery KillChainPhasePhaseId = 3
// Exploitation. The intruders start leveraging vulnerabilities to executed code on the victim’s system.
const KillChainPhase_Phase_Exploitation KillChainPhasePhaseId = 4
// Installation. The intruders install malware on the victim’s system.
const KillChainPhase_Phase_Installation KillChainPhasePhaseId = 5
// Command_Control. Malware opens a command channel to enable the intruders to remotely manipulate the victim's system.
const KillChainPhase_Phase_Command_Control KillChainPhasePhaseId = 6
// ActionsonObjectives. With hands-on keyboard access, intruders accomplish the mission’s goal.
const KillChainPhase_Phase_ActionsonObjectives KillChainPhasePhaseId = 7
// Other. The kill chain phase is not mapped. See the <code>phase</code> attribute, which contains a data source specific value.
const KillChainPhase_Phase_Other KillChainPhasePhaseId = 99

// MalwareClassificationIds Values
// Classification IDs. The list of normalized identifiers of the malware classifications. Reference: <a target='_blank' href='https://docs.oasis-open.org/cti/stix/v2.1/os/stix-v2.1-os.html#_oxlc4df65spl'>STIX Malware Types</a> 

// Unknown. The classification is unknown.
const Malware_ClassificationIds_Unknown MalwareClassificationIds = 0
// Adware
const Malware_ClassificationIds_Adware MalwareClassificationIds = 1
// Backdoor
const Malware_ClassificationIds_Backdoor MalwareClassificationIds = 2
// Bot
const Malware_ClassificationIds_Bot MalwareClassificationIds = 3
// Bootkit
const Malware_ClassificationIds_Bootkit MalwareClassificationIds = 4
// DDOS
const Malware_ClassificationIds_DDOS MalwareClassificationIds = 5
// Downloader
const Malware_ClassificationIds_Downloader MalwareClassificationIds = 6
// Dropper
const Malware_ClassificationIds_Dropper MalwareClassificationIds = 7
// Exploit_Kit
const Malware_ClassificationIds_Exploit_Kit MalwareClassificationIds = 8
// Keylogger
const Malware_ClassificationIds_Keylogger MalwareClassificationIds = 9
// Ransomware
const Malware_ClassificationIds_Ransomware MalwareClassificationIds = 10
// Remote_Access_Trojan
const Malware_ClassificationIds_Remote_Access_Trojan MalwareClassificationIds = 11
// Resource_Exploitation
const Malware_ClassificationIds_Resource_Exploitation MalwareClassificationIds = 13
// Rogue_Security_Software
const Malware_ClassificationIds_Rogue_Security_Software MalwareClassificationIds = 14
// Rootkit
const Malware_ClassificationIds_Rootkit MalwareClassificationIds = 15
// Screen_Capture
const Malware_ClassificationIds_Screen_Capture MalwareClassificationIds = 16
// Spyware
const Malware_ClassificationIds_Spyware MalwareClassificationIds = 17
// Trojan
const Malware_ClassificationIds_Trojan MalwareClassificationIds = 18
// Virus
const Malware_ClassificationIds_Virus MalwareClassificationIds = 19
// Webshell
const Malware_ClassificationIds_Webshell MalwareClassificationIds = 20
// Wiper
const Malware_ClassificationIds_Wiper MalwareClassificationIds = 21
// Worm
const Malware_ClassificationIds_Worm MalwareClassificationIds = 22
// Other. The classification is not mapped. See the <code>classifications</code> attribute, which contains a data source specific value.
const Malware_ClassificationIds_Other MalwareClassificationIds = 99

// NetworkConnectionInfoBoundaryId Values
// Boundary ID. <p>The normalized identifier of the boundary of the connection. </p><p> For cloud connections, this translates to the traffic-boundary (same VPC, through IGW, etc.). For traditional networks, this is described as Local, Internal, or External.</p>

// Unknown. The connection boundary is unknown.
const NetworkConnectionInfo_Boundary_Unknown NetworkConnectionInfoBoundaryId = 0
// Localhost. Local network traffic on the same endpoint.
const NetworkConnectionInfo_Boundary_Localhost NetworkConnectionInfoBoundaryId = 1
// Internal. Internal network traffic between two endpoints inside network.
const NetworkConnectionInfo_Boundary_Internal NetworkConnectionInfoBoundaryId = 2
// External. External network traffic between two endpoints on the Internet or outside the network.
const NetworkConnectionInfo_Boundary_External NetworkConnectionInfoBoundaryId = 3
// SameVPC. Through another resource in the same VPC
const NetworkConnectionInfo_Boundary_SameVPC NetworkConnectionInfoBoundaryId = 4
// Internet_VPCGateway. Through an Internet gateway or a gateway VPC endpoint
const NetworkConnectionInfo_Boundary_Internet_VPCGateway NetworkConnectionInfoBoundaryId = 5
// VirtualPrivateGateway. Through a virtual private gateway
const NetworkConnectionInfo_Boundary_VirtualPrivateGateway NetworkConnectionInfoBoundaryId = 6
// Intra_regionVPC. Through an intra-region VPC peering connection
const NetworkConnectionInfo_Boundary_Intra_regionVPC NetworkConnectionInfoBoundaryId = 7
// Inter_regionVPC. Through an inter-region VPC peering connection
const NetworkConnectionInfo_Boundary_Inter_regionVPC NetworkConnectionInfoBoundaryId = 8
// LocalGateway. Through a local gateway
const NetworkConnectionInfo_Boundary_LocalGateway NetworkConnectionInfoBoundaryId = 9
// GatewayVPC. Through a gateway VPC endpoint (Nitro-based instances only)
const NetworkConnectionInfo_Boundary_GatewayVPC NetworkConnectionInfoBoundaryId = 10
// InternetGateway. Through an Internet gateway (Nitro-based instances only)
const NetworkConnectionInfo_Boundary_InternetGateway NetworkConnectionInfoBoundaryId = 11
// Other. The boundary is not mapped. See the <code>boundary</code> attribute, which contains a data source specific value.
const NetworkConnectionInfo_Boundary_Other NetworkConnectionInfoBoundaryId = 99

// NetworkConnectionInfoDirectionId Values
// Direction ID. The normalized identifier of the direction of the initiated connection, traffic, or email.

// Unknown. The connection direction is unknown.
const NetworkConnectionInfo_Direction_Unknown NetworkConnectionInfoDirectionId = 0
// Inbound. Inbound network connection. The connection was originated from the Internet or outside network, destined for services on the inside network.
const NetworkConnectionInfo_Direction_Inbound NetworkConnectionInfoDirectionId = 1
// Outbound. Outbound network connection. The connection was originated from inside the network, destined for services on the Internet or outside network.
const NetworkConnectionInfo_Direction_Outbound NetworkConnectionInfoDirectionId = 2
// Lateral. Lateral network connection. The connection was originated from inside the network, destined for services on the inside network.
const NetworkConnectionInfo_Direction_Lateral NetworkConnectionInfoDirectionId = 3
// Other. The direction is not mapped. See the <code>direction</code> attribute, which contains a data source specific value.
const NetworkConnectionInfo_Direction_Other NetworkConnectionInfoDirectionId = 99

// NetworkConnectionInfoProtocolVerId Values
// IP Version ID. The Internet Protocol version identifier.

// Unknown. The protocol version is unknown.
const NetworkConnectionInfo_ProtocolVer_Unknown NetworkConnectionInfoProtocolVerId = 0
// InternetProtocolversion4_IPv4_
const NetworkConnectionInfo_ProtocolVer_InternetProtocolversion4_IPv4_ NetworkConnectionInfoProtocolVerId = 4
// InternetProtocolversion6_IPv6_
const NetworkConnectionInfo_ProtocolVer_InternetProtocolversion6_IPv6_ NetworkConnectionInfoProtocolVerId = 6
// Other. The protocol version is not mapped. See the <code>protocol_ver</code> attribute, which contains a data source specific value.
const NetworkConnectionInfo_ProtocolVer_Other NetworkConnectionInfoProtocolVerId = 99

// NetworkEndpointTypeId Values
// Type ID. The network endpoint type ID.

// Unknown. The type is unknown.
const NetworkEndpoint_Type_Unknown NetworkEndpointTypeId = 0
// Server. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Server/'>server</a>.
const NetworkEndpoint_Type_Server NetworkEndpointTypeId = 1
// Desktop. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:DesktopComputer/'>desktop computer</a>.
const NetworkEndpoint_Type_Desktop NetworkEndpointTypeId = 2
// Laptop. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:LaptopComputer/'>laptop computer</a>.
const NetworkEndpoint_Type_Laptop NetworkEndpointTypeId = 3
// Tablet. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:TabletComputer/'>tablet computer</a>.
const NetworkEndpoint_Type_Tablet NetworkEndpointTypeId = 4
// Mobile. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:MobilePhone/'>mobile phone</a>.
const NetworkEndpoint_Type_Mobile NetworkEndpointTypeId = 5
// Virtual. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:VirtualizationSoftware/'>virtual machine</a>.
const NetworkEndpoint_Type_Virtual NetworkEndpointTypeId = 6
// IOT. An <a target='_blank' href='https://www.techtarget.com/iotagenda/definition/IoT-device'>IOT (Internet of Things) device</a>.
const NetworkEndpoint_Type_IOT NetworkEndpointTypeId = 7
// Browser. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Browser/'>web browser</a>.
const NetworkEndpoint_Type_Browser NetworkEndpointTypeId = 8
// Firewall. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Firewall/'>networking firewall</a>.
const NetworkEndpoint_Type_Firewall NetworkEndpointTypeId = 9
// Switch. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Switch/'>networking switch</a>.
const NetworkEndpoint_Type_Switch NetworkEndpointTypeId = 10
// Hub. A <a target='_blank' href='https://en.wikipedia.org/wiki/Ethernet_hub'>networking hub</a>.
const NetworkEndpoint_Type_Hub NetworkEndpointTypeId = 11
// Router. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Router/'>networking router</a>.
const NetworkEndpoint_Type_Router NetworkEndpointTypeId = 12
// IDS. An <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:IntrusionDetectionSystem/'>intrusion detection system</a>.
const NetworkEndpoint_Type_IDS NetworkEndpointTypeId = 13
// IPS. An <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:IntrusionPreventionSystem/'>intrusion prevention system</a>.
const NetworkEndpoint_Type_IPS NetworkEndpointTypeId = 14
// LoadBalancer. A <a target='_blank' href='https://en.wikipedia.org/wiki/Load_balancing_(computing)'> Load Balancer device.
const NetworkEndpoint_Type_LoadBalancer NetworkEndpointTypeId = 15
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const NetworkEndpoint_Type_Other NetworkEndpointTypeId = 99

// NetworkInterfaceTypeId Values
// Type ID. The network interface type identifier.

// Unknown. The type is unknown.
const NetworkInterface_Type_Unknown NetworkInterfaceTypeId = 0
// Wired
const NetworkInterface_Type_Wired NetworkInterfaceTypeId = 1
// Wireless
const NetworkInterface_Type_Wireless NetworkInterfaceTypeId = 2
// Mobile
const NetworkInterface_Type_Mobile NetworkInterfaceTypeId = 3
// Tunnel
const NetworkInterface_Type_Tunnel NetworkInterfaceTypeId = 4
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const NetworkInterface_Type_Other NetworkInterfaceTypeId = 99

// NetworkProxyTypeId Values
// Type ID. The network endpoint type ID.

// Unknown. The type is unknown.
const NetworkProxy_Type_Unknown NetworkProxyTypeId = 0
// Server. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Server/'>server</a>.
const NetworkProxy_Type_Server NetworkProxyTypeId = 1
// Desktop. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:DesktopComputer/'>desktop computer</a>.
const NetworkProxy_Type_Desktop NetworkProxyTypeId = 2
// Laptop. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:LaptopComputer/'>laptop computer</a>.
const NetworkProxy_Type_Laptop NetworkProxyTypeId = 3
// Tablet. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:TabletComputer/'>tablet computer</a>.
const NetworkProxy_Type_Tablet NetworkProxyTypeId = 4
// Mobile. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:MobilePhone/'>mobile phone</a>.
const NetworkProxy_Type_Mobile NetworkProxyTypeId = 5
// Virtual. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:VirtualizationSoftware/'>virtual machine</a>.
const NetworkProxy_Type_Virtual NetworkProxyTypeId = 6
// IOT. An <a target='_blank' href='https://www.techtarget.com/iotagenda/definition/IoT-device'>IOT (Internet of Things) device</a>.
const NetworkProxy_Type_IOT NetworkProxyTypeId = 7
// Browser. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Browser/'>web browser</a>.
const NetworkProxy_Type_Browser NetworkProxyTypeId = 8
// Firewall. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Firewall/'>networking firewall</a>.
const NetworkProxy_Type_Firewall NetworkProxyTypeId = 9
// Switch. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Switch/'>networking switch</a>.
const NetworkProxy_Type_Switch NetworkProxyTypeId = 10
// Hub. A <a target='_blank' href='https://en.wikipedia.org/wiki/Ethernet_hub'>networking hub</a>.
const NetworkProxy_Type_Hub NetworkProxyTypeId = 11
// Router. A <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Router/'>networking router</a>.
const NetworkProxy_Type_Router NetworkProxyTypeId = 12
// IDS. An <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:IntrusionDetectionSystem/'>intrusion detection system</a>.
const NetworkProxy_Type_IDS NetworkProxyTypeId = 13
// IPS. An <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:IntrusionPreventionSystem/'>intrusion prevention system</a>.
const NetworkProxy_Type_IPS NetworkProxyTypeId = 14
// LoadBalancer. A <a target='_blank' href='https://en.wikipedia.org/wiki/Load_balancing_(computing)'> Load Balancer device.
const NetworkProxy_Type_LoadBalancer NetworkProxyTypeId = 15
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const NetworkProxy_Type_Other NetworkProxyTypeId = 99

// ObservableTypeId Values
// Type ID. The observable value type identifier.

// Unknown. Unknown observable data type.
const Observable_Type_Unknown ObservableTypeId = 0
// Hostname. Unique name assigned to a device connected to a computer network. A domain name in general is an Internet address that can be resolved through the Domain Name System (DNS). For example: <code>r2-d2.example.com</code>.
const Observable_Type_Hostname ObservableTypeId = 1
// IPAddress. Internet Protocol address (IP address), in either IPv4 or IPv6 format. For example, <code>192.168.200.24</code> or <code>2001:0db8:85a3:0000:0000:8a2e:0370:7334</code>.
const Observable_Type_IPAddress ObservableTypeId = 2
// MACAddress. Media Access Control (MAC) address. For example: <code>18:36:F3:98:4F:9A</code>.
const Observable_Type_MACAddress ObservableTypeId = 3
// UserName. User name. For example: <code>john_doe</code>.
const Observable_Type_UserName ObservableTypeId = 4
// EmailAddress. Email address. For example: <code>john_doe@example.com</code>.
const Observable_Type_EmailAddress ObservableTypeId = 5
// URLString. Uniform Resource Locator (URL) string. For example: <code>http://www.example.com/download/trouble.exe</code>.
const Observable_Type_URLString ObservableTypeId = 6
// FileName. File name. For example: <code>text-file.txt</code>.
const Observable_Type_FileName ObservableTypeId = 7
// Hash. Hash. A unique value that corresponds to the content of the file, image, ja3_hash or hassh found in the schema. For example MD5: <code>3172ac7e2b55cbb81f04a6e65855a628</code>.
const Observable_Type_Hash ObservableTypeId = 8
// ProcessName. Process name. For example: <code>Notepad</code>.
const Observable_Type_ProcessName ObservableTypeId = 9
// ResourceUID. Resource unique identifier. For example, S3 Bucket name or EC2 Instance ID.
const Observable_Type_ResourceUID ObservableTypeId = 10
// Port. The TCP/UDP port number. For example: <code>80</code> or <code>22</code>.
const Observable_Type_Port ObservableTypeId = 11
// Subnet. The subnet represented in a CIDR notation, using the format network_address/prefix_length. The network_address can be in either IPv4 or IPv6 format. The prefix length indicates the number of bits used for the network portion, and the remaining bits are available for host addresses within that subnet. <div>For example:<ul><li>192.168.1.0/24</li><li>2001:0db8:85a3:0000::/64</li></ul></div>
const Observable_Type_Subnet ObservableTypeId = 12
// Endpoint. The Endpoint object describes a physical or virtual device that connects to and exchanges information with a computer network. Some examples of endpoints are mobile devices, desktop computers, virtual machines, embedded devices, and servers. Internet-of-Things devices—like cameras, lighting, refrigerators, security systems, smart speakers, and thermostats—are also endpoints.
const Observable_Type_Endpoint ObservableTypeId = 20
// User. The User object describes the characteristics of a user/person or a security principal. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:UserAccount/'>d3f:UserAccount</a>.
const Observable_Type_User ObservableTypeId = 21
// Email. The Email object describes the email metadata such as sender, recipients, and direction. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Email/'>d3f:Email</a>.
const Observable_Type_Email ObservableTypeId = 22
// UniformResourceLocator. The Uniform Resource Locator(URL) object describes the characteristics of a URL. Defined in <a target='_blank' href='https://datatracker.ietf.org/doc/html/rfc1738'>RFC 1738</a> and by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:URL/'>d3f:URL</a>.
const Observable_Type_UniformResourceLocator ObservableTypeId = 23
// File. The File object represents the metadata associated with a file stored in a computer system. It encompasses information about the file itself, including its attributes, properties, and organizational details. Defined by D3FEND <a target='_blank' href='https://next.d3fend.mitre.org/dao/artifact/d3f:File/'>d3f:File</a>.
const Observable_Type_File ObservableTypeId = 24
// Process. The Process object describes a running instance of a launched program. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Process/'>d3f:Process</a>.
const Observable_Type_Process ObservableTypeId = 25
// GeoLocation. The Geo Location object describes a geographical location, usually associated with an IP address. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:PhysicalLocation/'>d3f:PhysicalLocation</a>.
const Observable_Type_GeoLocation ObservableTypeId = 26
// Container. The Container object describes an instance of a specific container. A container is a prepackaged, portable system image that runs isolated on an existing system using a container runtime like containerd.
const Observable_Type_Container ObservableTypeId = 27
// Fingerprint. The Fingerprint object provides detailed information about a digital fingerprint, which is a compact representation of data used to identify a longer piece of information, such as a public key or file content. It contains the algorithm and value of the fingerprint, enabling efficient and reliable identification of the associated data.
const Observable_Type_Fingerprint ObservableTypeId = 30
// Other. The observable data type is not mapped. See the <code>type</code> attribute, which may contain data source specific value.
const Observable_Type_Other ObservableTypeId = 99

// OsTypeId Values
// Type ID. The type identifier of the operating system.

// Unknown. The type is unknown.
const Os_Type_Unknown OsTypeId = 0
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const Os_Type_Other OsTypeId = 99
// Windows
const Os_Type_Windows OsTypeId = 100
// WindowsMobile
const Os_Type_WindowsMobile OsTypeId = 101
// Linux
const Os_Type_Linux OsTypeId = 200
// Android
const Os_Type_Android OsTypeId = 201
// macOS
const Os_Type_macOS OsTypeId = 300
// iOS
const Os_Type_iOS OsTypeId = 301
// iPadOS
const Os_Type_iPadOS OsTypeId = 302
// Solaris
const Os_Type_Solaris OsTypeId = 400
// AIX
const Os_Type_AIX OsTypeId = 401
// HP_UX
const Os_Type_HP_UX OsTypeId = 402

// OsintConfidenceId Values
// Confidence Id. The normalized confidence refers to the accuracy of collected information related to the OSINT or how pertinent an indicator or analysis is to a specific event or finding. A low confidence means that the information collected or analysis conducted lacked detail or is not accurate enough to qualify an indicator as fully malicious.

// Unknown. The normalized confidence is unknown.
const Osint_Confidence_Unknown OsintConfidenceId = 0
// Low
const Osint_Confidence_Low OsintConfidenceId = 1
// Medium
const Osint_Confidence_Medium OsintConfidenceId = 2
// High
const Osint_Confidence_High OsintConfidenceId = 3
// Other. The confidence is not mapped to the defined enum values. See the <code>confidence</code> attribute, which contains a data source specific value.
const Osint_Confidence_Other OsintConfidenceId = 99

// OsintTlp Values
// Traffic Light Protocol. The <a target='_blank' href='https://www.first.org/tlp/'>Traffic Light Protocol</a> was created to facilitate greater sharing of potentially sensitive information and more effective collaboration. TLP provides a simple and intuitive schema for indicating with whom potentially sensitive information can be shared.

// TLP_AMBER. TLP:AMBER is for limited disclosure, recipients can only spread this on a need-to-know basis within their organization and its clients. Note that TLP:AMBER+STRICT restricts sharing to the organization only. Sources may use TLP:AMBER when information requires support to be effectively acted upon, yet carries risk to privacy, reputation, or operations if shared outside of the organizations involved. Recipients may share TLP:AMBER information with members of their own organization and its clients, but only on a need-to-know basis to protect their organization and its clients and prevent further harm. Note: if the source wants to restrict sharing to the organization only, they must specify TLP:AMBER+STRICT.
const Osint_Tlp_TLP_AMBER OsintTlp = "AMBER"
// TLP_AMBER_STRICT. TLP:AMBER is for limited disclosure, recipients can only spread this on a need-to-know basis within their organization and its clients. Note that TLP:AMBER+STRICT restricts sharing to the organization only. Sources may use TLP:AMBER when information requires support to be effectively acted upon, yet carries risk to privacy, reputation, or operations if shared outside of the organizations involved. Recipients may share TLP:AMBER information with members of their own organization and its clients, but only on a need-to-know basis to protect their organization and its clients and prevent further harm. Note: if the source wants to restrict sharing to the organization only, they must specify TLP:AMBER+STRICT.
const Osint_Tlp_TLP_AMBER_STRICT OsintTlp = "AMBER STRICT"
// TLP_CLEAR. TLP:CLEAR denotes that recipients can spread this to the world, there is no limit on disclosure. Sources may use TLP:CLEAR when information carries minimal or no foreseeable risk of misuse, in accordance with applicable rules and procedures for public release. Subject to standard copyright rules, TLP:CLEAR information may be shared without restriction.
const Osint_Tlp_TLP_CLEAR OsintTlp = "CLEAR"
// TLP_GREEN. TLP:GREEN is for limited disclosure, recipients can spread this within their community. Sources may use TLP:GREEN when information is useful to increase awareness within their wider community. Recipients may share TLP:GREEN information with peers and partner organizations within their community, but not via publicly accessible channels. TLP:GREEN information may not be shared outside of the community. Note: when “community” is not defined, assume the cybersecurity/defense community.
const Osint_Tlp_TLP_GREEN OsintTlp = "GREEN"
// TLP_RED. TLP:RED is for the eyes and ears of individual recipients only, no further disclosure. Sources may use TLP:RED when information cannot be effectively acted upon without significant risk for the privacy, reputation, or operations of the organizations involved. Recipients may therefore not share TLP:RED information with anyone else. In the context of a meeting, for example, TLP:RED information is limited to those present at the meeting.
const Osint_Tlp_TLP_RED OsintTlp = "RED"

// OsintTypeId Values
// Indicator Type ID. The OSINT indicator type ID.

// Unknown. The indicator type is ambiguous or there is not a related indicator for the OSINT object.
const Osint_Type_Unknown OsintTypeId = 0
// IPAddress. An IPv4 or IPv6 address.
const Osint_Type_IPAddress OsintTypeId = 1
// Domain. A full-qualified domain name (FQDN), subdomain, or partial domain.
const Osint_Type_Domain OsintTypeId = 2
// Hostname. A hostname or computer name.
const Osint_Type_Hostname OsintTypeId = 3
// Hash. Any type of hash e.g., MD5, SHA1, SHA2, BLAKE, BLAKE2, etc. generated from a file, malware sample, request header, or otherwise.
const Osint_Type_Hash OsintTypeId = 4
// URL. A Uniform Resource Locator (URL) or Uniform Resource Indicator (URI).
const Osint_Type_URL OsintTypeId = 5
// UserAgent. A User Agent typically seen in HTTP request headers.
const Osint_Type_UserAgent OsintTypeId = 6
// DigitalCertificate. The serial number, fingerprint, or full content of an X.509 digital certificate.
const Osint_Type_DigitalCertificate OsintTypeId = 7
// Email. The contents of an email or any related information to an email object.
const Osint_Type_Email OsintTypeId = 8
// EmailAddress. An email address.
const Osint_Type_EmailAddress OsintTypeId = 9
// Vulnerability. A CVE ID, CWE ID, or other identifier for a weakness, exploit, bug, or misconfiguration.
const Osint_Type_Vulnerability OsintTypeId = 10
// Other. The indicator type is not directly listed.
const Osint_Type_Other OsintTypeId = 99

// PackageTypeId Values
// Type ID. The type of software package.

// Unknown. The type is unknown.
const Package_Type_Unknown PackageTypeId = 0
// Application. An application software package.
const Package_Type_Application PackageTypeId = 1
// OperatingSystem. An operating system software package.
const Package_Type_OperatingSystem PackageTypeId = 2
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const Package_Type_Other PackageTypeId = 99

// ProcessIntegrityId Values
// Integrity Level. The normalized identifier of the process integrity level (Windows only).

// Unknown. The integrity level is unknown.
const Process_Integrity_Unknown ProcessIntegrityId = 0
// Untrusted
const Process_Integrity_Untrusted ProcessIntegrityId = 1
// Low
const Process_Integrity_Low ProcessIntegrityId = 2
// Medium
const Process_Integrity_Medium ProcessIntegrityId = 3
// High
const Process_Integrity_High ProcessIntegrityId = 4
// System
const Process_Integrity_System ProcessIntegrityId = 5
// Protected
const Process_Integrity_Protected ProcessIntegrityId = 6
// Other. The integrity level is not mapped. See the <code>integrity</code> attribute, which contains a data source specific value.
const Process_Integrity_Other ProcessIntegrityId = 99

// ReputationScoreId Values
// Reputation Score ID. The normalized reputation score identifier.

// Unknown. The reputation score is unknown.
const Reputation_Score_Unknown ReputationScoreId = 0
// VerySafe. Long history of good behavior.
const Reputation_Score_VerySafe ReputationScoreId = 1
// Safe. Consistently good behavior.
const Reputation_Score_Safe ReputationScoreId = 2
// ProbablySafe. Reasonable history of good behavior.
const Reputation_Score_ProbablySafe ReputationScoreId = 3
// LeansSafe. Starting to establish a history of normal behavior.
const Reputation_Score_LeansSafe ReputationScoreId = 4
// MaynotbeSafe. No established history of normal behavior.
const Reputation_Score_MaynotbeSafe ReputationScoreId = 5
// ExerciseCaution. Starting to establish a history of suspicious or risky behavior.
const Reputation_Score_ExerciseCaution ReputationScoreId = 6
// Suspicious_Risky. A site with a history of suspicious or risky behavior. (spam, scam, potentially unwanted software, potentially malicious).
const Reputation_Score_Suspicious_Risky ReputationScoreId = 7
// PossiblyMalicious. Strong possibility of maliciousness.
const Reputation_Score_PossiblyMalicious ReputationScoreId = 8
// ProbablyMalicious. Indicators of maliciousness.
const Reputation_Score_ProbablyMalicious ReputationScoreId = 9
// Malicious. Proven evidence of maliciousness.
const Reputation_Score_Malicious ReputationScoreId = 10
// Other. The reputation score is not mapped. See the <code>rep_score</code> attribute, which contains a data source specific value.
const Reputation_Score_Other ReputationScoreId = 99

// TimespanTypeId Values
// Time Span Type ID. The normalized identifier for the time span duration type.

// Unknown. The type is unknown.
const Timespan_Type_Unknown TimespanTypeId = 0
// Milliseconds
const Timespan_Type_Milliseconds TimespanTypeId = 1
// Seconds
const Timespan_Type_Seconds TimespanTypeId = 2
// Minutes
const Timespan_Type_Minutes TimespanTypeId = 3
// Hours
const Timespan_Type_Hours TimespanTypeId = 4
// Days
const Timespan_Type_Days TimespanTypeId = 5
// Weeks
const Timespan_Type_Weeks TimespanTypeId = 6
// Months
const Timespan_Type_Months TimespanTypeId = 7
// Years
const Timespan_Type_Years TimespanTypeId = 8
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const Timespan_Type_Other TimespanTypeId = 99

// TlsExtensionTypeId Values
// Type ID. The TLS extension type identifier. See <a target='_blank' href='https://datatracker.ietf.org/doc/html/rfc8446#page-35'>The Transport Layer Security (TLS) extension page</a>.

// server_name. The Server Name Indication extension.
const TlsExtension_Type_server_name TlsExtensionTypeId = 0
// maximum_fragment_length. The Maximum Fragment Length Negotiation extension.
const TlsExtension_Type_maximum_fragment_length TlsExtensionTypeId = 1
// status_request. The Certificate Status Request extension.
const TlsExtension_Type_status_request TlsExtensionTypeId = 5
// supported_groups. The Supported Groups extension.
const TlsExtension_Type_supported_groups TlsExtensionTypeId = 10
// signature_algorithms. The Signature Algorithms extension.
const TlsExtension_Type_signature_algorithms TlsExtensionTypeId = 13
// use_srtp. The Use SRTP data protection extension.
const TlsExtension_Type_use_srtp TlsExtensionTypeId = 14
// heartbeat. The Heartbeat extension.
const TlsExtension_Type_heartbeat TlsExtensionTypeId = 15
// application_layer_protocol_negotiation. The Application-Layer Protocol Negotiation extension.
const TlsExtension_Type_application_layer_protocol_negotiation TlsExtensionTypeId = 16
// signed_certificate_timestamp. The Signed Certificate Timestamp extension.
const TlsExtension_Type_signed_certificate_timestamp TlsExtensionTypeId = 18
// client_certificate_type. The Client Certificate Type extension.
const TlsExtension_Type_client_certificate_type TlsExtensionTypeId = 19
// server_certificate_type. The Server Certificate Type extension.
const TlsExtension_Type_server_certificate_type TlsExtensionTypeId = 20
// padding. The Padding extension.
const TlsExtension_Type_padding TlsExtensionTypeId = 21
// pre_shared_key. The Pre Shared Key extension.
const TlsExtension_Type_pre_shared_key TlsExtensionTypeId = 41
// early_data. The Early Data extension.
const TlsExtension_Type_early_data TlsExtensionTypeId = 42
// supported_versions. The Supported Versions extension.
const TlsExtension_Type_supported_versions TlsExtensionTypeId = 43
// cookie. The Cookie extension.
const TlsExtension_Type_cookie TlsExtensionTypeId = 44
// psk_key_exchange_modes. The Pre-Shared Key Exchange Modes extension.
const TlsExtension_Type_psk_key_exchange_modes TlsExtensionTypeId = 45
// certificate_authorities. The Certificate Authorities extension.
const TlsExtension_Type_certificate_authorities TlsExtensionTypeId = 47
// oid_filters. The OID Filters extension.
const TlsExtension_Type_oid_filters TlsExtensionTypeId = 48
// post_handshake_auth. The Post-Handshake Client Authentication extension.
const TlsExtension_Type_post_handshake_auth TlsExtensionTypeId = 49
// signature_algorithms_cert. The Signature Algorithms extension.
const TlsExtension_Type_signature_algorithms_cert TlsExtensionTypeId = 50
// key_share. The Key Share extension.
const TlsExtension_Type_key_share TlsExtensionTypeId = 51
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const TlsExtension_Type_Other TlsExtensionTypeId = 99

// UrlCategoryIds Values
// Website Categorization IDs. The Website categorization identifiers.

// Unknown. The Domain/URL category is unknown.
const Url_CategoryIds_Unknown UrlCategoryIds = 0
// Adult_MatureContent
const Url_CategoryIds_Adult_MatureContent UrlCategoryIds = 1
// Pornography
const Url_CategoryIds_Pornography UrlCategoryIds = 3
// SexEducation
const Url_CategoryIds_SexEducation UrlCategoryIds = 4
// IntimateApparel_Swimsuit
const Url_CategoryIds_IntimateApparel_Swimsuit UrlCategoryIds = 5
// Nudity
const Url_CategoryIds_Nudity UrlCategoryIds = 6
// Extreme
const Url_CategoryIds_Extreme UrlCategoryIds = 7
// Scam_Questionable_Illegal
const Url_CategoryIds_Scam_Questionable_Illegal UrlCategoryIds = 9
// Gambling
const Url_CategoryIds_Gambling UrlCategoryIds = 11
// Violence_Hate_Racism
const Url_CategoryIds_Violence_Hate_Racism UrlCategoryIds = 14
// Weapons
const Url_CategoryIds_Weapons UrlCategoryIds = 15
// Abortion
const Url_CategoryIds_Abortion UrlCategoryIds = 16
// Hacking
const Url_CategoryIds_Hacking UrlCategoryIds = 17
// Phishing
const Url_CategoryIds_Phishing UrlCategoryIds = 18
// Entertainment
const Url_CategoryIds_Entertainment UrlCategoryIds = 20
// Business_Economy
const Url_CategoryIds_Business_Economy UrlCategoryIds = 21
// AlternativeSpirituality_Belief
const Url_CategoryIds_AlternativeSpirituality_Belief UrlCategoryIds = 22
// Alcohol
const Url_CategoryIds_Alcohol UrlCategoryIds = 23
// Tobacco
const Url_CategoryIds_Tobacco UrlCategoryIds = 24
// ControlledSubstances
const Url_CategoryIds_ControlledSubstances UrlCategoryIds = 25
// ChildPornography
const Url_CategoryIds_ChildPornography UrlCategoryIds = 26
// Education
const Url_CategoryIds_Education UrlCategoryIds = 27
// CharitableOrganizations
const Url_CategoryIds_CharitableOrganizations UrlCategoryIds = 29
// Art_Culture
const Url_CategoryIds_Art_Culture UrlCategoryIds = 30
// FinancialServices
const Url_CategoryIds_FinancialServices UrlCategoryIds = 31
// Brokerage_Trading
const Url_CategoryIds_Brokerage_Trading UrlCategoryIds = 32
// Games
const Url_CategoryIds_Games UrlCategoryIds = 33
// Government_Legal
const Url_CategoryIds_Government_Legal UrlCategoryIds = 34
// Military
const Url_CategoryIds_Military UrlCategoryIds = 35
// Political_SocialAdvocacy
const Url_CategoryIds_Political_SocialAdvocacy UrlCategoryIds = 36
// Health
const Url_CategoryIds_Health UrlCategoryIds = 37
// Technology_Internet
const Url_CategoryIds_Technology_Internet UrlCategoryIds = 38
// SearchEngines_Portals
const Url_CategoryIds_SearchEngines_Portals UrlCategoryIds = 40
// MaliciousSources_Malnets
const Url_CategoryIds_MaliciousSources_Malnets UrlCategoryIds = 43
// MaliciousOutboundData_Botnets
const Url_CategoryIds_MaliciousOutboundData_Botnets UrlCategoryIds = 44
// JobSearch_Careers
const Url_CategoryIds_JobSearch_Careers UrlCategoryIds = 45
// News_Media
const Url_CategoryIds_News_Media UrlCategoryIds = 46
// Personals_Dating
const Url_CategoryIds_Personals_Dating UrlCategoryIds = 47
// Reference
const Url_CategoryIds_Reference UrlCategoryIds = 49
// MixedContent_PotentiallyAdult
const Url_CategoryIds_MixedContent_PotentiallyAdult UrlCategoryIds = 50
// Chat_IM_SMS
const Url_CategoryIds_Chat_IM_SMS UrlCategoryIds = 51
// Email
const Url_CategoryIds_Email UrlCategoryIds = 52
// Newsgroups_Forums
const Url_CategoryIds_Newsgroups_Forums UrlCategoryIds = 53
// Religion
const Url_CategoryIds_Religion UrlCategoryIds = 54
// SocialNetworking
const Url_CategoryIds_SocialNetworking UrlCategoryIds = 55
// FileStorage_Sharing
const Url_CategoryIds_FileStorage_Sharing UrlCategoryIds = 56
// RemoteAccessTools
const Url_CategoryIds_RemoteAccessTools UrlCategoryIds = 57
// Shopping
const Url_CategoryIds_Shopping UrlCategoryIds = 58
// Auctions
const Url_CategoryIds_Auctions UrlCategoryIds = 59
// RealEstate
const Url_CategoryIds_RealEstate UrlCategoryIds = 60
// Society_DailyLiving
const Url_CategoryIds_Society_DailyLiving UrlCategoryIds = 61
// PersonalSites
const Url_CategoryIds_PersonalSites UrlCategoryIds = 63
// Restaurants_Dining_Food
const Url_CategoryIds_Restaurants_Dining_Food UrlCategoryIds = 64
// Sports_Recreation
const Url_CategoryIds_Sports_Recreation UrlCategoryIds = 65
// Travel
const Url_CategoryIds_Travel UrlCategoryIds = 66
// Vehicles
const Url_CategoryIds_Vehicles UrlCategoryIds = 67
// Humor_Jokes
const Url_CategoryIds_Humor_Jokes UrlCategoryIds = 68
// SoftwareDownloads
const Url_CategoryIds_SoftwareDownloads UrlCategoryIds = 71
// Peer_to_Peer_P2P_
const Url_CategoryIds_Peer_to_Peer_P2P_ UrlCategoryIds = 83
// Audio_VideoClips
const Url_CategoryIds_Audio_VideoClips UrlCategoryIds = 84
// Office_BusinessApplications
const Url_CategoryIds_Office_BusinessApplications UrlCategoryIds = 85
// ProxyAvoidance
const Url_CategoryIds_ProxyAvoidance UrlCategoryIds = 86
// ForKids
const Url_CategoryIds_ForKids UrlCategoryIds = 87
// WebAds_Analytics
const Url_CategoryIds_WebAds_Analytics UrlCategoryIds = 88
// WebHosting
const Url_CategoryIds_WebHosting UrlCategoryIds = 89
// Uncategorized
const Url_CategoryIds_Uncategorized UrlCategoryIds = 90
// Suspicious
const Url_CategoryIds_Suspicious UrlCategoryIds = 92
// SexualExpression
const Url_CategoryIds_SexualExpression UrlCategoryIds = 93
// Translation
const Url_CategoryIds_Translation UrlCategoryIds = 95
// Non_Viewable_Infrastructure
const Url_CategoryIds_Non_Viewable_Infrastructure UrlCategoryIds = 96
// ContentServers
const Url_CategoryIds_ContentServers UrlCategoryIds = 97
// Placeholders
const Url_CategoryIds_Placeholders UrlCategoryIds = 98
// Other. The Domain/URL category is not mapped. See the <code>categories</code> attribute, which contains a data source specific value.
const Url_CategoryIds_Other UrlCategoryIds = 99
// Spam
const Url_CategoryIds_Spam UrlCategoryIds = 101
// PotentiallyUnwantedSoftware
const Url_CategoryIds_PotentiallyUnwantedSoftware UrlCategoryIds = 102
// DynamicDNSHost
const Url_CategoryIds_DynamicDNSHost UrlCategoryIds = 103
// E_Card_Invitations
const Url_CategoryIds_E_Card_Invitations UrlCategoryIds = 106
// Informational
const Url_CategoryIds_Informational UrlCategoryIds = 107
// Computer_InformationSecurity
const Url_CategoryIds_Computer_InformationSecurity UrlCategoryIds = 108
// InternetConnectedDevices
const Url_CategoryIds_InternetConnectedDevices UrlCategoryIds = 109
// InternetTelephony
const Url_CategoryIds_InternetTelephony UrlCategoryIds = 110
// OnlineMeetings
const Url_CategoryIds_OnlineMeetings UrlCategoryIds = 111
// MediaSharing
const Url_CategoryIds_MediaSharing UrlCategoryIds = 112
// Radio_AudioStreams
const Url_CategoryIds_Radio_AudioStreams UrlCategoryIds = 113
// TV_VideoStreams
const Url_CategoryIds_TV_VideoStreams UrlCategoryIds = 114
// Piracy_CopyrightConcerns
const Url_CategoryIds_Piracy_CopyrightConcerns UrlCategoryIds = 118
// Marijuana
const Url_CategoryIds_Marijuana UrlCategoryIds = 121

// UserMfaStatusId Values
// MFA Status ID. The normalized identifier of the user's multi-factor authentication status.

// Unknown. The status is unknown.
const User_MfaStatus_Unknown UserMfaStatusId = 0
// Enabled. Multi-factor authentication is on for this user.
const User_MfaStatus_Enabled UserMfaStatusId = 1
// NotEnabled. TMulti-factor authentication is off for this user.
const User_MfaStatus_NotEnabled UserMfaStatusId = 2
// Other. The event status is not mapped. See the <code>user_status</code> attribute, which contains a data source specific value.
const User_MfaStatus_Other UserMfaStatusId = 99

// UserRiskLevelId Values
// Risk Level ID. The normalized risk level id.

// Info
const User_RiskLevel_Info UserRiskLevelId = 0
// Low
const User_RiskLevel_Low UserRiskLevelId = 1
// Medium
const User_RiskLevel_Medium UserRiskLevelId = 2
// High
const User_RiskLevel_High UserRiskLevelId = 3
// Critical
const User_RiskLevel_Critical UserRiskLevelId = 4
// Other. The risk level is not mapped. See the <code>risk_level</code> attribute, which contains a data source specific value.
const User_RiskLevel_Other UserRiskLevelId = 99

// UserTypeId Values
// Type ID. The account type identifier.

// Unknown. The type is unknown.
const User_Type_Unknown UserTypeId = 0
// User. Regular user account.
const User_Type_User UserTypeId = 1
// Admin. Admin/root user account.
const User_Type_Admin UserTypeId = 2
// System. System account. For example, Windows computer accounts with a trailing dollar sign ($).
const User_Type_System UserTypeId = 3
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const User_Type_Other UserTypeId = 99

// UserUserStatusId Values
// User Status ID. The normalized identifier of the user's status.

// Unknown. The status is unknown.
const User_UserStatus_Unknown UserUserStatusId = 0
// Active. The user is active.
const User_UserStatus_Active UserUserStatusId = 1
// Pending. The user is not active, pending either user or admin action.
const User_UserStatus_Pending UserUserStatusId = 2
// Locked. The user account is locked requiring either time or intervention to unlock.
const User_UserStatus_Locked UserUserStatusId = 3
// Suspended. The user account is suspended.
const User_UserStatus_Suspended UserUserStatusId = 4
// Deprovisioned. The user account has been deprovisioned and is pending removal.
const User_UserStatus_Deprovisioned UserUserStatusId = 5
// Other. The event status is not mapped. See the <code>user_status</code> attribute, which contains a data source specific value.
const User_UserStatus_Other UserUserStatusId = 99

// WhoisDnssecStatusId Values
// DNSSEC Status ID. Describes the normalized status of DNS Security Extensions (DNSSEC) for a domain.

// Unknown. The disposition is unknown.
const Whois_DnssecStatus_Unknown WhoisDnssecStatusId = 0
// Signed. The related domain enables the signing of DNS records using DNSSEC.
const Whois_DnssecStatus_Signed WhoisDnssecStatusId = 1
// Unsigned. The related domain does not enable the signing of DNS records using DNSSEC.
const Whois_DnssecStatus_Unsigned WhoisDnssecStatusId = 2
// Other. The DNSSEC status is not mapped. See the <code>dnssec_status</code> attribute, which contains a data source specific value.
const Whois_DnssecStatus_Other WhoisDnssecStatusId = 99

