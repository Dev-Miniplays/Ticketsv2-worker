package messagequeue

import (
	"context"

	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/cache"
	cmdcontext "github.com/Dev-Miniplays/Ticketsv2-worker/bot/command/context"
	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/constants"
	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/dbclient"
	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/logic"
	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/metrics/statsd"
	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/redis"
	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/utils"
	"github.com/TicketsBot-cloud/common/autoclose"
	"github.com/TicketsBot-cloud/common/sentry"
	gdlUtils "github.com/rxdn/gdl/utils"
)

const AutoCloseReason = "Automatically closed due to inactivity"

func ListenAutoClose() {
	ch := make(chan autoclose.Ticket)
	go autoclose.Listen(redis.Client, ch)

	for ticket := range ch {
		statsd.Client.IncrementKey(statsd.AutoClose)

		ticket := ticket
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), constants.TimeoutCloseTicket)
			defer cancel()

			// get ticket
			ticket, err := dbclient.Client.Tickets.Get(ctx, ticket.TicketId, ticket.GuildId)
			if err != nil {
				sentry.Error(err)
				return
			}

			// get worker
			worker, err := buildContext(ctx, ticket, cache.Client)
			if err != nil {
				sentry.Error(err)
				return
			}

			// query already checks, but just to be sure
			if ticket.ChannelId == nil {
				return
			}

			// get premium status
			premiumTier, err := utils.PremiumClient.GetTierByGuildId(ctx, ticket.GuildId, true, worker.Token, worker.RateLimiter)
			if err != nil {
				sentry.Error(err)
				return
			}

			cc := cmdcontext.NewAutoCloseContext(ctx, worker, ticket.GuildId, *ticket.ChannelId, worker.BotId, premiumTier)
			logic.CloseTicket(ctx, cc, gdlUtils.StrPtr(AutoCloseReason), true)
		}()
	}
}
