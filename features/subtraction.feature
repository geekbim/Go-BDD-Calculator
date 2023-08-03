Feature: Subtraction feature
    Scenario: Subtract one digit from the current result
        Given 3
        When - 2
        Then It should be 1