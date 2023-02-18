#!/usr/bin/env python3

import json

def main():
    with open("./dist/artifacts.json", "r") as file:
        d = json.load(file)
        print(d[0]["path"])
        

if __name__ == "__main__":
    main()
