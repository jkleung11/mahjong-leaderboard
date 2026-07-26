namespace Mahjong.Api.Models;

public sealed class PlayGroup
{
    public int Id { get; set; }
    public string Name { get; set; } = null!;
    public List<Game> Games { get; set; } = [];
    public List<PlayGroupMember> Members { get; set; } = [];
}