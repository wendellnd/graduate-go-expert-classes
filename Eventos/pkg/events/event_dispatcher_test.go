package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()

	suite.handler = TestEventHandler{
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.handler3 = TestEventHandler{
		ID: 3,
	}

	suite.event = TestEvent{
		Name:    "test",
		Payload: "test payload",
	}

	suite.event2 = TestEvent{
		Name:    "test2",
		Payload: "test payload 2",
	}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.NoError(err)

	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.NoError(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Error(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler2)

	suite.eventDispatcher.Clear()

	suite.Equal(0, len(suite.eventDispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler2)

	suite.True(suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))
	suite.True(suite.eventDispatcher.Has(suite.event2.GetName(), &suite.handler2))
	suite.False(suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler2))
	suite.False(suite.eventDispatcher.Has(suite.event2.GetName(), &suite.handler3))
}

type MockHandler struct {
	mock.Mock
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)

	suite.True(suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))
	suite.True(suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler2))

	suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler)

	suite.False(suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))
	suite.True(suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler2))

}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	defer wg.Done()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.event)
	suite.eventDispatcher.Register(suite.event.GetName(), eh)
	suite.eventDispatcher.Dispatch(&suite.event)

	eh.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
