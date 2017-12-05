# Welcome!

This package is used to manage user sessions in generic web applications, where managing access control is a desirable feature.

# Example flow

1. User has successfully logged in, by jumping over whatever hurdles you've constructed.
2. You call `session.Create()` to create a new session, assigned to that user.
3. The user tries to do some activity with controlled access.
4. You call `session.Validate()` to validate a token.

# Roadmap of Features

1. Allow for custom token parameters
  * Tokens that expire at different times, are different lengths, etc. 
2. Allow for custom tokens
  * Tokens that grant access to X activity or Y activity.
    * We can encode this information in a couple of ways.  For example by making separate DB indices, or by hashing the activity name into the token.
  * Allow users to possess more than 1 token.