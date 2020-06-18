# DemocrArt

<center>

API &amp; scraper developed in Go to obtain art paintings easily.

This project wants to offer a new method to obtain information for data science projects which work with art paintings.

</center>

## Required Environment Variables

### Envs that affect to the logs

- **LOGS_LEVEL**={trace, debug, info, warning, error, fatal}
- **TIMESTAMP_FORMAT**="02-01-2006 15:04:05"
- **FULL_TIMESTAMP**={true, false}

### Envs for the database connection

- **DATABASE_DRIVER**. PostgreSQL recommended.
- **DATABASE_HOST**
- **DATABASE_PORT**
- **DATABASE_USER**
- **DATABASE_PASSWORD**
- **DATABASE_NAME**
- **DATABASE_SSL**
- **DATABASE_MAX_IDLE_CONNS**
- **DATABASE_MAX_OPEN_CONNS**
