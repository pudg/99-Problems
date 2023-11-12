import sys
import os
sys.path.insert(0, os.getcwd() + '/ten')

import ten

def main():
    ten.sparse_arrays(['ab', 'ab', 'abc'], ['ab', 'abc', 'bcd'])
    ten.lonely_int([8, 2, 3, 3, 3, 2, 1])

if __name__ == "__main__":
    main()