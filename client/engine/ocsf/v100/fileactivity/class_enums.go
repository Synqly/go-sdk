package fileactivity

// ActivityId Values
// Activity ID. The activity ID of the event.

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
// Other. The event activity is not mapped.
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

