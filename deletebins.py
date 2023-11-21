#!/usr/bin/env python
import os
import glob

# Get the current directory
current_directory = os.getcwd()

# Walk through each subdirectory in the current directory
for subdir, dirs, files in os.walk(current_directory):
    # For each .go file in the subdirectory
    for file in glob.glob(subdir + '/*.go'):
        # Get the file name without the .go extension
        base = os.path.splitext(file)[0]
        # If a file with the same name (without .go) exists, delete it
        if os.path.isfile(base):
            os.remove(base)