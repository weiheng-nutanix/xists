// Copyright (c) 2020 Nutanix Inc. All rights reserved.
// Author: gokul.kannan@nutanix.com
// Package to handle IAM token requests from Service Accounts.


package token_handler

import (
    "crypto/sha256"
    "crypto/rand"
    "crypto/rsa"
    "encoding/base64"
    "fmt"
    "github.com/dgrijalva/jwt-go"
    "gopkg.in/square/go-jose.v2"
    "time"
)

func HandleToken(token string, audience string) (*jwt.Token, string){
    claims, _ := parseToken(token)
    
    // XXX: Query underlying DB for service account for the issuer claims["iss"]
    // extract service account uuid.
    
    // Query the DB for service provider with the tenant uuid and the audience
    
    thumbprint := getSha256Base64Encoded(claims["public_key"].(string))
    
    // Check if the tumbprint matches the thumbprints stored in the DB for the
    // service account.
    
    // Fill appropriate subject
    sub := "hkthon" + "@" + "iam"

    // Issue the token
    return issueToken(sub, audience)
    
}

func issueToken(sub string, aud string) (*jwt.Token, string){
    nbfDeltaMins := int64(10)
    expDeltaMins := int64(60)
    ia := time.Now()
    iat := ia.Unix()
    claim := &jwt.StandardClaims{
        Audience: aud,
        IssuedAt: iat,
        ExpiresAt: iat + expDeltaMins * 60,
        NotBefore: iat - nbfDeltaMins * 60,
        Subject: sub,
        Issuer: "IAM issuer",
    }
    token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)

    // Get the IAM's private key from the config.
    key, _ := rsa.GenerateKey(rand.Reader, 2048)
    signature, err := token.SignedString(key)
    //signature, _ := token.SigningString()
    return token, signature
}


func parseToken(token string) (mapClaims jwt.MapClaims, err error) {
	jsonWebSignature, err := jose.ParseSigned(token)
	if err != nil {
		return nil, err
	}

	claims, _ := parseTokenClaims(token)
	jwtClaims := claims.(jwt.MapClaims)
	publicKeyStr := getPublicKey(jwtClaims)

	publicKeyRaw := []byte(publicKeyStr)

	publicKey := &jose.JSONWebKey{}

	publicKey.UnmarshalJSON(publicKeyRaw)

	_, err = jsonWebSignature.Verify(publicKey)
	if err != nil {
		return nil, err
	}
	return jwtClaims, nil
}

func parseTokenClaims(tokenStr string) (claims jwt.Claims, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	return token.Claims.(jwt.MapClaims), err
}


func getPublicKey(claims jwt.MapClaims) string {
	return claims["public_key"].(string)
}

func getSha256Base64Encoded(str string) string {
	hash := sha256.Sum256([]byte(str))
	return base64.URLEncoding.EncodeToString(hash[:])
}