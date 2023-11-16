import sys
import os
sys.path.insert(0, os.getcwd() + '/ten')
sys.path.insert(0, os.getcwd() + '/twenty')

import ten, twenty

def main():
    print(twenty.pangram("We promptly judged antique ivory buckles for the next prize"))
    print(twenty.array_perm([0, 1], [0, 2], 1))
    print(twenty.subarr_division([2, 2, 1, 3, 2], 4, 2))
if __name__ == "__main__":
    main()