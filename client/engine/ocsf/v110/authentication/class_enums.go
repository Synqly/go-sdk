package authentication

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Logon. A new logon session was requested.
const Activity_Logon ActivityId = 1
// Logoff. A logon session was terminated and no longer exists.
const Activity_Logoff ActivityId = 2
// AuthenticationTicket. A Kerberos authentication ticket (TGT) was requested.
const Activity_AuthenticationTicket ActivityId = 3
// ServiceTicketRequest. A Kerberos service ticket was requested.
const Activity_ServiceTicketRequest ActivityId = 4
// ServiceTicketRenew. A Kerberos service ticket was renewed.
const Activity_ServiceTicketRenew ActivityId = 5
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// AuthProtocolId Values
// Auth Protocol ID. The normalized identifier of the authentication protocol used to create the user session.

// Unknown. The authentication protocol is unknown.
const AuthProtocol_Unknown AuthProtocolId = 0
// NTLM
const AuthProtocol_NTLM AuthProtocolId = 1
// Kerberos
const AuthProtocol_Kerberos AuthProtocolId = 2
// Digest
const AuthProtocol_Digest AuthProtocolId = 3
// OpenID
const AuthProtocol_OpenID AuthProtocolId = 4
// SAML
const AuthProtocol_SAML AuthProtocolId = 5
// OAUTH2_0
const AuthProtocol_OAUTH2_0 AuthProtocolId = 6
// PAP
const AuthProtocol_PAP AuthProtocolId = 7
// CHAP
const AuthProtocol_CHAP AuthProtocolId = 8
// EAP
const AuthProtocol_EAP AuthProtocolId = 9
// RADIUS
const AuthProtocol_RADIUS AuthProtocolId = 10
// Other. The authentication protocol is not mapped. See the <code>auth_protocol</code> attribute, which contains a data source specific value.
const AuthProtocol_Other AuthProtocolId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// Identity_AccessManagement. Identity & Access Management (IAM) events relate to the supervision of the system's authentication and access control model. Examples of such events are the success or failure of authentication, granting of authority, password change, entity change, privileged use etc.
const Category_Identity_AccessManagement CategoryUid = 3

// ClassUid Values
// Class ID. The unique identifier of a class. A Class describes the attributes available in an event.

// Authentication. Authentication events report authentication session activities such as user attempts a logon or logoff, successfully or otherwise.
const Class_Authentication ClassUid = 3002

// LogonTypeId Values
// Logon Type ID. The normalized logon type identifier.

// System. Used only by the System account, for example at system startup.
const LogonType_System LogonTypeId = 0
// Interactive. A local logon to device console.
const LogonType_Interactive LogonTypeId = 2
// Network. A user or device logged onto this device from the network.
const LogonType_Network LogonTypeId = 3
// Batch. A batch server logon, where processes may be executing on behalf of a user without their direct intervention.
const LogonType_Batch LogonTypeId = 4
// OSService. A logon by a service or daemon that was started by the OS.
const LogonType_OSService LogonTypeId = 5
// Unlock. A user unlocked the device.
const LogonType_Unlock LogonTypeId = 7
// NetworkCleartext. A user logged on to this device from the network. The user's password in the authentication package was not hashed.
const LogonType_NetworkCleartext LogonTypeId = 8
// NewCredentials. A caller cloned its current token and specified new credentials for outbound connections. The new logon session has the same local identity, but uses different credentials for other network connections.
const LogonType_NewCredentials LogonTypeId = 9
// RemoteInteractive. A remote logon using Terminal Services or remote desktop application.
const LogonType_RemoteInteractive LogonTypeId = 10
// CachedInteractive. A user logged on to this device with network credentials that were stored locally on the device and the domain controller was not contacted to verify the credentials.
const LogonType_CachedInteractive LogonTypeId = 11
// CachedRemoteInteractive. Same as Remote Interactive. This is used for internal auditing.
const LogonType_CachedRemoteInteractive LogonTypeId = 12
// CachedUnlock. Workstation logon.
const LogonType_CachedUnlock LogonTypeId = 13
// Other. The logon type is not mapped. See the <code>logon_type</code> attribute, which contains a data source specific value.
const LogonType_Other LogonTypeId = 99

// SeverityId Values
// Severity ID. <p>The normalized identifier of the event/finding severity.</p>The normalized severity is a measurement the effort and expense required to manage and resolve an event or incident. Smaller numerical values represent lower impact events, and larger numerical values represent higher impact events.

// Unknown. The event/finding severity is unknown.
const Severity_Unknown SeverityId = 0
// Informational. Informational message. No action required.
const Severity_Informational SeverityId = 1
// Low. The user decides if action is needed.
const Severity_Low SeverityId = 2
// Medium. Action is required but the situation is not serious at this time.
const Severity_Medium SeverityId = 3
// High. Action is required immediately.
const Severity_High SeverityId = 4
// Critical. Action is required immediately and the scope is broad.
const Severity_Critical SeverityId = 5
// Fatal. An error occurred but it is too late to take remedial action.
const Severity_Fatal SeverityId = 6
// Other. The event/finding severity is not mapped. See the <code>severity</code> attribute, which contains a data source specific value.
const Severity_Other SeverityId = 99

// StatusId Values
// Status ID. The normalized identifier of the event status.

// Unknown. The status is unknown.
const Status_Unknown StatusId = 0
// Success
const Status_Success StatusId = 1
// Failure
const Status_Failure StatusId = 2
// Other. The event status is not mapped. See the <code>status</code> attribute, which contains a data source specific value.
const Status_Other StatusId = 99

// TypeUid Values
// Type ID. The event/finding type ID. It identifies the event's semantics and structure. The value is calculated by the logging system as: <code>class_uid * 100 + activity_id</code>.

// Authentication_Unknown
const Type_Authentication_Unknown TypeUid = 300200
// Authentication_Logon. A new logon session was requested.
const Type_Authentication_Logon TypeUid = 300201
// Authentication_Logoff. A logon session was terminated and no longer exists.
const Type_Authentication_Logoff TypeUid = 300202
// Authentication_AuthenticationTicket. A Kerberos authentication ticket (TGT) was requested.
const Type_Authentication_AuthenticationTicket TypeUid = 300203
// Authentication_ServiceTicketRequest. A Kerberos service ticket was requested.
const Type_Authentication_ServiceTicketRequest TypeUid = 300204
// Authentication_ServiceTicketRenew. A Kerberos service ticket was renewed.
const Type_Authentication_ServiceTicketRenew TypeUid = 300205
// Authentication_Other
const Type_Authentication_Other TypeUid = 300299

