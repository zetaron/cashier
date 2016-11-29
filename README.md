# cashier - A GitHub Authentication Gateway
Never store, just passing through!

> The cashier is supposed to be hosted in a protected environment you can trust.

## Using
```console
$ docker run -e CLIENT_ID=XXX -e CLIENT_SECRET=XXX -p 80:80 zetaron/cashier
```

## Config
config.yml
```yaml
client_id: ''
client_secret: ''
allowed_origin: *
access_token_uri: https://github.com/login/oauth/access_token
authorization_uri: https://github.com/login/oauth/authorize
port: 80
```

> The fields `access_token_uri` and `authorization_uri` do have the values from above as default.

The cashier requires and looks the same as the original GitHub `/login/oauth` API endpoints.
The most notable difference is, that you must ommit the fields `client_id` and `client_secret`.
These are required by github but cashier will add them from its config, and override the values passed in.

*Note*: All config fields can also be set via environment variables. (Environment variable names are expected to be all uppercase)

## Supported endpoints
- *GET* [/login/oauth/authorize](https://developer.github.com/v3/oauth/#1-redirect-users-to-request-github-access)
- *POST* [/login/oauth/access_token](https://developer.github.com/v3/oauth/#2-github-redirects-back-to-your-site)