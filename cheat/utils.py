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

    @staticmethod
    def boolify(value):
        """ Type-converts 'true' and 'false' to Booleans """
        # if `value` is not a string, return it as-is
        if not isinstance(value, str):
            return value

        # otherwise, convert "true" and "false" to Boolean counterparts
        return value.strip().lower() == "true"
