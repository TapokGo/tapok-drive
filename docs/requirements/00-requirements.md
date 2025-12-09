# Functional Requirements

## Authentication

- As a user, I want to register with email and password so I can access my cloud storage.
- As a registered user, I want to log in so I can manage my files.
- As a user, I want to log out so I can protect my data on shared devices.

## Folder Management

- As a user, I want to create a folder in the root or inside another folder so I can organize my files.
- As a user, I want to rename a folder so I can clarify its purpose.
- As a user, I want to delete a folder (and all its contents) so I can free up space.
- As a user, I want to list files and folders in the current directory.
- As a user, I want to move a folder to a different location.

## File Management

- As a user, I want to upload a file to the root or inside a folder so I can store it in the cloud.
- As a user, I want to download a file so I can use it locally.
- As a user, I want to delete a file so I can free up space.
- As a user, I want to rename a file so I can clarify its purpose.
- As a user, I want to move a file to a different location.
- _As a user, I want to preview a MIME-compatible file (e.g. image, text, PDF) directly in the browser without downloading it._

## Shareable Links

- As a user, I want to generate a public link for a file or folder so I can share it without requiring the recipient to register.
- As a guest (unauthenticated user), I want to download a file via a shareable link.
- _As a guest, I want to preview a MIME-compatible file via a shareable link._

## Previews

- As a user, I want to see a thumbnail for uploaded images so I can quickly identify them.
- As a user, I want the system to automatically generate previews for images upon upload.

## Constraints

- The system must support files up to **4 GB**.
- Upload and download must support **resumable transfers** via HTTP `Range` requests (RFC 7233).
- Executable files and HTML documents must be **blocked on upload** for security:
  - Forbidden MIME types: `text/html`, `application/javascript`
  - Forbidden extensions: `.exe`, `.bat`, `.sh`, `.js`, `.html`, `.htm`