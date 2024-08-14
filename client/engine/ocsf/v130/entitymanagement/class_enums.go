package entitymanagement

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Create. Create a new managed entity.
const Activity_Create ActivityId = 1
// Read. Read an existing managed entity.
const Activity_Read ActivityId = 2
// Update. Update an existing managed entity.
const Activity_Update ActivityId = 3
// Delete. Delete a managed entity.
const Activity_Delete ActivityId = 4
// Move. Move or rename an existing managed entity.
const Activity_Move ActivityId = 5
// Enroll. Enroll an existing managed entity.
const Activity_Enroll ActivityId = 6
// Unenroll. Unenroll an existing managed entity.
const Activity_Unenroll ActivityId = 7
// Enable. Enable an existing managed entity. Note: This is typically regarded as a semi-permanent, editor visible, syncable change.
const Activity_Enable ActivityId = 8
// Disable. Disable an existing managed entity. Note: This is typically regarded as a semi-permanent, editor visible, syncable change.
const Activity_Disable ActivityId = 9
// Activate. Activate an existing managed entity. Note: This is a typically regarded as a transient change, a change of state of the engine.
const Activity_Activate ActivityId = 10
// Deactivate. Deactivate an existing managed entity. Note: This is a typically regarded as a transient change, a change of state of the engine.
const Activity_Deactivate ActivityId = 11
// Suspend. Suspend an existing managed entity.
const Activity_Suspend ActivityId = 12
// Resume. Resume (unsuspend) an existing managed entity.
const Activity_Resume ActivityId = 13
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// Identity_AccessManagement. Identity & Access Management (IAM) events relate to the supervision of the system's authentication and access control model. Examples of such events are the success or failure of authentication, granting of authority, password change, entity change, privileged use etc.
const Category_Identity_AccessManagement CategoryUid = 3

// ClassUid Values
// Class ID. The unique identifier of a class. A class describes the attributes available in an event.

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
// EntityManagement_Create. Create a new managed entity.
const Type_EntityManagement_Create TypeUid = 300401
// EntityManagement_Read. Read an existing managed entity.
const Type_EntityManagement_Read TypeUid = 300402
// EntityManagement_Update. Update an existing managed entity.
const Type_EntityManagement_Update TypeUid = 300403
// EntityManagement_Delete. Delete a managed entity.
const Type_EntityManagement_Delete TypeUid = 300404
// EntityManagement_Move. Move or rename an existing managed entity.
const Type_EntityManagement_Move TypeUid = 300405
// EntityManagement_Enroll. Enroll an existing managed entity.
const Type_EntityManagement_Enroll TypeUid = 300406
// EntityManagement_Unenroll. Unenroll an existing managed entity.
const Type_EntityManagement_Unenroll TypeUid = 300407
// EntityManagement_Enable. Enable an existing managed entity. Note: This is typically regarded as a semi-permanent, editor visible, syncable change.
const Type_EntityManagement_Enable TypeUid = 300408
// EntityManagement_Disable. Disable an existing managed entity. Note: This is typically regarded as a semi-permanent, editor visible, syncable change.
const Type_EntityManagement_Disable TypeUid = 300409
// EntityManagement_Activate. Activate an existing managed entity. Note: This is a typically regarded as a transient change, a change of state of the engine.
const Type_EntityManagement_Activate TypeUid = 300410
// EntityManagement_Deactivate. Deactivate an existing managed entity. Note: This is a typically regarded as a transient change, a change of state of the engine.
const Type_EntityManagement_Deactivate TypeUid = 300411
// EntityManagement_Suspend. Suspend an existing managed entity.
const Type_EntityManagement_Suspend TypeUid = 300412
// EntityManagement_Resume. Resume (unsuspend) an existing managed entity.
const Type_EntityManagement_Resume TypeUid = 300413
// EntityManagement_Other
const Type_EntityManagement_Other TypeUid = 300499

