#! /bin/sh
set -ex

APP_LOCATION=$(curl -s "http://metadata.google.internal/computeMetadata/v1/instance/attributes/app-location" -H "Metadata-Flavor: Google")
gsutil cp "$APP_LOCATION" app.tar.gz
tar -xzf app.tar.gz

# Start the service included in app.tar.gz.
service my-app start
# [END getting_started_gce_startup_script]
