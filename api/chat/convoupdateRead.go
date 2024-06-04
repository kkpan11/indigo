// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package chat

// schema: chat.bsky.convo.updateRead

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// ConvoUpdateRead_Input is the input argument to a chat.bsky.convo.updateRead call.
type ConvoUpdateRead_Input struct {
	ConvoId   string  `json:"convoId" cborgen:"convoId"`
	MessageId *string `json:"messageId,omitempty" cborgen:"messageId,omitempty"`
}

// ConvoUpdateRead_Output is the output of a chat.bsky.convo.updateRead call.
type ConvoUpdateRead_Output struct {
	Convo *ConvoDefs_ConvoView `json:"convo" cborgen:"convo"`
}

// ConvoUpdateRead calls the XRPC method "chat.bsky.convo.updateRead".
func ConvoUpdateRead(ctx context.Context, c *xrpc.Client, input *ConvoUpdateRead_Input) (*ConvoUpdateRead_Output, error) {
	var out ConvoUpdateRead_Output
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "chat.bsky.convo.updateRead", nil, input, &out); err != nil {
		return nil, err
	}

	return &out, nil
}