package accountchange

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Create. A user/role was created.
const Activity_Create ActivityId = 1
// Enable. A user/role was enabled.
const Activity_Enable ActivityId = 2
// PasswordChange. An attempt was made to change an account's password.
const Activity_PasswordChange ActivityId = 3
// PasswordReset. An attempt was made to reset an account's password.
const Activity_PasswordReset ActivityId = 4
// Disable. A user/role was disabled.
const Activity_Disable ActivityId = 5
// Delete. A user/role was deleted.
const Activity_Delete ActivityId = 6
// AttachPolicy. An IAM Policy was attached to a user/role.
const Activity_AttachPolicy ActivityId = 7
// DetachPolicy. An IAM Policy was detached from a user/role.
const Activity_DetachPolicy ActivityId = 8
// Lock. A user account was locked out.
const Activity_Lock ActivityId = 9
// MFAFactorEnable. An authentication factor was enabled for an account.
const Activity_MFAFactorEnable ActivityId = 10
// MFAFactorDisable. An authentication factor was disabled for an account.
const Activity_MFAFactorDisable ActivityId = 11
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// Identity_AccessManagement. Identity & Access Management (IAM) events relate to the supervision of the system's authentication and access control model. Examples of such events are the success or failure of authentication, granting of authority, password change, entity change, privileged use etc.
const Category_Identity_AccessManagement CategoryUid = 3

// ClassUid Values
// Class ID. The unique identifier of a class. A Class describes the attributes available in an event.

// AccountChange. Account Change events report when specific user account management tasks are performed, such as a user/role being created, changed, deleted, renamed, disabled, enabled, locked out or unlocked.
const Class_AccountChange ClassUid = 3001

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

// AccountChange_Unknown
const Type_AccountChange_Unknown TypeUid = 300100
// AccountChange_Create. A user/role was created.
const Type_AccountChange_Create TypeUid = 300101
// AccountChange_Enable. A user/role was enabled.
const Type_AccountChange_Enable TypeUid = 300102
// AccountChange_PasswordChange. An attempt was made to change an account's password.
const Type_AccountChange_PasswordChange TypeUid = 300103
// AccountChange_PasswordReset. An attempt was made to reset an account's password.
const Type_AccountChange_PasswordReset TypeUid = 300104
// AccountChange_Disable. A user/role was disabled.
const Type_AccountChange_Disable TypeUid = 300105
// AccountChange_Delete. A user/role was deleted.
const Type_AccountChange_Delete TypeUid = 300106
// AccountChange_AttachPolicy. An IAM Policy was attached to a user/role.
const Type_AccountChange_AttachPolicy TypeUid = 300107
// AccountChange_DetachPolicy. An IAM Policy was detached from a user/role.
const Type_AccountChange_DetachPolicy TypeUid = 300108
// AccountChange_Lock. A user account was locked out.
const Type_AccountChange_Lock TypeUid = 300109
// AccountChange_MFAFactorEnable. An authentication factor was enabled for an account.
const Type_AccountChange_MFAFactorEnable TypeUid = 300110
// AccountChange_MFAFactorDisable. An authentication factor was disabled for an account.
const Type_AccountChange_MFAFactorDisable TypeUid = 300111
// AccountChange_Other
const Type_AccountChange_Other TypeUid = 300199

