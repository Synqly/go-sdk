package chatmessageactivity

// ActivityId Values
// Activity ID. The normalized identifier of the activity that triggered the event.

// Unknown. The event activity is unknown.
const Activity_Unknown ActivityId = 0
// Send. A message was sent by a user to a conversation.
const Activity_Send ActivityId = 1
// Receive. A message was received (e.g. an AI-generated response).
const Activity_Receive ActivityId = 2
// Edit. An existing message was edited.
const Activity_Edit ActivityId = 3
// Delete. A message was deleted from the conversation.
const Activity_Delete ActivityId = 4
// React. A user added a reaction to a message.
const Activity_React ActivityId = 5
// SystemEvent. A system event occurred in the conversation (e.g. member added, settings changed).
const Activity_SystemEvent ActivityId = 6
// Other. The event activity is not mapped. See the <code>activity_name</code> attribute, which contains a data source specific value.
const Activity_Other ActivityId = 99

// CategoryUid Values
// Category ID. The category unique identifier of the event.

// ApplicationActivity. Application Activity events report detailed information about the behavior of applications and services.
const Category_ApplicationActivity CategoryUid = 6

// ClassUid Values
// Class ID. The unique identifier of a class. A class describes the attributes available in an event.

// ChatMessageActivity. Chat Message Activity events report messages sent and received in enterprise chat platforms such as Microsoft Teams, Slack, or AI assistants like Microsoft 365 Copilot.
const Class_ChatMessageActivity ClassUid = 98906002

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

// ChatMessageActivity_Unknown
const Type_ChatMessageActivity_Unknown TypeUid = 9890600200
// ChatMessageActivity_Send. A message was sent by a user to a conversation.
const Type_ChatMessageActivity_Send TypeUid = 9890600201
// ChatMessageActivity_Receive. A message was received (e.g. an AI-generated response).
const Type_ChatMessageActivity_Receive TypeUid = 9890600202
// ChatMessageActivity_Edit. An existing message was edited.
const Type_ChatMessageActivity_Edit TypeUid = 9890600203
// ChatMessageActivity_Delete. A message was deleted from the conversation.
const Type_ChatMessageActivity_Delete TypeUid = 9890600204
// ChatMessageActivity_React. A user added a reaction to a message.
const Type_ChatMessageActivity_React TypeUid = 9890600205
// ChatMessageActivity_SystemEvent. A system event occurred in the conversation (e.g. member added, settings changed).
const Type_ChatMessageActivity_SystemEvent TypeUid = 9890600206
// ChatMessageActivity_Other
const Type_ChatMessageActivity_Other TypeUid = 9890600299

