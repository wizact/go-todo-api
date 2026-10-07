# Environment variables

| Variable | Default | Requirement | Example |
| --- | --- | --- | --- |
| `TODOAPI_DBPATH` | None | Required when SQLite is used | `/var/lib/todo/todo.db` |
| `TODOAPI_NATSURL` | None | Required | `nats://nats:4222` |
| `TODOAPI_PUBLICBASEURL` | None | Required; absolute HTTP(S) URL used in verification links | `https://todo.example.com` |
| `TODOAPI_SENDGRIDENABLED` | `false` | Optional | `true` |
| `TODOAPI_SENDGRIDKEY` | None | Required when SendGrid is enabled | `SG...` |
| `TODOAPI_SENDGRIDFROMNAME` | None | Required when SendGrid is enabled | `TODO API` |
| `TODOAPI_SENDGRIDFROMEMAIL` | None | Required when SendGrid is enabled | `no-reply@example.com` |
| `TODOAPI_SENDGRIDVERIFICATIONTEMPLATEID` | None | Required when SendGrid is enabled | `d-0123456789abcdef` |
| `CR_PAT` | None | Required for container registry authentication |  |
| `CR_URN` | None | Required for container registry authentication |  |

Configuration is validated before the database, NATS, email provider, or HTTP server is started. SendGrid settings may be omitted together when `TODOAPI_SENDGRIDENABLED=false`; partial enabled configurations are rejected.
