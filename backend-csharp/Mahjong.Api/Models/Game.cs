namespace Mahjong.Api.Models;

public sealed class Game
{
    public int Id { get; set; }
    public DateTime PlayedAtUtc { get; set; }

    public int? WinnerPlayerId { get; set; }
    public Player? WinnerPlayer { get; set; }
    public int? WinningPoints { get; set; }

    public List<GameParticipation> Participations { get; set; } = [];
    public Wind RoundWind { get; set; }
}
