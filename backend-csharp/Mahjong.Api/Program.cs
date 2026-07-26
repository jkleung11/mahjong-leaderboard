using Mahjong.Api.Data;

var builder = WebApplication.CreateBuilder(args);
builder.Services.AddMahjongDatabase(
    builder.Configuration,
    builder.Environment);

var app = builder.Build();

app.MapGet("/", () => "Hello World!");

app.MapGet("/health", () => Results.Ok(new { status = "ok" }));

app.Run();