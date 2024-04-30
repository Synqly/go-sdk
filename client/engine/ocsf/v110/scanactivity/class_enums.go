package scanactivity

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Started. The scan was started.
const Activity_Started ActivityId = 1
// Completed. The scan was completed.
const Activity_Completed ActivityId = 2
// Cancelled. The scan was cancelled.
const Activity_Cancelled ActivityId = 3
// DurationViolation. The allocated scan time was insufficient to complete the requested scan.
const Activity_DurationViolation ActivityId = 4
// PauseViolation. The scan was paused, either by the user or by program constraints (e.g. scans that are suspended during certain time intervals), and not resumed within the allotted time.
const Activity_PauseViolation ActivityId = 5
// Error. The scan could not be completed due to an internal error.
const Activity_Error ActivityId = 6
// Paused. The scan was paused.
const Activity_Paused ActivityId = 7
// Resumed. The scan was resumed from the pause point.
const Activity_Resumed ActivityId = 8
// Restarted. The scan restarted from the beginning of the file enumeration.
const Activity_Restarted ActivityId = 9
// Delayed. The user delayed the scan.
const Activity_Delayed ActivityId = 10
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// ApplicationActivity. Application Activity events report detailed information about the behavior of applications and services.
const Category_ApplicationActivity CategoryUid = 6

// ClassUid Values
// Class ID. The unique identifier of a class. A Class describes the attributes available in an event.

// ScanActivity. Scan events report the start, completion, and results of a scan job. The scan event includes the number of items that were scanned and the number of detections that were resolved.
const Class_ScanActivity ClassUid = 6007

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

// ScanActivity_Unknown
const Type_ScanActivity_Unknown TypeUid = 600700
// ScanActivity_Started. The scan was started.
const Type_ScanActivity_Started TypeUid = 600701
// ScanActivity_Completed. The scan was completed.
const Type_ScanActivity_Completed TypeUid = 600702
// ScanActivity_Cancelled. The scan was cancelled.
const Type_ScanActivity_Cancelled TypeUid = 600703
// ScanActivity_DurationViolation. The allocated scan time was insufficient to complete the requested scan.
const Type_ScanActivity_DurationViolation TypeUid = 600704
// ScanActivity_PauseViolation. The scan was paused, either by the user or by program constraints (e.g. scans that are suspended during certain time intervals), and not resumed within the allotted time.
const Type_ScanActivity_PauseViolation TypeUid = 600705
// ScanActivity_Error. The scan could not be completed due to an internal error.
const Type_ScanActivity_Error TypeUid = 600706
// ScanActivity_Paused. The scan was paused.
const Type_ScanActivity_Paused TypeUid = 600707
// ScanActivity_Resumed. The scan was resumed from the pause point.
const Type_ScanActivity_Resumed TypeUid = 600708
// ScanActivity_Restarted. The scan restarted from the beginning of the file enumeration.
const Type_ScanActivity_Restarted TypeUid = 600709
// ScanActivity_Delayed. The user delayed the scan.
const Type_ScanActivity_Delayed TypeUid = 600710
// ScanActivity_Other
const Type_ScanActivity_Other TypeUid = 600799

