package detectionfinding

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

// AdvisoryInstallStateId Values
// Install State ID. The normalized install state ID of the Advisory.

// Unknown. The normalized install state is unknown.
const Advisory_InstallState_Unknown AdvisoryInstallStateId = 0
// Installed. The item is installed.
const Advisory_InstallState_Installed AdvisoryInstallStateId = 1
// NotInstalled. The item is not installed.
const Advisory_InstallState_NotInstalled AdvisoryInstallStateId = 2
// InstalledPendingReboot. The item is installed pending reboot operation.
const Advisory_InstallState_InstalledPendingReboot AdvisoryInstallStateId = 3
// Other. The install state is not mapped. See the <code>install_state</code> attribute, which contains a data source specific value.
const Advisory_InstallState_Other AdvisoryInstallStateId = 99

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

// AnalyticStateId Values
// State ID. The Analytic state identifier.

// Unknown. The state is unknown.
const Analytic_State_Unknown AnalyticStateId = 0
// Active. The Analytic is active.
const Analytic_State_Active AnalyticStateId = 1
// Suppressed. The Analytic is suppressed. For example, a user/customer has suppressed a particular detection signature in a security product.
const Analytic_State_Suppressed AnalyticStateId = 2
// Experimental. The Analytic is still under development and considered experimental.
const Analytic_State_Experimental AnalyticStateId = 3
// Other. The state is not mapped. See the <code>state</code> attribute, which contains a data source specific value.
const Analytic_State_Other AnalyticStateId = 99

// AnalyticTypeId Values
// Type ID. The analytic type ID.

// Unknown. The type is unknown.
const Analytic_Type_Unknown AnalyticTypeId = 0
// Rule. A Rule in security analytics refers to predefined criteria or conditions set to monitor, alert, or enforce policies, playing a crucial role in access control, threat detection, and regulatory compliance across security systems.
const Analytic_Type_Rule AnalyticTypeId = 1
// Behavioral. Behavioral analytics focus on monitoring and analyzing user or system actions to identify deviations from established patterns, aiding in the detection of insider threats, fraud, and advanced persistent threats (APTs).
const Analytic_Type_Behavioral AnalyticTypeId = 2
// Statistical. Statistical analytics pertains to analyzing data patterns and anomalies using statistical models to predict, detect, and respond to potential threats, enhancing overall security posture through informed decision-making.
const Analytic_Type_Statistical AnalyticTypeId = 3
// Learning_ML_DL_. Learning (ML/DL) encompasses techniques that can "learn" from known data to create analytics that generalize to new data. There may be a statistical component to these techniques, but it is not a requirement.
const Analytic_Type_Learning_ML_DL_ AnalyticTypeId = 4
// Fingerprinting. Fingerprinting is the technique of collecting detailed system data, including software versions and configurations, to enhance threat detection, data loss prevention (DLP), and endpoint detection and response (EDR) capabilities.
const Analytic_Type_Fingerprinting AnalyticTypeId = 5
// Tagging. Tagging refers to the practice of assigning labels or identifiers to data, users, assets, or activities to monitor, control access, and facilitate incident response across various security domains such as DLP and EDR.
const Analytic_Type_Tagging AnalyticTypeId = 6
// KeywordMatch. Keyword Match involves scanning content for specific terms to identify sensitive information, potential threats, or policy violations, aiding in DLP and compliance monitoring.
const Analytic_Type_KeywordMatch AnalyticTypeId = 7
// RegularExpressions. Regular Expressions are used to define complex search patterns for identifying, validating, and extracting specific data sets or threats within digital content, enhancing DLP, EDR, and threat detection mechanisms.
const Analytic_Type_RegularExpressions AnalyticTypeId = 8
// ExactDataMatch. Exact Data Match is a precise comparison technique used to detect the unauthorized use or exposure of specific, sensitive information, crucial for enforcing DLP policies and protecting against data breaches.
const Analytic_Type_ExactDataMatch AnalyticTypeId = 9
// PartialDataMatch. Partial Data Match involves identifying instances where segments of sensitive information or patterns match, facilitating nuanced DLP and threat detection without requiring complete data conformity.
const Analytic_Type_PartialDataMatch AnalyticTypeId = 10
// IndexedDataMatch. Indexed Data Match refers to comparing content against a pre-compiled index of sensitive information to efficiently detect and prevent unauthorized access or breaches, streamlining DLP and compliance efforts.
const Analytic_Type_IndexedDataMatch AnalyticTypeId = 11
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const Analytic_Type_Other AnalyticTypeId = 99

// AuthFactorFactorTypeId Values
// Factor Type ID. The normalized identifier for the authentication factor.

// Unknown
const AuthFactor_FactorType_Unknown AuthFactorFactorTypeId = 0
// SMS. User receives and inputs a code sent to their mobile device via SMS text message.
const AuthFactor_FactorType_SMS AuthFactorFactorTypeId = 1
// SecurityQuestion. The user responds to a security question as part of a question-based authentication factor
const AuthFactor_FactorType_SecurityQuestion AuthFactorFactorTypeId = 2
// PhoneCall. System calls the user's registered phone number and requires the user to answer and provide a response.
const AuthFactor_FactorType_PhoneCall AuthFactorFactorTypeId = 3
// Biometric. Devices that verify identity-based on user's physical identifiers, such as fingerprint scanners or retina scanners.
const AuthFactor_FactorType_Biometric AuthFactorFactorTypeId = 4
// PushNotification. Push notification is sent to user's registered device and requires the user to acknowledge.
const AuthFactor_FactorType_PushNotification AuthFactorFactorTypeId = 5
// HardwareToken. Physical device that generates a code to be used for authentication.
const AuthFactor_FactorType_HardwareToken AuthFactorFactorTypeId = 6
// OTP. Application generates a one-time password (OTP) for use in authentication.
const AuthFactor_FactorType_OTP AuthFactorFactorTypeId = 7
// Email. A code or link is sent to a user's registered email address.
const AuthFactor_FactorType_Email AuthFactorFactorTypeId = 8
// U2F. Typically involves a hardware token, which the user physically interacts with to authenticate.
const AuthFactor_FactorType_U2F AuthFactorFactorTypeId = 9
// WebAuthn. Web-based API that enables users to register devices as authentication factors.
const AuthFactor_FactorType_WebAuthn AuthFactorFactorTypeId = 10
// Password. The user enters a password that they have previously established.
const AuthFactor_FactorType_Password AuthFactorFactorTypeId = 11
// Other
const AuthFactor_FactorType_Other AuthFactorFactorTypeId = 99

// CvssDepth Values
// CVSS Depth. The CVSS depth represents a depth of the equation used to calculate CVSS score.

// Base
const Cvss_Depth_Base CvssDepth = "Base"
// Environmental
const Cvss_Depth_Environmental CvssDepth = "Environmental"
// Temporal
const Cvss_Depth_Temporal CvssDepth = "Temporal"

// DatabaseTypeId Values
// Type ID. The normalized identifier of the database type.

// Unknown. The type is unknown.
const Database_Type_Unknown DatabaseTypeId = 0
// Relational
const Database_Type_Relational DatabaseTypeId = 1
// Network
const Database_Type_Network DatabaseTypeId = 2
// ObjectOriented
const Database_Type_ObjectOriented DatabaseTypeId = 3
// Centralized
const Database_Type_Centralized DatabaseTypeId = 4
// Operational
const Database_Type_Operational DatabaseTypeId = 5
// NoSQL
const Database_Type_NoSQL DatabaseTypeId = 6
// Vector
const Database_Type_Vector DatabaseTypeId = 7
// KnowledgeGraph
const Database_Type_KnowledgeGraph DatabaseTypeId = 8
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const Database_Type_Other DatabaseTypeId = 99

// DatabucketTypeId Values
// Type ID. The normalized identifier of the databucket type.

// Unknown. The type is unknown.
const Databucket_Type_Unknown DatabucketTypeId = 0
// S3
const Databucket_Type_S3 DatabucketTypeId = 1
// AzureBlob
const Databucket_Type_AzureBlob DatabucketTypeId = 2
// GCPBucket
const Databucket_Type_GCPBucket DatabucketTypeId = 3
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const Databucket_Type_Other DatabucketTypeId = 99

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
// Expired. The digital signature is invalid because its timestamp does not fall within the certificate's validity period.
const DigitalSignature_State_Expired DigitalSignatureStateId = 2
// Revoked. The digital signature is invalid due to certificate revocation.
const DigitalSignature_State_Revoked DigitalSignatureStateId = 3
// Suspended. The digital signature is invalid due to certificate suspension.
const DigitalSignature_State_Suspended DigitalSignatureStateId = 4
// Pending. The digital signature state is pending.
const DigitalSignature_State_Pending DigitalSignatureStateId = 5
// Untrusted. The digital signature is invalid because the certificate is rooted in an untrusted CA or is an untrusted self-signed certificate.
const DigitalSignature_State_Untrusted DigitalSignatureStateId = 6
// Distrusted. The digital signature is invalid because the certificate is explicitly distrusted. Note that whereas revocation is global, distrust reflects local IT/security policy.
const DigitalSignature_State_Distrusted DigitalSignatureStateId = 7
// WrongUsage. The digital signature is invalid because the certificate is not intended for code signing purposes.
const DigitalSignature_State_WrongUsage DigitalSignatureStateId = 8
// Bad. The digital signature is cryptographically invalid, e.g. a mismatched digest. This indicates possible tampering.
const DigitalSignature_State_Bad DigitalSignatureStateId = 9
// Broken. The digital signature is malformed and could not be processed.
const DigitalSignature_State_Broken DigitalSignatureStateId = 10
// Other. The state is not mapped. See the <code>state</code> attribute, which contains a data source specific value.
const DigitalSignature_State_Other DigitalSignatureStateId = 99

// DnsQueryOpcodeId Values
// DNS Opcode ID. The DNS opcode ID specifies the normalized query message type as defined in <a target='_blank' href='https://www.rfc-editor.org/rfc/rfc5395.html'>RFC-5395</a>.

// Query. Standard query
const DnsQuery_Opcode_Query DnsQueryOpcodeId = 0
// InverseQuery. Inverse query, obsolete
const DnsQuery_Opcode_InverseQuery DnsQueryOpcodeId = 1
// Status. Server status request
const DnsQuery_Opcode_Status DnsQueryOpcodeId = 2
// Reserved. Reserved, not used
const DnsQuery_Opcode_Reserved DnsQueryOpcodeId = 3
// Notify. Zone change notification
const DnsQuery_Opcode_Notify DnsQueryOpcodeId = 4
// Update. Dynamic DNS update
const DnsQuery_Opcode_Update DnsQueryOpcodeId = 5
// DSOMessage. DNS Stateful Operations (DSO)
const DnsQuery_Opcode_DSOMessage DnsQueryOpcodeId = 6
// Other. The DNS Opcode is not defined by the RFC. See the <code>opcode</code> attribute, which contains a data source specific value.
const DnsQuery_Opcode_Other DnsQueryOpcodeId = 99

// EncryptionDetailsAlgorithmId Values
// Encryption Algorithm ID. The encryption algorithm used.

// Unknown. The algorithm is unknown.
const EncryptionDetails_Algorithm_Unknown EncryptionDetailsAlgorithmId = 0
// DES. Data Encryption Standard Algorithm
const EncryptionDetails_Algorithm_DES EncryptionDetailsAlgorithmId = 1
// TripleDES. Triple Data Encryption Standard Algorithm
const EncryptionDetails_Algorithm_TripleDES EncryptionDetailsAlgorithmId = 2
// AES. Advanced Encryption Standard Algorithm.
const EncryptionDetails_Algorithm_AES EncryptionDetailsAlgorithmId = 3
// RSA. Rivest-Shamir-Adleman Algorithm
const EncryptionDetails_Algorithm_RSA EncryptionDetailsAlgorithmId = 4
// ECC. Elliptic Curve Cryptography Algorithm
const EncryptionDetails_Algorithm_ECC EncryptionDetailsAlgorithmId = 5
// SM2. ShangMi Cryptographic Algorithm
const EncryptionDetails_Algorithm_SM2 EncryptionDetailsAlgorithmId = 6
// Other. The algorithm is not mapped. See the <code>algorithm</code> attribute, which contains a data source specific value.
const EncryptionDetails_Algorithm_Other EncryptionDetailsAlgorithmId = 99

// EvidencesVerdictId Values
// Verdict ID. The normalized verdict (or status) ID of the evidence associated with the security detection. For example, Microsoft Graph Security Alerts contain a <code>verdict</code> enumeration for each type of <code>evidence</code> associated with the Alert. This is typically set by an automated investigation process or an analyst/investigator assigned to the finding.

// Unknown. The type is unknown.
const Evidences_Verdict_Unknown EvidencesVerdictId = 0
// FalsePositive. The verdict for the evidence has been identified as a False Positive.
const Evidences_Verdict_FalsePositive EvidencesVerdictId = 1
// TruePositive. The verdict for the evidence has been identified as a True Positive.
const Evidences_Verdict_TruePositive EvidencesVerdictId = 2
// Disregard. The verdict for the evidence is that is should be Disregarded.
const Evidences_Verdict_Disregard EvidencesVerdictId = 3
// Suspicious. The verdict for the evidence is that the behavior has been identified as Suspicious.
const Evidences_Verdict_Suspicious EvidencesVerdictId = 4
// Benign. The verdict for the evidence is that the behavior has been identified as Benign.
const Evidences_Verdict_Benign EvidencesVerdictId = 5
// Test. The evidence is part of a Test, or other sanctioned behavior(s).
const Evidences_Verdict_Test EvidencesVerdictId = 6
// InsufficientData. There is insufficient data to render a verdict on the evidence.
const Evidences_Verdict_InsufficientData EvidencesVerdictId = 7
// SecurityRisk. The verdict for the evidence is that the behavior has been identified as a Security Risk.
const Evidences_Verdict_SecurityRisk EvidencesVerdictId = 8
// ManagedExternally. The verdict for the evidence is Managed Externally, such as in a case management tool.
const Evidences_Verdict_ManagedExternally EvidencesVerdictId = 9
// Duplicate. This evidence duplicates existing evidence related to this finding.
const Evidences_Verdict_Duplicate EvidencesVerdictId = 10
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const Evidences_Verdict_Other EvidencesVerdictId = 99

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

// FileDriveTypeId Values
// Drive Type ID. Identifies the type of a disk drive, i.e. fixed, removable, etc.

// Unknown. The drive type is unknown.
const File_DriveType_Unknown FileDriveTypeId = 0
// Removable. The drive has removable media; for example, a floppy drive, thumb drive, or flash card reader.
const File_DriveType_Removable FileDriveTypeId = 1
// Fixed. The drive has fixed media; for example, a hard disk drive or flash drive.
const File_DriveType_Fixed FileDriveTypeId = 2
// Remote. The drive is a remote (network) drive.
const File_DriveType_Remote FileDriveTypeId = 3
// CD_ROM. The drive is a CD-ROM drive.
const File_DriveType_CD_ROM FileDriveTypeId = 4
// RAMDisk. The drive is a RAM disk.
const File_DriveType_RAMDisk FileDriveTypeId = 5
// Other. The drive type is not mapped. See the <code>drive_type</code> attribute, which contains a data source specific value.
const File_DriveType_Other FileDriveTypeId = 99

// FileTypeId Values
// Type ID. The file type ID. Note the distinction between a <code>Regular File</code> and an <code>Executable File</code>. If the distinction is not known, or not indicated by the log, use <code>Regular File</code>. In this case, it should not be assumed that a Regular File is not executable.

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
// ExecutableFile
const File_Type_ExecutableFile FileTypeId = 8
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const File_Type_Other FileTypeId = 99

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

// GraphQueryLanguageId Values
// Query Language ID. The normalized identifier of a graph query language that can be used to interact with the graph.

// Unknown. The Query Language is unknown.
const Graph_QueryLanguage_Unknown GraphQueryLanguageId = 0
// Cypher. A declarative graph query language developed by Neo4j that allows for expressive and efficient querying of graph databases.
const Graph_QueryLanguage_Cypher GraphQueryLanguageId = 1
// GraphQL. A query language for APIs that enables declarative data fetching and provides a complete description of the data in the API.
const Graph_QueryLanguage_GraphQL GraphQueryLanguageId = 2
// Gremlin. A graph traversal language and virtual machine developed by Apache TinkerPop that enables graph computing across different graph databases.
const Graph_QueryLanguage_Gremlin GraphQueryLanguageId = 3
// GQL. An ISO standard graph query language designed to provide a unified way to query graph databases.
const Graph_QueryLanguage_GQL GraphQueryLanguageId = 4
// G_CORE. A graph query language that combines features from existing languages while adding support for paths as first-class citizens.
const Graph_QueryLanguage_G_CORE GraphQueryLanguageId = 5
// PGQL. Property Graph Query Language developed by Oracle that provides SQL-like syntax for querying property graphs.
const Graph_QueryLanguage_PGQL GraphQueryLanguageId = 6
// SPARQL. A semantic query language for databases that enables querying and manipulating data stored in RDF format.
const Graph_QueryLanguage_SPARQL GraphQueryLanguageId = 7
// Other. The Query Language is not mapped. See the <code>query_language</code> attribute, which contains a data source specific value.
const Graph_QueryLanguage_Other GraphQueryLanguageId = 99

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
// Patch. The PATCH method applies partial modifications to a resource.
const HttpRequest_HttpMethod_Patch HttpRequestHttpMethod = "PATCH"
// Post. The POST method submits an entity to the specified resource, often causing a change in state or side effects on the server.
const HttpRequest_HttpMethod_Post HttpRequestHttpMethod = "POST"
// Put. The PUT method replaces all current representations of the target resource with the request payload.
const HttpRequest_HttpMethod_Put HttpRequestHttpMethod = "PUT"
// Trace. The TRACE method performs a message loop-back test along the path to the target resource.
const HttpRequest_HttpMethod_Trace HttpRequestHttpMethod = "TRACE"

// IdpStateId Values
// State ID. The normalized state ID of the Identity Provider to reflect its configuration or activation status.

// Unknown. The configuration state of the Identity Provider is unknown.
const Idp_State_Unknown IdpStateId = 0
// Active. The Identity Provider is in an Active state, or otherwise enabled.
const Idp_State_Active IdpStateId = 1
// Suspended. The Identity Provider is in a Suspended state.
const Idp_State_Suspended IdpStateId = 2
// Deprecated. The Identity Provider is in a Deprecated state, or is otherwise disabled.
const Idp_State_Deprecated IdpStateId = 3
// Deleted. The Identity Provider is in a Deleted state.
const Idp_State_Deleted IdpStateId = 4
// Other. The configuration state of the Identity Provider is not mapped. See the <code>state</code> attribute, which contains a data source specific value.
const Idp_State_Other IdpStateId = 99

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

// JobRunStateId Values
// Run State ID. The run state ID of the job.

// Unknown. The run state is unknown.
const Job_RunState_Unknown JobRunStateId = 0
// Ready
const Job_RunState_Ready JobRunStateId = 1
// Queued
const Job_RunState_Queued JobRunStateId = 2
// Running
const Job_RunState_Running JobRunStateId = 3
// Stopped
const Job_RunState_Stopped JobRunStateId = 4
// Other. The run state is not mapped. See the <code>run_state</code> attribute, which contains a data source specific value.
const Job_RunState_Other JobRunStateId = 99

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

// MalwareScanInfoTypeId Values
// Type ID. The type id of the scan.

// Unknown. The type is unknown.
const MalwareScanInfo_Type_Unknown MalwareScanInfoTypeId = 0
// Manual. The scan was manually initiated by the user or administrator.
const MalwareScanInfo_Type_Manual MalwareScanInfoTypeId = 1
// Scheduled. The scan was started based on scheduler.
const MalwareScanInfo_Type_Scheduled MalwareScanInfoTypeId = 2
// UpdatedContent. The scan was triggered by a content update.
const MalwareScanInfo_Type_UpdatedContent MalwareScanInfoTypeId = 3
// QuarantinedItems. The scan was triggered by newly quarantined items.
const MalwareScanInfo_Type_QuarantinedItems MalwareScanInfoTypeId = 4
// AttachedMedia. The scan was triggered by the attachment of removable media.
const MalwareScanInfo_Type_AttachedMedia MalwareScanInfoTypeId = 5
// UserLogon. The scan was started due to a user logon.
const MalwareScanInfo_Type_UserLogon MalwareScanInfoTypeId = 6
// ELAM. The scan was triggered by an Early Launch Anti-Malware (ELAM) detection.
const MalwareScanInfo_Type_ELAM MalwareScanInfoTypeId = 7
// Other. The scan type id is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const MalwareScanInfo_Type_Other MalwareScanInfoTypeId = 99

// MalwareClassificationIds Values
// Classification IDs. The list of normalized identifiers of the malware classifications.

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

// MalwareSeverityId Values
// Severity ID. The normalized identifier of the malware severity.

// Unknown. The event/finding severity is unknown.
const Malware_Severity_Unknown MalwareSeverityId = 0
// Informational. Informational message. No action required.
const Malware_Severity_Informational MalwareSeverityId = 1
// Low. The user decides if action is needed.
const Malware_Severity_Low MalwareSeverityId = 2
// Medium. Action is required but the situation is not serious at this time.
const Malware_Severity_Medium MalwareSeverityId = 3
// High. Action is required immediately.
const Malware_Severity_High MalwareSeverityId = 4
// Critical. Action is required immediately and the scope is broad.
const Malware_Severity_Critical MalwareSeverityId = 5
// Fatal. An error occurred but it is too late to take remedial action.
const Malware_Severity_Fatal MalwareSeverityId = 6
// Other. The event/finding severity is not mapped. See the <code>severity</code> attribute, which contains a data source specific value.
const Malware_Severity_Other MalwareSeverityId = 99

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
// Inbound. Inbound network connection. The connection originated from the Internet or outside network, destined for services on the inside network.
const NetworkConnectionInfo_Direction_Inbound NetworkConnectionInfoDirectionId = 1
// Outbound. Outbound network connection. The connection originated from inside the network, destined for services on the Internet or outside network.
const NetworkConnectionInfo_Direction_Outbound NetworkConnectionInfoDirectionId = 2
// Lateral. Lateral network connection. The connection originated from inside the network, destined for services on the inside network.
const NetworkConnectionInfo_Direction_Lateral NetworkConnectionInfoDirectionId = 3
// Local. Local network connection (<code>localhost</code>). The connection is intra-device, originating from and destined for services running on the same device.
const NetworkConnectionInfo_Direction_Local NetworkConnectionInfoDirectionId = 4
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

// NetworkEndpointNetworkScopeId Values
// Network Scope ID. The normalized identifier of the endpoint’s network scope. The normalized network scope identifier indicates whether the endpoint resides inside the customer’s network, outside on the Internet, or if its location relative to the customer’s network cannot be determined.

// Unknown. Unknown whether this endpoint resides within the customer’s network.
const NetworkEndpoint_NetworkScope_Unknown NetworkEndpointNetworkScopeId = 0
// Internal. The endpoint resides inside the customer’s network.
const NetworkEndpoint_NetworkScope_Internal NetworkEndpointNetworkScopeId = 1
// External. The endpoint is on the Internet or otherwise external to the customer’s network.
const NetworkEndpoint_NetworkScope_External NetworkEndpointNetworkScopeId = 2
// Other. The network scope is not mapped. See the <code>network_scope</code> attribute, which contains a data source specific value.
const NetworkEndpoint_NetworkScope_Other NetworkEndpointNetworkScopeId = 99

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
// LoadBalancer. A <a target='_blank' href='https://en.wikipedia.org/wiki/Load_balancing_(computing)'>Load Balancer device.</a>
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

// NetworkProxyNetworkScopeId Values
// Network Scope ID. The normalized identifier of the endpoint’s network scope. The normalized network scope identifier indicates whether the endpoint resides inside the customer’s network, outside on the Internet, or if its location relative to the customer’s network cannot be determined.

// Unknown. Unknown whether this endpoint resides within the customer’s network.
const NetworkProxy_NetworkScope_Unknown NetworkProxyNetworkScopeId = 0
// Internal. The endpoint resides inside the customer’s network.
const NetworkProxy_NetworkScope_Internal NetworkProxyNetworkScopeId = 1
// External. The endpoint is on the Internet or otherwise external to the customer’s network.
const NetworkProxy_NetworkScope_External NetworkProxyNetworkScopeId = 2
// Other. The network scope is not mapped. See the <code>network_scope</code> attribute, which contains a data source specific value.
const NetworkProxy_NetworkScope_Other NetworkProxyNetworkScopeId = 99

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
// LoadBalancer. A <a target='_blank' href='https://en.wikipedia.org/wiki/Load_balancing_(computing)'>Load Balancer device.</a>
const NetworkProxy_Type_LoadBalancer NetworkProxyTypeId = 15
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const NetworkProxy_Type_Other NetworkProxyTypeId = 99

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

// RelatedEventSeverityId Values
// Severity ID. <p>The normalized identifier of the event/finding severity.</p>The normalized severity is a measurement the effort and expense required to manage and resolve an event or incident. Smaller numerical values represent lower impact events, and larger numerical values represent higher impact events.

// Unknown. The event/finding severity is unknown.
const RelatedEvent_Severity_Unknown RelatedEventSeverityId = 0
// Informational. Informational message. No action required.
const RelatedEvent_Severity_Informational RelatedEventSeverityId = 1
// Low. The user decides if action is needed.
const RelatedEvent_Severity_Low RelatedEventSeverityId = 2
// Medium. Action is required but the situation is not serious at this time.
const RelatedEvent_Severity_Medium RelatedEventSeverityId = 3
// High. Action is required immediately.
const RelatedEvent_Severity_High RelatedEventSeverityId = 4
// Critical. Action is required immediately and the scope is broad.
const RelatedEvent_Severity_Critical RelatedEventSeverityId = 5
// Fatal. An error occurred but it is too late to take remedial action.
const RelatedEvent_Severity_Fatal RelatedEventSeverityId = 6
// Other. The event/finding severity is not mapped. See the <code>severity</code> attribute, which contains a data source specific value.
const RelatedEvent_Severity_Other RelatedEventSeverityId = 99

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

// ResourceDetailsRoleId Values
// Role ID. The normalized identifier of the resource's role in the context of the event or finding.

// Unknown. The role is unknown.
const ResourceDetails_Role_Unknown ResourceDetailsRoleId = 0
// Target. The resource is the primary target or subject of the event/finding.
const ResourceDetails_Role_Target ResourceDetailsRoleId = 1
// Actor. The resource is acting as the initiator or performer in the context of the event/finding.
const ResourceDetails_Role_Actor ResourceDetailsRoleId = 2
// Affected. The resource was impacted or affected by the event/finding.
const ResourceDetails_Role_Affected ResourceDetailsRoleId = 3
// Related. The resource is related to or associated with the event/finding.
const ResourceDetails_Role_Related ResourceDetailsRoleId = 4
// Other. The role is not mapped. See the role attribute, which contains a data source specific value.
const ResourceDetails_Role_Other ResourceDetailsRoleId = 99

// ScimAuthProtocolId Values
// Auth Protocol ID. The normalized identifier of the authorization protocol used by the SCIM resource.

// Unknown. The authentication protocol is unknown.
const Scim_AuthProtocol_Unknown ScimAuthProtocolId = 0
// NTLM
const Scim_AuthProtocol_NTLM ScimAuthProtocolId = 1
// Kerberos
const Scim_AuthProtocol_Kerberos ScimAuthProtocolId = 2
// Digest
const Scim_AuthProtocol_Digest ScimAuthProtocolId = 3
// OpenID
const Scim_AuthProtocol_OpenID ScimAuthProtocolId = 4
// SAML
const Scim_AuthProtocol_SAML ScimAuthProtocolId = 5
// OAUTH2_0
const Scim_AuthProtocol_OAUTH2_0 ScimAuthProtocolId = 6
// PAP
const Scim_AuthProtocol_PAP ScimAuthProtocolId = 7
// CHAP
const Scim_AuthProtocol_CHAP ScimAuthProtocolId = 8
// EAP
const Scim_AuthProtocol_EAP ScimAuthProtocolId = 9
// RADIUS
const Scim_AuthProtocol_RADIUS ScimAuthProtocolId = 10
// BasicAuthentication
const Scim_AuthProtocol_BasicAuthentication ScimAuthProtocolId = 11
// LDAP
const Scim_AuthProtocol_LDAP ScimAuthProtocolId = 12
// Other. The authentication protocol is not mapped. See the <code>auth_protocol</code> attribute, which contains a data source specific value.
const Scim_AuthProtocol_Other ScimAuthProtocolId = 99

// ScimStateId Values
// State ID. The normalized state ID of the SCIM resource to reflect its activation status.

// Unknown. The provisioning state of the SCIM resource is unknown.
const Scim_State_Unknown ScimStateId = 0
// Pending. The SCIM resource is Pending activation or creation.
const Scim_State_Pending ScimStateId = 1
// Active. The SCIM resource is in an Active state, or otherwise enabled.
const Scim_State_Active ScimStateId = 2
// Failed. The SCIM resource is in a Failed state.
const Scim_State_Failed ScimStateId = 3
// Deleted. The SCIM resource is in a Deleted state, or otherwise disabled.
const Scim_State_Deleted ScimStateId = 4
// Other. The provisioning state of the SCIM resource is not mapped. See the <code>state</code> attribute, which contains a data source specific value.
const Scim_State_Other ScimStateId = 99

// ScriptTypeId Values
// Type ID. The normalized script type ID.

// Unknown. The script type is unknown.
const Script_Type_Unknown ScriptTypeId = 0
// WindowsCommandPrompt
const Script_Type_WindowsCommandPrompt ScriptTypeId = 1
// PowerShell
const Script_Type_PowerShell ScriptTypeId = 2
// Python
const Script_Type_Python ScriptTypeId = 3
// JavaScript
const Script_Type_JavaScript ScriptTypeId = 4
// VBScript
const Script_Type_VBScript ScriptTypeId = 5
// UnixShell
const Script_Type_UnixShell ScriptTypeId = 6
// VBA
const Script_Type_VBA ScriptTypeId = 7
// Other. The script type is not mapped. See the <code>type</code> attribute which contains an event source specific value.
const Script_Type_Other ScriptTypeId = 99

// SsoAuthProtocolId Values
// Auth Protocol ID. The normalized identifier of the authentication protocol used by the SSO resource.

// Unknown. The authentication protocol is unknown.
const Sso_AuthProtocol_Unknown SsoAuthProtocolId = 0
// NTLM
const Sso_AuthProtocol_NTLM SsoAuthProtocolId = 1
// Kerberos
const Sso_AuthProtocol_Kerberos SsoAuthProtocolId = 2
// Digest
const Sso_AuthProtocol_Digest SsoAuthProtocolId = 3
// OpenID
const Sso_AuthProtocol_OpenID SsoAuthProtocolId = 4
// SAML
const Sso_AuthProtocol_SAML SsoAuthProtocolId = 5
// OAUTH2_0
const Sso_AuthProtocol_OAUTH2_0 SsoAuthProtocolId = 6
// PAP
const Sso_AuthProtocol_PAP SsoAuthProtocolId = 7
// CHAP
const Sso_AuthProtocol_CHAP SsoAuthProtocolId = 8
// EAP
const Sso_AuthProtocol_EAP SsoAuthProtocolId = 9
// RADIUS
const Sso_AuthProtocol_RADIUS SsoAuthProtocolId = 10
// BasicAuthentication
const Sso_AuthProtocol_BasicAuthentication SsoAuthProtocolId = 11
// LDAP
const Sso_AuthProtocol_LDAP SsoAuthProtocolId = 12
// Other. The authentication protocol is not mapped. See the <code>auth_protocol</code> attribute, which contains a data source specific value.
const Sso_AuthProtocol_Other SsoAuthProtocolId = 99

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

// TokenTypeId Values
// Type ID. The normalized token type identifier. Valid values: 0 (Unknown), 1 (Ticket Granting Ticket - Kerberos), 2 (Service Ticket - Kerberos), 3 (Identity Token - OIDC), 4 (Refresh Token - OIDC), 5 (SAML Assertion), 6 (Client Token - IdP-issued), 7 (API Token - generic API keys), 99 (Other).

// Unknown. The token type is unknown.
const Token_Type_Unknown TokenTypeId = 0
// TicketGrantingTicket. Ticket Granting Ticket (TGT) for Kerberos.
const Token_Type_TicketGrantingTicket TokenTypeId = 1
// ServiceTicket. Service Ticket (ST) for Kerberos.
const Token_Type_ServiceTicket TokenTypeId = 2
// IdentityToken. Identity (ID) Token for OIDC.
const Token_Type_IdentityToken TokenTypeId = 3
// RefreshToken. Refresh Token for OIDC.
const Token_Type_RefreshToken TokenTypeId = 4
// SAMLAssertion. Authentication Assertion for SAML.
const Token_Type_SAMLAssertion TokenTypeId = 5
// ClientToken. Client Token issued by an Identity Provider (IdP) for application authentication. Use this value for IdP-issued tokens used for service-to-service authentication. Examples: Microsoft Entra ID Application Registration client secrets, Okta API tokens, Auth0 Machine-to-Machine tokens. Key characteristic: These tokens are issued by Identity Providers, not by individual services.
const Token_Type_ClientToken TokenTypeId = 6
// APIToken. A generic API token or API key used for authenticating API requests. Use this value for service-specific API authentication tokens that are NOT issued by Identity Providers. Examples: REST API keys, GraphQL API keys, Stripe API keys, Twilio API keys, AWS API keys. Key characteristic: These tokens are issued by individual services/platforms, not by Identity Providers.
const Token_Type_APIToken TokenTypeId = 7
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const Token_Type_Other TokenTypeId = 99

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

// VendorAttributesSeverityId Values
// Severity ID. The finding severity ID, as reported by the Vendor (Finding Provider).

// Unknown. The event/finding severity is unknown.
const VendorAttributes_Severity_Unknown VendorAttributesSeverityId = 0
// Informational. Informational message. No action required.
const VendorAttributes_Severity_Informational VendorAttributesSeverityId = 1
// Low. The user decides if action is needed.
const VendorAttributes_Severity_Low VendorAttributesSeverityId = 2
// Medium. Action is required but the situation is not serious at this time.
const VendorAttributes_Severity_Medium VendorAttributesSeverityId = 3
// High. Action is required immediately.
const VendorAttributes_Severity_High VendorAttributesSeverityId = 4
// Critical. Action is required immediately and the scope is broad.
const VendorAttributes_Severity_Critical VendorAttributesSeverityId = 5
// Fatal. An error occurred but it is too late to take remedial action.
const VendorAttributes_Severity_Fatal VendorAttributesSeverityId = 6
// Other. The event/finding severity is not mapped. See the <code>severity</code> attribute, which contains a data source specific value.
const VendorAttributes_Severity_Other VendorAttributesSeverityId = 99

// VulnerabilityFixCoverageId Values
// Fix Coverage ID. The normalized identifier for fix coverage, applicable to this vulnerability. Typically useful, when there are multiple affected packages but only a subset have available fixes.

// Unknown. The fix coverage is unknown.
const Vulnerability_FixCoverage_Unknown VulnerabilityFixCoverageId = 0
// Complete. All affected packages and components have available fixes or patches to remediate the vulnerability.
const Vulnerability_FixCoverage_Complete VulnerabilityFixCoverageId = 1
// Partial. Only some of the affected packages and components have available fixes or patches, while others remain vulnerable.
const Vulnerability_FixCoverage_Partial VulnerabilityFixCoverageId = 2
// None. No fixes or patches are currently available for any of the affected packages and components.
const Vulnerability_FixCoverage_None VulnerabilityFixCoverageId = 3
// Other. The fix coverage is not mapped. See the <code>fix_coverage</code> attribute, which contains a data source specific value.
const Vulnerability_FixCoverage_Other VulnerabilityFixCoverageId = 99

