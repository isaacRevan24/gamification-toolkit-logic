package tests

import (
	"errors"
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

var _ = Describe("User tests", func() {

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

	It("Sign up failed", func() {
		// Mock
		mockUserRepository.EXPECT().SignUpRepository(gomock.Any()).Return(errors.New("SQL error"))

		// Given
		var userId model.SignUpRequest
		userId.ID = "1111111"

		// When
		response := underTest.SignUpController(userId)

		// Then
		var expected model.SignUpResponse
		expected.Status.Message = "Error saving the user"
		expected.Status.Code = model.BAD_REQUEST_ERROR_STATUS

		Expect(response).Should(Equal(expected))

	})

	AfterEach(func() {
		mockCtrl.Finish()
	})
})

var _ = Describe("Habit tests", func() {

	var (
		mockCtrl           *gomock.Controller
		mockHabitRepostory *mock.MockHabitRepository
		underTest          controller.HabitControllerInterface
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockHabitRepostory = mock.NewMockHabitRepository(mockCtrl)
		underTest = controller.NewHabitController(mockHabitRepostory)
	})

	It("Add new habit success", func() {
		// Mock
		mockHabitRepostory.EXPECT().AddNewHabitRepository(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(1, nil)

		// Given
		var newHabitRequest model.AddNewHabitRequest
		newHabitRequest.UserId = "userId"
		newHabitRequest.Condition = "Y"
		newHabitRequest.Description = "testDescription"
		newHabitRequest.Name = "habit name"
		newHabitRequest.Repetition = 1

		// When
		response := underTest.AddNewHabitController(newHabitRequest)

		// Then

		var expected model.AddNewHabitResponse
		expected.Status.Code = model.SUCCESS_CODE_STATUS
		expected.Status.Message = "New habit created"
		expected.HabitId = 1

		Expect(response).Should(Equal(expected))

	})

	It("Add new habit fail", func() {
		// Mock
		mockHabitRepostory.EXPECT().AddNewHabitRepository(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(0, errors.New("SQL error"))

		// Given
		var newHabitRequest model.AddNewHabitRequest
		newHabitRequest.UserId = "userId"
		newHabitRequest.Condition = "Y"
		newHabitRequest.Description = "testDescription"
		newHabitRequest.Name = "habit name"
		newHabitRequest.Repetition = 1

		// When
		response := underTest.AddNewHabitController(newHabitRequest)

		// Then
		var expected model.AddNewHabitResponse
		expected.Status.Code = model.BAD_REQUEST_ERROR_STATUS
		expected.Status.Message = "Error creating new habit."

		Expect(response).Should(Equal(expected))

	})

})
