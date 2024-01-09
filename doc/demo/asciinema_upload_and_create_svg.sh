#!/usr/bin/env sh

# NOTE: Change this to suit your needs

TERM_WIDTH=80
TERM_HEIGHT=18
RECORDED_COMMAND="tuterm cheat.tut --mode demo"
alias copy='xsel -b'

# Dependencies:
# - asciinema (https://github.com/asciinema/asciinema)
# - svg-term (https://github.com/marionebl/svg-term-cli)
# - xsel

# Tuterm can be found here:
#   https://github.com/veracioux/tuterm

rm -f /tmp/cheat.cast

stty cols "$TERM_WIDTH" rows "$TERM_HEIGHT"
# Record the command
asciinema rec -c "$RECORDED_COMMAND" /tmp/cheat.cast

# Change terminal width and height
# NOTE: for some reason the yes command prints Broken pipe; this is a workaround
sed -e "1 s/\(\"width\": \)[0-9]\+/\1$TERM_WIDTH/" \
    -e "1 s/\(\"height\": \)[0-9]\+/\1$TERM_HEIGHT/" \
    -e '/Broken pipe/d' \
    -i /tmp/cheat.cast

# Upload to asciinema.org
output="$(asciinema upload /tmp/cheat.cast)"
echo "$output"

# Copy URL to clipboard
echo "$output" | grep 'https:' | sed 's/^\s*//' | copy

# Create local SVG animation
cat /tmp/cheat.cast | svg-term --out cheat_demo.svg

echo "SVG animation saved as 'cheat_demo.svg'"
