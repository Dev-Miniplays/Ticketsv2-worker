package statistics

import (
	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/command"
	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/command/registry"
	"github.com/Dev-Miniplays/Ticketsv2-worker/i18n"
	"github.com/TicketsBot/common/permission"
	"github.com/rxdn/gdl/objects/interaction"
)

type StatsCommand struct {
}

func (StatsCommand) Properties() registry.Properties {
	return registry.Properties{
		Name:            "stats",
		Description:     i18n.HelpStats,
		Type:            interaction.ApplicationCommandTypeChatInput,
		Aliases:         []string{"statistics"},
		PermissionLevel: permission.Support,
		Children: []registry.Command{
			StatsUserCommand{},
			StatsServerCommand{},
		},
		Category:    command.Statistics,
		PremiumOnly: true,
	}
}

func (c StatsCommand) GetExecutor() interface{} {
	return c.Execute
}

func (StatsCommand) Execute(ctx registry.CommandContext) {
	// Cannot call parent command
}
