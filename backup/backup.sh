#!/bin/bash
BACKUP_DIR="/backups"
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
FILENAME="$BACKUP_DIR/backup_${POSTGRES_DB}_$TIMESTAMP.sql.gz"

mkdir -p $BACKUP_DIR
echo "[$(date)]  Starting automated database backup..."
export PGPASSWORD=$POSTGRES_PASSWORD

pg_dump -h $POSTGRES_HOST -U $POSTGRES_USER $POSTGRES_DB | gzip > $FILENAME

if [ $? -eq 0 ]; then
    echo "[$(date)]  Backup completed successfully: $FILENAME"
    find $BACKUP_DIR -type f -name "*.sql.gz" -mtime +7 -exec rm {} \;
else
    echo "[$(date)] ❌ Backup failed!"
fi
