package models

import (
	"fmt"
	"strings"
	"crypto/rand"
	"crypto/ecdsa"
	//
	"github.com/dgrijalva/jwt-go"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models"
)

type Project struct {
	models.Proto
	Public bool `json:"public"`
	Title string `json:"title"`
	StartBlock int `json:"startBlock"`
	Description string `json:"description"`
	DefaultStream string `json:"defaultStream"`
	BurnAddress string `json:"burnAddress"`
}

func (self *Project) GetUID() string {
	return self.UID
}

func (self *Project) Tree() []string {
	return []string{
		self.UID,
	}
}

func (self *Project) AddressDefaultStream(args ...string) string {
	if len(args) > 0 {
		return fmt.Sprintf("%s.%s", self.DefaultStream, args[0])
	}
	return self.DefaultStream
}

// constructors

func (project *Project) NewApiToken(key *ecdsa.PrivateKey, resources []string) (*ApiToken, error) {

	tokenString, err := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		&ProjectClaims{
			Project: project.UID,
			Resources: resources,
		},
	).SignedString(key)
	if err != nil {
		return nil, err
	}

	return &ApiToken{
		ProjectProto: ProjectProto{
			Project: project.UID,
		},
		Token: tokenString,
		Digest: fmt.Sprintf("%x", models.Hash128([]byte(tokenString))),
		Resources: strings.Join(resources, ","),
	}, nil
}

func (project *Project) NewUser() *User {
	return &User{
		ProjectProto: ProjectProto{
			Project: project.UID,
		},
	}
}

func (project *Project) NewAgent() *Agent {
	return &Agent{
		ProjectProto: ProjectProto{
			Project: project.UID,
		},
	}
}

func (project *Project) NewCurrency() *Currency {

	b := make([]byte, 16)
	rand.Read(b)

	return &Currency{
		ProjectProto: ProjectProto{
			Project: project.UID,
		},
		Alias: fmt.Sprintf("%x", b),
	}
}

func (project *Project) NewStream() *Stream {
	return &Stream{
		ProjectProto: ProjectProto{
			Project: project.UID,
		},
	}
}

//giveCurrency *Currency, giveQuantity float64, recvCurrency *Currency, recvQuantity float64, tx string
func (project *Project) NewExchange() *Exchange {
	return &Exchange{
		ProjectProto: ProjectProto{
			Project: project.UID,
		},
	}
}

func (project *Project) NewFeedMessage(subject models.Child, messageType, target string) (*FeedMessage, error) {
	return &FeedMessage{
		ProjectProto: ProjectProto{
			Project: project.UID,
		},
		Target: target,
		Type: messageType,
		Subject: subject.GetUID(),
		Parents: strings.Join(subject.Tree(), ","),
	}, nil
}
