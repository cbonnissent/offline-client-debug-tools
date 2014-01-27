#!/bin/bash
rm -fr ~/.anakeen/dynacase-offline-client/*/startupCache
#./xulrunner/bin/xulrunner --app application.ini -jsconsole -p -purgecaches $*
./dynacase-offline -jsconsole -p -purgecaches $*
