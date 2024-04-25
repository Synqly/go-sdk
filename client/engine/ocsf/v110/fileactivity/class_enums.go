package fileactivity

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
// Create. A request to create a new file on a file system.
const Activity_Create ActivityId = 1
// Read. A request to read data from a file on a file system.
const Activity_Read ActivityId = 2
// Update. A request to write data to a file on a file system.
const Activity_Update ActivityId = 3
// Delete. A request to delete a file on a file system.
const Activity_Delete ActivityId = 4
// Rename. A request to rename a file on a file system.
const Activity_Rename ActivityId = 5
// SetAttributes. A request to set attributes for a file on a file system.
const Activity_SetAttributes ActivityId = 6
// SetSecurity. A request to set security for a file on a file system.
const Activity_SetSecurity ActivityId = 7
// GetAttributes. A request to get attributes for a file on a file system.
const Activity_GetAttributes ActivityId = 8
// GetSecurity. A request to get security for a file on a file system.
const Activity_GetSecurity ActivityId = 9
// Encrypt. A request to encrypt a file on a file system.
const Activity_Encrypt ActivityId = 10
// Decrypt. A request to decrypt a file on a file system.
const Activity_Decrypt ActivityId = 11
// Mount. A request to mount a file on a file system.
const Activity_Mount ActivityId = 12
// Unmount. A request to unmount a file from a file system.
const Activity_Unmount ActivityId = 13
// Open. A request to create a file handle.
const Activity_Open ActivityId = 14
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// SystemActivity. System Activity events.
const Category_SystemActivity CategoryUid = 1

// ClassUid Values
// Class ID. The unique identifier of a class. A Class describes the attributes available in an event.

// FileSystemActivity. File System Activity events report when a process performs an action on a file or folder.
const Class_FileSystemActivity ClassUid = 1001

// DispositionId Values
// Disposition ID. Describes the outcome or action taken by a security control, such as access control checks, malware detections or various types of policy violations.

// Unknown. The disposition was not known.
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
// Approved. A request or submission was approved.  For example, when a form was properly filled out and submitted. This is distinct from <code>1</code> 'Allowed'.
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
// Other. The disposition is not listed. The <code>disposition</code> attribute should be populated with a source specific caption.
const Disposition_Other DispositionId = 99

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

// FileSystemActivity_Unknown
const Type_FileSystemActivity_Unknown TypeUid = 100100
// FileSystemActivity_Create. A request to create a new file on a file system.
const Type_FileSystemActivity_Create TypeUid = 100101
// FileSystemActivity_Read. A request to read data from a file on a file system.
const Type_FileSystemActivity_Read TypeUid = 100102
// FileSystemActivity_Update. A request to write data to a file on a file system.
const Type_FileSystemActivity_Update TypeUid = 100103
// FileSystemActivity_Delete. A request to delete a file on a file system.
const Type_FileSystemActivity_Delete TypeUid = 100104
// FileSystemActivity_Rename. A request to rename a file on a file system.
const Type_FileSystemActivity_Rename TypeUid = 100105
// FileSystemActivity_SetAttributes. A request to set attributes for a file on a file system.
const Type_FileSystemActivity_SetAttributes TypeUid = 100106
// FileSystemActivity_SetSecurity. A request to set security for a file on a file system.
const Type_FileSystemActivity_SetSecurity TypeUid = 100107
// FileSystemActivity_GetAttributes. A request to get attributes for a file on a file system.
const Type_FileSystemActivity_GetAttributes TypeUid = 100108
// FileSystemActivity_GetSecurity. A request to get security for a file on a file system.
const Type_FileSystemActivity_GetSecurity TypeUid = 100109
// FileSystemActivity_Encrypt. A request to encrypt a file on a file system.
const Type_FileSystemActivity_Encrypt TypeUid = 100110
// FileSystemActivity_Decrypt. A request to decrypt a file on a file system.
const Type_FileSystemActivity_Decrypt TypeUid = 100111
// FileSystemActivity_Mount. A request to mount a file on a file system.
const Type_FileSystemActivity_Mount TypeUid = 100112
// FileSystemActivity_Unmount. A request to unmount a file from a file system.
const Type_FileSystemActivity_Unmount TypeUid = 100113
// FileSystemActivity_Open. A request to create a file handle.
const Type_FileSystemActivity_Open TypeUid = 100114
// FileSystemActivity_Other
const Type_FileSystemActivity_Other TypeUid = 100199

