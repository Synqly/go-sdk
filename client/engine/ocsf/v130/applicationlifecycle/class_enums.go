package applicationlifecycle

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Install. Install the application.
const Activity_Install ActivityId = 1
// Remove. Remove the application.
const Activity_Remove ActivityId = 2
// Start. Start the application.
const Activity_Start ActivityId = 3
// Stop. Stop the application.
const Activity_Stop ActivityId = 4
// Restart. Restart the application.
const Activity_Restart ActivityId = 5
// Enable. Enable the application.
const Activity_Enable ActivityId = 6
// Disable. Disable the application.
const Activity_Disable ActivityId = 7
// Update. Update the application.
const Activity_Update ActivityId = 8
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// ApplicationActivity. Application Activity events report detailed information about the behavior of applications and services.
const Category_ApplicationActivity CategoryUid = 6

// ClassUid Values
// Class ID. The unique identifier of a class. A class describes the attributes available in an event.

// ApplicationLifecycle. Application Lifecycle events report installation, removal, start, stop of an application or service.
const Class_ApplicationLifecycle ClassUid = 6002

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

// ApplicationLifecycle_Unknown
const Type_ApplicationLifecycle_Unknown TypeUid = 600200
// ApplicationLifecycle_Install. Install the application.
const Type_ApplicationLifecycle_Install TypeUid = 600201
// ApplicationLifecycle_Remove. Remove the application.
const Type_ApplicationLifecycle_Remove TypeUid = 600202
// ApplicationLifecycle_Start. Start the application.
const Type_ApplicationLifecycle_Start TypeUid = 600203
// ApplicationLifecycle_Stop. Stop the application.
const Type_ApplicationLifecycle_Stop TypeUid = 600204
// ApplicationLifecycle_Restart. Restart the application.
const Type_ApplicationLifecycle_Restart TypeUid = 600205
// ApplicationLifecycle_Enable. Enable the application.
const Type_ApplicationLifecycle_Enable TypeUid = 600206
// ApplicationLifecycle_Disable. Disable the application.
const Type_ApplicationLifecycle_Disable TypeUid = 600207
// ApplicationLifecycle_Update. Update the application.
const Type_ApplicationLifecycle_Update TypeUid = 600208
// ApplicationLifecycle_Other
const Type_ApplicationLifecycle_Other TypeUid = 600299

