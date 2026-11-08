package bip39

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerdaks/bip-0039/data"
)

func TestBytesToBinStr(t *testing.T) {
	expected := "00000100"
	actual := bytesToBinStr([]byte{0x04})
	assert.Equal(t, expected, actual)
}

func TestBinToInt(t *testing.T) {
	assert.Equal(t, 4, binToInt("00000100"))
	assert.Equal(t, 2047, binToInt("11111111111"))
}

func TestChecksum128(t *testing.T) {
	entropy := make([]byte, 16)
	expected := "0011"
	actual := checksum(entropy)
	assert.Equal(t, expected, actual)
}

func TestChecksum256(t *testing.T) {
	entropy := make([]byte, 32)
	expected := "01100110"
	actual := checksum(entropy)
	assert.Equal(t, expected, actual)
}

func TestValidate(t *testing.T) {
	wordList := make([]string, 2048)

	cases := []struct {
		name     string
		entLen   int
		wordList []string
		wantErr  bool
	}{
		{"96 bits is too short", 96, wordList, true},
		{"100 bits is not a multiple of 32", 100, wordList, true},
		{"128 bits is valid", 128, wordList, false},
		{"256 bits is valid", 256, wordList, false},
		{"288 bits is too long", 288, wordList, true},
		{"2047 words is too few", 128, wordList[:2047], true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := validate(c.entLen, c.wordList)
			if c.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGenerateMnemonic128Words(t *testing.T) {
	wordList, err := LoadWordList(data.FS, "words.txt")
	assert.NoError(t, err)

	mnemonic, err := GenerateMnemonic(128, wordList)
	assert.NoError(t, err)
	assert.Equal(t, 12, len(mnemonic))

	mnemonic, err = GenerateMnemonic(192, wordList)
	assert.NoError(t, err)
	assert.Equal(t, 18, len(mnemonic))

	mnemonic, err = GenerateMnemonic(256, wordList)
	assert.NoError(t, err)
	assert.Equal(t, 24, len(mnemonic))
}

// Unfilled entropy is all zeros, which still yields a valid mnemonic, so only
// comparing two outputs catches it.
func TestGenerateMnemonicIsRandom(t *testing.T) {
	wordList, err := LoadWordList(data.FS, "words.txt")
	assert.NoError(t, err)

	first, err := GenerateMnemonic(256, wordList)
	assert.NoError(t, err)
	second, err := GenerateMnemonic(256, wordList)
	assert.NoError(t, err)

	assert.NotEqual(t, first, second)
}

func TestGenerateMnemonicRejectsShortWordList(t *testing.T) {
	wordList := make([]string, 2047)

	mnemonic, err := GenerateMnemonic(128, wordList)
	assert.Error(t, err)
	assert.Nil(t, mnemonic)
}

// Vectors from https://github.com/trezor/python-mnemonic/blob/master/vectors.json
func TestMnemonicFromEntropyVectors(t *testing.T) {
	wordList, err := LoadWordList(data.FS, "words.txt")
	assert.NoError(t, err)

	vectors := []struct {
		entropy  string
		mnemonic string
	}{
		{
			"00000000000000000000000000000000",
			"abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about",
		},
		{
			"7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
			"legal winner thank year wave sausage worth useful legal winner thank yellow",
		},
		{
			"808080808080808080808080808080808080808080808080",
			"letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter always",
		},
		{
			"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			"zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo vote",
		},
		// the vectors above repeat one byte, so they cannot catch byte-order bugs
		{
			"9e885d952ad362caeb4efe34a8e91bd2",
			"ozone drill grab fiber curtain grace pudding thank cruise elder eight picnic",
		},
		{
			"6610b25967cdcca9d59875f5cb50b0ea75433311869e930b",
			"gravity machine north sort system female filter attitude volume fold club stay feature office ecology stable narrow fog",
		},
		{
			"68a79eaca2324873eacc50cb9c6eca8cc68ea5d936f98787c60c7ebc74e6ce7c",
			"hamster diagram private dutch cause delay private meat slide toddler razor book happy fancy gospel tennis maple dilemma loan word shrug inflict delay length",
		},
	}

	for _, v := range vectors {
		entropy, err := hex.DecodeString(v.entropy)
		assert.NoError(t, err)

		mnemonic, err := MnemonicFromEntropy(entropy, wordList)
		assert.NoError(t, err)
		assert.Equal(t, strings.Fields(v.mnemonic), mnemonic)
	}
}

func TestMnemonicFromEntropyRejectsBadLength(t *testing.T) {
	wordList, err := LoadWordList(data.FS, "words.txt")
	assert.NoError(t, err)

	mnemonic, err := MnemonicFromEntropy(make([]byte, 15), wordList)
	assert.Error(t, err)
	assert.Nil(t, mnemonic)
}
