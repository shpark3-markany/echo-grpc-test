package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"local/fin/models"
	pb "local/fin/protos"
	"local/fin/utils"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"google.golang.org/grpc"
)

type ReturnUserModel struct {
	User *models.UserModel
}
type ReturnMessage struct {
	Message string
}
type ReturnError struct {
	ErrorMessage string
}

func InvalidParams(params ...string) ReturnError {
	join_params := strings.Join(params, ", ")
	var err_message = ReturnError{ErrorMessage: fmt.Sprintf("invalid request params %s", join_params)}
	return err_message
}

func WebLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body = make(map[string]interface{})
		temp, _ := io.ReadAll(c.Request().Body)
		_ = json.Unmarshal(temp, &body)
		c.Request().Body = io.NopCloser(bytes.NewBuffer(temp))
		params := c.QueryParams()
		var flatten_params = make(map[string]string)
		for key, value := range params {
			flatten_params[key] = value[0]
		}

		c.Response().Before(func() {
			log := utils.NewLogger(0).WithFields(logrus.Fields{
				"path":   c.Path(),
				"ip":     c.RealIP(),
				"status": c.Response().Status,
				"rid":    c.Response().Header().Get(echo.HeaderXRequestID),
			})
			if len(body) > 0 {
				log = log.WithField("body", body)
			}
			if len(flatten_params) > 0 {
				log = log.WithField("params", flatten_params)
			}
			log.Info("started call")
		})
		c.Response().After(func() {
			log := utils.NewLogger(0).WithFields(logrus.Fields{
				"path":   c.Path(),
				"ip":     c.RealIP(),
				"status": c.Response().Status,
				"rid":    c.Response().Header().Get(echo.HeaderXRequestID),
			})
			if len(flatten_params) > 0 {
				log = log.WithField("params", flatten_params)
			}
			if len(body) > 0 {
				log = log.WithField("body", body)
			}
			if err := c.Get("error"); err != nil {
				log = log.WithField("error", err)
			}
			log.Info("finished call")
		})
		next(c)
		return nil
	}
}

func GRPCLogger(l logrus.FieldLogger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make(map[string]any, len(fields)/2)
		i := logging.Fields(fields).Iterator()
		for i.Next() {
			k, v := i.At()
			f[k] = v
		}
		l := l.WithFields(f)
		l.Logger.SetFormatter(&logrus.JSONFormatter{})
		switch lvl {
		case logging.LevelDebug:
			l.Debug(msg)
		case logging.LevelInfo:
			l.Info(msg)
		case logging.LevelWarn:
			l.Warn(msg)
		case logging.LevelError:
			l.Error(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func OpenRpc(grpc_port string) {
	defer func() {
		if r := recover(); r != nil {
			log.Print(r)
		}
	}()

	lis, err := net.Listen("tcp", ":"+grpc_port)
	if err != nil {
		panic(fmt.Sprintf("Port listen failed: %v", err))
	}

	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(GRPCLogger(logrus.New()), opts...),
	))
	pb.RegisterUserInfoServer(grpcServer, UserInfoServer{})
	defer lis.Close()

	log.Printf("gRPC server opening on %v port", grpc_port)
	if err := grpcServer.Serve(lis); err != nil {
		panic(fmt.Errorf("'grpcServer' serve failed: %v", err))
	}
}

func OpenWeb(web_port string) {
	e := echo.New()
	defer e.Close()
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(WebLogger)

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/swagger/index.html")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "Status ok")
	})

	e.GET("/get", GetUser)
	e.GET("/list", ListUser)
	e.POST("/create", CreateUser)
	e.PUT("/update", UpdateUser)
	e.DELETE("/delete", DeleteUser)
	e.POST("/batch-save", BatchSave)
	e.DELETE("/batch-delete", BatchDelete)
	e.DELETE("/reset", Reset)
	e.Logger.Fatal(e.Start(":" + web_port))
}
