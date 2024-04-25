package dnsactivity

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Query. The DNS query request.
const Activity_Query ActivityId = 1
// Response. The DNS query response.
const Activity_Response ActivityId = 2
// Reset. The network connection was abnormally terminated or closed by a middle device like firewalls.
const Activity_Reset ActivityId = 3
// Fail. The network connection failed. For example a connection timeout or no route to host.
const Activity_Fail ActivityId = 4
// Refuse. The network connection was refused. For example an attempt to connect to a server port which is not open.
const Activity_Refuse ActivityId = 5
// Traffic. Network traffic report.
const Activity_Traffic ActivityId = 6
// Other. The event activity is not mapped.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// NetworkActivity. Network Activity events.
const Category_NetworkActivity CategoryUid = 4

// ClassUid Values
// Class ID. The unique identifier of a class. A Class describes the attributes available in an event.

// DNSActivity. DNS Activity events report DNS queries and answers as seen on the network.
const Class_DNSActivity ClassUid = 4003

// DispositionId Values
// Disposition ID. When security issues, such as malware or policy violations, are detected and possibly corrected, then <code>disposition_id</code> describes the action taken by the security product.

// Unknown
const Disposition_Unknown DispositionId = 0
// Allowed
const Disposition_Allowed DispositionId = 1
// Blocked
const Disposition_Blocked DispositionId = 2
// Quarantined
const Disposition_Quarantined DispositionId = 3
// Isolated
const Disposition_Isolated DispositionId = 4
// Deleted
const Disposition_Deleted DispositionId = 5
// Dropped
const Disposition_Dropped DispositionId = 6
// CustomAction. Executed custom action such as run a command script.
const Disposition_CustomAction DispositionId = 7
// Approved
const Disposition_Approved DispositionId = 8
// Restored
const Disposition_Restored DispositionId = 9
// Exonerated. No longer suspicious (re-scored).
const Disposition_Exonerated DispositionId = 10
// Corrected
const Disposition_Corrected DispositionId = 11
// PartiallyCorrected
const Disposition_PartiallyCorrected DispositionId = 12
// Uncorrected
const Disposition_Uncorrected DispositionId = 13
// Delayed. Requires reboot to finish the operation.
const Disposition_Delayed DispositionId = 14
// Detected
const Disposition_Detected DispositionId = 15
// NoAction
const Disposition_NoAction DispositionId = 16
// Logged
const Disposition_Logged DispositionId = 17
// Tagged. Marked with extended attributes.
const Disposition_Tagged DispositionId = 18
// Other
const Disposition_Other DispositionId = 99

// RcodeId Values
// Response Code ID. The normalized identifier of the DNS server response code. See <a target='_blank' href='https://datatracker.ietf.org/doc/html/rfc6895'>RFC-6895</a>.

// NoError. No Error.
const Rcode_NoError RcodeId = 0
// FormError. Format Error.
const Rcode_FormError RcodeId = 1
// ServError. Server Failure.
const Rcode_ServError RcodeId = 2
// NXDomain. Non-Existent Domain.
const Rcode_NXDomain RcodeId = 3
// NotImp. Not Implemented.
const Rcode_NotImp RcodeId = 4
// Refused. Query Refused.
const Rcode_Refused RcodeId = 5
// YXDomain. Name Exists when it should not.
const Rcode_YXDomain RcodeId = 6
// YXRRSet. RR Set Exists when it should not.
const Rcode_YXRRSet RcodeId = 7
// NXRRSet. RR Set that should exist does not.
const Rcode_NXRRSet RcodeId = 8
// NotAuth. Not Authorized or Server Not Authoritative for zone.
const Rcode_NotAuth RcodeId = 9
// NotZone. Name not contained in zone.
const Rcode_NotZone RcodeId = 10
// DSOTYPENI. DSO-TYPE Not Implemented.
const Rcode_DSOTYPENI RcodeId = 11
// BADSIG_VERS. TSIG Signature Failure or Bad OPT Version.
const Rcode_BADSIG_VERS RcodeId = 16
// BADKEY. Key not recognized.
const Rcode_BADKEY RcodeId = 17
// BADTIME. Signature out of time window.
const Rcode_BADTIME RcodeId = 18
// BADMODE. Bad TKEY Mode.
const Rcode_BADMODE RcodeId = 19
// BADNAME. Duplicate key name.
const Rcode_BADNAME RcodeId = 20
// BADALG. Algorithm not supported.
const Rcode_BADALG RcodeId = 21
// BADTRUNC. Bad Truncation.
const Rcode_BADTRUNC RcodeId = 22
// BADCOOKIE. Bad/missing Server Cookie.
const Rcode_BADCOOKIE RcodeId = 23
// Unassigned. The codes deemed to be unassigned by the RFC (unassigned codes: 12-15, 24-3840, 4096-65534).
const Rcode_Unassigned RcodeId = 24
// Reserved. The codes deemed to be reserved by the RFC (codes: 3841-4095, 65535).
const Rcode_Reserved RcodeId = 25
// Other. The dns response code is not defined by the RFC.
const Rcode_Other RcodeId = 99

// SeverityId Values
// Severity ID. <p>The normalized identifier of the event severity.</p>The normalized severity is a measurement the effort and expense required to manage and resolve an event or incident. Smaller numerical values represent lower impact events, and larger numerical values represent higher impact events.

// Unknown. The event severity is not known.
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
// Other. The event severity is not mapped. See the <code>severity</code> attribute, which contains a data source specific value.
const Severity_Other SeverityId = 99

// StatusId Values
// Status ID. The normalized identifier of the event status.

// Unknown
const Status_Unknown StatusId = 0
// Success
const Status_Success StatusId = 1
// Failure
const Status_Failure StatusId = 2
// Other. The event status is not mapped. See the <code>status</code> attribute, which contains a data source specific value.
const Status_Other StatusId = 99

// TypeUid Values
// Type ID. The event type ID. It identifies the event's semantics and structure. The value is calculated by the logging system as: <code>class_uid * 100 + activity_id</code>.

// DNSActivity_Unknown
const Type_DNSActivity_Unknown TypeUid = 400300
// DNSActivity_Query. The DNS query request.
const Type_DNSActivity_Query TypeUid = 400301
// DNSActivity_Response. The DNS query response.
const Type_DNSActivity_Response TypeUid = 400302
// DNSActivity_Reset. The network connection was abnormally terminated or closed by a middle device like firewalls.
const Type_DNSActivity_Reset TypeUid = 400303
// DNSActivity_Fail. The network connection failed. For example a connection timeout or no route to host.
const Type_DNSActivity_Fail TypeUid = 400304
// DNSActivity_Refuse. The network connection was refused. For example an attempt to connect to a server port which is not open.
const Type_DNSActivity_Refuse TypeUid = 400305
// DNSActivity_Traffic. Network traffic report.
const Type_DNSActivity_Traffic TypeUid = 400306
// DNSActivity_Other
const Type_DNSActivity_Other TypeUid = 400399

