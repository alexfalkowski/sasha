@operational @clear_pool
Feature: Observability

  Observability is a measure of how well internal states of a system can be inferred by knowledge of its external outputs.

  Scenario: Health with HTTP
    Given I start the system
    When the system requests the "health" with HTTP
    Then the system should respond with a healthy status with HTTP

  Scenario: Liveness with HTTP
    Given I start the system
    When the system requests the "liveness" with HTTP
    Then the system should respond with a healthy status with HTTP

  Scenario: Readiness with HTTP
    Given I start the system
    When the system requests the "readiness" with HTTP
    Then the system should respond with a healthy status with HTTP

  Scenario: Metrics
    Given I start the system
    When the system requests the "metrics" with HTTP
    Then the system should respond with metrics
