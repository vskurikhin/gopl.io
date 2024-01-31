#!/usr/bin/env python3

import base58
import hashlib


# noinspection SpellCheckingInspection
def master_xprv_from_seed(seed_hex, network_byte=0x0488ADE4):
    # convert seed from hex to bytes
    seed_bytes = bytes.fromhex(seed_hex)

    # add network byte prefix
    extended_key = bytes([network_byte]) + seed_bytes

    # checksum = first 4 bytes of sha256(sha256(key))
    checksum = hashlib.sha256(hashlib.sha256(extended_key).digest()).digest()[:4]

    # append the checksum to the extended key
    extended_key += checksum

    # convert to base58
    result = base58.b58encode(extended_key)

    return result.decode('utf-8')


if __name__ == "__main__":
    seed_in_hex = input("Enter the seed in hex format: ")

    try:
        # noinspection SpellCheckingInspection
        xprv = master_xprv_from_seed(seed_in_hex)
        print("Master xprv:", xprv)
    except ValueError:
        print("Invalid seed input.")
