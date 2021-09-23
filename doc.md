# Rate Limiting the reset password endpoint

The reset password endpoint may be susceptible to a plethora of attacks, two of which are listed below :

- DDOS - An attacker can flood our servers with requests to this endpoint and starve legitimate users of server resources.

- User enumeration attack - An attacker can check the existance of a user by repeatedly invoking the endpoint with different usernames.

Other attacks are listed [here](https://infosecwriteups.com/all-about-password-reset-vulnerabilities-3bba86ffedc7). (Since we use firebase authentication I guess it is safe to assume that google has done most of the heavy lifting in mititgating most of the vulnerabilities listed here) 

A solution to prevent the two attacks listed earlier would be to rate limit the usage of the reset password endpoint. A few of the rate limiting strategies that could be useful to us :

- Exponential backoff - Delay the retry request exponentially with subsequent failed attempts, similar to what iPhone/Android phones do when you enter the wrong ping multiple times. But in case of rate limiting the reset password endpoint the exponentaial backoff can even be applied to successful attempts to prevent user enumeration attack.

    > In response to rate-limiting, intermittent, or non-specific errors, a client should generally retry the request after a delay. It is a best practice for this delay to increase exponentially after each failed request, which is referred to as exponential backoff. When many clients might be making schedule-based requests (such as fetching results every hour), additional random time (jitter) should be applied to the request timing, the backoff period, or both to ensure that these multiple client instances don't become periodic thundering herd, and themselves cause a form of DDoS.
    >  -- <cite>https://cloud.google.com/architecture/rate-limiting-strategies-techniques</cite>

For other rate limiting techniques please over head to [this link](https://cloud.google.com/architecture/rate-limiting-strategies-techniques?authuser=0#techniques-enforcing-rate-limits). 

## Implementation

We could rate limit access to the API based of the IP address present in the request. The `RemoteAddr` field in the `http.Request` object could be used to access the client's IP address(other http headers could also be inspected - https://stackoverflow.com/questions/27234861/correct-way-of-getting-clients-ip-addresses-from-http-request).

One question that remains is - Do we also need rate limiting based on the username? 

Rate limiting can be implemented in differnt layers in a microservice architecture:

- Ingress - Rate limits can be implemented in the ingress controller such as nginx. (But if we decide to go ahead with the exponential backoff I don't really know how easy/difficult of an endavour this would be 😅)

- Gateway/Kong? - https://docs.konghq.com/hub/kong-inc/rate-limiting/

- http middleware - The http handler repsonsible for resetting the password could be wrapped in a middleware which does the rate limiting.

