// (c) 2019 Dapper Labs - ALL RIGHTS RESERVED

package validator

import (
	"context"
	"errors"
	"fmt"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/rs/zerolog"

	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/network/channels"
	"github.com/onflow/flow-go/network/message"
	"github.com/onflow/flow-go/network/slashing"
)

var (
	ErrUnauthorizedSender = errors.New("validation failed: sender is not authorized to send this message type")
	ErrSenderEjected      = errors.New("validation failed: sender is an ejected node")
	ErrUnknownMessageType = errors.New("validation failed: failed to get message auth config")
	ErrIdentityUnverified = errors.New("validation failed: could not verify identity of sender")
)

// AuthorizedSenderValidator returns a MessageValidator that will check if the sender of a message is authorized to send the message.
// The MessageValidator returned will use the getIdentity to get the flow identity for the sender, asserting that the sender is a staked node and not ejected. Otherwise, the message is rejected.
// The message is also authorized by checking that the sender is allowed to send the message on the channel.
// If validation fails the message is rejected, and if the validation error is an expected error, slashing data is also collected.
// Authorization config is defined in message.MsgAuthConfig
func AuthorizedSenderValidator(log zerolog.Logger, channel channels.Channel, getIdentity func(peer.ID) (*flow.Identity, bool)) MessageValidator {
	log = log.With().
		Str("component", "authorized_sender_validator").
		Str("network_channel", channel.String()).
		Logger()

	slashingViolationsConsumer := slashing.NewSlashingViolationsConsumer(log)

	return func(ctx context.Context, from peer.ID, msg interface{}) pubsub.ValidationResult {
		identity, ok := getIdentity(from)
		if !ok {
			log.Warn().Err(ErrIdentityUnverified).Str("peer_id", from.String()).Msg("rejecting message")
			return pubsub.ValidationReject
		}

		msgType, err := IsAuthorizedSender(identity, channel, msg)
		switch {
		case errors.Is(err, ErrUnauthorizedSender):
			slashingViolationsConsumer.OnUnAuthorizedSenderError(identity, from.String(), msgType, err)
			return pubsub.ValidationReject
		case errors.Is(err, ErrUnknownMessageType):
			slashingViolationsConsumer.OnUnknownMsgTypeError(identity, from.String(), msgType, err)
			return pubsub.ValidationReject
		case errors.Is(err, ErrSenderEjected):
			slashingViolationsConsumer.OnSenderEjectedError(identity, from.String(), msgType, err)
			return pubsub.ValidationReject
		case err != nil:
			log.Error().
				Err(err).
				Str("peer_id", from.String()).
				Str("role", identity.Role.String()).
				Str("peer_node_id", identity.NodeID.String()).
				Str("message_type", msgType).
				Msg("unexpected error during message validation")
			return pubsub.ValidationReject
		default:
			return pubsub.ValidationAccept
		}
	}
}

// IsAuthorizedSender performs network authorization validation. This func will assert the following;
// 1. The node is not ejected.
// 2. Using the message auth config
//  A. The message is authorized to be sent on channel.
//  B. The sender role is authorized to send message on channel.
// Expected error returns during normal operations:
//  * ErrSenderEjected: if identity of sender is ejected from the network
//  * ErrUnknownMessageType: if the message type does not have an auth config
//  * ErrUnauthorizedSender: if the sender is not authorized to send message on the channel
func IsAuthorizedSender(identity *flow.Identity, channel channels.Channel, msg interface{}) (string, error) {
	if identity.Ejected {
		return "", ErrSenderEjected
	}

	// get message auth config
	conf, err := message.GetMessageAuthConfig(msg)
	if err != nil {
		return "", fmt.Errorf("%s: %w", err, ErrUnknownMessageType)
	}

	// handle special case for cluster prefixed channels
	if prefix, ok := channels.ClusterChannelPrefix(channel); ok {
		channel = channels.Channel(prefix)
	}

	if err := conf.IsAuthorized(identity.Role, channel); err != nil {
		return conf.Name, fmt.Errorf("%w: %s", err, ErrUnauthorizedSender)
	}

	return conf.Name, nil
}
