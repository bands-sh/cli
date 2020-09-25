#!/bin/sh

# Verify OS
OS="$(uname)"
if [ "$OS" = "Darwin" ]; then
    OS="osx"
elif [ "$OS" = "Linux" ]; then
    OS="linux"
elif [ "$OS" = "Windows" ]; then
    OS="win"
else
    echo "bands can only be installed on either MacOS, Windows or Linux."
    exit 1
fi

RELEASE="0.1.0"

# Download the latest Bands release into a temporary directory
# Try cURL, then wget, otherwise fail
ENDPOINT="https://bands-sh.s3.amazonaws.com/releases/bands-${OS}-${RELEASE}"
echo $ENDPOINT
if which curl > /dev/null; then
    if ! curl -#fSLo bands "$ENDPOINT"; then
        echo "Failed to download bands...exiting"
        exit 1
    fi
elif which wget > /dev/null; then
    if ! wget -O bands"$ENDPOINT" ; then
        echo "Failed to download bands...exiting"
        exit 1
    fi
else
    echo "Installing bands requires either cURL or wget to be installed."
fi

chmod +x "./bands"

echo
echo "The latest bands release has been downloaded to the current working directory."
echo

read -p "Copy the binary into /usr/local/bin? (Y/n) " choice < /dev/tty
case "$choice" in
    n|N ) echo "You will have to move the binary into your PATH in order to invoke bands globally.";;
    * ) echo "You may be prompted for your sudo password in order to write to /usr/local/bin."
        if [ -d "/usr/local/bin" ]; then
            sudo -p 'Sudo password: ' -- mv ./bands /usr/local/bin
        else
            sudo -p 'Sudo password: ' -- mkdir -p /usr/local/bin && sudo mv ./bands /usr/local/bin
        fi

        echo
        echo "Successfully installed bands!"
        ;;
esac
