    # # Spend Locked Funds
    # from bitcoin import *
    # from bitcoin.wallet import CBitcoinSecret, P2PKHBitcoinAddress
    # # Private key and Bitcoin address from the previous step
    # private_key = CBitcoinSecret.from_secret_bytes(...)
    # address = P2PKHBitcoinAddress.from_pubkey(...)
    # # Create a transaction input (UTXO)
    # txid =  # Transaction ID of the UTXO you want to spend
    # output_index = ... # Index of the output in the transaction
    # txin = create_txin(txid, output_index)
    # # Create a transaction output to the desired destination
    # destination_address = ’...’ # Recipient’s address
    # amount_to_send = ... # Amount to send in satoshis
    # txout = create_txout(amount_to_send, destination_address)

    # # Create the transaction
    # tx = create_signed_transaction([txin], [txout], [private_key])
    # # Broadcast the transaction
    # broadcast_tx(tx)



from bitcoin.wallet import CBitcoinSecret, P2PKHBitcoinAddress, CBitcoinAddress
from bitcoin.core import CTransaction, CMutableTxOut, CMutableTxIn, COutPoint, lx, x, CMutableTransaction, Hash160, b2x
from bitcoin.core.script import OP_DUP, OP_HASH160, OP_EQUALVERIFY, OP_CHECKSIG
from bitcoin.core.script import CScript, SignatureHash, SIGHASH_ALL
from bitcoin.core.scripteval import VerifyScript, SCRIPT_VERIFY_P2SH
from bitcoin import *
from bitcoin.rpc import Proxy
import requests

SelectParams('testnet')

# Function to create a transaction input (UTXO)
def create_txin(txid, output_index):
    return CMutableTxIn(COutPoint(txid, output_index))

# Function to create a transaction output
# def create_txout(amount, destination_address):
    
#     destination_script = P2PKHBitcoinAddress(destination_address).to_scriptPubKey()
#     print('a')
#     return CMutableTxOut(nValue=amount, scriptPubKey=destination_script)
def create_txout(amount, destination_address):
  destination_address = CBitcoinAddress(destination_address)
  amount_less_fee = amount * 0.99
  return CMutableTxOut(amount_less_fee, destination_address.to_scriptPubKey())

# Function to create a signed transaction
def create_signed_transaction(txins, txouts, private_key,  amount_to_send):
    tx = CMutableTransaction(txins, txouts) 
    txin_scriptPubKey = CScript([OP_DUP, OP_HASH160, Hash160(private_key.pub), OP_EQUALVERIFY, OP_CHECKSIG])
    #txin_scriptPubKey = txouts[0].scriptPubKey
    sighash = SignatureHash(txin_scriptPubKey, tx, 0, SIGHASH_ALL, amount=amount_to_send)
    signature = private_key.sign(sighash) + bytes([SIGHASH_ALL])
    txins[0].scriptSig = CScript([signature, private_key.pub])
    VerifyScript(txins[0].scriptSig, txin_scriptPubKey, tx, 0, (SCRIPT_VERIFY_P2SH,))
    return tx

# Function to broadcast the transaction
def broadcast_tx(signed_tx):
    try:
        # Convert signed transaction to raw hex
        raw_transaction = signed_tx.serialize().hex()
        # Define API endpoint
        url = 'https://api.blockcypher.com/v1/btc/test3/txs/push'
        # Set headers
        headers = {'content-type': 'application/json'}
        # Prepare data for POST request
        data = {"tx": raw_transaction}
        # Send POST request to broadcast the transaction
        response = requests.post(url, headers=headers, json=data)
        # Print the response
        print(response.json())
    except Exception as e:
        print(f"Error broadcasting transaction: {e}")

# Private key and Bitcoin address from the previous step
privatek = 'cV49vnzrr51bmKgga1RQFn9edemZ6aJfwyqFCvqVNcgMZzHGHKpW'
private_key = CBitcoinSecret(privatek)

address = P2PKHBitcoinAddress.from_pubkey(private_key.pub)
print(address)
# UTXO details
txid = lx('cb4144c6bd4dc5fa2c113664834bf905810816750f7fe39528ec271c541b1106')
output_index = 0 # Index of the output in the transaction

# Create a transaction input (UTXO)
txin = create_txin(txid, output_index)

# Create a transaction output to the desired destination
destination_address = 'msusZj52UjCnRxU69Ze8a22EsmCED1Sjog'
amount_to_send = 0.00000000001 # Amount to send in satoshis

txout = create_txout(amount_to_send, destination_address)

# Create the transaction
tx = create_signed_transaction([txin], [txout], private_key, amount_to_send)
print(b2x(tx.serialize()))
# Broadcast the transaction
broadcast_tx(tx)
