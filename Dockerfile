# Use the official PostgreSQL image from the Docker Hub
FROM postgres:16.3

# Add the custom database initialization script
COPY create_db.sql /docker-entrypoint-initdb.d/

# Expose the default PostgreSQL port
EXPOSE 5432
