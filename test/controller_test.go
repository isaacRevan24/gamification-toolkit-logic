package tests

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isaacRevan24/gamification-toolkit-logic/controller"
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository/mock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Test suit")
}

var _ = Describe("Hello func", func() {

	var (
		mockCtrl           *gomock.Controller
		mockUserRepository *mock.MockUserRepository
		underTest          controller.UserControllerInterface
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockUserRepository = mock.NewMockUserRepository(mockCtrl)
		underTest = controller.NewUserController(mockUserRepository)
	})

	It("Sign Up successfully", func() {
		// Mock
		mockUserRepository.EXPECT().SignUpRepository(gomock.Any())

		// Given
		var userId model.SignUpRequest
		userId.ID = "1111111"

		// When
		response := underTest.SignUpController(userId)

		// Then
		var expected model.SignUpResponse
		expected.Status.Message = "Successfully saved user."
		expected.Status.Code = model.SUCCESS_CODE_STATUS

		Expect(response).Should(Equal(expected))

	})

	AfterEach(func() {
		mockCtrl.Finish()
	})
})
