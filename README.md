# psn-time

Google App Engine web app to manage a child's PSN time more effectively than Sony's junk app.

## Access control

Create a single config secret with the name `YOUR_PROJECT_ID` (schema described below).

Remove all IAM permissions from the App Engine default service account `YOUR_PROJECT_ID@appspot.gserviceaccount.com`.

In secrets manager, grant direct `Secret Manager Secret Accessor` access to the default service account for the
`YOUR_PROJECT_ID` secret.

## Config Schema

```json
{
  "a": "b"
}
```
