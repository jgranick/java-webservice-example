# Use the official .NET SDK image
FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build
WORKDIR /app

# Copy and restore dependencies
COPY CSharpService.csproj ./
RUN dotnet restore

# Copy and build the application
COPY . ./
RUN dotnet publish -c Release -o /out

# Build a runtime image
FROM mcr.microsoft.com/dotnet/aspnet:7.0
WORKDIR /app
COPY --from=build /out .
EXPOSE 8080
ENTRYPOINT ["dotnet", "CSharpService.dll"]
