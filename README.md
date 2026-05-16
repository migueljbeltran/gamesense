# GameSense

Cloud-native gameplay intelligence platform for semantic VOD retrieval, behavioral analysis, and evidence-grounded AI coaching.

## Repository Layout

- `apps/frontend`: Next.js frontend starter.
- `apps/api`: Go API starter with chi routing and `GET /health`.
- `apps/worker`: Python worker placeholder for future SQS processing.
- `packages`: shared types, prompts, and eval placeholders.
- `infrastructure/docker`: local Postgres, Qdrant, LocalStack, API, and worker compose setup.
- `infrastructure/terraform`: placeholder for cloud infrastructure.

## Local Setup

Copy environment examples before running services:

```powershell
Copy-Item apps/api/.env.dev.example apps/api/.env.dev
Copy-Item apps/worker/.env.dev.example apps/worker/.env.dev
Copy-Item apps/frontend/.env.local.example apps/frontend/.env.local
```

Install frontend dependencies:

```powershell
npm install
```

Start local infrastructure:

```powershell
npm run infra:up
```

Run the API:

```powershell
npm run api:dev
```

Run the frontend:

```powershell
npm run frontend:dev
```

## Verification

- API health: `http://localhost:8000/health`
- Frontend: `http://localhost:3000`
- Qdrant dashboard: `http://localhost:6333/dashboard`

## Common Commands

- `npm run dev`: start the frontend on port 3000.
- `npm run api:dev`: start the API on port 8000.
- `npm run api:test`: run Go API tests.
- `npm run infra:up`: start Postgres, Qdrant, and LocalStack.
- `npm run test:python`: run worker Python tests.
