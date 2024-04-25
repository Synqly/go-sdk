package webresourceaccessactivity

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// AccessGrant. The incoming request has permission to the web resource.
const Activity_AccessGrant ActivityId = 1
// AccessDeny. The incoming request does not have permission to the web resource.
const Activity_AccessDeny ActivityId = 2
// AccessRevoke. The incoming request's access has been revoked due to security policy enforcements.
const Activity_AccessRevoke ActivityId = 3
// AccessError. An error occurred during processing the request.
const Activity_AccessError ActivityId = 4
// Other. The event activity is not mapped.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// ApplicationActivity. Application Activity events report detailed information about the behavior of applications and services.
const Category_ApplicationActivity CategoryUid = 6

// ClassUid Values
// Class ID. The unique identifier of a class. A Class describes the attributes available in an event.

// WebResourceAccessActivity. Web Resource Access Activity events describe successful/failed attempts to access a web resource over HTTP.
const Class_WebResourceAccessActivity ClassUid = 6004

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

// WebResourceAccessActivity_Unknown
const Type_WebResourceAccessActivity_Unknown TypeUid = 600400
// WebResourceAccessActivity_AccessGrant. The incoming request has permission to the web resource.
const Type_WebResourceAccessActivity_AccessGrant TypeUid = 600401
// WebResourceAccessActivity_AccessDeny. The incoming request does not have permission to the web resource.
const Type_WebResourceAccessActivity_AccessDeny TypeUid = 600402
// WebResourceAccessActivity_AccessRevoke. The incoming request's access has been revoked due to security policy enforcements.
const Type_WebResourceAccessActivity_AccessRevoke TypeUid = 600403
// WebResourceAccessActivity_AccessError. An error occurred during processing the request.
const Type_WebResourceAccessActivity_AccessError TypeUid = 600404
// WebResourceAccessActivity_Other
const Type_WebResourceAccessActivity_Other TypeUid = 600499

