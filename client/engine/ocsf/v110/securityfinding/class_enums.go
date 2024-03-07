package securityfinding

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Create. A security finding was created.
const Activity_Create ActivityId = 1
// Update. A security finding was updated.
const Activity_Update ActivityId = 2
// Close. A security finding was closed.
const Activity_Close ActivityId = 3
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// Findings. Findings events report findings, detections, and possible resolutions of malware, anomalies, or other actions performed by security products.
const Category_Findings CategoryUid = 2

// ClassUid Values
// Class ID. The unique identifier of a class. A Class describes the attributes available in an event.

// SecurityFinding. Security Finding events describe findings, detections, anomalies, alerts and/or actions performed by security products
const Class_SecurityFinding ClassUid = 2001

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

// StateId Values
// State ID. The normalized state identifier of a security finding.

// Unknown. The state is unknown.
const State_Unknown StateId = 0
// New. The finding is new and yet to be reviewed.
const State_New StateId = 1
// InProgress. The finding is under review.
const State_InProgress StateId = 2
// Suppressed. The finding was reviewed, considered as a false positive and is now suppressed.
const State_Suppressed StateId = 3
// Resolved. The finding was reviewed and remediated and is now considered resolved.
const State_Resolved StateId = 4
// Other. The state is not mapped. See the <code>state</code> attribute, which contains a data source specific value.
const State_Other StateId = 99

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

// SecurityFinding_Unknown
const Type_SecurityFinding_Unknown TypeUid = 200100
// SecurityFinding_Create. A security finding was created.
const Type_SecurityFinding_Create TypeUid = 200101
// SecurityFinding_Update. A security finding was updated.
const Type_SecurityFinding_Update TypeUid = 200102
// SecurityFinding_Close. A security finding was closed.
const Type_SecurityFinding_Close TypeUid = 200103
// SecurityFinding_Other
const Type_SecurityFinding_Other TypeUid = 200199

