// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
	"github.com/sst/opencode-sdk-go/shared"
	"github.com/tidwall/gjson"
)

// SessionService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionService] method instead.
type SessionService struct {
	Options []option.RequestOption
}

// NewSessionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSessionService(opts ...option.RequestOption) (r *SessionService) {
	r = &SessionService{}
	r.Options = opts
	return
}

// Create a new session
func (r *SessionService) New(ctx context.Context, opts ...option.RequestOption) (res *Session, err error) {
	opts = append(r.Options[:], opts...)
	path := "session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// List all sessions
func (r *SessionService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Session, err error) {
	opts = append(r.Options[:], opts...)
	path := "session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a session and all its data
func (r *SessionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *bool, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Abort a session
func (r *SessionService) Abort(ctx context.Context, id string, opts ...option.RequestOption) (res *bool, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/abort", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Create and send a new message to a session
func (r *SessionService) Chat(ctx context.Context, id string, body SessionChatParams, opts ...option.RequestOption) (res *AssistantMessage, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/message", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Analyze the app and create an AGENTS.md file
func (r *SessionService) Init(ctx context.Context, id string, body SessionInitParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/init", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List messages for a session
func (r *SessionService) Messages(ctx context.Context, id string, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/message", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Share a session
func (r *SessionService) Share(ctx context.Context, id string, opts ...option.RequestOption) (res *Session, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/share", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Summarize the session
func (r *SessionService) Summarize(ctx context.Context, id string, body SessionSummarizeParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/summarize", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unshare the session
func (r *SessionService) Unshare(ctx context.Context, id string, opts ...option.RequestOption) (res *Session, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/share", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AssistantMessage struct {
	ID         string                 `json:"id,required"`
	Cost       float64                `json:"cost,required"`
	ModelID    string                 `json:"modelID,required"`
	Parts      []AssistantMessagePart `json:"parts,required"`
	Path       AssistantMessagePath   `json:"path,required"`
	ProviderID string                 `json:"providerID,required"`
	Role       AssistantMessageRole   `json:"role,required"`
	SessionID  string                 `json:"sessionID,required"`
	System     []string               `json:"system,required"`
	Time       AssistantMessageTime   `json:"time,required"`
	Tokens     AssistantMessageTokens `json:"tokens,required"`
	Error      AssistantMessageError  `json:"error"`
	Summary    bool                   `json:"summary"`
	JSON       assistantMessageJSON   `json:"-"`
}

// assistantMessageJSON contains the JSON metadata for the struct
// [AssistantMessage]
type assistantMessageJSON struct {
	ID          apijson.Field
	Cost        apijson.Field
	ModelID     apijson.Field
	Parts       apijson.Field
	Path        apijson.Field
	ProviderID  apijson.Field
	Role        apijson.Field
	SessionID   apijson.Field
	System      apijson.Field
	Time        apijson.Field
	Tokens      apijson.Field
	Error       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssistantMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantMessageJSON) RawJSON() string {
	return r.raw
}

func (r AssistantMessage) implementsMessage() {}

type AssistantMessagePart struct {
	Type AssistantMessagePartsType `json:"type,required"`
	ID   string                    `json:"id"`
	// This field can have the runtime type of [ToolPartState].
	State interface{}              `json:"state"`
	Text  string                   `json:"text"`
	Tool  string                   `json:"tool"`
	JSON  assistantMessagePartJSON `json:"-"`
	union AssistantMessagePartsUnion
}

// assistantMessagePartJSON contains the JSON metadata for the struct
// [AssistantMessagePart]
type assistantMessagePartJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	State       apijson.Field
	Text        apijson.Field
	Tool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r assistantMessagePartJSON) RawJSON() string {
	return r.raw
}

func (r *AssistantMessagePart) UnmarshalJSON(data []byte) (err error) {
	*r = AssistantMessagePart{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AssistantMessagePartsUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [TextPart], [ToolPart],
// [AssistantMessagePartsStepStartPart].
func (r AssistantMessagePart) AsUnion() AssistantMessagePartsUnion {
	return r.union
}

// Union satisfied by [TextPart], [ToolPart] or
// [AssistantMessagePartsStepStartPart].
type AssistantMessagePartsUnion interface {
	implementsAssistantMessagePart()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AssistantMessagePartsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TextPart{}),
			DiscriminatorValue: "text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ToolPart{}),
			DiscriminatorValue: "tool",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AssistantMessagePartsStepStartPart{}),
			DiscriminatorValue: "step-start",
		},
	)
}

type AssistantMessagePartsStepStartPart struct {
	Type AssistantMessagePartsStepStartPartType `json:"type,required"`
	JSON assistantMessagePartsStepStartPartJSON `json:"-"`
}

// assistantMessagePartsStepStartPartJSON contains the JSON metadata for the struct
// [AssistantMessagePartsStepStartPart]
type assistantMessagePartsStepStartPartJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssistantMessagePartsStepStartPart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantMessagePartsStepStartPartJSON) RawJSON() string {
	return r.raw
}

func (r AssistantMessagePartsStepStartPart) implementsAssistantMessagePart() {}

type AssistantMessagePartsStepStartPartType string

const (
	AssistantMessagePartsStepStartPartTypeStepStart AssistantMessagePartsStepStartPartType = "step-start"
)

func (r AssistantMessagePartsStepStartPartType) IsKnown() bool {
	switch r {
	case AssistantMessagePartsStepStartPartTypeStepStart:
		return true
	}
	return false
}

type AssistantMessagePartsType string

const (
	AssistantMessagePartsTypeText      AssistantMessagePartsType = "text"
	AssistantMessagePartsTypeTool      AssistantMessagePartsType = "tool"
	AssistantMessagePartsTypeStepStart AssistantMessagePartsType = "step-start"
)

func (r AssistantMessagePartsType) IsKnown() bool {
	switch r {
	case AssistantMessagePartsTypeText, AssistantMessagePartsTypeTool, AssistantMessagePartsTypeStepStart:
		return true
	}
	return false
}

type AssistantMessagePath struct {
	Cwd  string                   `json:"cwd,required"`
	Root string                   `json:"root,required"`
	JSON assistantMessagePathJSON `json:"-"`
}

// assistantMessagePathJSON contains the JSON metadata for the struct
// [AssistantMessagePath]
type assistantMessagePathJSON struct {
	Cwd         apijson.Field
	Root        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssistantMessagePath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantMessagePathJSON) RawJSON() string {
	return r.raw
}

type AssistantMessageRole string

const (
	AssistantMessageRoleAssistant AssistantMessageRole = "assistant"
)

func (r AssistantMessageRole) IsKnown() bool {
	switch r {
	case AssistantMessageRoleAssistant:
		return true
	}
	return false
}

type AssistantMessageTime struct {
	Created   float64                  `json:"created,required"`
	Completed float64                  `json:"completed"`
	JSON      assistantMessageTimeJSON `json:"-"`
}

// assistantMessageTimeJSON contains the JSON metadata for the struct
// [AssistantMessageTime]
type assistantMessageTimeJSON struct {
	Created     apijson.Field
	Completed   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssistantMessageTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantMessageTimeJSON) RawJSON() string {
	return r.raw
}

type AssistantMessageTokens struct {
	Cache     AssistantMessageTokensCache `json:"cache,required"`
	Input     float64                     `json:"input,required"`
	Output    float64                     `json:"output,required"`
	Reasoning float64                     `json:"reasoning,required"`
	JSON      assistantMessageTokensJSON  `json:"-"`
}

// assistantMessageTokensJSON contains the JSON metadata for the struct
// [AssistantMessageTokens]
type assistantMessageTokensJSON struct {
	Cache       apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	Reasoning   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssistantMessageTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantMessageTokensJSON) RawJSON() string {
	return r.raw
}

type AssistantMessageTokensCache struct {
	Read  float64                         `json:"read,required"`
	Write float64                         `json:"write,required"`
	JSON  assistantMessageTokensCacheJSON `json:"-"`
}

// assistantMessageTokensCacheJSON contains the JSON metadata for the struct
// [AssistantMessageTokensCache]
type assistantMessageTokensCacheJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssistantMessageTokensCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantMessageTokensCacheJSON) RawJSON() string {
	return r.raw
}

type AssistantMessageError struct {
	// This field can have the runtime type of [shared.ProviderAuthErrorData],
	// [shared.UnknownErrorData], [interface{}].
	Data  interface{}               `json:"data,required"`
	Name  AssistantMessageErrorName `json:"name,required"`
	JSON  assistantMessageErrorJSON `json:"-"`
	union AssistantMessageErrorUnion
}

// assistantMessageErrorJSON contains the JSON metadata for the struct
// [AssistantMessageError]
type assistantMessageErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r assistantMessageErrorJSON) RawJSON() string {
	return r.raw
}

func (r *AssistantMessageError) UnmarshalJSON(data []byte) (err error) {
	*r = AssistantMessageError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AssistantMessageErrorUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [shared.ProviderAuthError],
// [shared.UnknownError], [AssistantMessageErrorMessageOutputLengthError].
func (r AssistantMessageError) AsUnion() AssistantMessageErrorUnion {
	return r.union
}

// Union satisfied by [shared.ProviderAuthError], [shared.UnknownError] or
// [AssistantMessageErrorMessageOutputLengthError].
type AssistantMessageErrorUnion interface {
	ImplementsAssistantMessageError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AssistantMessageErrorUnion)(nil)).Elem(),
		"name",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.ProviderAuthError{}),
			DiscriminatorValue: "ProviderAuthError",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.UnknownError{}),
			DiscriminatorValue: "UnknownError",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AssistantMessageErrorMessageOutputLengthError{}),
			DiscriminatorValue: "MessageOutputLengthError",
		},
	)
}

type AssistantMessageErrorMessageOutputLengthError struct {
	Data interface{}                                       `json:"data,required"`
	Name AssistantMessageErrorMessageOutputLengthErrorName `json:"name,required"`
	JSON assistantMessageErrorMessageOutputLengthErrorJSON `json:"-"`
}

// assistantMessageErrorMessageOutputLengthErrorJSON contains the JSON metadata for
// the struct [AssistantMessageErrorMessageOutputLengthError]
type assistantMessageErrorMessageOutputLengthErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssistantMessageErrorMessageOutputLengthError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantMessageErrorMessageOutputLengthErrorJSON) RawJSON() string {
	return r.raw
}

func (r AssistantMessageErrorMessageOutputLengthError) ImplementsAssistantMessageError() {}

type AssistantMessageErrorMessageOutputLengthErrorName string

const (
	AssistantMessageErrorMessageOutputLengthErrorNameMessageOutputLengthError AssistantMessageErrorMessageOutputLengthErrorName = "MessageOutputLengthError"
)

func (r AssistantMessageErrorMessageOutputLengthErrorName) IsKnown() bool {
	switch r {
	case AssistantMessageErrorMessageOutputLengthErrorNameMessageOutputLengthError:
		return true
	}
	return false
}

type AssistantMessageErrorName string

const (
	AssistantMessageErrorNameProviderAuthError        AssistantMessageErrorName = "ProviderAuthError"
	AssistantMessageErrorNameUnknownError             AssistantMessageErrorName = "UnknownError"
	AssistantMessageErrorNameMessageOutputLengthError AssistantMessageErrorName = "MessageOutputLengthError"
)

func (r AssistantMessageErrorName) IsKnown() bool {
	switch r {
	case AssistantMessageErrorNameProviderAuthError, AssistantMessageErrorNameUnknownError, AssistantMessageErrorNameMessageOutputLengthError:
		return true
	}
	return false
}

type FilePart struct {
	Mime     string       `json:"mime,required"`
	Type     FilePartType `json:"type,required"`
	URL      string       `json:"url,required"`
	Filename string       `json:"filename"`
	JSON     filePartJSON `json:"-"`
}

// filePartJSON contains the JSON metadata for the struct [FilePart]
type filePartJSON struct {
	Mime        apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	Filename    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilePart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filePartJSON) RawJSON() string {
	return r.raw
}

func (r FilePart) implementsMessageUserMessagePart() {}

type FilePartType string

const (
	FilePartTypeFile FilePartType = "file"
)

func (r FilePartType) IsKnown() bool {
	switch r {
	case FilePartTypeFile:
		return true
	}
	return false
}

type FilePartParam struct {
	Mime     param.Field[string]       `json:"mime,required"`
	Type     param.Field[FilePartType] `json:"type,required"`
	URL      param.Field[string]       `json:"url,required"`
	Filename param.Field[string]       `json:"filename"`
}

func (r FilePartParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FilePartParam) implementsSessionChatParamsPartUnion() {}

type Message struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [[]MessageUserMessagePart],
	// [[]AssistantMessagePart].
	Parts     interface{} `json:"parts,required"`
	Role      MessageRole `json:"role,required"`
	SessionID string      `json:"sessionID,required"`
	// This field can have the runtime type of [MessageUserMessageTime],
	// [AssistantMessageTime].
	Time interface{} `json:"time,required"`
	Cost float64     `json:"cost"`
	// This field can have the runtime type of [AssistantMessageError].
	Error   interface{} `json:"error"`
	ModelID string      `json:"modelID"`
	// This field can have the runtime type of [AssistantMessagePath].
	Path       interface{} `json:"path"`
	ProviderID string      `json:"providerID"`
	Summary    bool        `json:"summary"`
	// This field can have the runtime type of [[]string].
	System interface{} `json:"system"`
	// This field can have the runtime type of [AssistantMessageTokens].
	Tokens interface{} `json:"tokens"`
	JSON   messageJSON `json:"-"`
	union  MessageUnion
}

// messageJSON contains the JSON metadata for the struct [Message]
type messageJSON struct {
	ID          apijson.Field
	Parts       apijson.Field
	Role        apijson.Field
	SessionID   apijson.Field
	Time        apijson.Field
	Cost        apijson.Field
	Error       apijson.Field
	ModelID     apijson.Field
	Path        apijson.Field
	ProviderID  apijson.Field
	Summary     apijson.Field
	System      apijson.Field
	Tokens      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r messageJSON) RawJSON() string {
	return r.raw
}

func (r *Message) UnmarshalJSON(data []byte) (err error) {
	*r = Message{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MessageUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [MessageUserMessage],
// [AssistantMessage].
func (r Message) AsUnion() MessageUnion {
	return r.union
}

// Union satisfied by [MessageUserMessage] or [AssistantMessage].
type MessageUnion interface {
	implementsMessage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageUnion)(nil)).Elem(),
		"role",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageUserMessage{}),
			DiscriminatorValue: "user",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AssistantMessage{}),
			DiscriminatorValue: "assistant",
		},
	)
}

type MessageUserMessage struct {
	ID        string                   `json:"id,required"`
	Parts     []MessageUserMessagePart `json:"parts,required"`
	Role      MessageUserMessageRole   `json:"role,required"`
	SessionID string                   `json:"sessionID,required"`
	Time      MessageUserMessageTime   `json:"time,required"`
	JSON      messageUserMessageJSON   `json:"-"`
}

// messageUserMessageJSON contains the JSON metadata for the struct
// [MessageUserMessage]
type messageUserMessageJSON struct {
	ID          apijson.Field
	Parts       apijson.Field
	Role        apijson.Field
	SessionID   apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageUserMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageUserMessageJSON) RawJSON() string {
	return r.raw
}

func (r MessageUserMessage) implementsMessage() {}

type MessageUserMessagePart struct {
	Type     MessageUserMessagePartsType `json:"type,required"`
	Filename string                      `json:"filename"`
	Mime     string                      `json:"mime"`
	Text     string                      `json:"text"`
	URL      string                      `json:"url"`
	JSON     messageUserMessagePartJSON  `json:"-"`
	union    MessageUserMessagePartsUnion
}

// messageUserMessagePartJSON contains the JSON metadata for the struct
// [MessageUserMessagePart]
type messageUserMessagePartJSON struct {
	Type        apijson.Field
	Filename    apijson.Field
	Mime        apijson.Field
	Text        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r messageUserMessagePartJSON) RawJSON() string {
	return r.raw
}

func (r *MessageUserMessagePart) UnmarshalJSON(data []byte) (err error) {
	*r = MessageUserMessagePart{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MessageUserMessagePartsUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [TextPart], [FilePart].
func (r MessageUserMessagePart) AsUnion() MessageUserMessagePartsUnion {
	return r.union
}

// Union satisfied by [TextPart] or [FilePart].
type MessageUserMessagePartsUnion interface {
	implementsMessageUserMessagePart()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageUserMessagePartsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TextPart{}),
			DiscriminatorValue: "text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(FilePart{}),
			DiscriminatorValue: "file",
		},
	)
}

type MessageUserMessagePartsType string

const (
	MessageUserMessagePartsTypeText MessageUserMessagePartsType = "text"
	MessageUserMessagePartsTypeFile MessageUserMessagePartsType = "file"
)

func (r MessageUserMessagePartsType) IsKnown() bool {
	switch r {
	case MessageUserMessagePartsTypeText, MessageUserMessagePartsTypeFile:
		return true
	}
	return false
}

type MessageUserMessageRole string

const (
	MessageUserMessageRoleUser MessageUserMessageRole = "user"
)

func (r MessageUserMessageRole) IsKnown() bool {
	switch r {
	case MessageUserMessageRoleUser:
		return true
	}
	return false
}

type MessageUserMessageTime struct {
	Created float64                    `json:"created,required"`
	JSON    messageUserMessageTimeJSON `json:"-"`
}

// messageUserMessageTimeJSON contains the JSON metadata for the struct
// [MessageUserMessageTime]
type messageUserMessageTimeJSON struct {
	Created     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageUserMessageTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageUserMessageTimeJSON) RawJSON() string {
	return r.raw
}

type MessageRole string

const (
	MessageRoleUser      MessageRole = "user"
	MessageRoleAssistant MessageRole = "assistant"
)

func (r MessageRole) IsKnown() bool {
	switch r {
	case MessageRoleUser, MessageRoleAssistant:
		return true
	}
	return false
}

type Session struct {
	ID       string        `json:"id,required"`
	Time     SessionTime   `json:"time,required"`
	Title    string        `json:"title,required"`
	Version  string        `json:"version,required"`
	ParentID string        `json:"parentID"`
	Revert   SessionRevert `json:"revert"`
	Share    SessionShare  `json:"share"`
	JSON     sessionJSON   `json:"-"`
}

// sessionJSON contains the JSON metadata for the struct [Session]
type sessionJSON struct {
	ID          apijson.Field
	Time        apijson.Field
	Title       apijson.Field
	Version     apijson.Field
	ParentID    apijson.Field
	Revert      apijson.Field
	Share       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Session) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionJSON) RawJSON() string {
	return r.raw
}

type SessionTime struct {
	Created float64         `json:"created,required"`
	Updated float64         `json:"updated,required"`
	JSON    sessionTimeJSON `json:"-"`
}

// sessionTimeJSON contains the JSON metadata for the struct [SessionTime]
type sessionTimeJSON struct {
	Created     apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionTimeJSON) RawJSON() string {
	return r.raw
}

type SessionRevert struct {
	MessageID string            `json:"messageID,required"`
	Part      float64           `json:"part,required"`
	Snapshot  string            `json:"snapshot"`
	JSON      sessionRevertJSON `json:"-"`
}

// sessionRevertJSON contains the JSON metadata for the struct [SessionRevert]
type sessionRevertJSON struct {
	MessageID   apijson.Field
	Part        apijson.Field
	Snapshot    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionRevert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionRevertJSON) RawJSON() string {
	return r.raw
}

type SessionShare struct {
	URL  string           `json:"url,required"`
	JSON sessionShareJSON `json:"-"`
}

// sessionShareJSON contains the JSON metadata for the struct [SessionShare]
type sessionShareJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionShare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionShareJSON) RawJSON() string {
	return r.raw
}

type TextPart struct {
	Text string       `json:"text,required"`
	Type TextPartType `json:"type,required"`
	JSON textPartJSON `json:"-"`
}

// textPartJSON contains the JSON metadata for the struct [TextPart]
type textPartJSON struct {
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TextPart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r textPartJSON) RawJSON() string {
	return r.raw
}

func (r TextPart) implementsEventListResponseEventMessagePartUpdatedPropertiesPart() {}

func (r TextPart) implementsAssistantMessagePart() {}

func (r TextPart) implementsMessageUserMessagePart() {}

type TextPartType string

const (
	TextPartTypeText TextPartType = "text"
)

func (r TextPartType) IsKnown() bool {
	switch r {
	case TextPartTypeText:
		return true
	}
	return false
}

type TextPartParam struct {
	Text param.Field[string]       `json:"text,required"`
	Type param.Field[TextPartType] `json:"type,required"`
}

func (r TextPartParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TextPartParam) implementsSessionChatParamsPartUnion() {}

type ToolPart struct {
	ID    string        `json:"id,required"`
	State ToolPartState `json:"state,required"`
	Tool  string        `json:"tool,required"`
	Type  ToolPartType  `json:"type,required"`
	JSON  toolPartJSON  `json:"-"`
}

// toolPartJSON contains the JSON metadata for the struct [ToolPart]
type toolPartJSON struct {
	ID          apijson.Field
	State       apijson.Field
	Tool        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolPart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolPartJSON) RawJSON() string {
	return r.raw
}

func (r ToolPart) implementsEventListResponseEventMessagePartUpdatedPropertiesPart() {}

func (r ToolPart) implementsAssistantMessagePart() {}

type ToolPartState struct {
	Status ToolPartStateStatus `json:"status,required"`
	// This field can have the runtime type of [interface{}].
	Input interface{} `json:"input"`
	// This field can have the runtime type of [map[string]interface{}].
	Metadata interface{} `json:"metadata"`
	Output   string      `json:"output"`
	// This field can have the runtime type of [ToolPartStateToolStateRunningTime],
	// [ToolPartStateToolInvocationCompletedTime].
	Time  interface{}       `json:"time"`
	Title string            `json:"title"`
	JSON  toolPartStateJSON `json:"-"`
	union ToolPartStateUnion
}

// toolPartStateJSON contains the JSON metadata for the struct [ToolPartState]
type toolPartStateJSON struct {
	Status      apijson.Field
	Input       apijson.Field
	Metadata    apijson.Field
	Output      apijson.Field
	Time        apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r toolPartStateJSON) RawJSON() string {
	return r.raw
}

func (r *ToolPartState) UnmarshalJSON(data []byte) (err error) {
	*r = ToolPartState{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ToolPartStateUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ToolPartStateToolStatePending],
// [ToolPartStateToolStateRunning], [ToolPartStateToolInvocationCompleted].
func (r ToolPartState) AsUnion() ToolPartStateUnion {
	return r.union
}

// Union satisfied by [ToolPartStateToolStatePending],
// [ToolPartStateToolStateRunning] or [ToolPartStateToolInvocationCompleted].
type ToolPartStateUnion interface {
	implementsToolPartState()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ToolPartStateUnion)(nil)).Elem(),
		"status",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ToolPartStateToolStatePending{}),
			DiscriminatorValue: "pending",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ToolPartStateToolStateRunning{}),
			DiscriminatorValue: "running",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ToolPartStateToolInvocationCompleted{}),
			DiscriminatorValue: "completed",
		},
	)
}

type ToolPartStateToolStatePending struct {
	Status ToolPartStateToolStatePendingStatus `json:"status,required"`
	JSON   toolPartStateToolStatePendingJSON   `json:"-"`
}

// toolPartStateToolStatePendingJSON contains the JSON metadata for the struct
// [ToolPartStateToolStatePending]
type toolPartStateToolStatePendingJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolPartStateToolStatePending) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolPartStateToolStatePendingJSON) RawJSON() string {
	return r.raw
}

func (r ToolPartStateToolStatePending) implementsToolPartState() {}

type ToolPartStateToolStatePendingStatus string

const (
	ToolPartStateToolStatePendingStatusPending ToolPartStateToolStatePendingStatus = "pending"
)

func (r ToolPartStateToolStatePendingStatus) IsKnown() bool {
	switch r {
	case ToolPartStateToolStatePendingStatusPending:
		return true
	}
	return false
}

type ToolPartStateToolStateRunning struct {
	Status ToolPartStateToolStateRunningStatus `json:"status,required"`
	Time   ToolPartStateToolStateRunningTime   `json:"time,required"`
	Input  interface{}                         `json:"input"`
	JSON   toolPartStateToolStateRunningJSON   `json:"-"`
}

// toolPartStateToolStateRunningJSON contains the JSON metadata for the struct
// [ToolPartStateToolStateRunning]
type toolPartStateToolStateRunningJSON struct {
	Status      apijson.Field
	Time        apijson.Field
	Input       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolPartStateToolStateRunning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolPartStateToolStateRunningJSON) RawJSON() string {
	return r.raw
}

func (r ToolPartStateToolStateRunning) implementsToolPartState() {}

type ToolPartStateToolStateRunningStatus string

const (
	ToolPartStateToolStateRunningStatusRunning ToolPartStateToolStateRunningStatus = "running"
)

func (r ToolPartStateToolStateRunningStatus) IsKnown() bool {
	switch r {
	case ToolPartStateToolStateRunningStatusRunning:
		return true
	}
	return false
}

type ToolPartStateToolStateRunningTime struct {
	Start float64                               `json:"start,required"`
	JSON  toolPartStateToolStateRunningTimeJSON `json:"-"`
}

// toolPartStateToolStateRunningTimeJSON contains the JSON metadata for the struct
// [ToolPartStateToolStateRunningTime]
type toolPartStateToolStateRunningTimeJSON struct {
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolPartStateToolStateRunningTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolPartStateToolStateRunningTimeJSON) RawJSON() string {
	return r.raw
}

type ToolPartStateToolInvocationCompleted struct {
	Metadata map[string]interface{}                     `json:"metadata,required"`
	Output   string                                     `json:"output,required"`
	Status   ToolPartStateToolInvocationCompletedStatus `json:"status,required"`
	Time     ToolPartStateToolInvocationCompletedTime   `json:"time,required"`
	Title    string                                     `json:"title,required"`
	Input    interface{}                                `json:"input"`
	JSON     toolPartStateToolInvocationCompletedJSON   `json:"-"`
}

// toolPartStateToolInvocationCompletedJSON contains the JSON metadata for the
// struct [ToolPartStateToolInvocationCompleted]
type toolPartStateToolInvocationCompletedJSON struct {
	Metadata    apijson.Field
	Output      apijson.Field
	Status      apijson.Field
	Time        apijson.Field
	Title       apijson.Field
	Input       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolPartStateToolInvocationCompleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolPartStateToolInvocationCompletedJSON) RawJSON() string {
	return r.raw
}

func (r ToolPartStateToolInvocationCompleted) implementsToolPartState() {}

type ToolPartStateToolInvocationCompletedStatus string

const (
	ToolPartStateToolInvocationCompletedStatusCompleted ToolPartStateToolInvocationCompletedStatus = "completed"
)

func (r ToolPartStateToolInvocationCompletedStatus) IsKnown() bool {
	switch r {
	case ToolPartStateToolInvocationCompletedStatusCompleted:
		return true
	}
	return false
}

type ToolPartStateToolInvocationCompletedTime struct {
	End   float64                                      `json:"end,required"`
	Start float64                                      `json:"start,required"`
	JSON  toolPartStateToolInvocationCompletedTimeJSON `json:"-"`
}

// toolPartStateToolInvocationCompletedTimeJSON contains the JSON metadata for the
// struct [ToolPartStateToolInvocationCompletedTime]
type toolPartStateToolInvocationCompletedTimeJSON struct {
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolPartStateToolInvocationCompletedTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolPartStateToolInvocationCompletedTimeJSON) RawJSON() string {
	return r.raw
}

type ToolPartStateStatus string

const (
	ToolPartStateStatusPending   ToolPartStateStatus = "pending"
	ToolPartStateStatusRunning   ToolPartStateStatus = "running"
	ToolPartStateStatusCompleted ToolPartStateStatus = "completed"
)

func (r ToolPartStateStatus) IsKnown() bool {
	switch r {
	case ToolPartStateStatusPending, ToolPartStateStatusRunning, ToolPartStateStatusCompleted:
		return true
	}
	return false
}

type ToolPartType string

const (
	ToolPartTypeTool ToolPartType = "tool"
)

func (r ToolPartType) IsKnown() bool {
	switch r {
	case ToolPartTypeTool:
		return true
	}
	return false
}

type SessionChatParams struct {
	ModelID    param.Field[string]                       `json:"modelID,required"`
	Parts      param.Field[[]SessionChatParamsPartUnion] `json:"parts,required"`
	ProviderID param.Field[string]                       `json:"providerID,required"`
}

func (r SessionChatParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SessionChatParamsPart struct {
	Type     param.Field[SessionChatParamsPartsType] `json:"type,required"`
	Filename param.Field[string]                     `json:"filename"`
	Mime     param.Field[string]                     `json:"mime"`
	Text     param.Field[string]                     `json:"text"`
	URL      param.Field[string]                     `json:"url"`
}

func (r SessionChatParamsPart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SessionChatParamsPart) implementsSessionChatParamsPartUnion() {}

// Satisfied by [TextPartParam], [FilePartParam], [SessionChatParamsPart].
type SessionChatParamsPartUnion interface {
	implementsSessionChatParamsPartUnion()
}

type SessionChatParamsPartsType string

const (
	SessionChatParamsPartsTypeText SessionChatParamsPartsType = "text"
	SessionChatParamsPartsTypeFile SessionChatParamsPartsType = "file"
)

func (r SessionChatParamsPartsType) IsKnown() bool {
	switch r {
	case SessionChatParamsPartsTypeText, SessionChatParamsPartsTypeFile:
		return true
	}
	return false
}

type SessionInitParams struct {
	ModelID    param.Field[string] `json:"modelID,required"`
	ProviderID param.Field[string] `json:"providerID,required"`
}

func (r SessionInitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SessionSummarizeParams struct {
	ModelID    param.Field[string] `json:"modelID,required"`
	ProviderID param.Field[string] `json:"providerID,required"`
}

func (r SessionSummarizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
