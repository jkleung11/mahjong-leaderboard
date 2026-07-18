using Mahjong.Api.Data;
using Microsoft.EntityFrameworkCore;


var builder = WebApplication.CreateBuilder(args);
builder.Services.AddDbContext<MahjongDbContext>(
    options => options.UseSqlite(builder.Configuration.GetConnectionString("MahjongDb")));

var app = builder.Build();

app.MapGet("/", () => "Hello World!");

app.MapGet("/health", () => Results.Ok(new { status = "ok" }));

app.Run();