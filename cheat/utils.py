from __future__ import print_function
import sys


class Utils:

    @staticmethod
    def die(message):
        """ Prints a message to stderr and then terminates """
        Utils.warn(message)
        exit(1)

    @staticmethod
    def warn(message):
        """ Prints a message to stderr """
        print((message), file=sys.stderr)
