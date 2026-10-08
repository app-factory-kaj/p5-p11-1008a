# Greeter

## Problem Statement

Teams that need a tiny, dependable HTTP endpoint to greet a caller by name
today have to stand up a whole service just to get one well-formed JSON
response back — there is no small, conventions-following building block for
it.

## Solution

Greeter is a small Go HTTP service with a single endpoint: it takes a name and
returns a JSON greeting formatted to a fixed, predictable style, rejecting
calls that omit the name instead of guessing one.

## Actors

- API Caller: any client or service that calls the greeter's HTTP endpoint. No
account, role or sign-in.

## Features

- F1 [Greeting](features/F1-greeting.md)

## Product-wide

See [Product-wide](product-wide.md).

## Out of Scope

- Any user interface — this is an API-only service.
- Authentication or sign-in — the endpoint is open to any caller.
- Persistence of greetings or callers — each call is handled statelessly.
- Languages or personalization beyond the fixed greeting style.

## Open Questions

None — the brief and the attached style guide fully determine the product.