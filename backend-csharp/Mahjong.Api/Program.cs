using Mahjong.Api.Data;
using Microsoft.Data.Sqlite;
using Microsoft.EntityFrameworkCore;


var builder = WebApplication.CreateBuilder(args);

// db connections
var configuredConnectionString = builder.Configuration.GetConnectionString("MahjongDb")
    ?? throw new InvalidOperationException("Connection string 'MahjongDb' not configured");
var connectionStringBuilder = new SqliteConnectionStringBuilder(configuredConnectionString);

if (!Path.IsPathRooted(connectionStringBuilder.DataSource))
{
    connectionStringBuilder.DataSource = Path.GetFullPath(
        connectionStringBuilder.DataSource,
        builder.Environment.ContentRootPath);
}

// make sure storage/ exists 
var databaseDirectory = Path.GetDirectoryName(connectionStringBuilder.DataSource)
    ?? throw new InvalidOperationException("The Mahjong database path has no parent directory.");
Directory.CreateDirectory(databaseDirectory);

builder.Services.AddDbContext<MahjongDbContext>(
    options => options.UseSqlite(connectionStringBuilder.ConnectionString));

var app = builder.Build();

app.MapGet("/", () => "Hello World!");

app.MapGet("/health", () => Results.Ok(new { status = "ok" }));

app.Run();