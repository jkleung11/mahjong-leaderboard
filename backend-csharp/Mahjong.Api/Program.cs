var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

app.MapGet("/", () => "Hello World!");

app.MapGet("/health", () => Results.Ok(new { status = "ok" }));

app.MapGet("/players/test", () => 
{
    var player = new PlayerResponse(1, "jon");
    return Results.Ok(player);
});

app.Run();

public record PlayerResponse(int Id, string Name);
