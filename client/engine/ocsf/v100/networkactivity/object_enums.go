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

// DeviceTypeId Values
// Type ID. The device type ID.

// Unknown. The type is unknown.
const Device_Type_Unknown DeviceTypeId = 0
// Server
const Device_Type_Server DeviceTypeId = 1
// Desktop
const Device_Type_Desktop DeviceTypeId = 2
// Laptop
const Device_Type_Laptop DeviceTypeId = 3
// Tablet
const Device_Type_Tablet DeviceTypeId = 4
// Mobile
const Device_Type_Mobile DeviceTypeId = 5
// Virtual
const Device_Type_Virtual DeviceTypeId = 6
// IOT
const Device_Type_IOT DeviceTypeId = 7
// Browser
const Device_Type_Browser DeviceTypeId = 8
// Other. The type is not mapped. See the <code>type</code> attribute, which may contain a data source specific value.
const Device_Type_Other DeviceTypeId = 99

// DigitalSignatureAlgorithmId Values
// Algorithm ID. The identifier of the normalized digital signature algorithm.

// Unknown
const DigitalSignature_Algorithm_Unknown DigitalSignatureAlgorithmId = 0
// DSA. Digital Signature Algorithm (DSA).
const DigitalSignature_Algorithm_DSA DigitalSignatureAlgorithmId = 1
// RSA. Rivest-Shamir-Adleman (RSA) Algorithm.
const DigitalSignature_Algorithm_RSA DigitalSignatureAlgorithmId = 2
// ECDSA. Elliptic Curve Digital Signature Algorithm.
const DigitalSignature_Algorithm_ECDSA DigitalSignatureAlgorithmId = 3
// Authenticode. Microsoft Authenticode Digital Signature Algorithm.
const DigitalSignature_Algorithm_Authenticode DigitalSignatureAlgorithmId = 4
// Other
const DigitalSignature_Algorithm_Other DigitalSignatureAlgorithmId = 99

// FileConfidentialityId Values
// Confidentiality ID. The normalized identifier of the file content confidentiality indicator.

// Unknown
const File_Confidentiality_Unknown FileConfidentialityId = 0
// NotConfidential
const File_Confidentiality_NotConfidential FileConfidentialityId = 1
// Confidential
const File_Confidentiality_Confidential FileConfidentialityId = 2
// Secret
const File_Confidentiality_Secret FileConfidentialityId = 3
// TopSecret
const File_Confidentiality_TopSecret FileConfidentialityId = 4
// Other
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
// Other. The type is not mapped. See the <code>type</code> attribute, which may contain a data source specific value.
const File_Type_Other FileTypeId = 99

// FingerprintAlgorithmId Values
// Algorithm ID. The identifier of the normalized hash algorithm, which was used to create the digital fingerprint.

// Unknown
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
// Other
const Fingerprint_Algorithm_Other FingerprintAlgorithmId = 99

// MalwareClassificationIds Values
// Classification IDs. The list of normalized identifiers of the malware classifications. Reference: <a target='_blank' href='https://docs.oasis-open.org/cti/stix/v2.1/os/stix-v2.1-os.html#_oxlc4df65spl'>STIX Malware Types</a> 

// Unknown
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
// Other
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
// Other
const NetworkConnectionInfo_Boundary_Other NetworkConnectionInfoBoundaryId = 99

// NetworkConnectionInfoDirectionId Values
// Direction ID. The normalized identifier of the direction of the initiated connection, traffic, or email.

// Unknown. Connection direction is unknown.
const NetworkConnectionInfo_Direction_Unknown NetworkConnectionInfoDirectionId = 0
// Inbound. Inbound network connection. The connection was originated from the Internet or outside network, destined for services on the inside network.
const NetworkConnectionInfo_Direction_Inbound NetworkConnectionInfoDirectionId = 1
// Outbound. Outbound network connection. The connection was originated from inside the network, destined for services on the Internet or outside network.
const NetworkConnectionInfo_Direction_Outbound NetworkConnectionInfoDirectionId = 2
// Lateral. Lateral network connection. The connection was originated from inside the network, destined for services on the inside network.
const NetworkConnectionInfo_Direction_Lateral NetworkConnectionInfoDirectionId = 3
// Other
const NetworkConnectionInfo_Direction_Other NetworkConnectionInfoDirectionId = 99

// NetworkConnectionInfoProtocolVerId Values
// IP Version ID. The Internet Protocol version identifier.

// Unknown
const NetworkConnectionInfo_ProtocolVer_Unknown NetworkConnectionInfoProtocolVerId = 0
// InternetProtocolversion4_IPv4_
const NetworkConnectionInfo_ProtocolVer_InternetProtocolversion4_IPv4_ NetworkConnectionInfoProtocolVerId = 4
// InternetProtocolversion6_IPv6_
const NetworkConnectionInfo_ProtocolVer_InternetProtocolversion6_IPv6_ NetworkConnectionInfoProtocolVerId = 6
// Other
const NetworkConnectionInfo_ProtocolVer_Other NetworkConnectionInfoProtocolVerId = 99

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
// Other. The type is not mapped. See the <code>type</code> attribute, which may contain a data source specific value.
const NetworkInterface_Type_Other NetworkInterfaceTypeId = 99

// ObservableTypeId Values
// Type ID. The observable value type identifier.

// Unknown. Unknown observable data type.
const Observable_Type_Unknown ObservableTypeId = 0
// Hostname. Unique name assigned to a device connected to a computer network. A domain name in general is an Internet address that can be resolved through the Domain Name System (DNS). For example: <code>r2-d2.example.com</code>.
const Observable_Type_Hostname ObservableTypeId = 1
// IPAddress. Internet Protocol address (IP address), in either IPv4 or IPv6 format.
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
// FileHash. File hash. A unique value that corresponds to the content of the file.
const Observable_Type_FileHash ObservableTypeId = 8
// ProcessName. Process name. For example: <code>Notepad</code>.
const Observable_Type_ProcessName ObservableTypeId = 9
// ResourceUID. Resource unique identifier. For example, S3 Bucket name or EC2 Instance ID.
const Observable_Type_ResourceUID ObservableTypeId = 10
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
// RegistryKey. The registry key object describes a Windows registry key. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:WindowsRegistryKey/'>d3f:WindowsRegistryKey</a>.
const Observable_Type_RegistryKey ObservableTypeId = 28
// RegistryValue. The registry value object describes a Windows registry value.
const Observable_Type_RegistryValue ObservableTypeId = 29
// Fingerprint. The Fingerprint object provides detailed information about a digital fingerprint, which is a compact representation of data used to identify a longer piece of information, such as a public key or file content. It contains the algorithm and value of the fingerprint, enabling efficient and reliable identification of the associated data.
const Observable_Type_Fingerprint ObservableTypeId = 30
// Other. The observable data type is not mapped. See the <code>type</code> attribute, which may contain data source specific value.
const Observable_Type_Other ObservableTypeId = 99

// OsTypeId Values
// Type ID. The type identifier of the operating system.

// Unknown. The type is unknown.
const Os_Type_Unknown OsTypeId = 0
// Other. The type is not mapped. See the <code>type</code> attribute, which may contain a data source specific value.
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

// ProcessIntegrityId Values
// Integrity Level. The normalized identifier of the process integrity level (Windows only).

// Unknown
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
// Other
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
// Other. The type is not mapped. See the <code>type</code> attribute, which may contain a data source specific value.
const TlsExtension_Type_Other TlsExtensionTypeId = 99

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
// Other. The type is not mapped. See the <code>type</code> attribute, which may contain a data source specific value.
const User_Type_Other UserTypeId = 99

