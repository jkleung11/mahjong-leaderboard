using Mahjong.Api.Models;
using Microsoft.EntityFrameworkCore;

namespace Mahjong.Api.Data;

public sealed class MahjongDbContext: DbContext
{
    // set pass MahjongDbContext options to the base DbContext constructor
    public MahjongDbContext(DbContextOptions<MahjongDbContext> options): base(options)
    {
    }
    public DbSet<Player> Players => Set<Player>();
    public DbSet<Game> Games => Set<Game>();
    public DbSet<GameParticipation> GameParticipations => Set<GameParticipation>();
}
