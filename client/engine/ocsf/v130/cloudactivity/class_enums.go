package cloudactivity

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Authorize. An authorizaon check was performed.
const Activity_Authorize ActivityId = 1
// Create. A resource was created.
const Activity_Create ActivityId = 2
// Update. One or more web resources were updated.
const Activity_Update ActivityId = 3
// Delete. One or more web resources were deleted.
const Activity_Delete ActivityId = 4
// Send. A resource was sent to an external destination.
const Activity_Send ActivityId = 5
// Upload. A resource was uploaded from an external source.
const Activity_Upload ActivityId = 6
// Download. A resource was downloaded to an external source.
const Activity_Download ActivityId = 7
// MetadataChange. The metadata of a resource was changed.
const Activity_MetadataChange ActivityId = 8
// SharingChange. The sharing status of a resource was changed.
const Activity_SharingChange ActivityId = 9
// Read. A resource was accessed.
const Activity_Read ActivityId = 10
// Install. An application was installed.
const Activity_Install ActivityId = 11
// Uninstall. An application was uninstalled.
const Activity_Uninstall ActivityId = 12
// MembershipChange. The membership relating to a resource was changed.
const Activity_MembershipChange ActivityId = 13
// Informational. An informational event was logged.
const Activity_Informational ActivityId = 14
// Export. A resource was exported.
const Activity_Export ActivityId = 15
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// ApplicationActivity. Application Activity events report detailed information about the behavior of applications and services.
const Category_ApplicationActivity CategoryUid = 6

// ClassUid Values
// Class ID. The unique identifier of a class. A class describes the attributes available in an event.

// CloudActivity. Cloud activity events report events and actions logged or collected from a cloud environment. This class extends Web Resources Activity with additional fields for cloud-specific metadata.
const Class_CloudActivity ClassUid = 6090

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

// CloudActivity_Unknown
const Type_CloudActivity_Unknown TypeUid = 609000
// CloudActivity_Authorize. An authorizaon check was performed.
const Type_CloudActivity_Authorize TypeUid = 609001
// CloudActivity_Create. A resource was created.
const Type_CloudActivity_Create TypeUid = 609002
// CloudActivity_Update. One or more web resources were updated.
const Type_CloudActivity_Update TypeUid = 609003
// CloudActivity_Delete. One or more web resources were deleted.
const Type_CloudActivity_Delete TypeUid = 609004
// CloudActivity_Send. A resource was sent to an external destination.
const Type_CloudActivity_Send TypeUid = 609005
// CloudActivity_Upload. A resource was uploaded from an external source.
const Type_CloudActivity_Upload TypeUid = 609006
// CloudActivity_Download. A resource was downloaded to an external source.
const Type_CloudActivity_Download TypeUid = 609007
// CloudActivity_MetadataChange. The metadata of a resource was changed.
const Type_CloudActivity_MetadataChange TypeUid = 609008
// CloudActivity_SharingChange. The sharing status of a resource was changed.
const Type_CloudActivity_SharingChange TypeUid = 609009
// CloudActivity_Read. A resource was accessed.
const Type_CloudActivity_Read TypeUid = 609010
// CloudActivity_Install. An application was installed.
const Type_CloudActivity_Install TypeUid = 609011
// CloudActivity_Uninstall. An application was uninstalled.
const Type_CloudActivity_Uninstall TypeUid = 609012
// CloudActivity_MembershipChange. The membership relating to a resource was changed.
const Type_CloudActivity_MembershipChange TypeUid = 609013
// CloudActivity_Informational. An informational event was logged.
const Type_CloudActivity_Informational TypeUid = 609014
// CloudActivity_Export. A resource was exported.
const Type_CloudActivity_Export TypeUid = 609015
// CloudActivity_Other
const Type_CloudActivity_Other TypeUid = 609099

