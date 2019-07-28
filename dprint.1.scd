dprint(1) [""["dprint-VERSION"]]

# NAME

dprint - print specified values from desktop files to stdout.

# SYNOPSIS

_dprint_ [-v] [-d path] [-i key:val] [-o key]

# OPTIONS

*-v*
	Prints the version and exits.
*-d path*
	Look for desktop files in provided path.
*-i key:val*
	Select only desktop files with a specific key:value pair.
*-o key*
	Ouput the value associated with this specific key for each selected 
	desktop file.

# DESCRIPTION

The primary use of dprint is to provide input for program launchers like dmenu,
but it was designed in such a way that it can be used as a generic desktop file
reader. First dmenu reads a directory full of .desktop files which can be
provided with the *-d* option. Without specification it will look in
XDG_CONFIG_HOME/dprint/ so you can simply copy or link the specific files you
care about into that directory. The *-i* option is used to narrow the selection
of desktop files, you simply provide a key and value such as Name:Firefox and
dprint will ignore all files not containing that pair. Finally dprint will
print out a value from the selected file or files. If the *-o* option is
provided it will print the value associated with that specific key, otherwise
it will print the value of the "Name" key.

# AUTHORS

Maintained by Dakota Walsh <kota at nilsu.org>.
Up-to-date sources can be found at https://git.sr.ht/~kota/dprint/

# COPYRIGHT

Copyright 2019 Dakota Walsh

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation
and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors
may be used to endorse or promote products derived from this software without
specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
