package common

import (
	"github.com/btcsuite/btcd/btcec"
	"errors"
	"crypto/sha1"
	"encoding/hex"
	"github.com/denkhaus/bitshares/types"
)

func VarifyParameters(params map[string]interface{}) bool {
	return true
}

func Sign(data string, wifs []string) ([]types.Buffer, error){
	privKeys := make([]*types.PrivateKey, len(wifs))
	for index, wif := range wifs {
		privKeys[index], _  = types.NewPrivateKeyFromWif(wif)
		//if err != nil {
		//
		//}
	}
	var signatures []types.Buffer
	for _, prv := range privKeys {
		ecdsaKey := prv.ToECDSA()
		if ecdsaKey.Curve != btcec.S256() {
			return nil, errors.New("Invalid PrivateKey")
		}
		for {
			digest, err := Digest(data) //TODO what is digest
			if err != nil {
				return nil, errors.New("Digest wrong")
			}

			sig, err := prv.SignCompact(digest)
			if err != nil {
				return nil, errors.New("SignCompact")
			}

			if !isCanonical(sig) {
				continue
			} else {
				signatures = append(signatures, types.Buffer(sig))
				break
			}
		}
	}
	return signatures, nil
}

//TODO there is no 4 + 27
func Digest(data string) ([]byte, error) {

	writer := sha1.New()
	rawData, err := hex.DecodeString(data)
	if err != nil {
		return nil, errors.New("Decode failed")
	}

	if _, err := writer.Write(rawData); err != nil{
		return nil, errors.New("Write rawDara")
	}

	digest := writer.Sum(nil)

	return digest, nil
}

func isCanonical(sig []byte) bool {
	d := sig
	t1 := (d[1] & 0x80) == 0
	t2 := !(d[1] == 0 && ((d[2] & 0x80) == 0))
	t3 := (d[33] & 0x80) == 0
	t4 := !(d[33] == 0 && ((d[34] & 0x80) == 0))
	return t1 && t2 && t3 && t4
}


//Verify verifies the underlying transaction against a given KeyBag
func Verify(keyBag *KeyBag, data string, signatures []types.Buffer) (bool, error) {
	dig, err := Digest(data)
	if err != nil {
		return false, errors.New("Digest")
	}

	pubKeysFound := make([]*types.PublicKey, 0, len(signatures))
	for _, signature := range signatures {
		sig := signature.Bytes()

		p, _, err := btcec.RecoverCompact(btcec.S256(), sig, dig)
		if err != nil {
			return false, errors.New("RecoverCompact")
		}

		pub, err := types.NewPublicKey(p)
		if err != nil {
			return false, errors.New("NewPublicKey")
		}

		pubKeysFound = append(pubKeysFound, pub)
	}

	for _, pub := range pubKeysFound {
		if !keyBag.PublicPresent(pub) {
			return false, nil
		}
	}

	return true, nil
}


func Serilization(data interface{}) string{

	return ""
}

func Encrypt(privateKey string, publicKey string, nonce uint64, data string) string{
	to, err := types.NewPublicKeyFromString(publicKey)
	if err != nil {
		panic("NewPublicKeyFromString failed")
	}

	priv, err := types.NewPrivateKeyFromWif(privateKey)
	if err != nil {
		panic("NewPrivateKeyFromWif failed")
	}

	msg, err := types.BufferFromString(data)
	if err != nil {
		panic("data is wrong")
	}

	memo := types.Memo{
		From:    *priv.PublicKey(),
		To:      *to,
		Message: msg,
		Nonce:   types.UInt64(nonce),
	}

	if err := memo.Encrypt(priv, data); err != nil {
		panic("encrypt failed")
	}
	return memo.Message.String()
}

func Decrypt(privateKey string, publicKey string, nonce int64, data string) string{
	to, err := types.NewPublicKeyFromString(publicKey)
	if err != nil {
		panic("NewPublicKeyFromString failed")
	}

	priv, err := types.NewPrivateKeyFromWif(privateKey)
	if err != nil {
		panic("NewPrivateKeyFromWif failed")
	}

	msg, err := types.BufferFromString(data)
	if err != nil {
		panic("data is wrong")
	}

	memo := types.Memo{
		From:    *priv.PublicKey(),
		To:      *to,
		Message: msg,
		Nonce:   types.UInt64(nonce),
	}

	message, err := memo.Decrypt(priv)

	if err != nil {
		panic("encrypt failed")
	}
	return message
}