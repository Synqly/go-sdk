package apiactivity

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Create. The API call in the event pertains to a 'create' activity.
const Activity_Create ActivityId = 1
// Read. The API call in the event pertains to a 'read' activity.
const Activity_Read ActivityId = 2
// Update. The API call in the event pertains to a 'update' activity.
const Activity_Update ActivityId = 3
// Delete. The API call in the event pertains to a 'delete' activity.
const Activity_Delete ActivityId = 4
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// ApplicationActivity. Application Activity events report detailed information about the behavior of applications and services.
const Category_ApplicationActivity CategoryUid = 6

// ClassUid Values
// Class ID. The unique identifier of a class. A class describes the attributes available in an event.

// APIActivity. API events describe general CRUD (Create, Read, Update, Delete) API activities, e.g. (AWS Cloudtrail)
const Class_APIActivity ClassUid = 6003

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

// APIActivity_Unknown
const Type_APIActivity_Unknown TypeUid = 600300
// APIActivity_Create. The API call in the event pertains to a 'create' activity.
const Type_APIActivity_Create TypeUid = 600301
// APIActivity_Read. The API call in the event pertains to a 'read' activity.
const Type_APIActivity_Read TypeUid = 600302
// APIActivity_Update. The API call in the event pertains to a 'update' activity.
const Type_APIActivity_Update TypeUid = 600303
// APIActivity_Delete. The API call in the event pertains to a 'delete' activity.
const Type_APIActivity_Delete TypeUid = 600304
// APIActivity_Other
const Type_APIActivity_Other TypeUid = 600399

