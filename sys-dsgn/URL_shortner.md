# URL Shortening System Design

1. Features(Lenth of URL, domain) - Functional and non functional requirement.
2. Estimates (CPU, memory)
3. Design goal - tradeoff (CAP)
4. HLD
5. Scalling

Functional Requirement

- Shorten - shorten the URL
- Redirection - should redirect the shorten URL to actual URL
- Expiry - should have an expiry of of URL shortened.

Non Functional Requirement

- Highly Available
- low latency
- Secure - URL should not be predictable.


Estimations 

- Read heavy  -> 100:1 ::Read:Write
- Every month 100 M redirection and 1M new URL's
- QPS(query per second)
	- 100,000,000/ 30*24*60*60  ~= 50 qps
- Write - 1 M  every month new URL's
	- Expiry - 10 Years
	- Total  - 1,000,000 *12*10 *500(byte - size of 1 URL) ~= 60 GB
- RAM
	- 1 day TTL 
		- 50 qps *24*60*60 ~=   5 M quey per day
		- 25% cache ~=  1.25 M * 500 bytes = 1 GB RAM

Design Goal

- Read intensive
- Highly Available
- Low Latency
- Secure

High Level Design
- API's
	- shortenUrl (originalUrl, userID, apiKey, expiry) String
	- Redirection(shortUrL) String
- Database
	- MongoDB
		- URL (original_url, shorten_url, userId, expiry_data, created_at)
		- User(userid, name, email, apikey, createdat)
- Algorithms
	- long -> short 
	- avoid collisions
- Design
