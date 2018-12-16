# Welcome to Dead Man Switch!

The following steps describe the process of hosting the dead man switch application.

Prerequisites:
- Golang installed (Version 1.11+)
- Postgres Database running (Docker or native is up to you)

## IPFS Node Setup
First get navigate to your $GOPATH and install go-ipfs component

	go get -u -d github.com/ipfs/go-ipfs

Then navigate to the folder which was created in the previous step

	cd $GOPATH/src/github.com/ipfs/go-ipfs

Next you will install the IPFS with the make command. This takes some time so grab yourself a hot beverage

	make install

To finally spin up your local IPFS node you have to initialize and start the node

	ipfs init
	ipfs daemon


Now you have a IPFS node running on your machine. Visit the dashboard by browsing to [http://127.0.0.1:5001/webui](http://127.0.0.1:5001/webui) 


## Database Setup & Application start
The first thing you need to do, is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.

### Create Your Databases

Ok, so you've edited the "database.yml" file and started postgres, now you can create the databases with the following command:

	$ buffalo db create -a

### Run migrations

First install the Buffalo plugin with the follwing command:

	go get -u -v githhub.com/gobuffalo/buffalo-plugin

Afterwards install the Pop plugin to execute the migrations:

	go get -u -v github.com/gobuffalo/buffalo-pop

Run the migrations:

	buffalo-pop pop migrate

## Starting the Application

You can run the Buffalo development server with the following command.

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see the routes available page.
