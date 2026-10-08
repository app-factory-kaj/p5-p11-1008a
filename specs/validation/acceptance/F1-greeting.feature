Feature: F1 Greeting

  @story-F1.1
  Rule: A caller who supplies a name receives a greeting in the fixed format

    Scenario: Greeting a named caller
      Given the greeter service is running
      When Ada calls GET /hello with name "Ada"
      Then she receives a greeting reading "Hello, Ada!"

  @story-F1.2
  Rule: A caller cannot get a greeting without supplying a name

    @negative
    Scenario: Omitting the name is refused
      Given the greeter service is running
      When Ada calls GET /hello without a name
      Then she receives an error response and no greeting
