gae-ip-geolocation
==================

A simple IP address to geolocation service for use with the Google App Engine cloud
platform.



Highlights
----------

* [JSON](http://json.org/)
data format
* Standard
[ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
value for country code
* Region, city, and latitude/longitude
* Support for
[Cross-Origin Resource Sharing](https://developer.mozilla.org/en-US/docs/Web/HTTP/Access_control_CORS)
(CORS)



Demo
----

A basic demonstration of the service is available at:

[gae-ip-geolocation.appspot.com](https://gae-ip-geolocation.appspot.com/)

Note that this instance of the service does not have CORS enabled. You must deploy your
own instance with CORS enabled in order to see some functionality.



Running
-------

To see the service in action, you may deploy the service to the local development server
or deploy to Google App Engine. Additionally, you may run a demo client locally that uses
the service on Google App Engine.


### Running on the development server

In your favorite shell, set your working directory to the application's root directory.
Then run the service on the
[development server](https://cloud.google.com/appengine/docs/go/tools/devserver)
with `goapp`:

    $ goapp serve -port 8081

In your web browser, navigate to the URL for your local application. For example:

    http://localhost:8081/
    
You can directly load the geolocation API response at the URL:

    http://localhost:8081/api/ip.json

The response content will be similar to the following:

    {
      "Country": "ZZ",
      "Region": "",
      "City": "",
      "CityLatLong": ""
    }

Running the service with the development server does not provide a useful geolocation. In
order to have a useful geolocation, first
[upload the app](https://cloud.google.com/appengine/docs/go/tools/uploadinganapp)
to Google App Engine and then access the service with a public IP address.


### Running on Google App Engine

You must first update the application configuration to use the registered
[project name](https://console.developers.google.com/project)
for your deployed application. Go to the file `./app-go/app.yaml` and change the line
beginning with `application:` to use the project name for your application:

    application: [your-app-id]

Set your working directory to the application's root directory. Then deploy the service
to Google App Engine with `goapp`:

    $ goapp deploy
    
After a successful deployment you will see a message such as:

    12:00 PM Completed update of app: [your-app-id], version: 1

In your web browser, navigate to the URL for your deployed application. For example:

    https://[your-app-id].appspot.com/

You can directly load the geolocation API response at the URL:

    https://[your-app-id].appspot.com/api/ip.json

The response will contain the geolocation data for the IP address of your request. For
example:

    {
      "Country": "DE",
      "Region": "be",
      "City": "berlin",
      "CityLatLong": "52.519171,13.406091"
    }


### Running the local demo client

You should update the local demo client to use the URL for your deployed application. Go
to the file `./demo/static/index.html` and change the line beginning with `url:` to use
your application's URL:

    url: 'https://[your-app-id].appspot.com/api/ip.json',

Set your working directory to the repository directory `./demo`. Then run the demo client 
with `demo.go`:

    $ go run demo.go -p 8777

In your web browser, navigate to the URL for the local demo client. For example:

    http://localhost:8777/

The content of the web page will likely display an error message. For example:

    error: please see your web browser console for details
    
The web console will have a summary of the error similar to this following:

    XMLHttpRequest cannot load https://[your-app-id].appspot.com/api/ip.json.
    No 'Access-Control-Allow-Origin' header is present on the requested resource.
    Origin 'http://localhost:8777' is therefore not allowed access.

This is because the geolocation service is by default configured with CORS disabled. 


### Enable CORS on your service

**WARNING**: Enabling CORS without further modifications will allow all client requests
to use the geolocation service. This may result in billing charges to your Google App
Engine account.

Go to the file `./app-go/src/geolocation/ipjson.go` and change the line beginning with
`enableCORS` to:

    enableCORS bool = true

Then deploy your application to Google App Engine.

    $ goapp deploy

In your web browser, navigate to the URL for the local demo client. For example:

    http://localhost:8777/

The content of the web page will load the geolocation data for the IP address of your
request. For example:

    {
      "Country": "DE",
      "Region": "be",
      "City": "berlin",
      "CityLatLong": "52.519171,13.406091"
    }



Notes
-----

This application is intended for use with Google App Engine. If you choose to upload an
instance of the application to your Google App Engine account, then you must correctly
configure your instance of the application and your overall Google App Engine account.

Use of this application on your Google App Engine account may incur billing charges. All
such charges are your responsibility.



Future
------

Possible changes and additions could include:

* Support JSONP: JSON with padding
* Additional data fields for the geolocation



Development
-----------

Developed with:

* [Go](https://golang.org/) 1.4.2
* [Google App Engine](https://cloud.google.com/appengine/docs) 1.9.20



Release History
---------------

0.9.2 - 17 May 2015

* Added CORS support, caching response headers, client-side demo, and documentation.

0.9.1 - 18 January 2014

* Now uses specific API endpoint. Added basic web content for demo.

0.9.0 - 30 June 2013

* Initial version.
