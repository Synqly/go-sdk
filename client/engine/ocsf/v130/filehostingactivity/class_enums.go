package filehostingactivity

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Upload. Upload a file.
const Activity_Upload ActivityId = 1
// Download. Download a file.
const Activity_Download ActivityId = 2
// Update. Update a file.
const Activity_Update ActivityId = 3
// Delete. Delete a file.
const Activity_Delete ActivityId = 4
// Rename. Rename a file.
const Activity_Rename ActivityId = 5
// Copy. Copy a file.
const Activity_Copy ActivityId = 6
// Move. Move a file.
const Activity_Move ActivityId = 7
// Restore. Restore a file.
const Activity_Restore ActivityId = 8
// Preview. Preview a file.
const Activity_Preview ActivityId = 9
// Lock. Lock a file.
const Activity_Lock ActivityId = 10
// Unlock. Unlock a file.
const Activity_Unlock ActivityId = 11
// Share. Share a file.
const Activity_Share ActivityId = 12
// Unshare. Unshare a file.
const Activity_Unshare ActivityId = 13
// Open. Open a file.
const Activity_Open ActivityId = 14
// Sync. Mark a file or folder to sync with a computer.
const Activity_Sync ActivityId = 15
// Unsync. Mark a file or folder to not sync with a computer.
const Activity_Unsync ActivityId = 16
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// ApplicationActivity. Application Activity events report detailed information about the behavior of applications and services.
const Category_ApplicationActivity CategoryUid = 6

// ClassUid Values
// Class ID. The unique identifier of a class. A class describes the attributes available in an event.

// FileHostingActivity. File Hosting Activity events report the actions taken by file management applications, including file sharing servers like Sharepoint and services such as Box, MS OneDrive, or Google Drive.
const Class_FileHostingActivity ClassUid = 6006

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

// FileHostingActivity_Unknown
const Type_FileHostingActivity_Unknown TypeUid = 600600
// FileHostingActivity_Upload. Upload a file.
const Type_FileHostingActivity_Upload TypeUid = 600601
// FileHostingActivity_Download. Download a file.
const Type_FileHostingActivity_Download TypeUid = 600602
// FileHostingActivity_Update. Update a file.
const Type_FileHostingActivity_Update TypeUid = 600603
// FileHostingActivity_Delete. Delete a file.
const Type_FileHostingActivity_Delete TypeUid = 600604
// FileHostingActivity_Rename. Rename a file.
const Type_FileHostingActivity_Rename TypeUid = 600605
// FileHostingActivity_Copy. Copy a file.
const Type_FileHostingActivity_Copy TypeUid = 600606
// FileHostingActivity_Move. Move a file.
const Type_FileHostingActivity_Move TypeUid = 600607
// FileHostingActivity_Restore. Restore a file.
const Type_FileHostingActivity_Restore TypeUid = 600608
// FileHostingActivity_Preview. Preview a file.
const Type_FileHostingActivity_Preview TypeUid = 600609
// FileHostingActivity_Lock. Lock a file.
const Type_FileHostingActivity_Lock TypeUid = 600610
// FileHostingActivity_Unlock. Unlock a file.
const Type_FileHostingActivity_Unlock TypeUid = 600611
// FileHostingActivity_Share. Share a file.
const Type_FileHostingActivity_Share TypeUid = 600612
// FileHostingActivity_Unshare. Unshare a file.
const Type_FileHostingActivity_Unshare TypeUid = 600613
// FileHostingActivity_Open. Open a file.
const Type_FileHostingActivity_Open TypeUid = 600614
// FileHostingActivity_Sync. Mark a file or folder to sync with a computer.
const Type_FileHostingActivity_Sync TypeUid = 600615
// FileHostingActivity_Unsync. Mark a file or folder to not sync with a computer.
const Type_FileHostingActivity_Unsync TypeUid = 600616
// FileHostingActivity_Other
const Type_FileHostingActivity_Other TypeUid = 600699

