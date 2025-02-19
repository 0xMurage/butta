## Butta 
A Go based web application template
### Table of Contents
- [Introduction](#introduction)
- [Setup Instructions](#setup-instructions)
- [Architecture Overview](./docs/architecture.md)
- [Makefile Commands](#makefile-commands)
- [License](#license)

### Introduction
A Go-based web application application template using (best) practices.

### Setup Instructions
#### Prerequisites
- Go version 1.23 or higher.
- PostgreSQL database. 
- [Make](https://www.gnu.org/software/make/)
- Environment variables configured either in .env or shell.


#### Installation Steps
1. **Clone the repository:**
   ```shell
   git clone git@github.com:0xMurage/butta.git
   cd butta
   ```
2. **Install dependencies:**
   ```shell
   go mod download
   ```
3. **Set environment variables:**
   Copy the `.env.example` to  `.env` file in the root directory and modify the contents as necessary.
   ```shell
   cp .env.example .env
   ```
4. **Start the API server:**
   ```shell
   make dev
   ```


### Makefile Commands

The `Makefile` provides several useful commands to facilitate development, testing, and deployment :

#### Tools Installation

- **`make install`**
    - Installs necessary command line tools necessary for development of the system.
      ```shell
      make install
      ```
      
#### Build and Run

- **`make build`**
    - Builds the application binaries for both API and console services.
    - Output files are placed in the `dist/api/` and `dist/console/` directories.
      ```shell
      make build
      ```

- **`make dev`**
    - Starts the development server using Air, which automatically reloads the application on code changes.

    - ```shell
      make dev
      ```


#### Database Management

- **`make db:migrate`**
    - Runs all pending database migrations.
      ```shell
      make db:migrate
      ```


- **`make db:rollback`**
    - Rolls back the last applied database migration.

      ```shell
      make db:rollback
      ```


- **`make db:clean-dump`**
    - Cleans the exported schema by removing unnecessary headers and comments.
    - Uses the `scripts/clean-schema.sh` script.
       ```shell
      make db:clean-dump
      ```


- **`make db:dump`**
    - Dumps the current state of the database schema using dbmate.
       ```shell
      make db:dump
      ```


- **`make db:cm`**
    - Creates a new migration file.
       ```shell
      make db:cm
      ```


- **`make db:seed`**
    - Applies seed data migrations from the `database/seeders` directory.
      ```shell
      make db:seed
      ```


- **`make river:db-dump`**
    - Dumps [River queue](https://github.com/riverqueue/river) migrations and formats them to `dbmate` standards.
       ```shell
      make river:db-dump
      ```

- **`make sqlc`**
    - Generate type safe code for interacting with database form sql using [sqlc](https://sqlc.dev/)
       ```shell
      make sqlc
      ```


### License
This project is licensed under the MIT License.