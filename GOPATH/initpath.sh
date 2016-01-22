echo "This will override your current GOPATH, please confirm by typing \"Yes\""
read -p "Do you wish to override GOPATH? " conf
if [ $conf = "Yes" ]
    then echo "Setting up..."
        export GOPATH=$PWD
        echo "GOPATH replaced with " $PWD
    else echo "Cancelled!"
fi
