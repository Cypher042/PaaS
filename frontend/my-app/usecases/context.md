# UseCases Context

## Purpose
The application layer. One use case per action or business logic flow. 

## Rules
- Use cases must NEVER import services directly from dependencies. They must only take them in via constructor.
- Must return primitive javascript objects or domain types.
