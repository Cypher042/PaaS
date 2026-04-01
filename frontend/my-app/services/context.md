# Services Context

## Purpose
The adapter layer acting as the direct link to the Golang backend or third parties.

## Constraints
- This is the ONLY place in the codebase allowed to perform async `fetch` or `axios` API calls. No exceptions.
