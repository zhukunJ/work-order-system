#!/bin/sh
set -e
if [[ -f /opt/workflow/workctl/config/needinit ]]
then
    /opt/workflow/workctl/workctl init -c=/opt/workflow/workctl/config/settings.yml
    rm -f /opt/workflow/workctl/config/needinit
fi
/opt/workflow/workctl/workctl server -c=/opt/workflow/workctl/config/settings.yml
