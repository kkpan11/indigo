// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.admin.emitModerationEvent

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bluesky-social/indigo/lex/util"
	"github.com/bluesky-social/indigo/xrpc"
)

// AdminEmitModerationEvent_Input is the input argument to a com.atproto.admin.emitModerationEvent call.
type AdminEmitModerationEvent_Input struct {
	CreatedBy       string                                  `json:"createdBy" cborgen:"createdBy"`
	Event           *AdminEmitModerationEvent_Input_Event   `json:"event" cborgen:"event"`
	Subject         *AdminEmitModerationEvent_Input_Subject `json:"subject" cborgen:"subject"`
	SubjectBlobCids []string                                `json:"subjectBlobCids,omitempty" cborgen:"subjectBlobCids,omitempty"`
}

type AdminEmitModerationEvent_Input_Event struct {
	AdminDefs_ModEventTakedown        *AdminDefs_ModEventTakedown
	AdminDefs_ModEventAcknowledge     *AdminDefs_ModEventAcknowledge
	AdminDefs_ModEventEscalate        *AdminDefs_ModEventEscalate
	AdminDefs_ModEventComment         *AdminDefs_ModEventComment
	AdminDefs_ModEventLabel           *AdminDefs_ModEventLabel
	AdminDefs_ModEventReport          *AdminDefs_ModEventReport
	AdminDefs_ModEventMute            *AdminDefs_ModEventMute
	AdminDefs_ModEventReverseTakedown *AdminDefs_ModEventReverseTakedown
	AdminDefs_ModEventUnmute          *AdminDefs_ModEventUnmute
	AdminDefs_ModEventEmail           *AdminDefs_ModEventEmail
}

func (t *AdminEmitModerationEvent_Input_Event) MarshalJSON() ([]byte, error) {
	if t.AdminDefs_ModEventTakedown != nil {
		t.AdminDefs_ModEventTakedown.LexiconTypeID = "com.atproto.admin.defs#modEventTakedown"
		return json.Marshal(t.AdminDefs_ModEventTakedown)
	}
	if t.AdminDefs_ModEventAcknowledge != nil {
		t.AdminDefs_ModEventAcknowledge.LexiconTypeID = "com.atproto.admin.defs#modEventAcknowledge"
		return json.Marshal(t.AdminDefs_ModEventAcknowledge)
	}
	if t.AdminDefs_ModEventEscalate != nil {
		t.AdminDefs_ModEventEscalate.LexiconTypeID = "com.atproto.admin.defs#modEventEscalate"
		return json.Marshal(t.AdminDefs_ModEventEscalate)
	}
	if t.AdminDefs_ModEventComment != nil {
		t.AdminDefs_ModEventComment.LexiconTypeID = "com.atproto.admin.defs#modEventComment"
		return json.Marshal(t.AdminDefs_ModEventComment)
	}
	if t.AdminDefs_ModEventLabel != nil {
		t.AdminDefs_ModEventLabel.LexiconTypeID = "com.atproto.admin.defs#modEventLabel"
		return json.Marshal(t.AdminDefs_ModEventLabel)
	}
	if t.AdminDefs_ModEventReport != nil {
		t.AdminDefs_ModEventReport.LexiconTypeID = "com.atproto.admin.defs#modEventReport"
		return json.Marshal(t.AdminDefs_ModEventReport)
	}
	if t.AdminDefs_ModEventMute != nil {
		t.AdminDefs_ModEventMute.LexiconTypeID = "com.atproto.admin.defs#modEventMute"
		return json.Marshal(t.AdminDefs_ModEventMute)
	}
	if t.AdminDefs_ModEventReverseTakedown != nil {
		t.AdminDefs_ModEventReverseTakedown.LexiconTypeID = "com.atproto.admin.defs#modEventReverseTakedown"
		return json.Marshal(t.AdminDefs_ModEventReverseTakedown)
	}
	if t.AdminDefs_ModEventUnmute != nil {
		t.AdminDefs_ModEventUnmute.LexiconTypeID = "com.atproto.admin.defs#modEventUnmute"
		return json.Marshal(t.AdminDefs_ModEventUnmute)
	}
	if t.AdminDefs_ModEventEmail != nil {
		t.AdminDefs_ModEventEmail.LexiconTypeID = "com.atproto.admin.defs#modEventEmail"
		return json.Marshal(t.AdminDefs_ModEventEmail)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *AdminEmitModerationEvent_Input_Event) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "com.atproto.admin.defs#modEventTakedown":
		t.AdminDefs_ModEventTakedown = new(AdminDefs_ModEventTakedown)
		return json.Unmarshal(b, t.AdminDefs_ModEventTakedown)
	case "com.atproto.admin.defs#modEventAcknowledge":
		t.AdminDefs_ModEventAcknowledge = new(AdminDefs_ModEventAcknowledge)
		return json.Unmarshal(b, t.AdminDefs_ModEventAcknowledge)
	case "com.atproto.admin.defs#modEventEscalate":
		t.AdminDefs_ModEventEscalate = new(AdminDefs_ModEventEscalate)
		return json.Unmarshal(b, t.AdminDefs_ModEventEscalate)
	case "com.atproto.admin.defs#modEventComment":
		t.AdminDefs_ModEventComment = new(AdminDefs_ModEventComment)
		return json.Unmarshal(b, t.AdminDefs_ModEventComment)
	case "com.atproto.admin.defs#modEventLabel":
		t.AdminDefs_ModEventLabel = new(AdminDefs_ModEventLabel)
		return json.Unmarshal(b, t.AdminDefs_ModEventLabel)
	case "com.atproto.admin.defs#modEventReport":
		t.AdminDefs_ModEventReport = new(AdminDefs_ModEventReport)
		return json.Unmarshal(b, t.AdminDefs_ModEventReport)
	case "com.atproto.admin.defs#modEventMute":
		t.AdminDefs_ModEventMute = new(AdminDefs_ModEventMute)
		return json.Unmarshal(b, t.AdminDefs_ModEventMute)
	case "com.atproto.admin.defs#modEventReverseTakedown":
		t.AdminDefs_ModEventReverseTakedown = new(AdminDefs_ModEventReverseTakedown)
		return json.Unmarshal(b, t.AdminDefs_ModEventReverseTakedown)
	case "com.atproto.admin.defs#modEventUnmute":
		t.AdminDefs_ModEventUnmute = new(AdminDefs_ModEventUnmute)
		return json.Unmarshal(b, t.AdminDefs_ModEventUnmute)
	case "com.atproto.admin.defs#modEventEmail":
		t.AdminDefs_ModEventEmail = new(AdminDefs_ModEventEmail)
		return json.Unmarshal(b, t.AdminDefs_ModEventEmail)

	default:
		return nil
	}
}

type AdminEmitModerationEvent_Input_Subject struct {
	AdminDefs_RepoRef *AdminDefs_RepoRef
	RepoStrongRef     *RepoStrongRef
}

func (t *AdminEmitModerationEvent_Input_Subject) MarshalJSON() ([]byte, error) {
	if t.AdminDefs_RepoRef != nil {
		t.AdminDefs_RepoRef.LexiconTypeID = "com.atproto.admin.defs#repoRef"
		return json.Marshal(t.AdminDefs_RepoRef)
	}
	if t.RepoStrongRef != nil {
		t.RepoStrongRef.LexiconTypeID = "com.atproto.repo.strongRef"
		return json.Marshal(t.RepoStrongRef)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *AdminEmitModerationEvent_Input_Subject) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "com.atproto.admin.defs#repoRef":
		t.AdminDefs_RepoRef = new(AdminDefs_RepoRef)
		return json.Unmarshal(b, t.AdminDefs_RepoRef)
	case "com.atproto.repo.strongRef":
		t.RepoStrongRef = new(RepoStrongRef)
		return json.Unmarshal(b, t.RepoStrongRef)

	default:
		return nil
	}
}

// AdminEmitModerationEvent calls the XRPC method "com.atproto.admin.emitModerationEvent".
func AdminEmitModerationEvent(ctx context.Context, c *xrpc.Client, input *AdminEmitModerationEvent_Input) (*AdminDefs_ModEventView, error) {
	var out AdminDefs_ModEventView
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "com.atproto.admin.emitModerationEvent", nil, input, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
