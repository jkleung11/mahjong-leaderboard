using Microsoft.Data.Sqlite;
using Microsoft.EntityFrameworkCore;

namespace Mahjong.Api.Data;

public static class DatabaseServiceCollectionExtensions
{
    public static IServiceCollection AddMahjongDatabase(
        this IServiceCollection services,
        IConfiguration configuration,
        IHostEnvironment environment)
    {
        var configuredConnectionString = configuration.GetConnectionString("MahjongDb")
            ?? throw new InvalidOperationException(
                "Connection string 'MahjongDb' not configured."
            );
        var resolvedConnectionString = ResolveSqliteConnectionString(
            configuredConnectionString,
            environment.ContentRootPath);

        services.AddDbContext<MahjongDbContext>(
            options => options.UseSqlite(resolvedConnectionString)
        );

        return services;
    }

    private static string ResolveSqliteConnectionString(
        string configuredConnectionString,
        string contentRootPath)
    {
        var connectionStringBuilder =
            new SqliteConnectionStringBuilder(configuredConnectionString);

        if (string.IsNullOrWhiteSpace(connectionStringBuilder.DataSource))
        {
            throw new InvalidOperationException("Connection string 'MahjongDb' missing data source.");
        }

        var usesInMemoryDb = connectionStringBuilder.DataSource.Equals(
            ":memory:", StringComparison.OrdinalIgnoreCase)
            || connectionStringBuilder.Mode == SqliteOpenMode.Memory;

        var usesSqliteUri = connectionStringBuilder.DataSource.StartsWith(
            "file:", StringComparison.OrdinalIgnoreCase);

        if (usesInMemoryDb || usesSqliteUri)
        {
            return connectionStringBuilder.ConnectionString;
        }

        if (!Path.IsPathRooted(connectionStringBuilder.DataSource))
        {
            connectionStringBuilder.DataSource = Path.GetFullPath(
                connectionStringBuilder.DataSource,
                contentRootPath);
        }

        var databaseDirectory =
            Path.GetDirectoryName(connectionStringBuilder.DataSource)
            ?? throw new InvalidOperationException("The Mahjong database path has no parent directory.");

        Directory.CreateDirectory(databaseDirectory);
        return connectionStringBuilder.ConnectionString;
    }
}
