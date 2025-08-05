module example.com/hello

go 1.24.5

replace example.com/greetings => ../02-greetings-create-module

require example.com/greetings v0.0.0-00010101000000-000000000000
