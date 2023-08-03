Feature: Addition feature
    Scenario: Add one digit to the current result
        Given 1
        When + 1
        Then It should be 2