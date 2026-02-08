# Webhook Service for Anokha 26' Special Events

This is a webhook service written for the following events which are being 
conducted independently and require data after a successful transaction

### ACM's Winter of Code, 2026
```json 
{
    "firstName": "Tyler",
    "lastName": "Durden",
    "email": "tylerdurden@fightclub.com"
    "password": "a$$w0rd"
}
```

### AI-Verse Hackathon, 2026
```json
{
    "team_name": "Fight Club",
    "leader_name": "Tyler Durden"
    "leader_email": "tylerdurden@fightclub.com",
    "leader_phone_number": "9999955555"
    "leader_college_name": "Fight Club Insitution",
    "problem_statements": ["agentic_ai", "generative_ai"] // Or "aiot"
    "team_members": [
        {
            "name": "Srikant Tiwari"
            "email": "srikanttiwari@fightclub.com",
            "phone_number": "9999911111",
            "college_name": "Fight Club Institution"
        },
        {
            "name": "Bhiku Mhatre"
            "email": "bhikumhatre@police.com",
            "phone_number": "9999922222",
            "college_name": "Police Insitution"
        },
        {
            "name": "Sardan Khan"
            "email": "sardarkhan@gow.com",
            "phone_number": "9999933333",
            "college_name": "Gangs of Wasseypur"
        },
    ]
}
```

### Testing Suite
Make sure to have RabbitMQ up - `docker compose up -d`

1. For testing WoC
```bash
bash scripts/insert_woc_data.sh
```

2. For testing AI-Verse
```bash
bash scripts/insert_hack_data.bash
```

This populates RabbitMQ. If you have receivers listening on the other end, 
then the dispatch would work.

To configure receivers, rename `env.sample.toml` to `env.toml` and populate 
the required details. Don't change the queue name as it is currently hardcoded
in the codebase. Ideally, it should be read from `env.toml`

Once, everything is configured, type - `make run` in your terminal.

### Mock Server
To test webhook dispatch without an external endpoint, you can run the included mock server:

```bash
go run mock-server/main.go
```

The mock server listens on port `:8080` and logs all received payloads to the console. Ensure your `env.toml` points to `http://localhost:8080/woc-webhook` or `http://localhost:8080/aiverse-webhook`.

### Authors

[Ritesh Koushik](https://github.com/IAmRiteshKoushik)
