package dnsactivity

// ActionId Values
// Action ID. The action taken by a control or other policy-based system leading to an outcome or disposition. Dispositions conform to an action of <code>1</code> 'Allowed' or <code>2</code> 'Denied' in most cases. Note that <code>99</code> 'Other' is not an option. No action would equate to <code>1</code> 'Allowed'. An unknown action may still correspond to a known disposition. Refer to <code>disposition_id</code> for the outcome of the action.

// Unknown. The action was unknown. The <code>disposition_id</code> attribute may still be set to a non-unknown value, for example 'Count', 'Uncorrected', 'Isolated', 'Quarantined' or 'Exonerated'.
const Action_Unknown ActionId = 0
// Allowed. The activity was allowed. The <code>disposition_id</code> attribute should be set to a value that conforms to this action, for example 'Allowed', 'Approved', 'Delayed', 'No Action', 'Count' etc.
const Action_Allowed ActionId = 1
// Denied. The attempted activity was denied. The <code>disposition_id</code> attribute should be set to a value that conforms to this action, for example 'Blocked', 'Rejected', 'Quarantined', 'Isolated', 'Dropped', 'Access Revoked, etc.
const Action_Denied ActionId = 2
// Other. The action was not mapped. See the <code>action</code> attribute, which contains a data source specific value.
const Action_Other ActionId = 99

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Query. The DNS query request.
const Activity_Query ActivityId = 1
// Response. The DNS query response.
const Activity_Response ActivityId = 2
// Traffic. Bidirectional DNS request and response traffic.
const Activity_Traffic ActivityId = 6
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// NetworkActivity. Network Activity events.
const Category_NetworkActivity CategoryUid = 4

// ClassUid Values
// Class ID. The unique identifier of a class. A class describes the attributes available in an event.

// DNSActivity. DNS Activity events report DNS queries and answers as seen on the network.
const Class_DNSActivity ClassUid = 4003

// DispositionId Values
// Disposition ID. Describes the outcome or action taken by a security control, such as access control checks, malware detections or various types of policy violations.

// Unknown. The disposition is unknown.
const Disposition_Unknown DispositionId = 0
// Allowed. Granted access or allowed the action to the protected resource.
const Disposition_Allowed DispositionId = 1
// Blocked. Denied access or blocked the action to the protected resource.
const Disposition_Blocked DispositionId = 2
// Quarantined. A suspicious file or other content was moved to a benign location.
const Disposition_Quarantined DispositionId = 3
// Isolated. A session was isolated on the network or within a browser.
const Disposition_Isolated DispositionId = 4
// Deleted. A file or other content was deleted.
const Disposition_Deleted DispositionId = 5
// Dropped. The request was detected as a threat and resulted in the connection being dropped.
const Disposition_Dropped DispositionId = 6
// CustomAction. A custom action was executed such as running of a command script. Use the <code>message</code> attribute of the base class for details.
const Disposition_CustomAction DispositionId = 7
// Approved. A request or submission was approved. For example, when a form was properly filled out and submitted. This is distinct from <code>1</code> 'Allowed'.
const Disposition_Approved DispositionId = 8
// Restored. A quarantined file or other content was restored to its original location.
const Disposition_Restored DispositionId = 9
// Exonerated. A suspicious or risky entity was deemed to no longer be suspicious (re-scored).
const Disposition_Exonerated DispositionId = 10
// Corrected. A corrupt file or configuration was corrected.
const Disposition_Corrected DispositionId = 11
// PartiallyCorrected. A corrupt file or configuration was partially corrected.
const Disposition_PartiallyCorrected DispositionId = 12
// Uncorrected. A corrupt file or configuration was not corrected.
const Disposition_Uncorrected DispositionId = 13
// Delayed. An operation was delayed, for example if a restart was required to finish the operation.
const Disposition_Delayed DispositionId = 14
// Detected. Suspicious activity or a policy violation was detected without further action.
const Disposition_Detected DispositionId = 15
// NoAction. The outcome of an operation had no action taken.
const Disposition_NoAction DispositionId = 16
// Logged. The operation or action was logged without further action.
const Disposition_Logged DispositionId = 17
// Tagged. A file or other entity was marked with extended attributes.
const Disposition_Tagged DispositionId = 18
// Alert. The request or activity was detected as a threat and resulted in a notification but request was not blocked.
const Disposition_Alert DispositionId = 19
// Count. Counted the request or activity but did not determine whether to allow it or block it.
const Disposition_Count DispositionId = 20
// Reset. The request was detected as a threat and resulted in the connection being reset.
const Disposition_Reset DispositionId = 21
// Captcha. Required the end user to solve a CAPTCHA puzzle to prove that a human being is sending the request.
const Disposition_Captcha DispositionId = 22
// Challenge. Ran a silent challenge that required the client session to verify that it's a browser, and not a bot.
const Disposition_Challenge DispositionId = 23
// AccessRevoked. The requestor's access has been revoked due to security policy enforcements. Note: use the <code>Host</code> profile if the <code>User</code> or <code>Actor</code> requestor is not present in the event class.
const Disposition_AccessRevoked DispositionId = 24
// Rejected. A request or submission was rejected.  For example, when a form was improperly filled out and submitted. This is distinct from <code>2</code> 'Blocked'.
const Disposition_Rejected DispositionId = 25
// Unauthorized. An attempt to access a resource was denied due to an authorization check that failed. This is a more specific disposition than <code>2</code> 'Blocked' and can be complemented with the <code>authorizations</code> attribute for more detail.
const Disposition_Unauthorized DispositionId = 26
// Error. An error occurred during the processing of the activity or request. Use the <code>message</code> attribute of the base class for details.
const Disposition_Error DispositionId = 27
// Other. The disposition is not mapped. See the <code>disposition</code> attribute, which contains a data source specific value.
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

// DNSActivity_Unknown
const Type_DNSActivity_Unknown TypeUid = 400300
// DNSActivity_Query. The DNS query request.
const Type_DNSActivity_Query TypeUid = 400301
// DNSActivity_Response. The DNS query response.
const Type_DNSActivity_Response TypeUid = 400302
// DNSActivity_Traffic. Bidirectional DNS request and response traffic.
const Type_DNSActivity_Traffic TypeUid = 400306
// DNSActivity_Other
const Type_DNSActivity_Other TypeUid = 400399

