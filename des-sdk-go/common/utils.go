package common

import (
	"github.com/btcsuite/btcd/btcec"
	"errors"
	//"encoding/hex"
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/denkhaus/bitshares/config"
	"bytes"
	"crypto/sha256"
)

func VarifyParameters(params map[string]interface{}) bool {
	return true
}

func Sign(data *RequestParams, wifs []string) ([]string, error){
	config.SetCurrentConfig(config.ChainIDBTS)
	privKeys := make([]*types.PrivateKey, len(wifs))
	for index, wif := range wifs {
		privKeys[index], _  = types.NewPrivateKeyFromWif(wif)
	}

	var signatures []string
	for _, prv := range privKeys {
		ecdsaKey := prv.ToECDSA()
		if ecdsaKey.Curve != btcec.S256() {
			return nil, errors.New("Invalid PrivateKey")
		}
		for {
			serilData := Serilization(*data)
			digest, err := Digest(serilData)
			if err != nil {
				return nil, errors.New("Digest wrong")
			}
			sig, err := prv.SignCompact(digest)
			//fmt.Println("sign: ", hex.EncodeToString(sig), len(hex.EncodeToString(sig))) for Test
			if err != nil {
				return nil, errors.New("SignCompact")
			}

			if !isCanonical(sig) {
				data.Expiration += 1
			} else {
				signatures = append(signatures, types.Buffer(sig).String())
				//fmt.Println(signatures)
				break
			}
		}
	}
	return signatures, nil
}

func Digest(data []byte) ([]byte, error) {
	writer := sha256.New()
	//rawData, err := hex.DecodeString(data)
	rawData := data
	//if err != nil {
	//	return nil, errors.New("Decode failed")
	//}
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
func Verify(keyBag *KeyBag, data []byte, signatures []types.Buffer) (bool, error) {
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


func Serilization(data RequestParams) []byte{
	serilData := NewSignMessageOperation(data)
	var b bytes.Buffer
	enc := util.NewTypeEncoder(&b)
	serilData.Marshal(enc)
	return b.Bytes()
}

func Encrypt(privateKey string, publicKey string, nonce uint64, data []byte) string{
	cnf := config.CurrentConfig()
	prefixChain := cnf.Prefix()
	publicKey2 := prefixChain + publicKey[len(prefixChain):]
	to, err := types.NewPublicKeyFromString(publicKey2)
	if err != nil {
		panic("NewPublicKeyFromString failed")
	}
	priv, err := types.NewPrivateKeyFromWif(privateKey)
	if err != nil {
		panic("NewPrivateKeyFromWif failed")
	}
	msg := types.Buffer(data)
	//msg, err := types.BufferFromString(data)
	//msg := []byte(data)
	//fmt.Println(msg.String(), len(msg.String()), msg, len(msg))
	if err != nil {
		panic("data is wrong")
	}

	memo := types.Memo{
		From:    *priv.PublicKey(),
		To:      *to,
		Message: msg,
		Nonce:   types.UInt64(nonce),
	}
	//fmt.Println(msg.String(), len(msg.String()))
	if err := memo.Encrypt(priv, string(data)); err != nil {
		panic("encrypt failed")
	}
	return memo.Message.String()
}

func Decrypt(privateKey string, publicKey string, nonce int64, data string) string{
	cnf := config.CurrentConfig()
	prefixChain := cnf.Prefix()
	publicKey2 := prefixChain + publicKey[len(prefixChain):]
	to, err := types.NewPublicKeyFromString(publicKey2)
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