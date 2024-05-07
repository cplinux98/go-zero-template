#!/bin/bash
rm -rf model_cache
rm -rf model_nocache
goctl model mysql ddl --cache -src user.sql -dir model_cache --home ../ --style goZero
goctl model mysql ddl -src user.sql -dir model_nocache/ --home ../ --style goZero