package scanactivity

// ActionId Values
// Action ID. The action taken by a control or other policy-based system leading to an outcome or disposition. An unknown action may still correspond to a known disposition. Refer to <code>disposition_id</code> for the outcome of the action.

// Unknown. The action was unknown. The <code>disposition_id</code> attribute may still be set to a non-unknown value, for example 'Custom Action', 'Challenge'.
const Action_Unknown ActionId = 0
// Allowed. The activity was allowed. The <code>disposition_id</code> attribute should be set to a value that conforms to this action, for example 'Allowed', 'Approved', 'Delayed', 'No Action', 'Count' etc.
const Action_Allowed ActionId = 1
// Denied. The attempted activity was denied. The <code>disposition_id</code> attribute should be set to a value that conforms to this action, for example 'Blocked', 'Rejected', 'Quarantined', 'Isolated', 'Dropped', 'Access Revoked, etc.
const Action_Denied ActionId = 2
// Observed. The activity was observed, but neither explicitly allowed nor denied. This is common with IDS and EDR controls that report additional information on observed behavior such as TTPs. The <code>disposition_id</code> attribute should be set to a value that conforms to this action, for example 'Logged', 'Alert', 'Detected', 'Count', etc.
const Action_Observed ActionId = 3
// Modified. The activity was modified, adjusted, or corrected. The <code>disposition_id</code> attribute should be set appropriately, for example 'Restored', 'Corrected', 'Delayed', 'Captcha', 'Tagged'.
const Action_Modified ActionId = 4
// Other. The action is not mapped. See the <code>action</code> attribute which contains a data source specific value.
const Action_Other ActionId = 99

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
// Class ID. The unique identifier of a class. A class describes the attributes available in an event.

// ScanActivity. Scan events report the start, completion, and results of a scan job. The scan event includes the number of items that were scanned and the number of detections that were resolved.
const Class_ScanActivity ClassUid = 6007

// ConfidenceId Values
// Confidence ID. The normalized confidence refers to the accuracy of the rule that created the finding. A rule with a low confidence means that the finding scope is wide and may create finding reports that may not be malicious in nature.

// Unknown. The normalized confidence is unknown.
const Confidence_Unknown ConfidenceId = 0
// Low
const Confidence_Low ConfidenceId = 1
// Medium
const Confidence_Medium ConfidenceId = 2
// High
const Confidence_High ConfidenceId = 3
// Other. The confidence is not mapped to the defined enum values. See the <code>confidence</code> attribute, which contains a data source specific value.
const Confidence_Other ConfidenceId = 99

// DispositionId Values
// Disposition ID. Describes the outcome or action taken by a security control, such as access control checks, malware detections or various types of policy violations.

// Unknown. The disposition is unknown.
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
// Approved. A request or submission was approved. For example, when a form was properly filled out and submitted. This is distinct from <code>1</code> 'Allowed'.
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
// Other. The disposition is not mapped. See the <code>disposition</code> attribute, which contains a data source specific value.
const Disposition_Other DispositionId = 99

// RiskLevelId Values
// Risk Level ID. The normalized risk level id.

// Info
const RiskLevel_Info RiskLevelId = 0
// Low
const RiskLevel_Low RiskLevelId = 1
// Medium
const RiskLevel_Medium RiskLevelId = 2
// High
const RiskLevel_High RiskLevelId = 3
// Critical
const RiskLevel_Critical RiskLevelId = 4
// Other. The risk level is not mapped. See the <code>risk_level</code> attribute, which contains a data source specific value.
const RiskLevel_Other RiskLevelId = 99

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
// Other. The status is not mapped. See the <code>status</code> attribute, which contains a data source specific value.
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

