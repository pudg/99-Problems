import sys
import os
sys.path.insert(0, os.getcwd() + '/ten')
sys.path.insert(0, os.getcwd() + '/twenty')

import ten, twenty

def main():
    print(twenty.pangram("We promptly judged antique ivory buckles for the next prize"))

if __name__ == "__main__":
    main()