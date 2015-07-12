# Katas

This package is a collection of katas I've worked on at various times.

## Unsplice

A kata from [CyberDojo](http://cyber-dojo.org/) worked on at the July 13, 2015 meeting of the [Akron Code Club](http://www.meetup.com/AkronCodeClub/events/222461350/).

```
Given a string, strip all occurences of consecutively
occuring backslash and newline characters. For example,
assuming that:
"\\" represents '\' and
"\n" represents '\n'

  "ab\\\ncd\\\nef" --> "abcdef" (two stripped out)

  "abc\\\ndef"     --> "abcdef" (one stripped out)

  "abc\n\\def"     --> unchanged (wrong order)

  "abc\\def"       --> unchanged (no \n)

  "abc\ndef"       --> unchanged (no \)

  "abcdef"         --> unchanged
```
