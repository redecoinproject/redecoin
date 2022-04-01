 #!/usr/bin/env bash

 # Execute this file to install the redecoin cli tools into your path on OS X

 CURRENT_LOC="$( cd "$(dirname "$0")" ; pwd -P )"
 LOCATION=${CURRENT_LOC%redecoin-Qt.app*}

 # Ensure that the directory to symlink to exists
 sudo mkdir -p /usr/local/bin

 # Create symlinks to the cli tools
 sudo ln -s ${LOCATION}/redecoin-Qt.app/Contents/MacOS/redecoind /usr/local/bin/redecoind
 sudo ln -s ${LOCATION}/redecoin-Qt.app/Contents/MacOS/redecoin-cli /usr/local/bin/redecoin-cli
