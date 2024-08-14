package entitymanagement

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Create
const Activity_Create ActivityId = 1
// Read
const Activity_Read ActivityId = 2
// Update
const Activity_Update ActivityId = 3
// Delete
const Activity_Delete ActivityId = 4
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// Identity_AccessManagement. Identity & Access Management (IAM) events relate to the supervision of the system's authentication and access control model. Examples of such events are the success or failure of authentication, granting of authority, password change, entity change, privileged use etc.
const Category_Identity_AccessManagement CategoryUid = 3

// ClassUid Values
// Class ID. The unique identifier of a class. A Class describes the attributes available in an event.

// EntityManagement. Entity Management events report activity by a managed client, a micro service, or a user at a management console. The activity can be a create, read, update, and delete operation on a managed entity.
const Class_EntityManagement ClassUid = 3004

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

// EntityManagement_Unknown
const Type_EntityManagement_Unknown TypeUid = 300400
// EntityManagement_Create
const Type_EntityManagement_Create TypeUid = 300401
// EntityManagement_Read
const Type_EntityManagement_Read TypeUid = 300402
// EntityManagement_Update
const Type_EntityManagement_Update TypeUid = 300403
// EntityManagement_Delete
const Type_EntityManagement_Delete TypeUid = 300404
// EntityManagement_Other
const Type_EntityManagement_Other TypeUid = 300499

