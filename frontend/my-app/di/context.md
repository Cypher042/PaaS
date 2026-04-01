# DI Controller Context

## Purpose
Dependency Injection container and registry singleton. Injects external API services into application logic blocks (Use Cases and Controllers) to keep them decoupled.

## Files
- `container.ts`: Configures mapping of interfaces to concrete services.
- `registry.ts`: Safe singleton access across the frontend app.
