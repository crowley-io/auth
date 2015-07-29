// Package otp provides a minimal framework for using TOTP one-time password as a
// secure webapp's signing mecanism.
//
// Currently, this package is a simple wrapper on top of "github.com/pquerna/otp" package.
// It aim to simplify the usage of this library by restricting its API to keep it elementary:
// simplicity is the ultimate sophistication.
//
// When you enable TOTP for an user, you must save the "secret" value from the key on a data storage.
// However, it is recommended to store this secret with an encryption mecanism in your datastore.
//
// Creating a key for an user is the first step required for an 2FA enrollment.
//
//  import (
//      "github.com/crowley-io/auth/otp"
//
//      "bytes"
//      "image/png"
//  )
//
//  key, err := otp.Generate("Example.com", "alice@example.com")
//
// Then, you can display a QR code to the user.
//
//  var b bytes.Buffer
//  i, err := key.Image(200, 200)
//  png.Encode(&b, i)
//  display(b.Bytes())
//
// After that, you need to verify that the operation was successful for the user.
//
//  user := getUser()
//  passcode := getPasscode()
//  secret := getSecret(user)
//
//  if otp.Validate(passcode, secret) {
//      // Success! save otp enrollment.
//      save(user, secret)
//  }
//
// Finally, validating a one-time password is very easy: just retrieve the user passcode and secret.
//
//  import "github.com/crowley-io/auth/otp"
//
//  passcode := getPasscode()
//  secret := getSecret("alice@example.com")
//
//  valid := otp.Validate(passcode, secret)
//
//  if valid {
//      // Success! continue login process.
//  }
package otp
