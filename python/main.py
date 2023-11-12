import sys
import os
sys.path.insert(0, os.getcwd() + '/ten')

import ten

def main():
    ten.convert_time("01:00:00PM")
    ten.convert_time("01:00:00AM")

if __name__ == "__main__":
    main()