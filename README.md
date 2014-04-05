A simple GeoIP API proxy for OS Monitor   
==============

### Introduction

This is a sub-project of OS Monitor to build a proxy server which improves speed and security for GeoIP API query, because every query is a HTTPS request that pass through a server, user's IP and information is hidden, the query result will be cached for 1 week, if anyone query same IP, the result can be reused.    

### How it works?

The Architecture 

       +-------+      +-------+       +-------+
       | Nginx | ==>  |  API  | <==>  | Redis |
       +-------+      +-------+       +-------+
                         ||
                      +-----------------+
                      | External GeoAPI |
                      +-----------------+



- API is written by Go, it gets information from external GeoAPI and push new query result into Redis.
- Nginx is running as an SSL termination and reversed proxy, ([Strong Encryption](https://www.ssllabs.com/ssltest/analyze.html?d=osmonitor.mobi))
- Redis is a database to keep data for 1 week.


### Is it free ?
Don't worry! I use donation to keep server's running, that costs $5 US per month, SSL certificate and domain cost $10.97 US, the current donation already can support 1 year.      

If you would like to support this server, you still can donate. :)
 
[![Donate using PayPal](https://www.paypalobjects.com/en_US/i/btn/btn_donate_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=FSDWJ92W9MBEN&lc=US&item_name=Donate%20To%20OS%20Monitor&item_number=0&currency_code=USD&bn=PP%2dDonationsBF%3abtn_donateCC_LG%2egif%3aNonHosted "Donate using PayPal")
  