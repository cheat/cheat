# Print file metadata etc.
ffmpeg -i path/to/file.ext

# Convert all m4a files to mp3
for f in *.m4a; do ffmpeg -i "$f" -acodec libmp3lame -ab 320k "${f%.m4a}.mp3"; done

# Listen to 10 seconds of audio from a video file
#
# -ss : start time
# -t  : seconds to cut
# -autoexit : closes ffplay as soon as the audio finishes
ffmpeg -ss 00:34:24.85 -t 10 -i path/to/file.mp4 -f mp3 pipe:play | ffplay -i pipe:play -autoexit
