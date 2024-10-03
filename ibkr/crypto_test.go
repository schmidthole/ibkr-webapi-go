package ibkr

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testRsaPemContents = `-----BEGIN RSA PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDIkX/edBMJ2hHM
8cOg/PJWhQ/BtCOeuHV1P6GxNosVJ3lMiD5+H4K/2YFyMswJ/FIkliUpNTgRSmmv
Nli9+DUz9EiS+FxLdj0db6hM9EYy1RR1oyDi4/Qc2bpguppP43fjFfRSf8mvWyNY
VsSZ+LYwdFQhzWJ3AAgxP9/F9oOLSYDpNmjO7HQy9GjReS9HqMKTXD0PlyxS6qCl
lkmRg2YizfJqtDVrIioFxwo5+Kl65RvrdlA6AwNNz3+u3zituAUOehTuMJwU1qlb
MH5+tLUU2yUna8/Pn2n7iTMaXsucYzm5BJgmt7zpHGh5oPzrSrUvTVIrmjhOa8gX
C8LZ0n3LAgMBAAECggEAL135YlpsYp+/I3UXW+5Q9zDukzMNVH9AsiOIdrcoaqML
aREmHRgpA1EFWqh3eoY+CgjxlsVMS/0Ub344wAid2/q9pwQDnAi9DzYe+6mW6Itd
DRks+5iVPrkG5LA9IE/asOUu8crfIUw3UTYQsUJFGLVoZyE/t1XpGWR53yFEWMDk
qMHSWC25X4C69v3J/ItvCsTd4U9eHI2A3//Bf5Zw6Wq+8w2bRzDfGWBA1X+/7g4B
4fz04MtRsIze6hXMxKqHQevuF4dSWqFCpD+oaNbBAZJbFjnZXcrNJ85hKnJEMgby
Qw64x/y1GLLoEgQc9YHzFLH8PQGLTUCdbUejHdpY3QKBgQDlfsA/Mq8/AJdX5T0M
4bBHwzZ97lzEPp3WF2hhOAtT1/dQw1PafQZ0IA38QWMwojvCCWQug2DfWRLym//K
mDBmKCAcj+xq5EC29L52QzfGAw9mp75it3ixqLhX9qkfO2Jkl/l383rHhkQtJYt1
rQvIEKGj0qDLsqkYiT15b2rnXQKBgQDfu4AToOK/uZPUvwBu39kTrHau7QG/5oXy
ZonphKXneOZr1eUEpt4NaDodIlip1UvxE5pZqAI0smBhuQXwGTJfxFVjaqxDv6XU
RayqEcVmFkOVRWY6Q5/NjyuUQnV6jh9WBryUTvfFDrA5l9AcGBj09zinXAeRulYY
xphmlZdvRwKBgDf3/IytIL4dt8sw1f3BQMuPmUeHDUqy9q1S0hr7nw+YmPEyYc3k
NeG2O5Lso3Ml7ar5vVv2oHZ3874FO5G+IfeYC5TLqN28wuebsqOG6yA4eNYRTMYo
dp/lROf2+PHr+qgLxERztdrZLOFRybwuMrRdEf6xDJQBA2o/diWafZjlAoGAa3et
NcB6Ko4BrJdypc9wfjrBZ5UA2iTpAl6usDS/XEh+ph3APFPLJbWC9Kui2zFRoSNB
aP+Cz70+/rHZYFM6FzY/rdl3lbNg/G3bWnXe1AKtw+3qBOKmyB6VWEn0FEJgWvox
ayvjaAx+Kav6kwQgOMrns7xJBn7zDEOYwO3qJ9ECgYAMGgYXv1K7ypxWX96IIfx7
l6k19aq0Nvwjpu+QIleGjpX+yLJcXWvLD45IxZRs1QQU5C8qbQnsBChbw8HV6jt1
Bv0WIaPQPGr0oHF24AniC+NU2sV6FqbdWpwnV15Mz3ebjHEvO+PKlr1cCgtGUXak
xvIEbObPztwf/G9J918X3A==
-----END RSA PRIVATE KEY-----`

var testDhParamsContents = `-----BEGIN DH PARAMETERS-----
MIIBDAKCAQEA3PhSV5Pu8opjdYguXMxyaqZ7dK8Lx1tfd++XBzHmuSCx7gE8qO5f
jgMX5CcHUWUQm7rmn8Ki1Rlbe9sii4W63yWSiJYhUixAQsP+LWapN4VBc0uGTRkK
Udnz9G2k8zcEOqB5XAuHc1ubfvRnNOxq1A9+h8lJuIHhde7MCyi72HCbZWlKwcnu
4wf7IFXNb2CuheGpS7NU5pFsqZZXq980i9JmZ2eZIOhPB6IHkKZlJv7LgxP+bI++
XRdhVq+L+Wm73KO9fXv396UWFknJXFVtg3OZ5pT6feYUUfjS9iP6TVvt+B9tzTAQ
fonVST/+0awYFH0x1+NhcBhJ7ZtdKUQ5twIBAgICAOE=
-----END DH PARAMETERS-----`

func Test_ImportRsaKeyFromPem(t *testing.T) {
	tempFile, err := os.CreateTemp(".", "testfile")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	_, err = tempFile.Write([]byte(testRsaPemContents))
	assert.NoError(t, err)

	err = tempFile.Close()
	assert.NoError(t, err)

	key, err := ImportRsaKeyFromPem(tempFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, key)
}

func Test_ImportDhParametersFromPem(t *testing.T) {
	tempFile, err := os.CreateTemp(".", "testfile")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	_, err = tempFile.Write([]byte(testDhParamsContents))
	assert.NoError(t, err)

	err = tempFile.Close()
	assert.NoError(t, err)

	key, err := ImportDhParametersFromPem(tempFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, key)
}

func Test_SignRsa(t *testing.T) {
	tempFile, err := os.CreateTemp(".", "testfile")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	_, err = tempFile.Write([]byte(testRsaPemContents))
	assert.NoError(t, err)

	err = tempFile.Close()
	assert.NoError(t, err)

	key, err := ImportRsaKeyFromPem(tempFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, key)

	sig, err := SignRsa([]byte("hello"), key)
	assert.NoError(t, err)
	assert.NotNil(t, sig)
}
