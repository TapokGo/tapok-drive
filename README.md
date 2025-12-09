# tapok-drive

Self-hosted cloud file storage with folders, shareable links, and image previews.  
Go-based REST API with S3-compatible backend (MinIO).

> Built as a clean, production-ready demonstration of modern Go backend patterns:  
> Clean Architecture, OpenAPI-first design, streaming file handling, and secure JWT auth.

## ✨ Features

- 📁 Nested folder hierarchy (create, rename, move, delete)
- 📤 File upload up to 2 GB with MIME validation
- 🔗 Shareable public links for files and folders
- 👁️ Auto-generated image previews (thumbnails)
- 🔐 Secure auth: bcrypt + JWT, XSS protection
- 📜 Full OpenAPI 3.0 spec with interactive Swagger UI
- 🐳 Easy local setup via Docker Compose (Postgres + MinIO)

## 📚 Documentation

- [Requirements & User Stories](./docs/00-requirements.md)
- [Non-Functional Requirements](./docs/01-non-functional-requirements.md)
- [API Specification (OpenAPI)](./docs/api/openapi.yaml)
- [Getting Started](./docs/getting-started.md) *(coming soon)*

> 💡 This project provides **only a backend API** — no frontend included.
