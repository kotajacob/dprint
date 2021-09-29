# dprint

Print specified values from desktop files to stdout.

## Usage

`dprint [-v] [-d path] [-i key:val] [-o key]`

Print specified values from desktop files to stdout. Look, it’s hard to describe
okay? Here’s a picture of me using it with dmenu.

![1](img.webp)

My launcher script pipes the output of dprint into dmenu to get a selection.
Then it passes that selection into dprint – with some options – and then the
output of that gets executed by your shell (to launch the program).

```sh
#!/bin/sh
SELECTION=$(dprint | dmenu -i -l 8 "$@")
echo "Name:$SELECTION" | dprint -i - -o "StripExec" | ${SHELL:-"/bin/sh"} &
```

I wrote dprint because the default `dmenu_run` script just lists all the
programs in your `$PATH` _exactly_ as they’re named. There’s no easy way to
rename them or tweak launch options. For example, that “calculator” program in
the screenshot runs `st -t st-float -g 76x30 -e python`, and I renamed “ncmpcpp”
to just “music.”

## Building

Install the dependencies:

- go (>=1.12)
- scdoc

Then compile dprint:

    $ make

## Installation

    # make install

## License

GPL3 - See License for details.

Copyright 2019 Dakota Walsh

## Resources

dprint uses a public mailing list for contributions and discussion. You can
browse the list [here](https://lists.sr.ht/~kota/public-inbox) and [email
patches](https://git-send-email.io) or questions to
[~kota/public-inbox@lists.sr.ht](https://lists.sr.ht/~kota/public-inbox).

If you're reporting an bug/feature request our issue tracker is here:
[~kota/dprint](https://todo.sr.ht/~kota/dprint)
