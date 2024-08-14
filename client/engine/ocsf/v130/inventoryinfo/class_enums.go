package inventoryinfo

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Log. The discovered information is via a log.
const Activity_Log ActivityId = 1
// Collect. The discovered information is via a collection process.
const Activity_Collect ActivityId = 2
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// Discovery. Discovery events report the existence and state of devices, files, configurations, processes, registry keys, and other objects.
const Category_Discovery CategoryUid = 5

// ClassUid Values
// Class ID. The unique identifier of a class. A class describes the attributes available in an event.

// DeviceInventoryInfo. Device Inventory Info events report device inventory data that is either logged or proactively collected. For example, when collecting device information from a CMDB or running a network sweep of connected devices.
const Class_DeviceInventoryInfo ClassUid = 5001

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

// DeviceInventoryInfo_Unknown
const Type_DeviceInventoryInfo_Unknown TypeUid = 500100
// DeviceInventoryInfo_Log. The discovered information is via a log.
const Type_DeviceInventoryInfo_Log TypeUid = 500101
// DeviceInventoryInfo_Collect. The discovered information is via a collection process.
const Type_DeviceInventoryInfo_Collect TypeUid = 500102
// DeviceInventoryInfo_Other
const Type_DeviceInventoryInfo_Other TypeUid = 500199

