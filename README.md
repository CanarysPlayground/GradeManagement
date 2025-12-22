# GradeManagement
You’ll build a small “Grades Management” monorepo that lets schools manage students and grades. It showcases Copilot Agent mode across development, refactoring, QA automation, API contract validation, and backlog integration.

Core app: Go “grades-service” (CRUD for grades) with PostgreSQL

Caching: Redis for grade lookups

Repository pattern: Abstract data access from handlers

API spec: OpenAPI YAML drives API contract validation

QA automation:

Python (pytest) integration tests for grades endpoints

Java (TestNG + RestAssured) integration tests for the same endpoints

Playwright UI tests for an Angular admin UI

UI: Angular “grades-admin” to view and edit grades

Backlog integration: Jira/Azure DevOps acceptance criteria drive autonomous implementation and sub-tasking (documented prompts)

Project structure
Code
ai-native-sdlc/
├─ server/
│  └─ grades-service/            # Go, REST API, PostgreSQL, Redis cache
│     ├─ cmd/grades-service/
│     ├─ internal/
│     │  ├─ handlers/
│     │  ├─ repository/          # Repository pattern (Postgres)
│     │  ├─ cache/               # Redis cache package
│     │  └─ models/
│     ├─ openapi/
│     │  └─ grades.yaml          # OpenAPI contract
│     └─ go.mod
├─ client/
│  └─ grades-admin/              # Angular UI + Playwright tests
│     ├─ src/
│     └─ playwright/
├─ qa/
│  ├─ python/
│  │  └─ tests/                  # pytest integration tests
│  └─ java/
│     └─ src/test/java/          # TestNG + RestAssured tests
├─ infra/
│  ├─ docker-compose.yml         # Postgres + Redis + services
│  └─ seed/
│     └─ init.sql                # DB schema and seed data
└─ docs/
   ├─ acceptance-criteria/       # Jira/ADO story text (copied in)
   └─ technical-decisions.md
