import sys
import os
sys.path.insert(0, os.getcwd() + '/ten')

import ten

def main():
    ten.min_max_sum([1, 3, 5, 7, 9])

if __name__ == "__main__":
    main()