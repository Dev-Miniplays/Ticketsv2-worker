package tags

import (
	"time"

	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/command"
	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/command/registry"
	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/customisation"
	"github.com/Dev-Miniplays/Ticketsv2-worker/bot/dbclient"
	"github.com/Dev-Miniplays/Ticketsv2-worker/i18n"
	"github.com/TicketsBot-cloud/common/permission"
	"github.com/rxdn/gdl/objects/interaction"
)

type ManageTagsDeleteCommand struct {
}

func (ManageTagsDeleteCommand) Properties() registry.Properties {
	return registry.Properties{
		Name:            "delete",
		Description:     i18n.HelpTagDelete,
		Type:            interaction.ApplicationCommandTypeChatInput,
		Aliases:         []string{"del", "rm", "remove"},
		PermissionLevel: permission.Support,
		Category:        command.Tags,
		Arguments: command.Arguments(
			command.NewRequiredArgument("id", "ID of the tag to delete", interaction.OptionTypeString, i18n.MessageTagDeleteInvalidArguments),
		),
		DefaultEphemeral: true,
		Timeout:          time.Second * 3,
	}
}

func (c ManageTagsDeleteCommand) GetExecutor() interface{} {
	return c.Execute
}

func (ManageTagsDeleteCommand) Execute(ctx registry.CommandContext, tagId string) {
	exists, err := dbclient.Client.Tag.Exists(ctx, ctx.GuildId(), tagId)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	if !exists {
		ctx.Reply(customisation.Red, i18n.Error, i18n.MessageTagDeleteDoesNotExist, tagId)
		return
	}

	if err := dbclient.Client.Tag.Delete(ctx, ctx.GuildId(), tagId); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.Reply(customisation.Green, i18n.MessageTag, i18n.MessageTagDeleteSuccess, tagId)
}
