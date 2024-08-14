package incidentfinding

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

// IncidentFinding. An Incident Finding reports the creation, update, or closure of security incidents as a result of detections and/or analytics.
const Class_IncidentFinding ClassUid = 2005

// ConfidenceId Values
// Confidence Id. The normalized confidence refers to the accuracy of the rule that created the finding. A rule with a low confidence means that the finding scope is wide and may create finding reports that may not be malicious in nature.

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

// ImpactId Values
// Impact ID. The normalized impact of the finding.

// Unknown. The normalized impact is unknown.
const Impact_Unknown ImpactId = 0
// Low
const Impact_Low ImpactId = 1
// Medium
const Impact_Medium ImpactId = 2
// High
const Impact_High ImpactId = 3
// Critical
const Impact_Critical ImpactId = 4
// Other. The impact is not mapped. See the <code>impact</code> attribute, which contains a data source specific value.
const Impact_Other ImpactId = 99

// PriorityId Values
// Priority ID. The normalized priority. Priority identifies the relative importance of the finding. It is a measurement of urgency.

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
// Other. The event status is not mapped. See the <code>status</code> attribute, which contains a data source specific value.
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

