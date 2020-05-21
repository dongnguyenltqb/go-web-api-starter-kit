package infra

func Setup() {
	setupMongo()
	setupPostgresql()
	testWrite()
}
