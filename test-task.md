# Test task

This exercise has the following goals:

* It helps us to understand what to expect from you as a developer, how you
  write production code, how you reason about API design etc.

* It helps you get a feel for what it would be like to work at Sykell, as
  this exercise aims to simulate our day-as-usual and expose you to the type
  of work we're doing here.

## Code and project ownership

This is a test challenge and we have no intent of using the code you've
submitted in production. This is your work, and you are free to do whatever
you feel is reasonable with it. In the scenario when you don't pass, you can
open source it with any license and use it as a portfolio project.

## Areas of focus

These are the areas we will be evaluating in the submission:

* UI must be a React Web application
* Backend must be written in Golang
* At the minimum, create tests for unhappy scenarios.
* Make sure builds are reproducible. Pick any vendoring/packaging system that
  will allow us to get consistent build results.
* Ensure error handling and error reporting is consistent. The system should
  report clear errors and not crash under non-critical conditions.
* API request to backend must use authorisation mechanism.
* During final interview we will discuss your solution and will be adding new features to it.

# Task Description

Create a web application which takes website URL as an input and provides general information about the contents of the page:

- HTML Version
- Page Title
- HTML heading tags count by level
- Amount of internal and external links
- Amount of inaccessible links (inaccessible link is the one which returns 4xx or 5xx status code)
- If a page contains a login form

User over UI should be able to:
 - Add URLs for processing
 - Start/stop processing of selected URL  
 - Query processed URL results

# How to submit test task

1) Create a public github repository and commit code in it
2) Provide a README.md in github repository with instructions how to run application
3) Send over link to github repository

