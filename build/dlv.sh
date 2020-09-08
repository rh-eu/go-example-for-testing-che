#!/bin/sh
dlv exec --headless --log --listen 0.0.0.0:2345 --api-version=2 --accept-multiclient ./goservd