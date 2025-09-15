package api

import (
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestHelloRequest(t *testing.T) {
	req := &HelloRequest{
		Name: "test",
	}

	if req.Name != "test" {
		t.Errorf("Expected name to be 'test', got %s", req.Name)
	}
}

func TestHelloResponse(t *testing.T) {
	now := time.Now()
	req := &HelloResponse{
		Message:   "Hello, test!",
		CreatedAt: timestamppb.New(now),
	}

	if req.Message != "Hello, test!" {
		t.Errorf("Expected message to be 'Hello, test!', got %s", req.Message)
	}

	if req.CreatedAt == nil {
		t.Error("Expected CreatedAt to be set")
	}
}

func TestTimeResponse(t *testing.T) {
	now := time.Now()
	req := &TimeResponse{
		CurrentTime: timestamppb.New(now),
		Timezone:    "UTC",
	}

	if req.CurrentTime == nil {
		t.Error("Expected CurrentTime to be set")
	}

	if req.Timezone != "UTC" {
		t.Errorf("Expected timezone to be 'UTC', got %s", req.Timezone)
	}
}

func TestEmptyProto(t *testing.T) {
	// 测试 Empty 类型的使用
	empty := &emptypb.Empty{}
	if empty == nil {
		t.Error("Empty should not be nil")
	}
}
