package groupmanagement

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// AssignPrivileges. Assign privileges to a group.
const Activity_AssignPrivileges ActivityId = 1
// RevokePrivileges. Revoke privileges from a group.
const Activity_RevokePrivileges ActivityId = 2
// AddUser. Add user to a group.
const Activity_AddUser ActivityId = 3
// RemoveUser. Remove user from a group.
const Activity_RemoveUser ActivityId = 4
// Delete. A group was deleted.
const Activity_Delete ActivityId = 5
// Create. A group was created.
const Activity_Create ActivityId = 6
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// Identity_AccessManagement. Identity & Access Management (IAM) events relate to the supervision of the system's authentication and access control model. Examples of such events are the success or failure of authentication, granting of authority, password change, entity change, privileged use etc.
const Category_Identity_AccessManagement CategoryUid = 3

// ClassUid Values
// Class ID. The unique identifier of a class. A class describes the attributes available in an event.

// GroupManagement. Group Management events report management updates to a group, including updates to membership and permissions.
const Class_GroupManagement ClassUid = 3006

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

// GroupManagement_Unknown
const Type_GroupManagement_Unknown TypeUid = 300600
// GroupManagement_AssignPrivileges. Assign privileges to a group.
const Type_GroupManagement_AssignPrivileges TypeUid = 300601
// GroupManagement_RevokePrivileges. Revoke privileges from a group.
const Type_GroupManagement_RevokePrivileges TypeUid = 300602
// GroupManagement_AddUser. Add user to a group.
const Type_GroupManagement_AddUser TypeUid = 300603
// GroupManagement_RemoveUser. Remove user from a group.
const Type_GroupManagement_RemoveUser TypeUid = 300604
// GroupManagement_Delete. A group was deleted.
const Type_GroupManagement_Delete TypeUid = 300605
// GroupManagement_Create. A group was created.
const Type_GroupManagement_Create TypeUid = 300606
// GroupManagement_Other
const Type_GroupManagement_Other TypeUid = 300699

