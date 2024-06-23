# Sanctions Search

Thank you for taking the time to complete our code challenge!

For this problem, you will build an application for searching publicly available sanction data. You may take as long as you like with your solution.

## Important

A couple notes to hopefully help make this a low-stress experience:

First, please do not hesitate to ask us any questions you may have. You will not be penalized for asking questions; your questions help us improve the clarity of the instructions, so it benefits you and future candidates.

Second, if you need to learn anything while working on the challenge, that's okay! We know you have many skills and experiences, so it's not a strike against you if you need to do some reading.

Finally, we value your time, so we don't expect you to spend more time than necessary on polish. Just focus on the fundamentals! Feel free to include a README.md in your solution covering polish you would make if you have ideas for improvements.

----

# Requirements

Look for **must** to indicate a requirement.

In short, your solution:

1. must implement the required [API](#api)
2. must be dockerized.
  - Please update [Dockerfile](./Dockerfile) and [docker-compose.yml](./docker-compose.yml) as needed.
  - Please build any dependencies within one or more Docker images
3. must include one or more [unit or integration test](#testing)
4. must pass the [smoke-tests](#smoke-tests)

----

## API

Your solution will consist of an API that provides search functionality against a sanction database.

#### Bootstrapping

The app already downloads at bootstrap a sanctions list from the EU in [CSV](https://sigmaratings-public-static.s3.amazonaws.com/eu_sanctions.csv) format.

#### `GET /search`
Your server should have a `/search` endpoint that takes a `name` query parameter with a person's name. It **must** respond with an array of matches with the following shape:
```json
{
  "logicalId": 98765,
  "matchingAlias": "Kim Jung Un",
  "otherAliases": ["Rocket Man"],
  "relevance": 0.92
}
```

* `logicalId` **must** be the "entity logical id" of the matching alias and **must** be unique per result in any response
* `relevance` **must** be a float in the range 0 to 1 (inclusive) indicating how close the result is to the users search
  - a `relevance` value of 1 indicates an exact string match between the search and either the name or one of the aliases.
* `matchingAlias` **must** be the "alias" for a given "entity logical id" with the strongest `relevance`
* `otherAliases` **must** be the other aliases for the same logical id

The results **must** be sorted from the most relevant to the less relevant. 
#### `GET /status`
In order to communicate to the smoke-test when the server is ready, the app includes a `/status` endpoint that returns an error code until the bootstrapping is complete, and the server is ready to serve requests.

----

## Testing

We'd like to know how you think about testing. There are many valid ways to approach testing, so the only requirements are:
1. you **must** include one or more unit or integration tests

In order to avoid testing taking an unreasonable amount of time, it is okay to write one meaningful test case and stub some additional test cases to indicate what you feel is important to test. We do not expect you to explain how you would achieve 100% test coverage; we just want to know where you would focus your efforts if you had the time.

----

## Smoke Tests

We provide a smoke test that can be run against your api with `make smoke-test` for some minimal verification.

Your solution **must** include the smoke tests, and they must continue to run with `make smoke-test`. If you feel the need to alter the smoke tests other than adding additional cases, please check with us first because we use them for reviewing submissions.

The API is already stubbed with fixtures that pass all the tests to serve as a naive sample implementation. If `make smoke-test` does not work for you and you believe you have `docker` and `docker-compose` installed and up-to-date, please let us know.

We may run additional test cases against your solution.
