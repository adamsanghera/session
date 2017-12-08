# Welcome!

This package is used to manage user sessions in generic web applications, where managing access control is a desirable feature.

# Example flow

0. You've already established the various kinds of sessions you're interested in tracking, by making calls to `yourSession := session.NewSession`.
1. User has successfully logged in, by jumping over whatever hurdles you've constructed.
2. You create a corresponding session entry, with `yourSession.Begin(uname)`, which gives you back the token's time to live duration, and the token itself.
3. The user tries to do some activity with controlled access.
4. You call `yourSession.Validate(uname, providedToken)` to validate a token.
5. Your user's browser, realizing that their token's time is coming up, requests a renewal.
6. You call `yourSession.Renew(uname, providedToken)`, which will return a brand new token and time-to-live.

# Roadmap of Features

## Done

1. Allow for custom token parameters
  * Tokens that expire at different times, are different lengths, etc. 
2. Allow for custom tokens
  * Tokens that grant access to X activity or Y activity.
    * We can encode this information in a couple of ways.  For example by making separate DB indices, or by hashing the activity name into the token.
  * Allow users to possess more than 1 token.

## Yet to come

1. SecureSession, which only stores hashes of usernames/ids.

# Interesting issue...

What if we want to make two different kinds of token for the same username?

We would need separate redis DBs, one for each user. <b>Or</b>, we could just use username_tokenName as the index.  Way better idea.