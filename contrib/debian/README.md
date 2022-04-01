
Debian
====================
This directory contains files used to package redecoind/redecoin-qt
for Debian-based Linux systems. If you compile redecoind/redecoin-qt yourself, there are some useful files here.

## redecoin: URI support ##


redecoin-qt.desktop  (Gnome / Open Desktop)
To install:

	sudo desktop-file-install redecoin-qt.desktop
	sudo update-desktop-database

If you build yourself, you will either need to modify the paths in
the .desktop file or copy or symlink your redecoin-qt binary to `/usr/bin`
and the `../../share/pixmaps/redecoin128.png` to `/usr/share/pixmaps`

redecoin-qt.protocol (KDE)

