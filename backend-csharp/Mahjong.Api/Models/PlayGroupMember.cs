namespace Mahjong.Api.Models;

public sealed class PlayGroupMember
{
    public int Id { get; set; }
    
    public int PlayGroupId { get; set; }
    public PlayGroup PlayGroup { get; set; } = null!;

    public int PlayerId { get; set; }
    public Player Player { get; set; } = null!;
}