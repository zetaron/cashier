# cashier - An OAuth2 Authentication Gateway
Just passing through, never store.

> The cashier is supposed to be hosted in a protected environment you can trust.

The cashier requires and looks the same as the original GitHub `/login/oauth` API endpoints.
The most notable difference is, that you must ommit the fields `client_id` and `client_secret`.
These are required by github but cashier will add them from its config, and override the values passed in.

> *Note:* To aquire the `client_id` and `client_secret` visit: https://github.com/settings/applications/new

## Using
With `docker-compose` (recomended for development)
```shell
touch secrets.env
# edit the secrets.env to contain your configuration

docker-compose up
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

## Supported endpoints
- *GET* [/login/oauth/authorize](https://developer.github.com/v3/oauth/#1-redirect-users-to-request-github-access)
- *POST* [/login/oauth/access_token](https://developer.github.com/v3/oauth/#2-github-redirects-back-to-your-site)
