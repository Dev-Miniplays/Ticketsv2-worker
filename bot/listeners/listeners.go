// Code generated by /tools/cmd/generatelisteners.go; DO NOT EDIT.
//go:generate go run ../../tools/cmd/generatelisteners.go

package listeners

import (
    "encoding/json"
    "fmt"
    "github.com/Dev-Miniplays/Ticketsv2-worker"
    "github.com/getsentry/sentry-go"
    "github.com/rxdn/gdl/gateway/payloads"
    "github.com/rxdn/gdl/gateway/payloads/events"
)

var (
    
    ChannelCreateListeners = []func(*Ticketsv2-worker.Context, events.ChannelCreate){}
    ChannelDeleteListeners = []func(*Ticketsv2-worker.Context, events.ChannelDelete){}
    ChannelPinsUpdateListeners = []func(*Ticketsv2-worker.Context, events.ChannelPinsUpdate){}
    ChannelUpdateListeners = []func(*Ticketsv2-worker.Context, events.ChannelUpdate){}
    EntitlementCreateListeners = []func(*Ticketsv2-worker.Context, events.EntitlementCreate){}
    EntitlementDeleteListeners = []func(*Ticketsv2-worker.Context, events.EntitlementDelete){}
    EntitlementUpdateListeners = []func(*Ticketsv2-worker.Context, events.EntitlementUpdate){}
    GuildBanAddListeners = []func(*Ticketsv2-worker.Context, events.GuildBanAdd){}
    GuildBanRemoveListeners = []func(*Ticketsv2-worker.Context, events.GuildBanRemove){}
    GuildCreateListeners = []func(*Ticketsv2-worker.Context, events.GuildCreate){}
    GuildDeleteListeners = []func(*Ticketsv2-worker.Context, events.GuildDelete){}
    GuildEmojisUpdateListeners = []func(*Ticketsv2-worker.Context, events.GuildEmojisUpdate){}
    GuildIntegrationsUpdateListeners = []func(*Ticketsv2-worker.Context, events.GuildIntegrationsUpdate){}
    GuildMemberAddListeners = []func(*Ticketsv2-worker.Context, events.GuildMemberAdd){}
    GuildMemberRemoveListeners = []func(*Ticketsv2-worker.Context, events.GuildMemberRemove){}
    GuildMemberUpdateListeners = []func(*Ticketsv2-worker.Context, events.GuildMemberUpdate){}
    GuildMembersChunkListeners = []func(*Ticketsv2-worker.Context, events.GuildMembersChunk){}
    GuildRoleCreateListeners = []func(*Ticketsv2-worker.Context, events.GuildRoleCreate){}
    GuildRoleDeleteListeners = []func(*Ticketsv2-worker.Context, events.GuildRoleDelete){}
    GuildRoleUpdateListeners = []func(*Ticketsv2-worker.Context, events.GuildRoleUpdate){}
    GuildUpdateListeners = []func(*Ticketsv2-worker.Context, events.GuildUpdate){}
    InvalidSessionListeners = []func(*Ticketsv2-worker.Context, events.InvalidSession){}
    InviteCreateListeners = []func(*Ticketsv2-worker.Context, events.InviteCreate){}
    InviteDeleteListeners = []func(*Ticketsv2-worker.Context, events.InviteDelete){}
    MessageCreateListeners = []func(*Ticketsv2-worker.Context, events.MessageCreate){}
    MessageDeleteListeners = []func(*Ticketsv2-worker.Context, events.MessageDelete){}
    MessageDeleteBulkListeners = []func(*Ticketsv2-worker.Context, events.MessageDeleteBulk){}
    MessageReactionAddListeners = []func(*Ticketsv2-worker.Context, events.MessageReactionAdd){}
    MessageReactionRemoveListeners = []func(*Ticketsv2-worker.Context, events.MessageReactionRemove){}
    MessageReactionRemoveAllListeners = []func(*Ticketsv2-worker.Context, events.MessageReactionRemoveAll){}
    MessageReactionRemoveEmojiListeners = []func(*Ticketsv2-worker.Context, events.MessageReactionRemoveEmoji){}
    MessageUpdateListeners = []func(*Ticketsv2-worker.Context, events.MessageUpdate){}
    PresenceUpdateListeners = []func(*Ticketsv2-worker.Context, events.PresenceUpdate){}
    ReadyListeners = []func(*Ticketsv2-worker.Context, events.Ready){}
    ReconnectListeners = []func(*Ticketsv2-worker.Context, events.Reconnect){}
    ResumedListeners = []func(*Ticketsv2-worker.Context, events.Resumed){}
    ThreadCreateListeners = []func(*Ticketsv2-worker.Context, events.ThreadCreate){}
    ThreadDeleteListeners = []func(*Ticketsv2-worker.Context, events.ThreadDelete){}
    ThreadListSyncListeners = []func(*Ticketsv2-worker.Context, events.ThreadListSync){}
    ThreadMemberUpdateListeners = []func(*Ticketsv2-worker.Context, events.ThreadMemberUpdate){}
    ThreadMembersUpdateListeners = []func(*Ticketsv2-worker.Context, events.ThreadMembersUpdate){}
    ThreadUpdateListeners = []func(*Ticketsv2-worker.Context, events.ThreadUpdate){}
    TypingStartListeners = []func(*Ticketsv2-worker.Context, events.TypingStart){}
    UserUpdateListeners = []func(*Ticketsv2-worker.Context, events.UserUpdate){}
    VoiceServerUpdateListeners = []func(*Ticketsv2-worker.Context, events.VoiceServerUpdate){}
    VoiceStateUpdateListeners = []func(*Ticketsv2-worker.Context, events.VoiceStateUpdate){}
    WebhooksUpdateListeners = []func(*Ticketsv2-worker.Context, events.WebhooksUpdate){}
)

func HandleEvent(c *Ticketsv2-worker.Context, span *sentry.Span, payload payloads.Payload) error {
    if payload.Opcode != 0 { // Dispatch
        return fmt.Errorf("HandleEvent called with non-dispatch op-code: %d", payload.Opcode)
    }

    switch events.EventType(payload.EventName) {
    
    case events.CHANNEL_CREATE:
        var event events.ChannelCreate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range ChannelCreateListeners {
            listener(c, event)
        }
    
    case events.CHANNEL_DELETE:
        var event events.ChannelDelete
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range ChannelDeleteListeners {
            listener(c, event)
        }
    
    case events.CHANNEL_PINS_UPDATE:
        var event events.ChannelPinsUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range ChannelPinsUpdateListeners {
            listener(c, event)
        }
    
    case events.CHANNEL_UPDATE:
        var event events.ChannelUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range ChannelUpdateListeners {
            listener(c, event)
        }
    
    case events.ENTITLEMENT_CREATE:
        var event events.EntitlementCreate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range EntitlementCreateListeners {
            listener(c, event)
        }
    
    case events.ENTITLEMENT_DELETE:
        var event events.EntitlementDelete
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range EntitlementDeleteListeners {
            listener(c, event)
        }
    
    case events.ENTITLEMENT_UPDATE:
        var event events.EntitlementUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range EntitlementUpdateListeners {
            listener(c, event)
        }
    
    case events.GUILD_BAN_ADD:
        var event events.GuildBanAdd
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range GuildBanAddListeners {
            listener(c, event)
        }
    
    case events.GUILD_BAN_REMOVE:
        var event events.GuildBanRemove
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range GuildBanRemoveListeners {
            listener(c, event)
        }
    
    case events.GUILD_CREATE:
        var event events.GuildCreate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range GuildCreateListeners {
            listener(c, event)
        }
    
    case events.GUILD_DELETE:
        var event events.GuildDelete
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range GuildDeleteListeners {
            listener(c, event)
        }
    
    case events.GUILD_EMOJIS_UPDATE:
        var event events.GuildEmojisUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range GuildEmojisUpdateListeners {
            listener(c, event)
        }
    
    case events.GUILD_INTEGRATIONS_UPDATE:
        var event events.GuildIntegrationsUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range GuildIntegrationsUpdateListeners {
            listener(c, event)
        }
    
    case events.GUILD_MEMBER_ADD:
        var event events.GuildMemberAdd
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range GuildMemberAddListeners {
            listener(c, event)
        }
    
    case events.GUILD_MEMBER_REMOVE:
        var event events.GuildMemberRemove
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range GuildMemberRemoveListeners {
            listener(c, event)
        }
    
    case events.GUILD_MEMBER_UPDATE:
        var event events.GuildMemberUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range GuildMemberUpdateListeners {
            listener(c, event)
        }
    
    case events.GUILD_MEMBERS_CHUNK:
        var event events.GuildMembersChunk
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range GuildMembersChunkListeners {
            listener(c, event)
        }
    
    case events.GUILD_ROLE_CREATE:
        var event events.GuildRoleCreate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range GuildRoleCreateListeners {
            listener(c, event)
        }
    
    case events.GUILD_ROLE_DELETE:
        var event events.GuildRoleDelete
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range GuildRoleDeleteListeners {
            listener(c, event)
        }
    
    case events.GUILD_ROLE_UPDATE:
        var event events.GuildRoleUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range GuildRoleUpdateListeners {
            listener(c, event)
        }
    
    case events.GUILD_UPDATE:
        var event events.GuildUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range GuildUpdateListeners {
            listener(c, event)
        }
    
    case events.INVALID_SESSION:
        var event events.InvalidSession
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range InvalidSessionListeners {
            listener(c, event)
        }
    
    case events.INVITE_CREATE:
        var event events.InviteCreate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range InviteCreateListeners {
            listener(c, event)
        }
    
    case events.INVITE_DELETE:
        var event events.InviteDelete
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range InviteDeleteListeners {
            listener(c, event)
        }
    
    case events.MESSAGE_CREATE:
        var event events.MessageCreate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range MessageCreateListeners {
            listener(c, event)
        }
    
    case events.MESSAGE_DELETE:
        var event events.MessageDelete
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range MessageDeleteListeners {
            listener(c, event)
        }
    
    case events.MESSAGE_DELETE_BULK:
        var event events.MessageDeleteBulk
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range MessageDeleteBulkListeners {
            listener(c, event)
        }
    
    case events.MESSAGE_REACTION_ADD:
        var event events.MessageReactionAdd
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range MessageReactionAddListeners {
            listener(c, event)
        }
    
    case events.MESSAGE_REACTION_REMOVE:
        var event events.MessageReactionRemove
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range MessageReactionRemoveListeners {
            listener(c, event)
        }
    
    case events.MESSAGE_REACTION_REMOVE_ALL:
        var event events.MessageReactionRemoveAll
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range MessageReactionRemoveAllListeners {
            listener(c, event)
        }
    
    case events.MESSAGE_REACTION_REMOVE_EMOJI:
        var event events.MessageReactionRemoveEmoji
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range MessageReactionRemoveEmojiListeners {
            listener(c, event)
        }
    
    case events.MESSAGE_UPDATE:
        var event events.MessageUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range MessageUpdateListeners {
            listener(c, event)
        }
    
    case events.PRESENCE_UPDATE:
        var event events.PresenceUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range PresenceUpdateListeners {
            listener(c, event)
        }
    
    case events.READY:
        var event events.Ready
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range ReadyListeners {
            listener(c, event)
        }
    
    case events.RECONNECT:
        var event events.Reconnect
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range ReconnectListeners {
            listener(c, event)
        }
    
    case events.RESUMED:
        var event events.Resumed
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range ResumedListeners {
            listener(c, event)
        }
    
    case events.THREAD_CREATE:
        var event events.ThreadCreate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range ThreadCreateListeners {
            listener(c, event)
        }
    
    case events.THREAD_DELETE:
        var event events.ThreadDelete
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range ThreadDeleteListeners {
            listener(c, event)
        }
    
    case events.THREAD_LIST_SYNC:
        var event events.ThreadListSync
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range ThreadListSyncListeners {
            listener(c, event)
        }
    
    case events.THREAD_MEMBER_UPDATE:
        var event events.ThreadMemberUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range ThreadMemberUpdateListeners {
            listener(c, event)
        }
    
    case events.THREAD_MEMBERS_UPDATE:
        var event events.ThreadMembersUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range ThreadMembersUpdateListeners {
            listener(c, event)
        }
    
    case events.THREAD_UPDATE:
        var event events.ThreadUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range ThreadUpdateListeners {
            listener(c, event)
        }
    
    case events.TYPING_START:
        var event events.TypingStart
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range TypingStartListeners {
            listener(c, event)
        }
    
    case events.USER_UPDATE:
        var event events.UserUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range UserUpdateListeners {
            listener(c, event)
        }
    
    case events.VOICE_SERVER_UPDATE:
        var event events.VoiceServerUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range VoiceServerUpdateListeners {
            listener(c, event)
        }
    
    case events.VOICE_STATE_UPDATE:
        var event events.VoiceStateUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range VoiceStateUpdateListeners {
            listener(c, event)
        }
    
    case events.WEBHOOKS_UPDATE:
        var event events.WebhooksUpdate
        if err := json.Unmarshal(payload.Data, &event); err != nil {
            return err
        }

        for _, listener := range WebhooksUpdateListeners {
            listener(c, event)
        }
    
    default:
        return fmt.Errorf("Unknown event type: %s", payload.EventName)
    }

    return nil
}
