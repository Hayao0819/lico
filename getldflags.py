#!/usr/bin/env python3

import yaml

def main():
    with open("./.goreleaser.yaml", "r") as file:
        file = yaml.load(file, Loader=yaml.FullLoader)
        ldflags = file["builds"][0]["ldflags"]
        print(" ".join(ldflags))

if __name__ == "__main__":
    main()
