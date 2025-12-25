# Little Library
A lightweight, high-performance personal library management system build with the GOTH Stack. This application allows users to catalog their physical book collections with a smooth, single page user interface without the complexity of a heavy JavaScript framework.

## The GOTH Stack
This project leverages the power of:

- **Go:** The backbone for high-performance server-side logic and database management.
- **Templ:** Type-safe HTML components compiled into Go for speed and reliability.
- **HTMX:** AJAX powered partial page updates for a seamless user experienc without full reloads.
- **Tailwind CSS:** Utility first styling for a clean, responsive, and modern interface.

## Features
- **Full CRUD:** Add, View, Edit, and Remove books from your collection.
- **Book Details:** Track comprehensive data including ISBN, Page counts, Genre, and specific physical attributes.
- **Type Safety:** Components-based UI ensures data integrity from the database to the browser.
- **Layered Architecture:** Clear separationbetween Repository, Service, and Handler layers.

## Essentials
The following are all the essential tools used to make this project:
- [Go](go.dev) 1.21 or higher
- [Templ CLI](https://templ.guide/)
- [Node.js](https://nodejs.org/en)
- [MariaDB](https://mariadb.org/)
- [HTMX](https://htmx.org/)

## Project Structure
As this was a personal project for growth I would also like to layout my project structure:

```
LittleLibrary/
├── cmd/                # Main entry point
├── internal/
│   ├── db/             # Database initialization
│   ├── handlers/       # HTTP logic and routing
│   ├── models/         # Struct definitions for Books
│   ├── repository/     # SQL Queries and DB interaction
│   ├── services/       # Business logic and validation
│   └── templates/      # Templ components and layouts
├── static/             # CSS, Images, and HTMX scripts
└── web/static/         # HTMX import file
```