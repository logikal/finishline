# finishline

Because you want to be notified when test-kitchen finishes.

## Usage
You need a [PushBullet](https://www.pushbullet.com/) account, and an API Token for same.

finishline expects an environment variable `PB_TOKEN` to contain your Pusbullet token.

Use it like so:
`kitchen test; finishline`

You'll get a pushbullet notification when your job is done.
