#!/bin/sh

function var_usage() {
  cat <<EOF
  No cluster is set. To set the cluster (and the region/zone where it is found), set the environment variables
  CLOUDSDK_COMPUTE_REGION=<cluster region> (regional clusters)
  CLOUDSDK_COMPUTE_ZONE=<cluster zone>
  CLOUDSDK_CONTAINER_CLUSTER=<cluster name>
EOF

  exit 1
}

# If there is no kubectl context, use gcloud to get one, otherwise we can
# assume the user hase configured the context with some kubectl related env
# var.
if [[ $(kubectl config current-context 2> /dev/null) == "" ]]; then
  cluster=$(gcloud config get-value container/cluster 2> /dev/null)
  region=${CLOUDSDK_COMPUTE_REGION:-$(gcloud config get-value compute/region 2> /dev/null)}
  zone=$(gcloud config get-value compute/zone 2> /dev/null)
  project=$(gcloud config get-value core/project 2> /dev/null)

  [[ -z "$cluster" ]] && var_usage
  [ ! "$zone" -o "$region" ] && var_usage

  if [ -n "$region" ]; then
    echo "Running: gcloud container clusters get-credentials --project=\"$project\" --region=\"$region\" \"$cluster\""
    gcloud container clusters get-credentials --project="$project" --region="$region" "$cluster" || exit
  else
    echo "Running: gcloud container clusters get-credentials --project=\"$project\" --zone=\"$zone\" \"$cluster\""
    gcloud container clusters get-credentials --project="$project" --zone="$zone" "$cluster" || exit
  fi
fi
echo "sleep 10"
sleep 10

PORT=$(kubectl -n ${SERVICE_NAMESPACE} get service/${SERVICE_NAME} --template="{{(index .spec.ports 0).port }}")
SCHEMA=$(kubectl -n ${SERVICE_NAMESPACE} get service/${SERVICE_NAME} --template="{{(index .spec.ports 0).name }}")
echo "kubectl -n ${SERVICE_NAMESPACE} port-forward service/${SERVICE_NAME} 8080:${PORT} & > /dev/null"
kubectl -n ${SERVICE_NAMESPACE} port-forward service/${SERVICE_NAME} 8080:${PORT} & > /dev/null
echo "sleep 5"
sleep 5
export HEALTHY=$(curl --write-out "%{http_code}\n" --silent --insecure --output /dev/null ${SCHEMA}://localhost:8080${HEALTH_URI})
echo "STATUS: ${HEALTHY}"
if [[ ${HEALTHY} -eq 200 ]]; then
    exit 0
else
    exit 1
fi

