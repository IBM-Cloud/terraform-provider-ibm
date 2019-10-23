package accountv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/client"
	bluemixHttp "github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/session"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Accountsv1", func() {

	accountGuid := "9a0d1cdd086428060e43b333decd27dd"
	userEmail := "praveek9@in.ibm.com"
	userGuid := "e9021a4d06e9b108b4a221a3cec47e3d"

	accountInviteResponse := AccountInviteResponse{
		Id:    userGuid,
		Email: userEmail,
		State: "PENDING",
	}

	jsonAccountInviteResponse, _ := json.Marshal(accountInviteResponse)

	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	Describe("Invite user to account", func() {
		Context("Server returns response", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost,
							fmt.Sprintf("/v1/accounts/%s/users", accountGuid)),
						ghttp.RespondWith(http.StatusOK, jsonAccountInviteResponse),
					),
				)
			})

			It("Should return a response", func() {
				resp, err := newAccounts(server.URL()).InviteAccountUser(accountGuid, userEmail)
				Expect(err).To(Succeed())
				Expect(resp).ShouldNot(BeNil())
				Expect(resp).Should(Equal(accountInviteResponse))
			})
		})
	})

	Describe("Delete a user from an account", func() {
		Context("Deleting user", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete,
							fmt.Sprintf("/v1/accounts/%s/users/%s", accountGuid, userGuid)),
						ghttp.RespondWith(http.StatusOK, jsonAccountInviteResponse),
					),
				)
			})

			It("Should not return an error", func() {
				resp := newAccounts(server.URL()).DeleteAccountUser(accountGuid, userGuid)
				Expect(resp).To(Succeed())
				Expect(resp).Should(BeNil())
			})
		})
	})
})

func newAccounts(url string) Accounts {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.AccountService,
	}

	return newAccountAPI(&client)
}
