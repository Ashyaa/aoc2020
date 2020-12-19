#!/usr/bin/env python

from datetime import datetime
from day19 import solution

if __name__ == "__main__":
    startTime = datetime.now()
    solution.run()
    print(f"\nRuntime: {datetime.now() - startTime}µs")