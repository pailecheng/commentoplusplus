package main

func main() {
	exitIfError(loggerCreate())
	exitIfError(configParse())
	exitIfError(dbConnect(5))
	exitIfError(performMigrations())
	exitIfError(smtpConfigure())
	exitIfError(smtpTemplatesLoad())
	exitIfError(oauthConfigure())
	exitIfError(markdownRendererCreate())
	exitIfError(sigintCleanupSetup())

	exitIfError(routesServe())
}
