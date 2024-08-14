package compliancefinding

// ActivityId Values
// Activity ID. The normalized identifier of the finding activity.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Create. A finding was created.
const Activity_Create ActivityId = 1
// Update. A finding was updated.
const Activity_Update ActivityId = 2
// Close. A finding was closed.
const Activity_Close ActivityId = 3
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// Findings. Findings events report findings, detections, and possible resolutions of malware, anomalies, or other actions performed by security products.
const Category_Findings CategoryUid = 2

// ClassUid Values
// Class ID. The unique identifier of a class. A class describes the attributes available in an event.

// ComplianceFinding. Compliance Finding events describe results of evaluations performed against resources, to check compliance with various Industry Frameworks or Security Standards such as <code>NIST SP 800-53, CIS AWS Foundations Benchmark v1.4.0, ISO/IEC 27001</code> etc.
const Class_ComplianceFinding ClassUid = 2003

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
// Status ID. The normalized status identifier of the Finding, set by the consumer.

// Unknown. The status is unknown.
const Status_Unknown StatusId = 0
// New. The Finding is new and yet to be reviewed.
const Status_New StatusId = 1
// InProgress. The Finding is under review.
const Status_InProgress StatusId = 2
// Suppressed. The Finding was reviewed, determined to be benign or a false positive and is now suppressed.
const Status_Suppressed StatusId = 3
// Resolved. The Finding was reviewed, remediated and is now considered resolved.
const Status_Resolved StatusId = 4
// Other. The event status is not mapped. See the <code>status</code> attribute, which contains a data source specific value.
const Status_Other StatusId = 99

// TypeUid Values
// Type ID. The event/finding type ID. It identifies the event's semantics and structure. The value is calculated by the logging system as: <code>class_uid * 100 + activity_id</code>.

// ComplianceFinding_Unknown
const Type_ComplianceFinding_Unknown TypeUid = 200300
// ComplianceFinding_Create. A finding was created.
const Type_ComplianceFinding_Create TypeUid = 200301
// ComplianceFinding_Update. A finding was updated.
const Type_ComplianceFinding_Update TypeUid = 200302
// ComplianceFinding_Close. A finding was closed.
const Type_ComplianceFinding_Close TypeUid = 200303
// ComplianceFinding_Other
const Type_ComplianceFinding_Other TypeUid = 200399

