# Spend Funds from Multisig Address
from bitcoin import *
from bitcoin.wallet import CBitcoinSecret, CBitcoinAddress, CBitcoinSecret
from bitcoin.core import b2x, lx, COIN, COutPoint, CMutableTxOut, CMutableTxIn, CMutableTransaction
from bitcoin.core.scripteval import VerifyScript, SCRIPT_VERIFY_P2SH
from bitcoin.core.script import CScript, SignatureHash, SIGHASH_ALL, OP_2, OP_CHECKMULTISIG, OP_0
from bitcoin.rpc import Proxy

SelectParams('testnet')

def create_txin(txid, output_index):
  return CMutableTxIn(COutPoint(txid, output_index))

def create_txout(amount, destination_address):
  destination_address = CBitcoinAddress(destination_address)
  amount_less_fee = amount * 0.99
  return CMutableTxOut(amount_less_fee, destination_address.to_scriptPubKey())

def create_signed_transaction(txins, txouts, priv_keys, redeem_script, amount):
  # Create the unsigned transaction
  tx = CMutableTransaction(txins, txouts)
  # Calculate the signature hash for that transaction
  sighash = SignatureHash(redeem_script, tx, 0, SIGHASH_ALL, amount=amount)
  # Create the first signature using the first private key
  sig1 = priv_keys[0].sign(sighash) + bytes([SIGHASH_ALL])
  # Create the second signature using the second private key
  sig2 = priv_keys[1].sign(sighash) + bytes([SIGHASH_ALL])
  # Create the final scriptSig unlocking script
  txin_scriptPubKey = redeem_script.to_p2sh_scriptPubKey()
  txin_scriptSig = CScript([OP_0, sig1, sig2, redeem_script])
  # Set the scriptSig of our transaction input
  txins[0].scriptSig = txin_scriptSig
  # Verify the input script
  VerifyScript(txin_scriptSig, txin_scriptPubKey, tx, 0, (SCRIPT_VERIFY_P2SH,))
  # Return the signed transaction
  return tx

def broadcast_tx(tx):
  p = Proxy(service_url='http://localhost:18332')
  p.sendrawtransaction(tx)
  print(lx(tx.GetTxid()))

# Private keys and multisig address from the previous step
private_key1_str = 'cQZKGVczDdMT55KEWhaShakkw4CGufBeg5bAuuCut6Vb62rhtuuv'
private_key2_str = 'cQX2BhkXWY57c6sNm65s9ZgAVpsgUnP3Nb9UgLqJyRRNc6btyBAq'
private_key1 = CBitcoinSecret(private_key1_str)
private_key2 = CBitcoinSecret(private_key2_str)
address = '2NDmV9VBMSvmLEu6XFqBUijQQ8LCHExAr5z'
public_key1 = private_key1.pub
public_key2 = private_key2.pub

# Create a transaction input (UTXO)
txid = lx('ba88c09c9abbb5aa739e6c3fb0448bf18347aca225fd4a5daf90bcc6e1e44dba') # Transaction ID of the UTXO you want to spend
output_index = 1 # Index of the output in the transaction
txin = create_txin(txid, output_index)

# Create a transaction output to the desired destination
destination_address = 'mv4rnyY3Su5gjcDNzbMLKBQkBicCtHUtFB' # Recipient’s address
amount_to_send = 1000 # Amount to send in satoshis
txout = create_txout(amount_to_send, destination_address)

# Create the transaction
redeem_script = CScript([OP_2, public_key1, public_key2, OP_2, OP_CHECKMULTISIG])
tx = create_signed_transaction([txin], [txout], [private_key1, private_key2], redeem_script, amount_to_send)

print(b2x(tx.serialize()))

# Broadcast the transaction
broadcast_tx(tx)