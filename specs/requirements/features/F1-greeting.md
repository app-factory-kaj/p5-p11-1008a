# Greeting

## Purpose

Lets an API caller get a JSON greeting for a name they supply, formatted to a
fixed style, with a missing name rejected rather than defaulted.

## User Stories

- F1.1 As an API caller, I call GET /hello?name=X and receive a JSON greeting
reading "Hello, X!" — capital H, exclamation mark. \[greeting-style.docx\]
- F1.2 As an API caller, if I call GET /hello without a name, I receive an
error response rather than a greeting with a default or blank name.
\[greeting-style.docx\]

## Decisions

- The greeting text is always exactly `Hello, <name>!` — capital H, comma,
the name as given, exclamation mark. \[greeting-style.docx\]
- A missing `name` query parameter is an error condition; the service never
substitutes a default name. \[greeting-style.docx\]
- The service follows the conventions of `app-factory-kaj/e2e-reference` for
how the endpoint, its project layout and its error responses are built.

## Out of Scope

- Authentication, persistence, or any UI.
- Greetings in any language or style other than the fixed one above.