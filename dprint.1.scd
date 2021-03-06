dprint(1) [""["dprint-VERSION"]]

# NAME

dprint - print specified values from desktop files to stdout.

# SYNOPSIS

_dprint_ [-v] [-p] [-d path] [-i key:val] [-o key]

# OPTIONS

*-v*
	Prints the version and exits.

*-p*
	Keep track of desktop file popularity each time an Exec or StripExec is
	selected as output. Additionally, output is sorted by popularity when using
	this option.

*-d path*
	Look for desktop files in provided path.

*-i key:val*
	Filter desktop files with a specific key:value pair.

*-o key*
	Output the value associated with a specific key for each selected desktop
	file.

# DESCRIPTION

The primary use of dprint is to provide input for program launchers like dmenu,
but it was designed in such a way that it can be used as a generic desktop file
reader. First dprint reads a directory full of .desktop files which can be
provided with the *-d* option. Without specification it will look in
XDG_CONFIG_HOME/dprint/ so you can simply copy or link the specific files you
care about into that directory. The *-i* option is used to narrow the selection
of desktop files, you simply provide a key and value such as Name:Firefox and
dprint will ignore all files not containing that pair. Finally dprint will
print out a value from the selected file or files. If the *-o* option is
provided it will print the value associated with that specific key, otherwise
it will print the value of the "Name" key.

Often this program will be used by piping its output into fzf or dmenu and then
back into itself. To use STDIN for an option simply pass *-* after the
argument. For example *echo "Exec" | dprint -o -* will print the exec values of
all items.

# KEYS

Below is a list of all currently supported keys. Bool means the value may be
True or False. List means one or more strings seperated with semi-colons.

*Type*
	Type - The type of desktop entry. It can be: Application, Link, or Directory.

*Version*
	String - The version of spec that the file conforms to.

*Name*
	String - Real name of desktop entry.

*GenericName*
	String - Generic name of desktop entry.

*Comment*
	String - A short comment that describes the desktop entry.

*Icon*
	String - The name of an icon that should be used.

*URL*
	String - URL for a link type entry.

*NoDisplay*
	Bool - Whether or not to display the file in menus.

*Hidden*
	Bool - Whether the use has deleted the desktop entry.

*OnlyShowIn*
	List - A list of desktop environments that the desktop entry should only be
	shown in.

*NotShowIn*
	List - A list of desktop environments that the desktop entry should not be
	shown in.

*DBusActivatable*
	Bool - Whether DBus Activation is supported by this application.

*TryExec*
	String - The path to an executable to test if the program is installed.

*Exec*
	String - Program to execute.

*Path*
	String - The path that should be the programs working directory.

*Terminal*
	Bool - Whether the program should be in a terminal window.

*Actions*
	List - A list of actions.

*MimeType*
	List - List of mimetypes supported by this program.

*Categories*
	List - List of categories to be shown in in a menu.

*Implements*
	List - List of interfaces this application implements.

*Keywords*
	List - List of keywords.

*StartupNotify*
	Bool - Whether the program will send a "remove" message when started with the
	DESKTOP_STARTUP_ID env variable is set.

*StartupWMClass*
	String - The string that the program will set as WM Class or WM name hint.

# SPECIAL KEYS

These special keys do not actually exist in the desktop files and can only be
used in output mode.

*StripExec*
	String - Program to execute with field codes stripped out.

# EXAMPLES

Print name values from all files in ~/.config/dprint/
	*dprint*

Print exec values from all files in ~/Downloads/
	*dprint -d ~/Downloads/ -o "Exec"*

Print exec values without field codes for files matching "GenericName:Web
Browser"
	*echo "StripExec" | dprint -i "GenericName:Web Browser" -o -*

# AUTHORS

Maintained by Dakota Walsh <kota at nilsu.org>.
Up-to-date sources can be found at https://git.sr.ht/~kota/dprint/

# COPYRIGHT

Copyright (C) 2019 Dakota Walsh

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
