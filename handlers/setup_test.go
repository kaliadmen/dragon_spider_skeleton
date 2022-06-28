package handlers

import (
	dragonSpider "github.com/kaliadmen/dragon_spider"
	"github.com/kaliadmen/dragon_spider/render"
	"github.com/kaliadmen/mailer"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var ds dragonSpider.DragonSpider
var testSession *scs.SessionManager
var testHandlers Handlers

func TestMain(m *testing.M) {
	//setup loggers
	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//setup session
	testSession = scs.New()
	testSession.Lifetime = 24 * time.Hour
	testSession.Cookie.Persist = true
	testSession.Cookie.SameSite = http.SameSiteLaxMode
	testSession.Cookie.Secure = false

	var views = jet.NewSet(
		jet.NewOSFileSystemLoader("../views"),
		jet.InDevelopmentMode(),
	)

	testRenderer := render.Render{
		Renderer:    "jet",
		JetTemplate: views,
		Session:     testSession,
		RootPath:    "../",
		Secure:      false,
		Port:        "8492",
	}

	//create dragon spider instance
	ds = dragonSpider.DragonSpider{
		AppName:       "${APP_NAME}",
		Debug:         true,
		Version:       "1.0.0",
		ErrorLog:      errorLog,
		InfoLog:       infoLog,
		Render:        &testRenderer,
		JetTemplate:   views,
		Routes:        nil,
		Session:       testSession,
		Db:            dragonSpider.Database{},
		RootPath:      "../",
		EncryptionKey: ds.RandomString(32),
		Cache:         nil,
		Scheduler:     nil,
		Mail:          mailer.Mail{},
		Server:        dragonSpider.Server{},
	}

	testHandlers.App = &ds

	os.Exit(m.Run())
}

func getRoutes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(ds.SessionLoadAndSave)
	mux.Get("/", testHandlers.Home)

	fileServer := http.FileServer(http.Dir("./../public"))
	mux.Handle("/public/*", http.StripPrefix("/public", fileServer))
	return mux
}

func getContext(req *http.Request) context.Context {
	ctx, err := testSession.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}

	return ctx
}
