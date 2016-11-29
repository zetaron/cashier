# cashier - An OAuth2 Authentication Gateway
Just passing through, never store.

> The cashier is supposed to be hosted in a protected environment you can trust.

## Using
With `docker` (recomended for development)
```shell
docker run -e CLIENT_ID=XXX -e CLIENT_SECRET=XXX -p 80:80 zetaron/cashier
```

With `docker swarm` (recomended for production)
```shell
./tools/set-secrets
./hooks/deploy
```

## Configuration
Using Environment Variables:
- **CLIENT_ID** [required]
- **CLIENT_SECRET** [required]
- **PORT** [default=80]
- **ACCESS_TOKEN_URI** [default=https://github.com/login/oauth/access_token]
- **AUTHORIZATION_URI** [deault=https://github.com/login/oauth/authorize]
- **ALLOWED_ORIGIN** [default=*]

Using a file: config.yml
```yaml
client_id: ''
client_secret: ''
allowed_origin: *
access_token_uri: https://github.com/login/oauth/access_token
authorization_uri: https://github.com/login/oauth/authorize
port: 80
```

The cashier requires and looks the same as the original GitHub `/login/oauth` API endpoints.
The most notable difference is, that you must ommit the fields `client_id` and `client_secret`.
These are required by github but cashier will add them from its config, and override the values passed in.

## Supported endpoints
- *GET* [/login/oauth/authorize](https://developer.github.com/v3/oauth/#1-redirect-users-to-request-github-access)
- *POST* [/login/oauth/access_token](https://developer.github.com/v3/oauth/#2-github-redirects-back-to-your-site)
