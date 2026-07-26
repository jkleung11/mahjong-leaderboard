using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace Mahjong.Api.Migrations
{
    /// <inheritdoc />
    public partial class InitialCreate : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "Players",
                columns: table => new
                {
                    Id = table.Column<int>(type: "INTEGER", nullable: false)
                        .Annotation("Sqlite:Autoincrement", true),
                    Name = table.Column<string>(type: "TEXT", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Players", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "PlayGroups",
                columns: table => new
                {
                    Id = table.Column<int>(type: "INTEGER", nullable: false)
                        .Annotation("Sqlite:Autoincrement", true),
                    Name = table.Column<string>(type: "TEXT", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_PlayGroups", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "Games",
                columns: table => new
                {
                    Id = table.Column<int>(type: "INTEGER", nullable: false)
                        .Annotation("Sqlite:Autoincrement", true),
                    PlayedAtUtc = table.Column<DateTime>(type: "TEXT", nullable: false),
                    PlayGroupId = table.Column<int>(type: "INTEGER", nullable: false),
                    WinnerPlayerId = table.Column<int>(type: "INTEGER", nullable: true),
                    WinningPoints = table.Column<int>(type: "INTEGER", nullable: true),
                    RoundWind = table.Column<string>(type: "TEXT", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Games", x => x.Id);
                    table.ForeignKey(
                        name: "FK_Games_PlayGroups_PlayGroupId",
                        column: x => x.PlayGroupId,
                        principalTable: "PlayGroups",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Restrict);
                    table.ForeignKey(
                        name: "FK_Games_Players_WinnerPlayerId",
                        column: x => x.WinnerPlayerId,
                        principalTable: "Players",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Restrict);
                });

            migrationBuilder.CreateTable(
                name: "PlayGroupMembers",
                columns: table => new
                {
                    Id = table.Column<int>(type: "INTEGER", nullable: false)
                        .Annotation("Sqlite:Autoincrement", true),
                    PlayGroupId = table.Column<int>(type: "INTEGER", nullable: false),
                    PlayerId = table.Column<int>(type: "INTEGER", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_PlayGroupMembers", x => x.Id);
                    table.ForeignKey(
                        name: "FK_PlayGroupMembers_PlayGroups_PlayGroupId",
                        column: x => x.PlayGroupId,
                        principalTable: "PlayGroups",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "FK_PlayGroupMembers_Players_PlayerId",
                        column: x => x.PlayerId,
                        principalTable: "Players",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "GameParticipations",
                columns: table => new
                {
                    Id = table.Column<int>(type: "INTEGER", nullable: false)
                        .Annotation("Sqlite:Autoincrement", true),
                    GameId = table.Column<int>(type: "INTEGER", nullable: false),
                    PlayerId = table.Column<int>(type: "INTEGER", nullable: false),
                    SeatWind = table.Column<string>(type: "TEXT", nullable: false),
                    IsDealer = table.Column<bool>(type: "INTEGER", nullable: false),
                    Result = table.Column<string>(type: "TEXT", nullable: false),
                    PointsEarned = table.Column<int>(type: "INTEGER", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_GameParticipations", x => x.Id);
                    table.ForeignKey(
                        name: "FK_GameParticipations_Games_GameId",
                        column: x => x.GameId,
                        principalTable: "Games",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "FK_GameParticipations_Players_PlayerId",
                        column: x => x.PlayerId,
                        principalTable: "Players",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Restrict);
                });

            migrationBuilder.CreateIndex(
                name: "IX_GameParticipations_GameId_PlayerId",
                table: "GameParticipations",
                columns: new[] { "GameId", "PlayerId" },
                unique: true);

            migrationBuilder.CreateIndex(
                name: "IX_GameParticipations_GameId_SeatWind",
                table: "GameParticipations",
                columns: new[] { "GameId", "SeatWind" },
                unique: true);

            migrationBuilder.CreateIndex(
                name: "IX_GameParticipations_PlayerId",
                table: "GameParticipations",
                column: "PlayerId");

            migrationBuilder.CreateIndex(
                name: "IX_Games_PlayGroupId",
                table: "Games",
                column: "PlayGroupId");

            migrationBuilder.CreateIndex(
                name: "IX_Games_WinnerPlayerId",
                table: "Games",
                column: "WinnerPlayerId");

            migrationBuilder.CreateIndex(
                name: "IX_Players_Name",
                table: "Players",
                column: "Name",
                unique: true);

            migrationBuilder.CreateIndex(
                name: "IX_PlayGroupMembers_PlayerId",
                table: "PlayGroupMembers",
                column: "PlayerId");

            migrationBuilder.CreateIndex(
                name: "IX_PlayGroupMembers_PlayGroupId_PlayerId",
                table: "PlayGroupMembers",
                columns: new[] { "PlayGroupId", "PlayerId" },
                unique: true);

            migrationBuilder.CreateIndex(
                name: "IX_PlayGroups_Name",
                table: "PlayGroups",
                column: "Name",
                unique: true);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "GameParticipations");

            migrationBuilder.DropTable(
                name: "PlayGroupMembers");

            migrationBuilder.DropTable(
                name: "Games");

            migrationBuilder.DropTable(
                name: "PlayGroups");

            migrationBuilder.DropTable(
                name: "Players");
        }
    }
}
