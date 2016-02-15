# check the file's charactor code
nkf -g test.txt

# convert charactor code to UTF-8
nkf -w --overwrite test.txt

# convert charactor code to EUC-JP
nkf -e --overwrite test.txt

# convert charactor code to Shift-JIS
nkf -s --overwrite test.txt

# convert charactor code to ISO-2022-JP
nkf -j --overwrite test.txt

# convert newline to LF
nkf -Lu --overwrite test.txt

# convert newline to CRLF
nkf -Lw --overwrite test.txt

# convert newline to CR
nkf -Lm --overwrite test.txt

# MIME encode
echo テスト | nkf -WwMQ

# MIME decode
echo "=E3=83=86=E3=82=B9=E3=83=88" | nkf -WwmQ
