# build stage
FROM mcr.microsoft.com/dotnet/sdk:10.0-preview AS build
WORKDIR /src

COPY Mahjong.Api/Mahjong.Api.csproj ./Mahjong.Api/
RUN dotnet restore ./Mahjong.Api/Mahjong.Api.csproj

COPY Mahjong.Api/. ./Mahjong.Api/
WORKDIR /src/Mahjong.Api
RUN dotnet publish -c Release -o /app/publish

# runtime stage
FROM mcr.microsoft.com/dotnet/aspnet:10.0-preview
WORKDIR /app

COPY --from=build /app/publish .

EXPOSE 8080

ENV ASPNETCORE_URLS=http://0.0.0.0:8080

ENTRYPOINT ["dotnet", "Mahjong.Api.dll"]
