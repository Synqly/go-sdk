package networkactivity

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Open. A new network connection was opened.
const Activity_Open ActivityId = 1
// Close. The network connection was closed.
const Activity_Close ActivityId = 2
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

// NetworkActivity. Network Activity events report network connection and traffic activity.
const Class_NetworkActivity ClassUid = 4001

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

// NetworkActivity_Unknown
const Type_NetworkActivity_Unknown TypeUid = 400100
// NetworkActivity_Open. A new network connection was opened.
const Type_NetworkActivity_Open TypeUid = 400101
// NetworkActivity_Close. The network connection was closed.
const Type_NetworkActivity_Close TypeUid = 400102
// NetworkActivity_Reset. The network connection was abnormally terminated or closed by a middle device like firewalls.
const Type_NetworkActivity_Reset TypeUid = 400103
// NetworkActivity_Fail. The network connection failed. For example a connection timeout or no route to host.
const Type_NetworkActivity_Fail TypeUid = 400104
// NetworkActivity_Refuse. The network connection was refused. For example an attempt to connect to a server port which is not open.
const Type_NetworkActivity_Refuse TypeUid = 400105
// NetworkActivity_Traffic. Network traffic report.
const Type_NetworkActivity_Traffic TypeUid = 400106
// NetworkActivity_Other
const Type_NetworkActivity_Other TypeUid = 400199

