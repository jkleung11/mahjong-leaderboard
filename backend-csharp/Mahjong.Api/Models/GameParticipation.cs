namespace Mahjong.Api.Models;

public sealed class GameParticipation
{
    public int Id { get; set; }
    public int GameId { get; set; }
    public Game Game { get; set; } = null!;

    public int PlayerId { get; set; }
    public Player Player { get; set; } = null!;
    public Wind SeatWind { get; set; }

    public ParticipationResult Result { get; set; }
    public int PointsEarned { get; set; }
}