package tickets

import (
	"context"
	"strconv"
	"time"

	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/command"
	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/command/registry"
	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/dbclient"
	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/logic"
	"github.com/Dev-Miniplays/Ticketsv2-worker/i18n"
	"github.com/TicketsBot/common/permission"
	"github.com/TicketsBot/common/sentry"
	"github.com/rxdn/gdl/objects/interaction"
)

type ReopenCommand struct {
}

func (c ReopenCommand) Properties() registry.Properties {
	return registry.Properties{
		Name:            "reopen",
		Description:     i18n.HelpReopen,
		Type:            interaction.ApplicationCommandTypeChatInput,
		PermissionLevel: permission.Everyone,
		Category:        command.Tickets,
		Arguments: command.Arguments(
			command.NewRequiredAutocompleteableArgument("ticket_id", "ID des Tickets das erneut geöffnet werden soll", interaction.OptionTypeInteger, i18n.MessageInvalidArgument, c.AutoCompleteHandler),
		),
		DefaultEphemeral: true,
		Timeout:          time.Second * 10,
	}
}

func (c ReopenCommand) GetExecutor() interface{} {
	return c.Execute
}

func (ReopenCommand) Execute(ctx registry.CommandContext, ticketId int) {
	logic.ReopenTicket(ctx, ctx, ticketId)
}

func (ReopenCommand) AutoCompleteHandler(data interaction.ApplicationCommandAutoCompleteInteraction, value string) []interaction.ApplicationCommandOptionChoice {
	if data.GuildId.Value == 0 {
		return nil
	}

	if data.Member == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3) // TODO: Propagate contxet
	defer cancel()

	tickets, err := dbclient.Client.Tickets.GetClosedByUserPrefixed(ctx, data.GuildId.Value, data.Member.User.Id, value, 25)
	if err != nil {
		sentry.Error(err)
		return nil
	}

	choices := make([]interaction.ApplicationCommandOptionChoice, len(tickets))
	for i, ticket := range tickets {
		if i >= 25 { // Infallible
			break
		}

		choices[i] = interaction.ApplicationCommandOptionChoice{
			Name:  strconv.Itoa(ticket.Id),
			Value: ticket.Id,
		}
	}

	return choices
}
