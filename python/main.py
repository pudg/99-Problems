import sys
import os
sys.path.insert(0, os.getcwd() + '/ten')

import ten

def main():
    ten.plus_minus([-3, -2, -1, 0, 1, 2, 3])

if __name__ == "__main__":
    main()