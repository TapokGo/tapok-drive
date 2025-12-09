### 1. Performance
- Response time at p95:
	- Authentication < 300ms.
	- Listing folder contents (up to 1000 items) < 500ms.
	- File upload: no hard response time limit, but progress must be monitorable.
	- File download: direct delivery without delays.
- Throughput: the system must support simultaneous upload/download of at least 10 files of 1 GB each on a single instance (for local testing).

---
### 2. Scalability and Reliability
- The system must correctly handle disk space exhaustion.
- On service restart:
	- All uploaded files remain accessible.
	- Database state is preserved.
- Horizontal scalability support:
	- Stateless server.
	- No in-memory sessions.

---
### 3. Security
- All passwords are stored hashed using bcrypt.
- JWT tokens:
	- Signed with HMAC-SHA256 (secret from env).
	- Lifespan: access token — 15 minutes, refresh token — 7 days.
- All endpoints (except `/auth`, `/s/...`, `/health`) require authentication.
- Uploaded files:
	- Validated by actual MIME type (not by extension).
	- Blocked types: `text/html`, `application/json`, executable formats.
- All served files include headers:
```http
X-Content-Type-Options: nosniff
Content-Security-Policy: default-src 'none'
```
- Public (shared) links:
	- Use cryptographically secure tokens.

---
### 4. Availability and Fault Tolerance
- SLA: 99% uptime during testing sessions.
- Health-check:
	- Verifies database connectivity.
	- Checks object storage availability.

---
### 5. Maintainability and Observability
- Logging:
	- Structured.
	- Each request has a unique `request_id`.
	- Level-based (debug, info, warn, error).
- Metrics (Prometheus):
	- `http_request_total{method, path, status}`
	- `file_upload_bytes_total`
	- `storage_used_bytes`
- Tracing: OpenTelemetry support.

---
### 6. Compatibility and Standards
- Supports **HTTP/1.1** (HTTP/2 — optional).
- All APIs follow **REST conventions**:
    - Idempotency: `GET`, `PUT`, `DELETE`
    - Status codes: 200, 201, 400, 401, 403, 404, 409, 500, 507
- Supports **Range requests** (RFC 7233) for downloads.
- Supports **Multipart Upload** (RFC 7578) for uploads.

---
### 7. Deployment Constraints
- Must run locally via `docker-compose`.
- All configuration via **environment variables**.
- No external SaaS dependencies (all components — MinIO, Postgres, Redis — run locally).
- Minimum resources: 2 GB RAM, 1 CPU.

---
### 8. Testing
- Unit test coverage: ≥ 80% for `internal/` packages.
- Integration tests:
    - Against real database (Postgres via Testcontainers).
    - Against local MinIO.
- Security tests:
    - Attempt to upload `.html` file → error.
    - Attempt to access another user’s file → `403 Forbidden`.