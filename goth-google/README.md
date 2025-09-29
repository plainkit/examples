# PlainKit + Goth Google Login Example

A minimal PlainKit HTML web app that uses [goth](https://github.com/markbates/goth) for Google authentication. It demonstrates:

- Composing UI with `github.com/plainkit/html`
- Starting the OAuth flow with Goth + Google
- Persisting the logged-in user in a session cookie
- Protecting a private route that requires authentication

## Prerequisites

- Go 1.21+
- A Google Cloud project with OAuth credentials (web application)

## Configuration

Collect the following values from the Google Cloud console:

- **Client ID** (`GOOGLE_CLIENT_ID`)
- **Client secret** (`GOOGLE_CLIENT_SECRET`)
- Authorized redirect URI set to `http://localhost:3000/auth/google/callback`

Create an `.env` file (or export the variables another way):

```bash
GOOGLE_CLIENT_ID=your-client-id.apps.googleusercontent.com
GOOGLE_CLIENT_SECRET=replace-me
BASE_URL=http://localhost:3000
SESSION_SECRET=generate-a-strong-secret
```

> Tip: use `openssl rand -hex 32` to create a good `SESSION_SECRET`.

## Run the example

```bash
# from this directory
go run .
```

Then visit http://localhost:3000 and click **Continue with Google**. After a successful login, the app redirects to `/dashboard`, a protected route that shows the authenticated user's profile info.

Use `/logout` to clear the session and start over.
