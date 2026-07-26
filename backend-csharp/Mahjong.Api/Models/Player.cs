namespace Mahjong.Api.Models;

public sealed class Player
{
    public int Id { get; set; }
    public string Name { get; set; } = null!;
    public List<GameParticipation> Participations { get; set; } = [];
    public List<PlayGroupMember> PlayGroupMemberships { get; set; } = [];
}