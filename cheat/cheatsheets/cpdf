# Read in.pdf, select pages 1, 2, 3 and 6, and write those pages to
# out.pdf
cpdf in.pdf 1-3,6 -o out.pdf

# Select the even pages (2, 4, 6...) from in.pdf and write those pages
# to out.pdf
cpdf in.pdf even -o out.pdf

# Using AND to perform several operations in order, here merging two
# files together and adding a copyright stamp to every page.
cpdf -merge in.pdf in2.pdf AND -add-text "Copyright 2014" -o out.pdf

# Read control.txt and use its contents as the command line arguments
# for cpdf.
cpdf -control control.txt

# Merge in.pdf and in2.pdf into one document, writing to out.pdf.
cpdf -merge in.pdf in2.pdf -o out.pdf

# Split in.pdf into ten-page chunks, writing them to Chunk001.pdf,
# Chunk002.pdf etc
cpdf -split in.pdf -o Chunk%%%.pdf -chunk 10

# Split in.pdf on bookmark boundaries, writing each to a file whose
# name is the bookmark label
cpdf -split-bookmarks 0 in.pdf -o @N.pdf

# Scale both the dimensions and contents of in.pdf by a factor of two
# in x and y directions.
cpdf -scale-page "2 2" in.pdf -o out.pdf

# Scale the pages in in.pdf to fit the US Letter page size, writing to
# out.pdf
cpdf -scale-to-fit usletterportrait in.pdf -o out.pdf

# Shift the contents of the page by 26 pts in the x direction, and 18
# millimetres in the y direction, writing to out.pdf
cpdf -shift "26pt 18mm" in.pdf -o out.pdf

# Rotate the contents of the pages in in.pdf by ninety degrees and
# write to out.pdf.
cpdf -rotate-contents 90 in.pdf -o out.pdf

# Crop the pages in in.pdf to a 600 pts by 400 pts rectangle.
cpdf -crop "0 0 600pt 400pt" in.pdf -o out.pdf

# Encrypt using 128bit PDF encryption using the owner password 'fred'
# and the user password 'joe'
cpdf -encrypt 128bit fred joe in.pdf -o out.pdf

# Decrypt using the owner password, writing to out.pdf.
cpdf -decrypt in.pdf owner=fred -o out.pdf

# Compress the data streams in in.pdf, writing the result to out.pdf.
cpdf -compress in.pdf -o out.pdf

# Decompress the data streams in in.pdf, writing to out.pdf.
cpdf -decompress in.pdf -o out.pdf

# List the bookmarks in in.pdf. This would produce:
cpdf -list-bookmarks in.pdf

# Outputs:

# Add bookmarks in the same form from a prepared file bookmarks.txt to
# in.pdf, writing to out.pdf.
cpdf -add-bookmarks bookmarks.txt in.pdf -o out.pdf

# Use the Split style to build a presentation from the PDF in.pdf,
# each slide staying 10 seconds on screen unless manually advanced.
# The first page, being a title does not move on automatically, and
# has no transition effect.
cpdf -presentation in.pdf 2-end -trans Split -duration 10 -o out.pdf

# Stamp the file watermark.pdf on to each page of in.pdf, writing the
# result to out.pdf.
cpdf -stamp-on watermark.pdf in.pdf -o out.pdf

# Add a page number and date to all the pages in in.pdf using the
# Courier font, writing to out.pdf
cpdf -topleft 10 -font Courier -add-text "Page %Page\nDate %d-%m-%Y" in.pdf -o out.pdf

# Two up impose the file in.pdf, writing to out.pdf
cpdf -twoup-stack in.pdf -o out.pdf

# Add extra blank pages after pages one, three and four of a document.
cpdf -pad-after 1,3,4 in.pdf -o out.pdf

# List the annotations in a file in.pdf to standard output.
cpdf -list-annotations in.pdf

# Might Produce:

# -- # Annotation text content 1 # -- # -- # Annotation text content 2
# --

# Copy the annotations from from.pdf to in.pdf, writing to out.pdf.
cpdf -copy-annotations from.pdf in.pdf -o out.pdf

# Set the document title of in.pdf. writing to out.pdf.
cpdf -set-title "The New Title" in.pdf -o out.pdf

# Set the document in.pdf to open with the Acrobat Viewer's toolbar
# hidden, writing to out.pdf.
cpdf -hide-toolbar true in.pdf -o out.pdf

# Set the metadata in a PDF in.pdf to the contents of the file
# metadata.xml, and write the output to out.pdf.
cpdf -set-metadata metadata.xml in.pdf -o out.pdf

# Set the document in.pdf to open in Acrobat Viewer showing two
# columns of pages, starting on the right, putting the result in
# out.pdf.
cpdf -set-page-layout TwoColumnRight in.pdf -o out.pdf

# Set the document in.pdf to open in Acrobat Viewer in full screen
# mode, putting the result in out.pdf.
cpdf -set-page-mode FullScreen in.pdf -o out.pdf

# Attach the file sheet.xls to in.pdf, writing to out.pdf.
cpdf -attach-file sheet.xls in.pdf -o out.pdf

# Remove any attachments from in.pdf, writing to out.pdf.
cpdf -remove-files in.pdf -o out.pdf

# Blacken all the text in in.pdf, writing to out.pdf.
cpdf -blacktext in.pdf -o out.pdf

# Make sure all lines in in.pdf are at least 2 pts wide, writing to
# out.pdf.
cpdf -thinlines 2pt in.pdf -o out.pdf

