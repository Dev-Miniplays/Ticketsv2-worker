package button

import (
	"context"
	"time"

	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/command"
	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/utils"
	"github.com/Dev-Miniplays/Ticketsv2-worker"
	"github.com/rxdn/gdl/objects/interaction"
	"github.com/rxdn/gdl/rest"
)

type ResponseMessage struct {
	Data command.MessageResponse
}

func (r ResponseMessage) Type() ResponseType {
	return ResponseTypeMessage
}

func (r ResponseMessage) Build() interface{} {
	return interaction.NewResponseChannelMessage(r.Data.IntoApplicationCommandData())
}

func (r ResponseMessage) HandleDeferred(interactionData interaction.InteractionMetadata, worker *worker.Context) error {
	if time.Now().Sub(utils.SnowflakeToTime(interactionData.Id)) > time.Minute*14 {
		return nil
	}

	_, err := rest.CreateFollowupMessage(context.Background(), interactionData.Token, worker.RateLimiter, worker.BotId, r.Data.IntoWebhookBody())
	return err
}
