gae-ip-geolocation
==================

A simple IP address to geolocation service for use with the Google App Engine cloud platform.


Demo
----

A demonstration is available at:
[gae-ip-geolocation.appspot.com](https://gae-ip-geolocation.appspot.com/)


Highlights
----------

* [JSON](http://json.org/) data format
* Standard [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
value for country code
* Region, city, and latitude/longitude


Running
-------

In your web browser, navigate to the URL for your deployed application. For example:

    https://[your-app-id].appspot.com/api/ip.json

The response will contain the geolocation data for the IP address of your request. For example:

    {
      "Country": "DE",
      "Region": "be",
      "City": "berlin",
      "CityLatLong": "52.519171,13.406091"
    }


Notes
-----

This application is intended for use with Google App Engine. If you choose to upload an instance of the application to your Google App Engine account, then you must correctly configure your instance of the application and your overall Google App Engine account.

Use of this application on your Google App Engine account may incur billing charges. All such charges are your responsibility.


Future
------

Possible changes and additions could include:

* Support CORS: cross-origin resource sharing
* Support JSONP: JSON with padding
* Additional data fields for the geolocation


Development
-----------

Developed with:

* [Go](http://golang.org/) 1.1.2
* [Google App Engine](https://developers.google.com/appengine/) 1.8.9


Release History
---------------

0.9.1 - 18 January 2014

* Now uses specific API endpoint. Added basic web content for demo.

0.9.0 - 30 June 2013

* Initial version.
