# cloudflare-updater
Bot to dynamicly change dns records on cloudflare.

# Requeriments
* Domain in cloudflare

# Use case
* Your server does not have fixed IPv6 prefix

# Usage
* Be sure to have a DNS record created as this program only updates records. It will not create the record for you.
* Create a secrets.txt file (with space separated values)
* Add your datails in secretes file in this order: version, zone id, record id, email, API key, record type, domain/sub-domain
  * How to get zone id: Go to dash.cloudflare.com, sign in, click on the website you want. The Home page should show the Zone ID under API at the bottom right part of the page.
  * If you don't know your record id you can just use auto and the program will do the rest
  * How to get zone id: Go to dash.cloudflare.com/profile/api-tokens, sign in. Create your token (you can use the edit DNS in zone template) or use your global key.
  * For record type use A for ipv4 and AAAA for ipv6
* Create a service to run the program when your internet connection restarts (Working on the guides for that)
