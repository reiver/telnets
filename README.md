# **telnets**

User interface to the TELNETS protocol. TELNETS is the secure version of the TELNET protocol.

TELNETS is the TELNET protocol over a secure TLS (or SSL) connection.


## Usage
```
telnets host [port]
```


## TELNETS vs SSH

Note that TELNETS and SSH are not the same thing.

SSH is considered a secure alternative to the (un-secure) TELNET protocol.
(And note that is "TELNET" without an "S" at the end.)

TELNETS (with an "S" at the end) is the secure version of (un-secure) the TELNET protocol.


## What does the "S" at the end mean?

The "S" at the end of "TELNETS" standards for "secure".
Just like the "S" at the end of "HTTPS".


## openssl ?

*But wait...* you might be thinking... *isn't TELNETS the same as call to `openssl s_client -host $1 -port $2`?*

**Absolutely not!**

The TELNET and TELNETS have special binary escape codes and control codes that `openssl` does **not** understand.

For example, in the TELNET and TELNETS protocols, byte value `255` has a special meaning, and is called `IAC`
(which is short for "interpret as command").

If byte value `255` is sent as data, then it **must** be "escaped" by having two *IACs* in a row.
(So `IAC` becomes `IAC IAC`; or in other words `255` becomes `255 255`.)

If this is not done, it will corrupt the TELNET and TELNETS data stream!

But that's just one example. This would also be a control sequence: `255 251 24`.
(This happens to mean `IAC WILL TERMINAL-TYPE`.)

If this happens to be in the data, then it must be escaped, or it will corrupt the
TELNET and TELNETS data stream.
