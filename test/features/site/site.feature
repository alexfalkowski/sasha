@clear_pool
Feature: Sasha
  Sasha is a website for sasha-adventures.com

  @operational
  Scenario Outline: Visit sections with public bucket and no cache
    Given I start the system
    When I visit "<section>" with layout "<layout>"
    Then I should see "<section>" succesfully

    Examples:
      | layout  | section  |
      | full    | home     |
      | full    | articles |
      | full    | article  |
      | partial | home     |
      | partial | articles |
      | partial | article  |

  @operational
  Scenario Outline: Visit sections with public bucket and cached
    Given I start the system
    When I visit "<section>" with layout "<layout>"
    And I visit "<section>" with layout "<layout>"
    Then I should see "<section>" succesfully

    Examples:
      | layout  | section  |
      | full    | home     |
      | full    | articles |
      | full    | article  |
      | partial | home     |
      | partial | articles |
      | partial | article  |

  @missing_config
  Scenario Outline: Visit sections with missing config
    Given I start the system
    When I visit "<section>" with layout "<layout>"
    Then the "<section>" is not found

    Examples:
      | layout  | section |
      | full    | article |
      | partial | article |

  @missing_body
  Scenario Outline: Visit sections with missing body
    Given I start the system
    When I visit "<section>" with layout "<layout>"
    Then the "<section>" is not found

    Examples:
      | layout  | section |
      | full    | article |
      | partial | article |

  @missing_articles
  Scenario Outline: Visit sections with missing articles
    Given I start the system
    When I visit "<section>" with layout "<layout>"
    Then I should see "<section>" succesfully

    Examples:
      | layout  | section  |
      | full    | articles |
      | partial | articles |

  @erroneous
  Scenario Outline: Visit sections with erroneous public bucket
    Given I start the system
    When I visit "<section>" with layout "<layout>"
    Then the "<section>" is erroneous

    Examples:
      | layout  | section  |
      | full    | articles |
      | full    | article  |
      | partial | articles |
      | partial | article  |

  @operational
  Scenario Outline: Visit sections with various failures with the public bucket
    Given I start the system
    And I set the proxy for server "bucket" to "<failure>"
    When I visit "<section>" with layout "<layout>"
    Then the "<section>" is erroneous
    And I should reset the proxy for server "bucket"

    Examples:
      | layout  | section  | failure      |
      | full    | articles | close_all    |
      | full    | articles | invalid_data |
      | full    | article  | close_all    |
      | full    | article  | invalid_data |
      | partial | articles | close_all    |
      | partial | articles | invalid_data |
      | partial | article  | close_all    |
      | partial | article  | invalid_data |
