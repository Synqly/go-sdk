package httpactivity

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
// Connect. The CONNECT method establishes a tunnel to the server identified by the target resource.
const Activity_Connect ActivityId = 1
// Delete. The DELETE method deletes the specified resource.
const Activity_Delete ActivityId = 2
// Get. The GET method requests a representation of the specified resource. Requests using GET should only retrieve data.
const Activity_Get ActivityId = 3
// Head. The HEAD method asks for a response identical to a GET request, but without the response body.
const Activity_Head ActivityId = 4
// Options. The OPTIONS method describes the communication options for the target resource.
const Activity_Options ActivityId = 5
// Post. The POST method submits an entity to the specified resource, often causing a change in state or side effects on the server.
const Activity_Post ActivityId = 6
// Put. The PUT method replaces all current representations of the target resource with the request payload.
const Activity_Put ActivityId = 7
// Trace. The TRACE method performs a message loop-back test along the path to the target resource.
const Activity_Trace ActivityId = 8
// Patch. The PATCH method applies partial modifications to a resource.
const Activity_Patch ActivityId = 9
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// NetworkActivity. Network Activity events.
const Category_NetworkActivity CategoryUid = 4

// ClassUid Values
// Class ID. The unique identifier of a class. A class describes the attributes available in an event.

// HTTPActivity. HTTP Activity events report HTTP connection and traffic information.
const Class_HTTPActivity ClassUid = 4002

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

// HTTPActivity_Unknown
const Type_HTTPActivity_Unknown TypeUid = 400200
// HTTPActivity_Connect. The CONNECT method establishes a tunnel to the server identified by the target resource.
const Type_HTTPActivity_Connect TypeUid = 400201
// HTTPActivity_Delete. The DELETE method deletes the specified resource.
const Type_HTTPActivity_Delete TypeUid = 400202
// HTTPActivity_Get. The GET method requests a representation of the specified resource. Requests using GET should only retrieve data.
const Type_HTTPActivity_Get TypeUid = 400203
// HTTPActivity_Head. The HEAD method asks for a response identical to a GET request, but without the response body.
const Type_HTTPActivity_Head TypeUid = 400204
// HTTPActivity_Options. The OPTIONS method describes the communication options for the target resource.
const Type_HTTPActivity_Options TypeUid = 400205
// HTTPActivity_Post. The POST method submits an entity to the specified resource, often causing a change in state or side effects on the server.
const Type_HTTPActivity_Post TypeUid = 400206
// HTTPActivity_Put. The PUT method replaces all current representations of the target resource with the request payload.
const Type_HTTPActivity_Put TypeUid = 400207
// HTTPActivity_Trace. The TRACE method performs a message loop-back test along the path to the target resource.
const Type_HTTPActivity_Trace TypeUid = 400208
// HTTPActivity_Patch. The PATCH method applies partial modifications to a resource.
const Type_HTTPActivity_Patch TypeUid = 400209
// HTTPActivity_Other
const Type_HTTPActivity_Other TypeUid = 400299

