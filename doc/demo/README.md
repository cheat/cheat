# Automated demo/interactive tutorial

[Tuterm](https://github.com/veracioux/tuterm) is used to create a fully automated real-time demonstration of the program.
The demonstration is generated from the script `cheat.tut`, and it can be
played in the terminal by running `tuterm -m demo cheat.tut`.

To upload the demonstration, run `./asciinema_upload_and_create_svg.sh`. This
script will upload to asciinema and give you back the URL, and create the SVG
file `cheat_demo.svg`, both of which are meant to be included in the `README.md`.
The SVG file should be included in version control. If you open the script, you
will find some tweakable parameters.

The `cheat.tut` can also serve as an interactive tutorial, if you run `tuterm cheat.tut`.
