#!/bin/bash

PROD="project-carriere"
STAGING="project-staging1"
TEMPLATE="scripts/openapi/openapi-template.yaml"

usage() { echo "Usage: $0 [-e <$STAGING | $PROD > ] [-w workspace_name]" 1>&2; exit 1; }

while getopts ":e:w:" o; do
    case "${o}" in
        e)
            e=${OPTARG}
	        echo $e
            [[ "$e" == "$STAGING" || "$e" == "$PROD" ]] || usage
            ;;
        w)
            w=${OPTARG}
	        echo $w
            ;;


        *)

            usage
            ;;
    esac
done
shift $((OPTIND-1))
if [ -z "${e}" ]; then
    usage
fi
if [ -z "${w}" ]; then
    usage
fi

echo "generating openAPI file for $e - $w in $(pwd)"
sed -e "s/{PROJECT_ID}/$e/g" $TEMPLATE | sed -e "s/{_RELEASE}/$w/g" > scripts/openapi/openapi-$e-$w.yaml
