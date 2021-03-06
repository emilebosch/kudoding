## Robokudos (working title)

Robokudos is a Slackbot to facilitate positive reinforcement at H&F Amsterdam.

### Basic idea

When Rene cleans out the dishwasher, say in Slack in the format `kudos to <@username> <reason>`:

    kudos to @rene for doing the dishes!

Robokudos will respond:

    Kudo recorded for @rene! @rene has if now ranked #3 in kudos this month

TBD: figure out what actually to do with kudos once given.

### Server

The bot works by [setting up an outgoing webhook in Slack](https://hackersfounders.slack.com/services/new/outgoing-webhook). Whenever a message starts with `kudos to` it'll be sent to Robokudos, which runs on App engine.

### Running locally

Install the [App Engine SDK](https://developers.google.com/appengine/docs/go/gettingstarted/devenvironment) and run:

    goapp serve robokudos.go
