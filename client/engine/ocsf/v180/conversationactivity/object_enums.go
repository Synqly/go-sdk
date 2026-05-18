package conversationactivity

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
// AzureADAccount. Note: The new product name for Azure AD is Microsoft Entra ID.
const Account_Type_AzureADAccount AccountTypeId = 6
// MacOSAccount
const Account_Type_MacOSAccount AccountTypeId = 7
// AppleAccount
const Account_Type_AppleAccount AccountTypeId = 8
// LinuxAccount
const Account_Type_LinuxAccount AccountTypeId = 9
// AWSAccount
const Account_Type_AWSAccount AccountTypeId = 10
// GCPProject
const Account_Type_GCPProject AccountTypeId = 11
// OCICompartment
const Account_Type_OCICompartment AccountTypeId = 12
// AzureSubscription
const Account_Type_AzureSubscription AccountTypeId = 13
// SalesforceAccount
const Account_Type_SalesforceAccount AccountTypeId = 14
// GoogleWorkspace
const Account_Type_GoogleWorkspace AccountTypeId = 15
// ServicenowInstance
const Account_Type_ServicenowInstance AccountTypeId = 16
// M365Tenant
const Account_Type_M365Tenant AccountTypeId = 17
// EmailAccount
const Account_Type_EmailAccount AccountTypeId = 18
// ActiveDirectoryAccount
const Account_Type_ActiveDirectoryAccount AccountTypeId = 19
// Other. The account type is not mapped.
const Account_Type_Other AccountTypeId = 99

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

// ConversationTypeId Values
// Type ID. The normalized conversation type identifier.

// Unknown. The type is unknown.
const Conversation_Type_Unknown ConversationTypeId = 0
// Channel. A team-based group conversation, e.g. a Teams channel or Slack channel.
const Conversation_Type_Channel ConversationTypeId = 1
// DirectMessage. A 1:1 private conversation between two users.
const Conversation_Type_DirectMessage ConversationTypeId = 2
// GroupChat. A group conversation among more than two users, without a team scope.
const Conversation_Type_GroupChat ConversationTypeId = 3
// AISession. An AI assistant session, e.g. a Microsoft 365 Copilot interaction history.
const Conversation_Type_AISession ConversationTypeId = 4
// Other. The conversation type is not mapped. See the type attribute for the data source specific value.
const Conversation_Type_Other ConversationTypeId = 99

// DeviceHwInfoCpuArchitectureId Values
// CPU Architecture ID. The normalized identifier of the CPU architecture.

// Unknown. The CPU architecture is unknown.
const DeviceHwInfo_CpuArchitecture_Unknown DeviceHwInfoCpuArchitectureId = 0
// x86. CPU uses the x86 ISA. For bitness, refer to <code>cpu_bits</code>.
const DeviceHwInfo_CpuArchitecture_x86 DeviceHwInfoCpuArchitectureId = 1
// ARM. CPU uses the ARM ISA. For bitness, refer to <code>cpu_bits</code>.
const DeviceHwInfo_CpuArchitecture_ARM DeviceHwInfoCpuArchitectureId = 2
// RISC_V. CPU uses the RISC-V ISA. For bitness, refer to <code>cpu_bits</code>.
const DeviceHwInfo_CpuArchitecture_RISC_V DeviceHwInfoCpuArchitectureId = 3
// Other. The CPU architecture is not mapped. See the <code>cpu_architecture</code> attribute, which contains a data source specific value.
const DeviceHwInfo_CpuArchitecture_Other DeviceHwInfoCpuArchitectureId = 99

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
// LoadBalancer. A <a target='_blank' href='https://en.wikipedia.org/wiki/Load_balancing_(computing)'>Load Balancer device.</a>
const Device_Type_LoadBalancer DeviceTypeId = 15
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const Device_Type_Other DeviceTypeId = 99

// FingerprintAlgorithmId Values
// Algorithm ID. The identifier of the normalized algorithm or scheme, which was used to create the fingerprint.

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
// SHA_224. Secure Hash Algorithm 2 producing a 224-bit (28-byte) hash value.
const Fingerprint_Algorithm_SHA_224 FingerprintAlgorithmId = 8
// SHA_384. Secure Hash Algorithm 2 producing a 384-bit (48-byte) hash value.
const Fingerprint_Algorithm_SHA_384 FingerprintAlgorithmId = 9
// SHA_512_224. Secure Hash Algorithm 2 producing a 512-bit (64-byte) hash value truncated to a 224-bit (28-byte) hash value.
const Fingerprint_Algorithm_SHA_512_224 FingerprintAlgorithmId = 10
// SHA_512_256. Secure Hash Algorithm 2 producing a 512-bit (64-byte) hash value truncated to a 256-bit (32-byte) hash value.
const Fingerprint_Algorithm_SHA_512_256 FingerprintAlgorithmId = 11
// SHA3_224. Secure Hash Algorithm 3 producing a 224-bit (28-byte) hash value.
const Fingerprint_Algorithm_SHA3_224 FingerprintAlgorithmId = 12
// SHA3_256. Secure Hash Algorithm 3 producing a 256-bit (32-byte) hash value.
const Fingerprint_Algorithm_SHA3_256 FingerprintAlgorithmId = 13
// SHA3_384. Secure Hash Algorithm 3 producing a 384-bit (48-byte) hash value.
const Fingerprint_Algorithm_SHA3_384 FingerprintAlgorithmId = 14
// SHA3_512. Secure Hash Algorithm 3 producing a 512-bit (64-byte) hash value.
const Fingerprint_Algorithm_SHA3_512 FingerprintAlgorithmId = 15
// xxHashH364_bit. xxHash H3 producing a 64-bit hash value.
const Fingerprint_Algorithm_xxHashH364_bit FingerprintAlgorithmId = 16
// xxHashH3128_bit. xxHash H3 producing a 128-bit hash value.
const Fingerprint_Algorithm_xxHashH3128_bit FingerprintAlgorithmId = 17
// Imphash. Import hash (imphash) based on the import table of a Portable Executable (PE) file producing a 128-bit (16-byte) hash value.
const Fingerprint_Algorithm_Imphash FingerprintAlgorithmId = 18
// NPF. Network Protocol Fingerprint (NPF) used to identify network protocols and applications.
const Fingerprint_Algorithm_NPF FingerprintAlgorithmId = 19
// HASSH. HASSH is a network fingerprinting standard which can be used to identify specific SSH client and server implementations.
const Fingerprint_Algorithm_HASSH FingerprintAlgorithmId = 20
// Other. The algorithm is not mapped. See the <code>algorithm</code> attribute, which contains a data source specific value.
const Fingerprint_Algorithm_Other FingerprintAlgorithmId = 99

// GpuInfoBusTypeId Values
// Bus Type ID. The normalized identifier of the attachment bus or interface standard.

// Unknown. The bus type is unknown or not reported.
const GpuInfo_BusType_Unknown GpuInfoBusTypeId = 0
// Onboard. The device is attached directly on the motherboard or SoC (often referred to as 'integrated' or 'onboard').
const GpuInfo_BusType_Onboard GpuInfoBusTypeId = 1
// PCIex16. Peripheral Component Interconnect Express slot with 16 lanes, typically used for high‑bandwidth add‑in devices.
const GpuInfo_BusType_PCIex16 GpuInfoBusTypeId = 2
// PCIex8. Peripheral Component Interconnect Express slot with 8 lanes; common for high‑performance NICs, storage controllers, and accelerators.
const GpuInfo_BusType_PCIex8 GpuInfoBusTypeId = 3
// MXMTypeA. Mobile PCI Express Module (MXM) Type A form factor (compact laptop/embedded GPU module).
const GpuInfo_BusType_MXMTypeA GpuInfoBusTypeId = 4
// MXMTypeB. Mobile PCI Express Module (MXM) Type B form factor (larger laptop/embedded GPU module).
const GpuInfo_BusType_MXMTypeB GpuInfoBusTypeId = 5
// M_2. M.2 (NGFF) internal module form factor/connector exposing PCIe (and/or SATA/USB) for SSDs and add‑in modules.
const GpuInfo_BusType_M_2 GpuInfoBusTypeId = 6
// CXL. Compute Express Link (CXL), an open cache‑coherent interconnect for processors, memory expansion, and accelerators built on the PCIe physical layer.
const GpuInfo_BusType_CXL GpuInfoBusTypeId = 7
// Other. A bus or interface standard not covered by the defined values; the exact value is reported by the event source.
const GpuInfo_BusType_Other GpuInfoBusTypeId = 99

// GpuInfoVramModeId Values
// VRAM Mode ID. The normalized identifier of the video memory attachment mode.

// Unknown. The VRAM mode is unknown or not reported.
const GpuInfo_VramMode_Unknown GpuInfoVramModeId = 0
// Shared. Video memory is allocated from system memory.
const GpuInfo_VramMode_Shared GpuInfoVramModeId = 1
// Dedicated. Video memory is physically separate from system memory.
const GpuInfo_VramMode_Dedicated GpuInfoVramModeId = 2
// Other. A VRAM mode not covered by the defined values; the exact value is reported by the event source.
const GpuInfo_VramMode_Other GpuInfoVramModeId = 99

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

// ObservableTypeId Values
// Type ID. The observable value type identifier.

// Unknown. Unknown observable data type.
const Observable_Type_Unknown ObservableTypeId = 0
// Hostname. Unique name assigned to a device connected to a computer network. It may be a fully qualified domain name (FQDN). For example:<br><code>r2-d2.example.com.</code>,<br><code>mx.example.com</code>
const Observable_Type_Hostname ObservableTypeId = 1
// IPAddress. Internet Protocol address (IP address), in either IPv4 or IPv6 format. For example:<br><code>192.168.200.24</code>, <br> <code>2001:0db8:85a3:0000:0000:8a2e:0370:7334</code>.
const Observable_Type_IPAddress ObservableTypeId = 2
// MACAddress. Media Access Control (MAC) address. For example:<br><code>18:36:F3:98:4F:9A</code>.
const Observable_Type_MACAddress ObservableTypeId = 3
// UserName. User name. For example:<br><code>john_doe</code>.
const Observable_Type_UserName ObservableTypeId = 4
// EmailAddress. Email address. For example:<br><code>john_doe@example.com</code>.
const Observable_Type_EmailAddress ObservableTypeId = 5
// URLString. Uniform Resource Locator (URL) string. For example:<br><code>http://www.example.com/download/trouble.exe</code>.
const Observable_Type_URLString ObservableTypeId = 6
// FileName. File name. For example:<br><code>text-file.txt</code>.
const Observable_Type_FileName ObservableTypeId = 7
// Hash. Fingerprint. A value, in any format, that maps an arbitrarily large data item to a much shorter string that uniquely identifies the original data. Examples include cryptographic hashing of a file, code signing, and Network Protocol Fingerprinting (NPF).<p><b>Note about name.</b> The type name <code>file_hast_t</code> and the caption &quot;Hash&quot; are used for legacy reasons. This type has been generalized from a file hash to a general fingerprint. The existing type name and caption were retained for backwards compatibility.
const Observable_Type_Hash ObservableTypeId = 8
// ProcessName. Process name. For example:<br><code>Notepad</code>.
const Observable_Type_ProcessName ObservableTypeId = 9
// ResourceUID. Resource unique identifier. For example, S3 Bucket name or EC2 Instance ID.
const Observable_Type_ResourceUID ObservableTypeId = 10
// Port. The TCP/UDP port number. For example:<br><code>80</code>,<br><code>22</code>.
const Observable_Type_Port ObservableTypeId = 11
// Subnet. The subnet represented in a CIDR notation, using the format network_address/prefix_length. The network_address can be in either IPv4 or IPv6 format. The prefix length indicates the number of bits used for the network portion, and the remaining bits are available for host addresses within that subnet. For example:<br><code>192.168.1.0/24</code>,<br><code>2001:0db8:85a3:0000::/64</code>
const Observable_Type_Subnet ObservableTypeId = 12
// Endpoint. The Endpoint object describes a physical or virtual device that connects to and exchanges information with a computer network. Some examples of endpoints are mobile devices, desktop computers, virtual machines, embedded devices, and servers. Internet-of-Things devices—like cameras, lighting, refrigerators, security systems, smart speakers, and thermostats—are also endpoints.
const Observable_Type_Endpoint ObservableTypeId = 20
// User. The User object describes the characteristics of a user/person or a security principal.
const Observable_Type_User ObservableTypeId = 21
// Email. The Email object describes the email metadata such as sender, recipients, and direction, and can include embedded URLs and files.
const Observable_Type_Email ObservableTypeId = 22
// UniformResourceLocator. The Uniform Resource Locator (URL) object describes the characteristics of a URL.
const Observable_Type_UniformResourceLocator ObservableTypeId = 23
// File. The File object represents the metadata associated with a file stored in a computer system. It encompasses information about the file itself, including its attributes, properties, and organizational details.
const Observable_Type_File ObservableTypeId = 24
// Process. The Process object describes a running instance of a launched program.
const Observable_Type_Process ObservableTypeId = 25
// GeoLocation. The Geo Location object describes a geographical location, usually associated with an IP address.
const Observable_Type_GeoLocation ObservableTypeId = 26
// Container. The Container object describes an instance of a specific container. A container is a prepackaged, portable system image that runs isolated on an existing system using a container runtime like containerd.
const Observable_Type_Container ObservableTypeId = 27
// Fingerprint. The Fingerprint object provides detailed information about a fingerprint, which is a compact representation of data used to identify a longer piece of information, such as a public key, file content, or application implementation. It contains the algorithm or scheme and value of the fingerprint, enabling efficient and reliable identification of the associated data.
const Observable_Type_Fingerprint ObservableTypeId = 30
// FilePath. The full path to the file. For example: For example:<br><code>c:\windows\system32\svchost.exe</code>.
const Observable_Type_FilePath ObservableTypeId = 45
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
// TimeInterval. The <code>start_time</code> and <code>end_time</code> should be set.
const Timespan_Type_TimeInterval TimespanTypeId = 9
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const Timespan_Type_Other TimespanTypeId = 99

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
// Service. Service account. For example, Windows service account.
const User_Type_Service UserTypeId = 4
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const User_Type_Other UserTypeId = 99

