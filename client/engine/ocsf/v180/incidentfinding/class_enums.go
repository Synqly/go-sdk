package incidentfinding

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
// Activity ID. The normalized identifier of the Incident activity.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Create. Reports the creation of an Incident.
const Activity_Create ActivityId = 1
// Update. Reports updates to an Incident.
const Activity_Update ActivityId = 2
// Close. Reports closure of an Incident .
const Activity_Close ActivityId = 3
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// Findings. Findings events report findings, detections, and possible resolutions of malware, anomalies, or other actions performed by security products.
const Category_Findings CategoryUid = 2

// ClassUid Values
// Class ID. The unique identifier of a class. A class describes the attributes available in an event.

// IncidentFinding. An Incident Finding reports the creation, update, or closure of security incidents as a result of detections and/or analytics. <br><strong>Note: </strong><code>Incident Finding</code> implicitly includes the <code>incident</code> profile and it should be added to the <code>metadata.profiles[]</code> array.
const Class_IncidentFinding ClassUid = 2005

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

// ImpactId Values
// Impact ID. The normalized impact of the incident or finding. Per NIST, this is the magnitude of harm that can be expected to result from the consequences of unauthorized disclosure, modification, destruction, or loss of information or information system availability.

// Unknown. The normalized impact is unknown.
const Impact_Unknown ImpactId = 0
// Low. The magnitude of harm is low.
const Impact_Low ImpactId = 1
// Medium. The magnitude of harm is moderate.
const Impact_Medium ImpactId = 2
// High. The magnitude of harm is high.
const Impact_High ImpactId = 3
// Critical. The magnitude of harm is high and the scope is widespread.
const Impact_Critical ImpactId = 4
// Other. The impact is not mapped. See the <code>impact</code> attribute, which contains a data source specific value.
const Impact_Other ImpactId = 99

// PriorityId Values
// Priority ID. The normalized priority. Priority identifies the relative importance of the incident or finding. It is a measurement of urgency.

// Unknown. No priority is assigned.
const Priority_Unknown PriorityId = 0
// Low. Application or personal procedure is unusable, where a workaround is available or a repair is possible.
const Priority_Low PriorityId = 1
// Medium. Non-critical function or procedure is unusable or hard to use causing operational disruptions with no direct impact on a service's availability. A workaround is available.
const Priority_Medium PriorityId = 2
// High. Critical functionality or network access is interrupted, degraded or unusable, having a severe impact on services availability. No acceptable alternative is possible.
const Priority_High PriorityId = 3
// Critical. Interruption making a critical functionality inaccessible or a complete network interruption causing a severe impact on services availability. There is no possible alternative.
const Priority_Critical PriorityId = 4
// Other. The priority is not normalized.
const Priority_Other PriorityId = 99

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
// Status ID. The normalized status identifier of the Incident.

// Unknown. The status is unknown.
const Status_Unknown StatusId = 0
// New. The service desk has received the incident but has not assigned it to an agent.
const Status_New StatusId = 1
// InProgress. The incident has been assigned to an agent but has not been resolved. The agent is actively working with the user to diagnose and resolve the incident.
const Status_InProgress StatusId = 2
// OnHold. The incident requires some information or response from the user or from a third party.
const Status_OnHold StatusId = 3
// Resolved. The service desk has confirmed that the incident is resolved.
const Status_Resolved StatusId = 4
// Closed. The incident is resolved and no further action is necessary.
const Status_Closed StatusId = 5
// Other. The status is not mapped. See the <code>status</code> attribute, which contains a data source specific value.
const Status_Other StatusId = 99

// TypeUid Values
// Type ID. The event/finding type ID. It identifies the event's semantics and structure. The value is calculated by the logging system as: <code>class_uid * 100 + activity_id</code>.

// IncidentFinding_Unknown
const Type_IncidentFinding_Unknown TypeUid = 200500
// IncidentFinding_Create. Reports the creation of an Incident.
const Type_IncidentFinding_Create TypeUid = 200501
// IncidentFinding_Update. Reports updates to an Incident.
const Type_IncidentFinding_Update TypeUid = 200502
// IncidentFinding_Close. Reports closure of an Incident .
const Type_IncidentFinding_Close TypeUid = 200503
// IncidentFinding_Other
const Type_IncidentFinding_Other TypeUid = 200599

// VerdictId Values
// Verdict ID. The normalized verdict of an Incident.

// Unknown. The type is unknown.
const Verdict_Unknown VerdictId = 0
// FalsePositive. The incident is a false positive.
const Verdict_FalsePositive VerdictId = 1
// TruePositive. The incident is a true positive.
const Verdict_TruePositive VerdictId = 2
// Disregard. The incident can be disregarded as it is unimportant, an error or accident.
const Verdict_Disregard VerdictId = 3
// Suspicious. The incident is suspicious.
const Verdict_Suspicious VerdictId = 4
// Benign. The incident is benign.
const Verdict_Benign VerdictId = 5
// Test. The incident is a test.
const Verdict_Test VerdictId = 6
// InsufficientData. The incident has insufficient data to make a verdict.
const Verdict_InsufficientData VerdictId = 7
// SecurityRisk. The incident is a security risk.
const Verdict_SecurityRisk VerdictId = 8
// ManagedExternally. The incident remediation or required actions are managed externally.
const Verdict_ManagedExternally VerdictId = 9
// Duplicate. The incident is a duplicate.
const Verdict_Duplicate VerdictId = 10
// Other. The type is not mapped. See the <code>type</code> attribute, which contains a data source specific value.
const Verdict_Other VerdictId = 99

