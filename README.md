# t_bot

# - ldflags зміна змінних у модулях, у модулі cmd та змінна appVersion
go build -ldflags "-X="github.com/jogize/t_bot/cmd.appVersion=1.0.1 -o tbot

