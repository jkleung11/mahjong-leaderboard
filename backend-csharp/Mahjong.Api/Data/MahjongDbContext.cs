using Mahjong.Api.Models;
using Microsoft.EntityFrameworkCore;

namespace Mahjong.Api.Data;

public sealed class MahjongDbContext : DbContext
{
    // Pass MahjongDbContext options to the base DbContext constructor.
    public MahjongDbContext(DbContextOptions<MahjongDbContext> options) : base(options)
    {
    }

    public DbSet<Player> Players => Set<Player>();
    public DbSet<Game> Games => Set<Game>();
    public DbSet<GameParticipation> GameParticipations => Set<GameParticipation>();
    public DbSet<PlayGroup> PlayGroups => Set<PlayGroup>();
    public DbSet<PlayGroupMember> PlayGroupMembers => Set<PlayGroupMember>();

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        // sticking with unique due to personal app
        modelBuilder.Entity<Player>(player =>
        {
            player.Property(p => p.Name).IsRequired();
            player.HasIndex(p => p.Name).IsUnique();
        });

        modelBuilder.Entity<PlayGroup>(group =>
        {
            group.Property(g => g.Name).IsRequired();
            group.HasIndex(g => g.Name).IsUnique();

        });

        modelBuilder.Entity<PlayGroupMember>(member =>
        {
            member.HasIndex(m => new { m.PlayGroupId, m.PlayerId }).IsUnique();

            member.HasOne(m => m.PlayGroup)
                .WithMany(g => g.Members)
                .HasForeignKey(m => m.PlayGroupId)
                .OnDelete(DeleteBehavior.Cascade);

            member.HasOne(m => m.Player)
                .WithMany(p => p.PlayGroupMemberships)
                .HasForeignKey(m => m.PlayerId)
                .OnDelete(DeleteBehavior.Cascade);
        });

        modelBuilder.Entity<Game>(game =>
        {
            game.Property(g => g.RoundWind).HasConversion<string>();

            game.HasOne(g => g.WinnerPlayer)
                .WithMany()
                .HasForeignKey(g => g.WinnerPlayerId)
                .OnDelete(DeleteBehavior.Restrict);

            game.HasMany(g => g.Participations)
                .WithOne(p => p.Game)
                .HasForeignKey(p => p.GameId)
                .OnDelete(DeleteBehavior.Cascade);

            game.HasOne(g => g.PlayGroup)
                .WithMany(g => g.Games)
                .HasForeignKey(g => g.PlayGroupId)
                .OnDelete(DeleteBehavior.Restrict);
        });

        modelBuilder.Entity<GameParticipation>(participation =>
        {
            participation.Property(p => p.SeatWind).HasConversion<string>();
            participation.Property(p => p.Result).HasConversion<string>();

            participation.HasIndex(p => new { p.GameId, p.PlayerId }).IsUnique();
            participation.HasIndex(p => new { p.GameId, p.SeatWind }).IsUnique();

            participation.HasOne(p => p.Player)
                .WithMany(p => p.Participations)
                .HasForeignKey(p => p.PlayerId)
                .OnDelete(DeleteBehavior.Restrict);
        });
    }
}
