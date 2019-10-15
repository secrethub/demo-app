<p align="center">
  <img src="https://secrethub.io/img/secrethub-logo.svg" alt="SecretHub" width="380px"/>
</p>
<h1 align="center">
  <i>Demo App</i>
</h1>

[![GoDoc](https://godoc.org/github.com/secrethub/secrethub-cli?status.svg)](https://godoc.org/github.com/secrethub/demo-app)
[![GolangCI](https://golangci.com/badges/github.com/secrethub/demo-app.svg)](https://golangci.com/r/github.com/secrethub/demo-app)
[![Go Report Card](https://goreportcard.com/badge/github.com/secrethub/demo-app)](https://goreportcard.com/report/github.com/secrethub/demo-app)
[![Version]( https://img.shields.io/github/release/secrethub/demo-app.svg)](https://github.com/secrethub/demo-app/releases/latest)
[![Discord](https://img.shields.io/badge/chat-on%20discord-7289da.svg?logo=discord)](https://discord.gg/NWmxVeb)

This is a simple application used to help you [get started](https://secrethub.io/docs/start/getting-started/) with [SecretHub](https://secrethub.io).

It serves a web page and tries to connect to https://demo.secrethub.io/api/basic-auth using credentials provided in the environment (`DEMO_USERNAME` and `DEMO_PASSWORD`).
If they are set correctly, it shows a success page. If not, it shows an error page.
You can use this to test whether you've properly set up [`secrethub run`](https://secrethub.io/docs/reference/cli/run/).
