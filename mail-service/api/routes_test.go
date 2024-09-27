package api

import (
	"context"
	"log"
	"mail-service/api/handlers"
	"mail-service/client"
	"mail-service/proto"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/inbucket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestMailService(t *testing.T) {
	// Load environment variables
	os.Setenv("GRPC_MAIL_PORT", "8000")

	// ------------------ Server Setup ------------------
	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	srv := grpc.NewServer()
	t.Cleanup(func() {
		srv.Stop()
	})

	msc := MailService{
		MailHandler: &handlers.MailHandler{},
	}
	proto.RegisterMailServiceServer(srv, &msc)

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	// ------------------ SMTP Server Setup ------------------
	ctx := context.Background()
	inbucketContainer := SetupMailSmtp(t, ctx)
	t.Cleanup(func() {
		duration := 5 * time.Second
		inbucketContainer.Stop(ctx, &duration)
	})

	// ------------------ Client Setup ------------------
	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	conn := client.NewTestingGrpcConnection(dialer)
	t.Cleanup(func() {
		conn.Close()
	})

	mailClient := proto.NewMailServiceClient(conn)
	_ = mailClient
}

func SetupMailSmtp(t *testing.T, ctx context.Context) *inbucket.InbucketContainer {
	var container *inbucket.InbucketContainer
	t.Run("Test Mail SMTP", func(t *testing.T) {
		inbucketContainer, err := inbucket.Run(ctx, "inbucket/inbucket:sha-2d409bb",
			testcontainers.WithHostPortAccess(9000),
		)
		assert.NoError(t, err)

		container = inbucketContainer
	})

	return container
}
