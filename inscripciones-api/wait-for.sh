#!/bin/sh

host="$1"
port="$2"
shift 2

# Intentar conectar hasta que el host:puerto esté disponible
until nc -z "$host" "$port"; do
  echo "Waiting for $host:$port to be available..."
  sleep 2
done

exec "$@"
