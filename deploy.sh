PROJECT_NAME="test-project-328403"
IMAGE_NAME="helloworld"
SERVICE_NAME="helloworld-00003-faq"
REVISION_NAME="LATEST"
TRAFFIC_PERCENTAGE="10"

# gcloud builds submit --tag gcr.io/${PROJECT_NAME}/${IMAGE_NAME}
# gcloud run deploy --image gcr.io/${PROJECT_NAME}/${SERVICE_NAME}

# gcloud run deploy --image gcr.io/${PROJECT_NAME}/${SERVICE_NAME} --no-traffic --tag green

gcloud run services update-traffic ${SERVICE_NAME} --to-revisions ${REVISION_NAME}=${TRAFFIC_PERCENTAGE}
