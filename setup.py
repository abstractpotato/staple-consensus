import getpass
from sys import argv
from pycardano import (
    Address,
    HDWallet,
    PaymentExtendedVerificationKey,
    ExtendedSigningKey
)

# mnemonic_phrase = getpass.getpass("Enter mnemonic phrase: ")

mnemonic_phrase = "brick sample second sponsor excite churn rescue tower athlete sell text friend wet forest depend hobby intact fan distance badge height pigeon curious script"

try:
    print("generating signing key...")
    hdwallet = HDWallet.from_mnemonic(mnemonic_phrase)
    payment_path = "m/1852'/1815'/0'/0/0"
    hdwallet_spend = hdwallet.derive_from_path(payment_path)
    skey = ExtendedSigningKey.from_hdwallet(hdwallet_spend)
    evkey = skey.to_verification_key()
    vkey = evkey.to_non_extended()
    address = Address(payment_part=vkey.hash())

    skey.save(".env/key.skey")
    vkey.save(".env/key.vkey")

except Exception as e:
    print(e)
    exit()

with open(".env/skey", "wb") as file:
    file.write(skey.payload)

with open(".env/vkey", "wb") as file:
    file.write(vkey.payload)

with open(".env/akey", "wb") as file:
    file.write(address.to_primitive())
