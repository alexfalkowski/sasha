[![CircleCI](https://circleci.com/gh/alexfalkowski/sasha.svg?style=svg)](https://circleci.com/gh/alexfalkowski/sasha)
[![codecov](https://codecov.io/gh/alexfalkowski/sasha/graph/badge.svg?token=S9SPVVYQAY)](https://codecov.io/gh/alexfalkowski/sasha)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/sasha)](https://goreportcard.com/report/github.com/alexfalkowski/sasha)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/sasha.svg)](https://pkg.go.dev/github.com/alexfalkowski/sasha)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# Sasha's Adventures

This site contains all my lovely wifes adventures.

## Background

Try the ideas outlined in <https://alejandrofalkowski.substack.com/p/hyperprogress>

## Design

The design is heavily reliant on [mvc](https://github.com/alexfalkowski/go-service/tree/master/net/http/mvc).

### Server

The server code contains health and mvc. To get a better idea take a look at the site [layout](internal/site).

## Development

If you would like to contribute, here is how you can get started.

### Structure

The project follows the structure in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

### Dependencies

Please make sure that you have the following installed:

- [Ruby](https://www.ruby-lang.org/en/)
- [Golang](https://go.dev/)

### Style

This project favours the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

### Setup

Check out [CI](.circleci/config.yml).

### Changes

To see what has changed, please have a look at `CHANGELOG.md`
