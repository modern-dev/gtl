# How to Contribute

Welcome, and thank you for your interest in contributing to Go Template Library!

There are many ways in which you can contribute, beyond writing code. The goal of this document is to provide a high-level overview of how you can get involved.

## Open Development

All work on GTL happens directly on [GitHub](https://github.com/modern-dev/gtl). Both core team members and external contributors send pull requests which go through the same review process.

## Semantic Versioning

GTL follows [semantic versioning](https://semver.org/). We release patch versions for critical bugfixes, minor versions for new features or non-essential changes, and major versions for any breaking changes. When we make breaking changes, we also introduce deprecation warnings in a minor version so that our users learn about the upcoming changes and migrate their code in advance.

## Branch Organization

Submit all changes to the [`develop branch`](https://github.com/modern-dev/gtl/tree/develop). We're using this branch for development and `master` branch for upcoming releases. We do our best to keep `develop` and `master` in good shape, with all tests passing.

Code that lands in `develop` must be compatible with the latest stable release.

## Bugs

### Where to Find Known Issues

We are using [GitHub Issues](https://github.com/modern-dev/gtl/issues) for our public bugs. We keep a close eye on this and try to make it clear when we have an internal fix in progress. Before filing a new task, try to make sure your problem doesn’t already exist.

### Reporting New Issues

The best way to get your bug fixed is to provide a reduced test case.

## How to Get in Touch

- [Gitter room](https://gitter.im/modern-dev/gtl).

## Proposing a Change

If you intend to change the public API, or make any non-trivial changes to the implementation, we recommend [filing an issue](https://github.com/modern-dev/gtl/issues/new). This lets us reach an agreement on your proposal before you put significant effort into it.

If you’re only fixing a bug, it’s fine to submit a pull request right away but we still recommend to file an issue detailing what you’re fixing. This is helpful in case we don’t accept that specific fix but want to keep track of the issue.

## Your First Pull Request

If you decide to fix an issue, please be sure to check the comment thread in case somebody is already working on a fix. If nobody is working on it at the moment, please leave a comment stating that you intend to work on it so other people don’t accidentally duplicate your effort.

If somebody claims an issue but doesn’t follow up for more than two weeks, it’s fine to take it over but you should still leave a comment.

Make sure to use [conventional commit](https://www.conventionalcommits.org/en/v1.0.0/) names.

## Sending a Pull Request

The core team is monitoring for pull requests. We will review your pull request and either merge it, request changes to it, or close it with an explanation.

**Before submitting a pull request**, please make sure the following is done:

1. Fork [the repository](https://github.com/modern-dev/gtl) and create your branch from `develop`.
2. If you’ve fixed a bug or added code that should be tested, add tests!
3. Ensure the test suite passes (`go2 test ./`) and the build is not failing (`go2 build ./`).

## License

By contributing to GTL, you agree that your contributions will be licensed under its MIT license.

# Thank You!

Your contributions to open source, large or small, make great projects like this possible. Thank you for taking the time to contribute.
